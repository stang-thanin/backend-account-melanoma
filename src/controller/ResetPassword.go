package controller

import (
	"errors"
	"fmt"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type NewPassword struct {
	NewPassword string `json:"newPassword"`
}

func ResetPassword(c *fiber.Ctx) error {

	resetToken := c.Query("resetToken")
	newPassword := new(NewPassword)
	err := c.BodyParser(&newPassword)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"code":    1,
			"message": "body data is invalidate/incorrect format",
		})
	}

	if resetToken == "" {
		return c.Status(400).JSON(fiber.Map{
			"code":    1,
			"message": "body data is invalidate/incorrect format",
		})
	}

	token, _ := jwt.ParseWithClaims(resetToken, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(jwtSecret), nil
	})

	myClaims := token.Claims.(*MyClaims)

	if time.Now().After(time.Unix(int64(myClaims.Exp), 0)) {
		return c.Status(400).JSON(fiber.Map{
			"code":    1,
			"message": "body data is invalidate/incorrect format",
		})
	} else {
		password_hashed, err := bcrypt.GenerateFromPassword([]byte(newPassword.NewPassword), 10)
		if err != nil {
			return err
		}

		username := strings.Split(myClaims.Iss, " ")[0]

		updateNewPassword(string(password_hashed), username)

		return c.Status(200).SendString("OK")
	}

}

func updateNewPassword(password_hashed string, username string) error {

	db, err := OpenDatabase()
	if err != nil {
		return err
	}

	query := "update user set password=? where username=?"
	result, err := db.Exec(query, password_hashed, username)
	if err != nil {
		return err
	}
	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affected <= 0 {
		return errors.New("cannot update")
	}

	fmt.Println("update na ja")

	return nil
}
