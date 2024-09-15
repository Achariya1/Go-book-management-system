package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/jwt/v2"
	"github.com/Achariya1/go-books-crud/middleware"

)

func Initalize(router *fiber.App) {


	router.Use("/test",jwtware.New(jwtware.Config{
		SigningKey: []byte("secret"),
	}))

	router.Post("/signup",middleware.Singup)
	router.Post("/loginn",middleware.Loginn)
	router.Post("/test",middleware.Test)


	
	/*
	router.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("secret"),
	}))

	router.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).SendString("Hello, World!")
	})*/







	


}