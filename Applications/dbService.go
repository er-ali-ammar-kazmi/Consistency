package app

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	_ "modernc.org/sqlite"
)

var DB *gorm.DB

type custom interface {
	User | Blog | Role | Address
}

type CRUD[T custom] interface {
	Create() (T, error)
	Update(int32) (T, error)
	Delete(int32) (T, error)
	RetrieveOne(int32) (T, error)
	RetrieveAll() ([]T, error)
}

type User struct {
	gorm.Model
	Id      uint   `gorm:"id; primary key; autoIncrement" json:"id"`
	Name    string `gorm:"name; index" json:"name"`
	Email   string `gorm:"email" json:"email"`
	Blogs   []Blog
	Address Address
	Roles   []Role `gorm:"many2many:users_roles"`
}

type Blog struct {
	Id      uint   `gorm:"id; primary key; autoIncrement" json:"id"`
	Title   string `gorm:"title; index" json:"title"`
	Content string `gorm:"content" json:"content"`
	UserId  uint
}

type Address struct {
	Id          uint   `gorm:"id; primary key; autoIncrement" json:"id"`
	PinCode     uint   `gorm:"pincode" json:"pincode"`
	Street      string `gorm:"street" json:"street"`
	HomeAddress string `gorm:"address" json:"address"`
	UserId      uint
}

type Role struct {
	Id   uint   `gorm:"id; primary key; autoIncrement" json:"id"`
	Role string `gorm:"role" json:"role"`
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

	conn.AutoMigrate(&Blog{}, &User{}, &Address{}, &Role{})

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

func (db dbService) Run() {
	fmt.Println("Started!")
}
