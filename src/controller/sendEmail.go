package controller

import (
	"bytes"
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/gmail/v1"
	"google.golang.org/api/option"
)

var GmailService *gmail.Service

func SendEmail(receiver_email string, receiver_name string, purpose string, params []string) {

	err := godotenv.Load("src/email_templates/.env")
	if err != nil {
		fmt.Printf("Error getting env, %v", err)
	}

	OAuthGmailService()
	log.Print("Email sent to ", receiver_email, " [", purpose, "]")

	err = SendEmailOAUTH2(receiver_email, receiver_name, purpose, params)
	if err != nil {
		fmt.Println(err)
	}

}

func OAuthGmailService() {

	config := oauth2.Config{
		ClientID:     os.Getenv("CLIENT_ID"),
		ClientSecret: os.Getenv("CLIENT_SECRET"),
		Endpoint:     google.Endpoint,
		RedirectURL:  "http://localhost",
	}

	token := oauth2.Token{
		AccessToken:  os.Getenv("ACCESS_TOKEN"),
		RefreshToken: os.Getenv("REFRESH_TOKEN"),
		TokenType:    "Bearer",
		Expiry:       time.Now(),
	}

	var tokenSource = config.TokenSource(context.Background(), &token)

	srv, err := gmail.NewService(context.Background(), option.WithTokenSource(tokenSource))
	if err != nil {
		fmt.Printf("Unable to retrieve Gmail client: %v", err)
	}

	GmailService = srv

}

func SendEmailOAUTH2(receiver_email string, receiver_name string, purpose string, params []string) error {

	var templateFileName, email_subject string
	switch purpose {
	case "register":
		templateFileName = "register_template.txt"
		email_subject = "Your account on Melanoma.com was approved"
	case "forgot_password":
		templateFileName = "forgot_password_template.txt"
		email_subject = "Reset your password for Melanoma.com"
	default:
		return errors.New("email purpose is invalid")
	}

	emailBody, err := parseTemplate(templateFileName, receiver_name, params)
	if err != nil {
		return errors.New("unable to parse email template")
	}

	var message gmail.Message

	emailTo := "To: " + receiver_email + "\r\n"
	subject := "Subject: " + email_subject + "\n"
	mime := "MIME-version: 1.0;\nContent-Type: text/plain; charset=\"UTF-8\";\n\n"
	msg := []byte(emailTo + subject + mime + "\n" + emailBody)

	message.Raw = base64.URLEncoding.EncodeToString(msg)

	// Send the message
	_, err = GmailService.Users.Messages.Send("me", &message).Do()
	if err != nil {
		return err
	}
	fmt.Println("Send successfully")
	fmt.Println()
	return nil

}

func parseTemplate(templateFileName string, receiver_name string, params []string) (string, error) {

	templatePath, err := filepath.Abs(fmt.Sprintf("src/email_templates/%s", templateFileName))
	if err != nil {
		return "", errors.New("invalid template name")
	}

	t, err := template.ParseFiles(templatePath)
	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)
	data := struct {
		ReceiverName string
		Params       string
	}{
		ReceiverName: receiver_name,
		Params:       params[0],
	}
	if err = t.Execute(buf, data); err != nil {
		return "", err
	}

	body := buf.String()
	return body, nil

}
