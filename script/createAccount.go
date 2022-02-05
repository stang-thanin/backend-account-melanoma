package main

import (
	"database/sql"
	"encoding/csv"
	"encoding/hex"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

var db *sql.DB

type UserData struct {
	active     bool
	created_at string
	created_by string `default:"admin"`
	updated_at string
	update_by  string `default:"admin"`
	email      string
	fname      string
	lname      string
	password   string
	role       string
	tel        string
	username   string
}

func main() {
	var err error
	db, err = sql.Open("mysql", "********:********@tcp(********)/********")

	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	csvFile, err := os.Open("../user_data.csv")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened CSV file")
	defer csvFile.Close()

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println(err)
	}

	for _, line := range csvLines {
		user := UserData{
			active:     true,
			created_at: time.Now().Format("2006-01-02"),
			created_by: "admin",
			updated_at: time.Now().Format("2006-01-02"),
			update_by:  "admin",
			email:      line[0],
			fname:      line[1],
			lname:      line[2],
			password:   GenerateFirstTimePassword(16),
			role:       line[3],
			tel:        line[4],
			username:   "",
		}

		user.email = strings.Replace(strings.TrimSpace(user.email), "\ufeff", "", -1)
		user.username = user.email

		password_hashed, err := bcrypt.GenerateFromPassword([]byte(user.password), 10)
		if err != nil {
			fmt.Println("cannot generate")
		}

		err = AddUserToDB(user, password_hashed)
		if err != nil {
			fmt.Println(err)
			continue
		}

		name := user.fname + " " + user.lname
		SendEmail(user.email, name, "register")

	}

}

func GenerateFirstTimePassword(n_digit int) string {
	bytes := make([]byte, n_digit)
	if _, err := rand.Read(bytes); err != nil {
		return ""
	}
	fmt.Println(hex.EncodeToString(bytes))
	return hex.EncodeToString(bytes)
}

func AddUserToDB(user UserData, password_hashed []byte) error {

	if !IsEmailExist(user.email) {
		query := "INSERT INTO user VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?)"

		result, err := db.Exec(query, GenerateId(user.role), user.active, user.created_at, user.created_by, user.updated_at, user.update_by, user.email, user.fname, user.lname, string(password_hashed), user.role, user.tel, user.username)
		if err != nil {
			return err
		}

		affected, err := result.RowsAffected()
		if err != nil {
			return err
		}

		if affected <= 0 {
			return errors.New("cannot insert")
		}

		return nil
	} else {
		return errors.New(fmt.Sprint("email ", user.email, " is exist"))
	}
}

func GenerateId(role string) string {

	query := "SELECT count(*) FROM user WHERE role = ?;"

	resultRows, err := db.Query(query, role)
	if err != nil {
		fmt.Println(err)
	}
	defer resultRows.Close()

	var currentNumberOfUserRole int
	for resultRows.Next() {
		err = resultRows.Scan(&currentNumberOfUserRole)
		if err != nil {
			fmt.Println(err)
		}
	}

	var Abrv string

	switch role {
	case "แพทย์":
		Abrv = "D"
	case "อสม.": // Public health technical officer
		Abrv = "P"
	case "พยาบาล":
		Abrv = "N"
	case "ผู้ช่วยพยาบาล": // Nurse assistance
		Abrv = "A"
	case "นักศึกษาแพทย์":
		Abrv = "S"
	default:
		Abrv = "X"
	}

	tmp := strconv.Itoa(currentNumberOfUserRole + 1)
	id := Abrv + (strings.Repeat("0", 5-len(tmp)) + tmp)

	return id
}

func IsEmailExist(email string) bool {
	query := "SELECT username, tel FROM user WHERE email=?"
	rows := db.QueryRow(query, email)
	user := UserData{}
	err := rows.Scan(&user.username, &user.tel)
	return err == nil
}
