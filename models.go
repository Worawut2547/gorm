package main

import (
	"log"
	"fmt"

	"gorm.io/gorm"
)

type Book struct{
	ID 			int		
	Name 		string 	`json:"name"`
	Author 		string	`json:"author"`
	Description string	`json:"description"`
	Price 		uint	`json:"price"`
}

func createBook(db *gorm.DB , book *Book){
	result := db.Create(book)

	if result.Error != nil{
		log.Fatalf("Error creating book: %v",result.Error)
	}

	fmt.Println("Create Book Successful")
}

func getBook(db *gorm.DB , id int) *Book{
	
	var book Book
	result := db.First(&book , id)
	if result.Error != nil{
		log.Fatalf("Error get book: %v",result.Error)
	}
	return &book
}

func getBooks(db *gorm.DB) []Book{
	var books []Book
	result := db.Find(&books)
	if result.Error != nil{
		log.Fatalf("Error get book: %v",result.Error)
	}
	return books
}

func updateBook(db *gorm.DB , book *Book) error{
	result := db.Save(&book)

	if result.Error != nil{
		return result.Error
	}
	return nil
}

func deleteBook(db *gorm.DB , id int)  error{
	var book Book
	result := db.Delete(&book , id)

	if result.Error != nil{
		return result.Error
	}

	return nil
}