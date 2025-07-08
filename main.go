package main

import (

	"fmt"
	"log"
	"strconv"
	//"time"

	"github.com/gofiber/fiber/v2"
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
	//log.Println("Auto Migration completed")
	//------------------------------------------------------------ 


	//Set up fiber
	//------------------------------------------------------------ 
	app := fiber.New()
	//------------------------------------------------------------ 

	// Create
	//------------------------------------------------------------ 
	app.Post("/book",func(c *fiber.Ctx) error {
		book := new(Book)

		if err := c.BodyParser(book); err != nil{
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}

		createBook(db,book)
		return c.JSON(book)
	})
	//------------------------------------------------------------ 

	// Read All
	//------------------------------------------------------------ 
	app.Get("/books" ,func(c *fiber.Ctx) error {
		return c.JSON(getBooks(db))
	})

	// Read id
	//------------------------------------------------------------ 
	app.Get("/book/:id",func(c *fiber.Ctx) error{
		bookId , err := strconv.Atoi(c.Params("id"))
		if err != nil{
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}
		book := getBook(db,bookId)
		return c.JSON(book)
	})
	//------------------------------------------------------------ 

	// Update
	//------------------------------------------------------------ 
	app.Put("/book/:id",func(c *fiber.Ctx) error{
		bookId , err := strconv.Atoi(c.Params("id"))

		if err != nil{
			return c.SendStatus(fiber.StatusBadRequest)
		}
		
		book := new(Book)

		if err := c.BodyParser(book); err != nil{
			return c.SendStatus(fiber.StatusBadRequest)
		}

		book.ID = bookId
		err = updateBook(db,book)

		if err != nil{
			return c.SendStatus(fiber.StatusBadRequest)
		}
		return c.JSON(book)
	})
	//------------------------------------------------------------
	
	// Delete
	//------------------------------------------------------------
	app.Delete("/book/:id",func(c *fiber.Ctx) error{
		id , err := strconv.Atoi(c.Params("id"))
		if err != nil{
			return c.SendStatus(fiber.StatusBadRequest)
		}

		err = deleteBook(db,id)

		if err != nil{
			return c.SendStatus(fiber.StatusBadRequest)
		}

		return c.JSON(fiber.Map{
			"Message": "Delete Complete",
		})
	})
	//------------------------------------------------------------

	log.Println("Server is Running on Port 6000")
	app.Listen(":6000")
}