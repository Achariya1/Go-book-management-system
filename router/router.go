package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/jwt/v2"
	"github.com/Achariya1/go-books-crud/handlers"
	"os"

)

func Initalize(router *fiber.App) {


	router.Post("/signup",handlers.Singup)
	router.Post("/login",handlers.Login)


	router.Get("/book",handlers.GetBooks)
	router.Get("/book/:id",handlers.GetBookByID)

	router.Get("/author",handlers.GetAuthors)
	router.Get("/author/:id",handlers.GetAuthorByID)



	router.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte(os.Getenv("JWT_SECRET")),
	}))


	router.Post("/book",handlers.CreateBook)
	router.Put("/book/:id",handlers.UpdateBook)
	router.Delete("/book/:id",handlers.DeleteBook)

	router.Post("/author",handlers.CreateAuthor)
	router.Put("/author/:id",handlers.UpdateAuthor)
	router.Delete("/author/:id",handlers.DeleteAuthor)











	


}