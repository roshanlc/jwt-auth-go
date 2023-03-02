package handlers

import (
	"encoding/json"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/roshanlc/jwt-auth-go/initializers"
	"github.com/roshanlc/jwt-auth-go/models"
	"github.com/roshanlc/jwt-auth-go/utilities"
)

func LoginHandler(c *fiber.Ctx) error {
	// login handler logic
	return c.JSON("welcome")
}

func RegistrationHandler(c *fiber.Ctx) error {
	user := models.User{}

	err := json.Unmarshal(c.Body(), &user)
	if err != nil {
		return c.Status(400).JSON(
			fiber.Map{"message": "Mailformed json object"},
		)
	}
	dbConn := initializers.ConnectDB()
	defer initializers.CloseDBConnection(dbConn)
	var dbUser models.User
	// query the user table
	dbConn.Where("email = ?", user.Email).First(&dbUser)

	if dbUser.Email != "" {
		// return conflict status code
		return c.Status(409).JSON(
			fiber.Map{"message": "Email is already in use."},
		)
	}

	user.Password, err = utilities.GenerateHashPassword(user.Password)
	if err != nil {
		log.Println(err) // see the error message
		return c.Status(500).JSON(
			fiber.Map{"message": "Internal server error."},
		)
	}

	dbConn.Create(&user)
	return c.Status(201).JSON(
		fiber.Map{"message": "Account was created successfully."},
	)
}
