package database

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetDatabaseConnection() *gorm.DB {

	// Read the env file.
	err := godotenv.Load()

	if err != nil {
		log.Fatal("The .env file is not loaded.")
	}

	// Get the credentials, read the variables from the .env file
	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	database := os.Getenv("MYSQL_DATABASE")
	username := os.Getenv("MYSQL_USERNAME")
	password := os.Getenv("MYSQL_PASSWORD")

	// Address
	dsn := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + database + "?charset=utf8&parseTime=true&loc=Local"

	// Initiate a new DB connection
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("[Database] failed connecting to DB: " + dsn + ", err: " + err.Error())
	}

	// Logger to say it is connected!
	log.Println("[Database] Database is connected.")

	return db
}
