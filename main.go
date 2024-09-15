package main

import (
	"github.com/Achariya1/go-books-crud/database"
	"github.com/gofiber/fiber/v2"
	"github.com/Achariya1/go-books-crud/router"
)

func main (){
	
	database.ConnetDB()
	app := fiber.New()


	router.Initalize(app)
	app.Listen("127.0.0.1:8080")


	
}