package main

import (
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/roshanlc/jwt-auth-go/handlers"
	"github.com/roshanlc/jwt-auth-go/initializers"
	"github.com/roshanlc/jwt-auth-go/utilities"
)

// tutorial url : https://www.bacancytechnology.com/blog/golang-jwt

func init() {
	log.Println("Loading .env file")
	// read .env vars
	initializers.LoadEnvVariables()
	// setup secret key
	utilities.SigningKey = []byte(os.Getenv("SECRET_KEY"))
	// initial migration and population
	initializers.InitialMigration()

}
func main() {
	router := fiber.New()

	// routes registration
	router.Post("/login", handlers.LoginHandler)
	router.Post("/register", handlers.RegistrationHandler)
	router.Get("/home", handlers.IsAuthorized, handlers.HomeHandler)

	// Run the router in separate goroutine
	log.Fatal(router.Listen(":" + os.Getenv("PORT")))

}
