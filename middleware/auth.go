
package middleware

import (
	"fmt"
	"time"
	"github.com/Achariya1/go-books-crud/database"
	"github.com/Achariya1/go-books-crud/model"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func Singup(c *fiber.Ctx) error {

	request := model.User{}
	err := c.BodyParser(&request)
	if err != nil {
		//return fiber.ErrBadRequest
		return c.SendString("humm")
	}

	if request.Username == "" || request.Password == "" {
		return fiber.NewError(fiber.StatusUnprocessableEntity, "dog")
	}

	password, err := bcrypt.GenerateFromPassword([]byte(request.Password), 10)
	if err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, err.Error())

	}
	//change to hashed password
	fmt.Println(request.Password)
	request.Password = string(password)

	db := database.DB

	result := db.Create(&request)
	if result.Error != nil {
		return c.SendString("Error creating user: " + result.Error.Error())
	}

	fmt.Println(request.Password)

	return c.JSON(request)
}

func Loginn(c *fiber.Ctx) error {

	request := model.User{}
	if err := c.BodyParser(&request); err != nil {
		return fiber.NewError(fiber.StatusBadRequest)
	}

	db := database.DB

	fmt.Println("check")

	user := model.User{}
	result := db.Where("username = ? ", request.Username).First(&user)

	fmt.Println(" this is user : ", user)

	if result.Error != nil {
		return fiber.NewError(fiber.StatusNotFound, "Invalid username or password")
	}

	// Check if password matches
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password)); err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid username or password")
	}


	// Create token
	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = user.Username
	claims["role"] = "admin"
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{
		"username" : user.Username,
		"password" : user.Password,
		"jwtToken" : t,
	})
}

func Test(c *fiber.Ctx) error{
	return c.SendString("hello world")
}




func Hello(c *fiber.Ctx) error {

	fmt.Println("1111111")
	book := model.Book{}
	if err := c.BodyParser(&book); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
	}
	fmt.Println("2222222")

	db := database.DB


	fmt.Println("33333333")

	result := db.Create(&book)
	if result.Error != nil {
		return fiber.NewError(fiber.StatusInternalServerError, "Can't create data: "+result.Error.Error())
	}
	fmt.Println("4444444")

	return c.JSON(book)

}

