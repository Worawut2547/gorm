package main

import (

	"fmt"
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)


const (
	username = "root"
	password = "1234"
	hostname = "127.0.0.1:3306"
	dbname = "book"
)

func dsn(dbName string) string{
	return fmt.Sprintf("%s:%s@tcp(%s)/%s",username,password,hostname,dbName)
}

type Book struct{
	gorm.Model
	ID int `gorm:"primaryKey"`
	Title string
	Author string
	CreateedAt time.Time
}

func main() {
	// Gorm Connected DB
	//------------------------------------------------------------
	db , err := gorm.Open(mysql.Open(dsn(dbname)),&gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil{
		log.Printf("Failed to connect to DB: %v", err)
	}
	log.Println("Connected to DB successfully")
	//------------------------------------------------------------

	//AutoMigrate , Create Table
	//------------------------------------------------------------
	err = db.AutoMigrate(&Book{})
	if err != nil{
		log.Printf("Failed to migrate: %v", err)
	}
	log.Println("Auto Migration completed")
	//------------------------------------------------------------

	//Insert
	db.Create(&Book{
		Title: "Ironman",
		Author: "Marvel",
		CreateedAt: time.Now(),
	})
}