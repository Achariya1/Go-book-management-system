package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/Achariya1/go-books-crud/database"
	"github.com/Achariya1/go-books-crud/model"
)



func CreateBook(c *fiber.Ctx) error {
    // Parse the request body into the book model
    book := model.Book{}
    if err := c.BodyParser(&book); err != nil {
        return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
    }

    db := database.DB

    // Check if there are authors provided in the request
    if len(book.Author) == 0 {
        return fiber.NewError(fiber.StatusBadRequest, "A book must have at least one author")
    }

    // To store authors (new and existing)
    var authors []model.Author

    // Loop through the provided authors
    for _, author := range book.Author {
        // If the author has an ID, assume they are an existing author
        if author.ID != 0 {
            var existingAuthor model.Author
            result := db.First(&existingAuthor, author.ID)
            if result.Error != nil {
                return fiber.NewError(fiber.StatusInternalServerError, "Error finding author: "+result.Error.Error())
            }
            authors = append(authors, existingAuthor)
        } else {
            // If no ID is provided, assume it's a new author
            authors = append(authors, author)
        }
    }

    // Assign the resolved authors to the book
    book.Author = authors

    // Create the book with the authors relationship
    if err := db.Create(&book).Error; err != nil {
        return fiber.NewError(fiber.StatusInternalServerError, "Failed to create book: "+err.Error())
    }

    return c.JSON(book)
}


func GetBooks(c *fiber.Ctx) error {


	db := database.DB

	var book []model.Book
	err := db.Model(&book).Preload("Author").Find(&book).Error
	if err != nil{
		return fiber.NewError(fiber.StatusNotFound,"Book not found")
	}




	return c.JSON(book)

}


func GetBookByID(c *fiber.Ctx) error {
	id := c.Params("id")


	db := database.DB

	var book model.Book
	err := db.Model(&book).Preload("Author").First(&book,id).Error
	if err != nil{
		return fiber.NewError(fiber.StatusNotFound,"Book not found")
	}

	return c.JSON(book)

}



func UpdateBook(c *fiber.Ctx) error {
    id := c.Params("id")

    var book model.Book
    db := database.DB

    // Find the book by its ID
    if err := db.Preload("Author").First(&book, id).Error; err != nil {
        return fiber.NewError(fiber.StatusNotFound, "Book not found")
    }

    // Parse the request body into the book model
    updateData := model.Book{}
    if err := c.BodyParser(&updateData); err != nil {
        return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
    }

    // Clear current authors and update with new ones
    db.Model(&book).Association("Author").Clear()

    // Handle the authors (existing and new)
    var authors []model.Author
    for _, author := range updateData.Author {
        if author.ID != 0 {
            // If the author has an ID, find the existing author
            var existingAuthor model.Author
            if err := db.First(&existingAuthor, author.ID).Error; err != nil {
                return fiber.NewError(fiber.StatusInternalServerError, "Error finding author: "+err.Error())
            }
            authors = append(authors, existingAuthor)
        } else {
            // If no ID, treat it as a new author
            authors = append(authors, author)
        }
    }

    // Update the book data and authors
    book.Title = updateData.Title
    book.Price = updateData.Price
    book.Author = authors

    if err := db.Save(&book).Error; err != nil {
        return fiber.NewError(fiber.StatusInternalServerError, "Failed to update book: "+err.Error())
    }

    return c.JSON(book)
}





func DeleteBook(c *fiber.Ctx) error {

	id := c.Params("id")
	book := model.Book{}

	db := database.DB

	if err := db.First(&book,id).Error ; err != nil{
		return fiber.NewError(fiber.StatusNotFound,"Can't not found")
	}

	
	err := db.Model(&book).Association("Author").Clear()
	if err != nil{
		return fiber.NewError(fiber.StatusInternalServerError,err.Error())
	}


	if err := db.Delete(&book).Error ; err != nil{
		return fiber.NewError(fiber.StatusInternalServerError, "Can't delete book: ")
	}

	return c.SendStatus(fiber.StatusNoContent)
}




