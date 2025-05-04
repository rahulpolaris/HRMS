package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

type MongoInstance struct {
	Client string
	Db     string
}

type Employee struct {
	ID     string
	Name   string
	Salary float32
	Age    uint8
}

func Connect() error {

}

//TYPICAL MONGO URI IS LIKE============
// mongodb://username:password@hostname:port/database?authSource=admin

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error reading .env file \n", err)
	}
	// var mg MongoInstance

	//loading DB user creds from .env file and  DB name
	dbname := os.Getenv("DB_NAME")
	dbuser := os.Getenv("DB_USER")
	dbpassword := os.Getenv("DB_USER_PWD")
	dbport := os.Getenv("DB_PORT")
	dbhost := os.Getenv("DB_HOST")

	mongoURI := fmt.Sprintf("mongodb://%s:%s@%s:%s/%s?authSource=admin", dbuser, dbpassword, dbhost, dbport, dbname)
	fmt.Println("mongoURI is:", mongoURI)

	app := fiber.New()

	app.Get("/employee")
	app.Post("/employee")
	app.Put("/employee/:id")
	app.Delete("/employee/:id")

	// const mongoURI = "mongodb://hrmslocalhost:27017"

}
