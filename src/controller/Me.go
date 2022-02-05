package controller

import (
	"errors"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func Me(c *fiber.Ctx) error {

	request := new(TokenRequest)
	err := c.BodyParser(&request)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"code":    1,
			"message": "body data is invalidate/incorrect format",
		})
	}

	tokenString := request.Token

	token, _ := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(jwtSecret), nil
	})

	myClaims := token.Claims.(*MyClaims)

	email := strings.Split(myClaims.Iss, " ")[0]
	role := strings.Split(myClaims.Iss, " ")[1]

	db, err := OpenDatabase()
	if err != nil {
		return err
	}

	query := "SELECT id, active, fname, lname, tel FROM user WHERE email=?"
	rows := db.QueryRow(query, email)
	user := User{}
	err = rows.Scan(&user.Id, &user.Active, &user.Fname, &user.Lname, &user.Tel)
	if err != nil {
		return err
	}
	user.Email = email
	user.Username = email
	user.Role = role

	return c.Status(200).JSON(fiber.Map{
		"id":       user.Id,
		"active":   user.Active,
		"fname":    user.Fname,
		"lname":    user.Lname,
		"tel":      user.Tel,
		"email":    user.Email,
		"username": user.Username,
		"role":     user.Role,
	})

}
