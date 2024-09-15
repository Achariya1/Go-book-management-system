package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"fmt"
	"log"
	"os"
	"time"
	"github.com/Achariya1/go-books-crud/model"

)

const (
		host     = "localhost"  // or the Docker service name if running in another container
		port     = 5432         // default PostgreSQL port
		user     = "myuser"     // as defined in docker-compose.yml
		password = "mypassword" // as defined in docker-compose.yml
		dbname   = "mydatabase" // as defined in docker-compose.yml
)


func ConnetDB() {
	var err error // define error here to prevent overshadowing the global DB


	dsn := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	
	

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Enable color
		},
	)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: newLogger})
	if err != nil {
		//panic("failed to connect to database")
		log.Fatalf("Failed to connect to database: %v", err)
	}

	DB.AutoMigrate(&model.Book{},&model.User{})
	print("Conneting database Successfull")
	

}
