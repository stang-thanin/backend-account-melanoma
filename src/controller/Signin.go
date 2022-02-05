package controller

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type SigninRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Signin(c *fiber.Ctx) error {

	db, err := OpenDatabase()
	if err != nil {
		return err
	}

	request := new(SigninRequest)
	err = c.BodyParser(&request)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"code":    1,
			"message": "body data is invalidate/incorrect format",
		})
	}

	if request.Username == "" || request.Password == "" {
		return c.Status(400).JSON(fiber.Map{
			"code":    1,
			"message": "body data is invalidate/incorrect format",
		})
	}

	user := User{}
	query := "SELECT id, username, password, role FROM user WHERE username=?"
	err = db.Get(&user, query, request.Username)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"code":    1,
			"message": "body data is invalidate/incorrect format",
		})
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"code":    1,
			"message": "body data is invalidate/incorrect format",
		})
	}

	cliams := jwt.StandardClaims{
		Issuer:    user.Username + " " + user.Role,
		ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, cliams)
	token, err := jwtToken.SignedString([]byte(jwtSecret))
	if err != nil {
		return fiber.ErrInternalServerError
	}

	return c.JSON(fiber.Map{
		"jwtToken": token,
	})
}
