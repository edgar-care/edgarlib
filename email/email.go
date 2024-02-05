package email

import (
	"fmt"
	"io/ioutil"
	"net/smtp"
	"os"
	"strconv"
	"strings"

	"github.com/edgar-care/edgarlib"
)

type Email struct {
	To      string
	Subject string
	Body    string
}

func readHTMLFile(filePath string) (string, error) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func SendEmail(emailInfo Email) error {
	sesSMTPServer := os.Getenv("SES_SMTP_URL")
	sesSMTPPort, err := strconv.Atoi(os.Getenv("SES_SMTP_PORT"))
	edgarlib.CheckError(err)
	sesUsername := os.Getenv("SES_USERNAME")
	sesPassword := os.Getenv("SES_PASSWORD")

	fromAddress := os.Getenv("SENDER_EMAIL")

	htmlFile, err := readHTMLFile("./email/email_template.html")
	if err != nil {
		return err
	}
	htmlContent := strings.Replace(htmlFile, "{{.Body}}", emailInfo.Body, -1)

	message := fmt.Sprintf("Subject: %s\r\n", emailInfo.Subject)
	message += "MIME-version: 1.0;\r\n"
	message += "Content-Type: text/html; charset=\"UTF-8\";\r\n\r\n"
	message += htmlContent
	// message := fmt.Sprintf("Subject: %s\r\n\r\n%s", emailInfo.Subject, emailInfo.Body)

	auth := smtp.PlainAuth("", sesUsername, sesPassword, sesSMTPServer)
	err = smtp.SendMail(
		fmt.Sprintf("%s:%d", sesSMTPServer, sesSMTPPort),
		auth,
		fromAddress,
		[]string{emailInfo.To},
		[]byte(message),
	)
	if err != nil {
		fmt.Println("Error sending email:", err)
		return err
	}

	return nil
}
