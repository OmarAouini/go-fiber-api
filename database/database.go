package database

import (
	"fmt"
	"os/user"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	Pool *gorm.DB
)

func InitDatabase() {
	//connection config
	dsn := "root:root@tcp(127.0.0.1:3306)/shop?charset=utf8mb4&parseTime=True&loc=Local"
	conn, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	//assign db pool to global variable
	Pool = conn
	fmt.Println("database connection opened!")

	//auto migration
	//add here new structs to manage migration
	Pool.AutoMigrate(&user.User{})
	fmt.Println("database migrated!")
}
