package app

import (
	"fmt"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	_ "modernc.org/sqlite"
)

var DB *gorm.DB

type User struct {
	Id    uint   `gorm:"id;primaryKey;autoIncrement" json:"id"`
	Name  string `gorm:"name;size:50;notNull; index" json:"name"`
	Email string `gorm:"email;notNull;unique" json:"email"`

	Blogs   []Blog  `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"blogs"`
	Address Address `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"address"`
	Roles   []Role  `gorm:"many2many:users_roles;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"roles"`
}

type Blog struct {
	Id      uint   `gorm:"id;primaryKey;autoIncrement" json:"id"`
	Title   string `gorm:"title;size:255;notNull;index" json:"title"`
	Content string `gorm:"content;type:text;notNull" json:"content"`
	UserId  uint   `json:"user_id"`
}

type Address struct {
	Id          uint   `gorm:"id;primaryKey;autoIncrement" json:"id"`
	PinCode     uint   `gorm:"pincode" json:"pincode"`
	HomeAddress string `gorm:"address" json:"address"`
	UserId      uint   `json:"user_id"`
}

type Role struct {
	Role string `gorm:"role;primarykey;size:25" json:"role"`
}

type dbService struct{}

func DbService() dbService {
	return dbService{}
}

func (r dbService) ConnectSqlite() {
	defer fmt.Println("Connection with SqliteDataBase Initialized!")

	conn, err := gorm.Open(sqlite.Dialector{
		DriverName: "sqlite",
		DSN:        "file:SqliteDataBase.db?cache=shared&mode=rwc"},
		&gorm.Config{})
	if err != nil {
		fmt.Printf("Error: %v", err.Error())
		os.Exit(1)
	}

	conn.Exec("PRAGMA foreign_keys = ON;")
	m := conn.Migrator()

	if !m.HasTable(&User{}) || !m.HasTable(&Role{}) {
		m.CreateTable(&User{}, &Role{})
	}
	if !m.HasTable(&Blog{}) {
		m.CreateTable(&Blog{})
	}
	if !m.HasTable(&Address{}) {
		m.CreateTable(&Address{})
	}

	DB = conn
}

func (r dbService) CloseSqlite() {
	defer fmt.Println("Connection with SqliteDataBase Closed!")

	if sqliteDb, err := DB.DB(); err != nil {
		fmt.Println(err.Error())
	} else {
		if err := sqliteDb.Close(); err != nil {
			fmt.Println(err.Error())
		}
	}
}

func (r dbService) InsertUser(record *User) (bool, error) {
	result := DB.Create(record)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

func (r dbService) RetrieveUser(id uint) (User, error) {
	var record User
	result := DB.Preload("Blogs").Preload("Address").Preload("Roles").First(&record, id)
	if result.Error != nil {
		return record, result.Error
	}
	return record, nil
}

func (r dbService) RetrieveAllUsers() ([]User, error) {
	var records []User
	result := DB.Preload("Blogs").Preload("Address").Preload("Roles").Find(&records)
	if result.Error != nil {
		return records, result.Error
	}
	return records, nil
}

func (r dbService) UpdateUser(input *User) (User, error) {
	result := DB.Save(input)
	if result.Error != nil {
		return *input, result.Error
	}
	return *input, nil
}

func (r dbService) DeleteUser(id uint) (bool, error) {
	result := DB.Delete(&User{}, id)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

func (r dbService) InsertBlog(record *Blog) (bool, error) {
	result := DB.Create(record)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

func (r dbService) RetrieveBlog(id uint) (Blog, error) {
	var record Blog
	result := DB.First(&record, id)
	if result.Error != nil {
		return record, result.Error
	}
	return record, nil
}

func (r dbService) RetrieveAllBlogs() ([]Blog, error) {
	var records []Blog
	result := DB.Find(&records)
	if result.Error != nil {
		return records, result.Error
	}
	return records, nil
}

func (r dbService) UpdateBlog(input *Blog) (Blog, error) {
	result := DB.Save(input)
	if result.Error != nil {
		return *input, result.Error
	}
	return *input, nil
}

func (r dbService) DeleteBlog(id uint) (bool, error) {
	result := DB.Delete(&Blog{}, id)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

func (r dbService) RetrieveAddress(id uint) (Address, error) {
	var record Address
	result := DB.First(&record, id)
	if result.Error != nil {
		return record, result.Error
	}
	return record, nil
}

func (r dbService) UpdateAddress(input *Address) (Address, error) {
	result := DB.Save(input)
	if result.Error != nil {
		return *input, result.Error
	}
	return *input, nil
}

func (r dbService) DeleteAddress(id uint) (bool, error) {
	result := DB.Delete(&Address{}, id)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

func (r dbService) InsertRole(record *Role) (bool, error) {
	result := DB.Create(record)
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}

func (r dbService) RetrieveAllRoles() ([]Role, error) {
	var record []Role
	result := DB.Find(&record)
	if result.Error != nil {
		return record, result.Error
	}
	return record, nil
}

func (r dbService) DeleteRole(key string) (bool, error) {
	result := DB.Where("Role = ?", key).Delete(&Role{})
	if result.Error != nil {
		return false, result.Error
	}
	return true, nil
}
