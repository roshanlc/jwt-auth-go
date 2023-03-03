package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/roshanlc/jwt-auth-go/initializers"
	"github.com/roshanlc/jwt-auth-go/models"
	"github.com/roshanlc/jwt-auth-go/utilities"
)

// handler for the "/register" endpoint
func RegistrationHandler(c *fiber.Ctx) error {
	user := models.User{}

	err := json.Unmarshal(c.Body(), &user)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(
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
		return c.Status(http.StatusConflict).JSON(
			fiber.Map{"message": "Email is already in use."},
		)
	}

	user.Password, err = utilities.GenerateHashPassword(user.Password)
	if err != nil {
		log.Println(err) // see the error message
		return c.Status(http.StatusInternalServerError).JSON(
			fiber.Map{"message": "Internal server error."},
		)
	}

	dbConn.Create(&user)
	return c.Status(201).JSON(
		fiber.Map{"message": "Account was created successfully."},
	)
}

// handler for the "/login" endpoint
func LoginHandler(c *fiber.Ctx) error {
	var authDetails models.Authentication
	err := json.Unmarshal(c.Body(), &authDetails)
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(
			fiber.Map{"message": "Mailformed json object"},
		)
	}
	// check for empty email or password field
	if len(authDetails.Email) == 0 || len(authDetails.Password) == 0 {
		return c.Status(http.StatusBadRequest).JSON(
			fiber.Map{"message": "Please provide valid credentails."},
		)
	}

	dbConn := initializers.ConnectDB()
	defer initializers.CloseDBConnection(dbConn)

	var authUser models.User
	dbConn.Where("email = ?", authDetails.Email).First(&authUser)
	if authUser.Email == "" {
		return c.Status(http.StatusUnauthorized).JSON(
			fiber.Map{"message": "Please provide valid credentails."},
		)
	}

	check := utilities.CheckPasswordHash(authDetails.Password, authUser.Password)
	if !check {
		return c.Status(http.StatusForbidden).JSON(
			fiber.Map{"message": "Please provide valid credentails."},
		)
	}
	validToken, err := utilities.GenerateJWT(authUser.Email, authUser.Role)
	if err != nil {
		log.Println(err) // see the error message
		return c.Status(http.StatusInternalServerError).JSON(
			fiber.Map{"message": "Internal server error."},
		)
	}

	token := models.Token{
		Email:       authUser.Email,
		Role:        authUser.Role,
		TokenString: validToken,
	}
	return c.Status(http.StatusOK).JSON(token)

}

// handler for the "/home" endpoint
func HomeHandler(c *fiber.Ctx) error {
	return c.Status(http.StatusOK).JSON(
		fiber.Map{"message": "Welcome home"},
	)
}
