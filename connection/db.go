package main

import (
	"nujsua/thirgolang/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:mysql@tcp(127.0.0.1:3306)/test_demo1?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&model.User{})

	// Create
	db.Create(&model.User{Fname: "John", Lname: "Doe", Username: "johndoe", Avatar: "https://www.google.com/images/branding/googlelogo/1x/googlelogo_color_272x92dp.png"})

}
