package handlers

import (
	"os"
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
        return fiber.NewError(fiber.StatusBadRequest, "Invalid request body format")
	}

	if request.Username == "" || request.Password == "" {
        return fiber.NewError(fiber.StatusUnprocessableEntity, "Username and Password are required")
	}

	password, err := bcrypt.GenerateFromPassword([]byte(request.Password), 10)
	if err != nil {
		return fiber.NewError(fiber.StatusUnprocessableEntity, err.Error())

	}

	request.Password = string(password)

	db := database.DB

	result := db.Create(&request)
	if result.Error != nil {
        return fiber.NewError(fiber.StatusInternalServerError, "Error creating user: " + result.Error.Error())
	}

	return c.JSON(request)
}


func Login(c *fiber.Ctx) error {

	request := model.User{}
	if err := c.BodyParser(&request); err != nil {
        return fiber.NewError(fiber.StatusBadRequest, "Invalid request body format")
	}

	db := database.DB


	user := model.User{}
	result := db.Where("username = ? ", request.Username).First(&user)


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
	t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.JSON(fiber.Map{
		"username": user.Username,
		"password": user.Password,
		"jwtToken": t,
	})
}



