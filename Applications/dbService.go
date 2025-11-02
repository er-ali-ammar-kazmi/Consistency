package app

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	_ "modernc.org/sqlite"
)

var DB *gorm.DB

type User struct {
	Id      uint    `gorm:"id;primarykey;autoIncrement" json:"id"`
	Name    string  `gorm:"name;size:50;not null; index" json:"name"`
	Email   string  `gorm:"email;notnull;unique" json:"email"`
	Blogs   []Blog  `json:"blogs"`
	Address Address `json:"address"`
	Roles   []Role  `gorm:"many2many:users_roles" json:"roles"`
}

type Blog struct {
	Id      uint   `gorm:"id;primarykey;autoIncrement" json:"id"`
	Title   string `gorm:"title;size:255;notnull;index" json:"title"`
	Content string `gorm:"content;type:text;notnull" json:"content"`
	UserId  uint   `json:"user_id"`
}

type Address struct {
	Id          uint   `gorm:"id;primarykey;autoIncrement" json:"id"`
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
	conn, err := gorm.Open(sqlite.Dialector{
		DriverName: "sqlite",
		DSN:        "file:SqliteDataBase.db?cache=shared&mode=rwc"},
		&gorm.Config{})
	if err != nil {
		fmt.Printf("Error: %v", err.Error())
	}

	db := conn.Migrator()
	if !db.HasTable(&User{}) {
		db.CreateTable(&User{})
	}
	if !db.HasTable(&Blog{}) {
		db.CreateTable(&Blog{})
	}
	if !db.HasTable(&Address{}) {
		db.CreateTable(&Address{})
	}
	if !db.HasTable(&Role{}) {
		db.CreateTable(&Role{})
	}

	DB = conn
	fmt.Println("Connection with SqliteDataBase Initialized!")
}

func (r dbService) CloseSqlite() {
	if sqliteDb, err := DB.DB(); err != nil {
		fmt.Println(err.Error())
	} else {
		if err := sqliteDb.Close(); err != nil {
			fmt.Println(err.Error())
		}
	}
	fmt.Println("Connection with SqliteDataBase Closed!")
}

func (r dbService) InsertUser(record User) User {
	result := DB.Create(&record)
	if result.Error != nil {
		fmt.Println("Error: ", result.Error.Error())
	}
	return record
}

func (r dbService) RetrieveUser(id uint) User {
	var record User
	result := DB.Preload("Blogs").Preload("Address").Preload("Roles").First(&record, id)
	if result.Error != nil {
		fmt.Println("Error: ", result.Error.Error())
	}
	return record
}

func (r dbService) RetrieveAllUsers() []User {
	var record []User
	result := DB.Preload("Blogs").Preload("Address").Preload("Roles").Find(&record)
	if result.Error != nil {
		fmt.Println("Error: ", result.Error.Error())
	}
	return record
}

func (r dbService) UpdateUser(id uint, input User) User {
	result := DB.Save(input)
	if result.Error != nil {
		fmt.Println("Error: ", result.Error.Error())
	}
	return input
}

func (r dbService) DeleteUser(id uint) User {
	var record User
	result := DB.Delete(&record, id)
	if result.Error != nil {
		fmt.Println("Error: ", result.Error.Error())
	}
	return record
}

func (r dbService) InsertBlog(record Blog) Blog {
	result := DB.Create(&record)
	if result.Error != nil {
		fmt.Println("Error: ", result.Error.Error())
	}
	return record
}

func (r dbService) RetrieveBlog(id uint) Blog {
	var record Blog
	result := DB.First(&record, id)
	if result.Error != nil {
		fmt.Println("Error: ", result.Error.Error())
	}
	return record
}

func (r dbService) RetrieveAllBlogs() []Blog {
	var record []Blog
	result := DB.Find(&record)
	if result.Error != nil {
		fmt.Println("Error: ", result.Error.Error())
	}
	return record
}

func (r dbService) UpdateBlog(id uint, input Blog) Blog {
	result := DB.Save(input)
	if result.Error != nil {
		fmt.Println("Error: ", result.Error.Error())
	}
	return input
}

func (r dbService) DeleteBlog(id uint) Blog {
	var record Blog
	result := DB.Delete(&record, id)
	if result.Error != nil {
		fmt.Println("Error: ", result.Error.Error())
	}
	return record
}

func (r dbService) InsertAddress(record Address) Address {
	result := DB.Create(&record)
	if result.Error != nil {
		fmt.Println("Error: ", result.Error.Error())
	}
	return record
}

func (r dbService) RetrieveAddress(id uint) Address {
	var record Address
	result := DB.First(&record, id)
	if result.Error != nil {
		fmt.Println("Error: ", result.Error.Error())
	}
	return record
}

func (r dbService) UpdateAddress(id uint, input Address) Address {
	result := DB.Save(input)
	if result.Error != nil {
		fmt.Println("Error: ", result.Error.Error())
	}
	return input
}

func (r dbService) DeleteAddress(id uint) Address {
	var record Address
	result := DB.Delete(&record, id)
	if result.Error != nil {
		fmt.Println("Error: ", result.Error.Error())
	}
	return record
}

func (r dbService) InsertRole(record Role) Role {
	result := DB.Create(&record)
	if result.Error != nil {
		fmt.Println("Error: ", result.Error.Error())
	}
	return record
}

func (r dbService) RetrieveRole(id uint) Role {
	var record Role
	result := DB.First(&record, id)
	if result.Error != nil {
		fmt.Println("Error: ", result.Error.Error())
	}
	return record
}

func (r dbService) RetrieveAllRoles() []Role {
	var record []Role
	result := DB.Find(&record)
	if result.Error != nil {
		fmt.Println("Error: ", result.Error.Error())
	}
	return record
}

func (r dbService) DeleteRole(id uint) Role {
	var record Role
	result := DB.Delete(&record, id)
	if result.Error != nil {
		fmt.Println("Error: ", result.Error.Error())
	}
	return record
}
