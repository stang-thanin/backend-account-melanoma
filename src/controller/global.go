package controller

import (
	"github.com/golang-jwt/jwt"
	"github.com/jmoiron/sqlx"
)

type User struct {
	Id         string `db:"id" json:"id"`
	Active     bool   `db:"active"`
	Created_at string `db:"created_at"`
	Created_by string `db:"created_by"`
	Updated_at string `db:"updated_at"`
	Update_by  string `db:"update_by"`
	Email      string `db:"email"`
	Fname      string `db:"fname"`
	Lname      string `db:"lname"`
	Password   string `db:"password" json:"password"`
	Role       string `db:"role"`
	Tel        string `db:"tel"`
	Username   string `db:"username" json:"username"`
}

type MyClaims struct {
	jwt.StandardClaims
	Iss string `json:"iss"`
	Exp int    `json:"exp"`
}

type TokenRequest struct {
	Token string `json:"token"`
}

const jwtSecret = "********"

func OpenDatabase() (*sqlx.DB, error) {

	var db *sqlx.DB
	var err error

	db, err = sqlx.Open("mysql", "********:********@********(********)/********")
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
