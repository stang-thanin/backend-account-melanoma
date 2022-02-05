package controller

import (
	"errors"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func Auth(c *fiber.Ctx) error {

	request := new(TokenRequest)
	err := c.BodyParser(&request)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"code":    1,
			"message": "body data is invalidate/incorrect format",
		})
	}

	tokenString := request.Token
	if tokenString == "" {
		return c.Status(401).JSON(fiber.Map{
			"code":    1,
			"message": "body data is invalidate/incorrect format",
		})
	}

	token, _ := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(jwtSecret), nil
	})
	if token == nil {
		return c.Status(400).JSON(fiber.Map{
			"code":    1,
			"message": "body data is invalidate/incorrect format",
		})
	}

	myClaims := token.Claims.(*MyClaims)

	if time.Now().After(time.Unix(int64(myClaims.Exp), 0)) {
		return c.Status(400).JSON(fiber.Map{
			"code":    1,
			"message": "body data is invalidate/incorrect format",
		})
	} else {
		return c.Status(200).JSON(fiber.Map{
			"token":  tokenString,
			"status": "valid",
		})
	}

}
