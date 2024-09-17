package main

import (
	
	"github.com/Achariya1/go-books-crud/database"
	"github.com/gofiber/fiber/v2"
	"github.com/Achariya1/go-books-crud/router"
	"github.com/joho/godotenv"
	"log"
)

func main (){
	

	err := godotenv.Load(".env")
  	if err != nil {
    	log.Fatalf("Error loading .env file")
  	}

	

	database.ConnetDB()
	app := fiber.New()


	router.Initalize(app)
	app.Listen("127.0.0.1:8080")


	
}