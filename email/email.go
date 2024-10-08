package email

import (
	"embed"
	"fmt"
	"io/ioutil"
	"net/smtp"
	"os"
	"strconv"
	"strings"

	_ "embed"

	"github.com/edgar-care/edgarlib/v2"
)

//go:embed templates/*.html
var embeddedTemplates embed.FS

type Email struct {
	To            string
	Subject       string
	Body          string
	Template      string
	TemplateInfos map[string]interface{}
}

func readHTMLFile(filePath string) (string, error) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func replaceHtmlContent(content string, replacements map[string]interface{}) string {
	for key, value := range replacements {
		content = strings.Replace(content, "{{."+key+"}}", fmt.Sprintf("%v", value), -1)
	}
	return content
}

func SendEmail(emailInfo Email) error {
	sesSMTPServer := os.Getenv("SES_SMTP_URL")
	sesSMTPPort, err := strconv.Atoi(os.Getenv("SES_SMTP_PORT"))
	edgarlib.CheckError(err)
	sesUsername := os.Getenv("SES_USERNAME")
	sesPassword := os.Getenv("SES_PASSWORD")

	fromAddress := os.Getenv("SENDER_EMAIL")
	var message string

	message = fmt.Sprintf("Subject: %s\r\n", emailInfo.Subject)
	message += "MIME-version: 1.0;\r\n"
	message += "Content-Type: text/html; charset=\"UTF-8\";\r\n\r\n"
	if emailInfo.Template == "" {
		message += emailInfo.Body
	} else {
		templateFilePath := fmt.Sprintf("templates/%s.html", emailInfo.Template)
		templateContent, err := embeddedTemplates.ReadFile(templateFilePath)
		if err != nil {
			message += emailInfo.Body
		} else {
			htmlContent := replaceHtmlContent(string(templateContent), emailInfo.TemplateInfos)

			message += htmlContent
		}
	}
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
