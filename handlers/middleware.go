package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// IsAuthorized is a middleware for authorizatio
func IsAuthorized(c *fiber.Ctx) error {
	tokenHeader := c.GetReqHeaders()["Authorization"]
	// extract actual token
	tokens := strings.Split(tokenHeader, "Bearer ")

	// Check for Authorization header
	if c.GetReqHeaders()["Authorization"] == "" || len(tokens) != 2 {
		return c.Status(http.StatusBadRequest).JSON(
			fiber.Map{"message": "Token not found in request."},
		)
	}

	token, err := jwt.Parse(tokens[1], func(token *jwt.Token) (interface{}, error) {
		fmt.Println(token.Method.Alg())
		if token.Method.Alg() != "HS512" {
			return nil, fmt.Errorf("there was error while parsing")
		}
		return &token, nil
	})

	if err != nil {
		log.Println(err)
		return c.Status(http.StatusUnauthorized).JSON(
			fiber.Map{"message": "Expired or invalid token"},
		)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if claims["role"] != "user" {
			return c.Status(http.StatusForbidden).JSON(
				fiber.Map{"message": "Unsupported role"},
			)
		}
	}
	// Pass to next handler
	c.Next()
	return nil
}
