package controller

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func GetReset(c *fiber.Ctx) error {

	email := c.Params("username")
	db, err := OpenDatabase()
	if err != nil {
		return err
	}

	query := "SELECT fname, lname, username, role FROM user WHERE email=?"
	rows := db.QueryRow(query, email)
	user := User{}
	err = rows.Scan(&user.Fname, &user.Lname, &user.Username, &user.Role)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"code":    1,
			"message": "body data is invalidate/incorrect format",
		})
	}
	receiverName := user.Fname + " " + user.Lname

	cliams := jwt.StandardClaims{
		Issuer:    user.Username + " " + user.Role,
		ExpiresAt: time.Now().Add(time.Minute * 5).Unix(),
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, cliams)
	token, err := jwtToken.SignedString([]byte(jwtSecret))
	if err != nil {
		return fiber.ErrInternalServerError
	}
	var emailParam []string
	emailParam = append(emailParam, token)

	SendEmail(email, receiverName, "forgot_password", emailParam)

	return c.Status(200).SendString("OK")

}
