package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID		Unit	`gorm:"primarykey"`
	Name	string
	Email	string	`gorm:"UniqueIndex"`
}

func main()  {
	dsn := "username:userpassword@tcp(localhost:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&User{})

	//以降はCRUD操作記述欄
}