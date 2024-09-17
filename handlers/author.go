package handlers

import (

	"github.com/Achariya1/go-books-crud/database"
	"github.com/Achariya1/go-books-crud/model"
	"github.com/gofiber/fiber/v2"
)




func CreateAuthor(c *fiber.Ctx) error {
    // Parse the request body into the author model
    author := model.Author{}
    if err := c.BodyParser(&author); err != nil {
        return fiber.NewError(fiber.StatusBadRequest, "Invalid request body")
    }


    db := database.DB

    // Check if there are books provided in the request
    if len(author.Book) == 0 {
        return fiber.NewError(fiber.StatusBadRequest, "An author must has one book")
    }

    // For each book, check if they already exist in the database
    var books []model.Book
    for _, book := range author.Book {
        var existingBook model.Book
        if book.ID != 0 {
			result := db.First(&existingBook,book.ID)
			if result.Error != nil{
				return fiber.NewError(fiber.StatusNotFound, "Book not found: "+book.Title) 
			}else{
				books = append(books, existingBook)
			}
			
		}else {
            // If book not exists, create new book
            books = append(books, book)
        }
        
    }

    // Set the author's books
    author.Book = books

    // Create the author along with the relationship to the books
    if err := db.Create(&author).Error; err != nil {
        return fiber.NewError(fiber.StatusInternalServerError, "Failed to create author: "+err.Error())
    }

    return c.JSON(author)
}



func GetAuthors(c *fiber.Ctx) error {


	db := database.DB

	var author []model.Author
	err := db.Model(&author).Preload("Book").Find(&author).Error
	if err != nil{
		return fiber.NewError(fiber.StatusNotFound,"Author not found")
	}


	return c.JSON(author)

}


func GetAuthorByID(c *fiber.Ctx) error {
	id := c.Params("id")


	db := database.DB

	var author model.Author
	err := db.Model(&author).Preload("Book").First(&author,id).Error
	if err != nil{
		return fiber.NewError(fiber.StatusNotFound,"Author not found")
	}

	return c.JSON(author)

}


func UpdateAuthor(c *fiber.Ctx) error{

	id := c.Params("id")
	var author model.Author


	db := database.DB

	if err := db.First(&author,id).Error ; err != nil{
		return fiber.NewError(fiber.StatusNotFound, "Author not found")
	}

	var updateAuthor model.Author
	if err := c.BodyParser(&updateAuthor) ; err != nil{
		return fiber.NewError(fiber.StatusBadRequest)
	}

	db.Model(&author).Association("Book").Clear()


	var books []model.Book
	for _, book := range updateAuthor.Book {
		if book.ID != 0{

			var existingBook model.Book
			if err := db.First(&existingBook,book.ID).Error ; err != nil{
				return fiber.NewError(fiber.StatusInternalServerError, "Error finding author: "+err.Error())
			
			}
			books = append(books, existingBook)
		}else{
			books = append(books, book)
		}

	}


	author.Name = updateAuthor.Name
	author.Age = updateAuthor.Age
	author.Book = books

	if err := db.Save(&author).Error ;  err != nil{
		return fiber.NewError(fiber.StatusInternalServerError, "Failed to update author: "+err.Error()) 
	}



	return c.JSON(author)


}


func DeleteAuthor(c *fiber.Ctx) error {

	id := c.Params("id")
	db := database.DB

	var author model.Author
	if err := db.First(&author,id).Error; err != nil{
		return fiber.NewError(fiber.StatusNotFound,"Author not found")
	}

	if err := db.Model(&author).Association("Book").Clear(); err != nil {
        return fiber.NewError(fiber.StatusInternalServerError, "Can't clear book associations: "+err.Error())
    }


	if err := db.Delete(&author).Error; err != nil{
		return fiber.NewError(fiber.StatusInternalServerError, "Can't delete author: ")
	}

	return c.SendStatus(fiber.StatusNoContent)
}