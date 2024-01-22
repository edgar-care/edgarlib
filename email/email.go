package email

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
)

type Email struct {
	To      string
	Subject string
	Body    string
}

func SendEmail(emailInfo Email) error {
	// smtpServer := os.Getenv("SMTP_SERVER")
	// smtpPort, err := strconv.Atoi(os.Getenv("SMTP_PORT"))

	// edgarlib.CheckError(err)

	// senderEmail := os.Getenv("SENDER_EMAIL")
	// senderPassword := os.Getenv("SENDER_PASSWORD")
	// recipientEmail := emailInfo.To

	// emailMessage := []byte(
	// 	"Subject: " + emailInfo.Subject + "\r\n" +
	// 		"To: " + recipientEmail + "\r\n" +
	// 		"Content-Type: text/plain; charset=UTF-8\r\n" +
	// 		"\r\n" + emailInfo.Body,
	// )
	// auth := smtp.PlainAuth("", senderEmail, senderPassword, smtpServer)
	// err = smtp.SendMail(smtpServer+":"+strconv.Itoa(smtpPort), auth, senderEmail, []string{recipientEmail}, emailMessage)
	// if err != nil {
	// 	log.Printf("Email sent sucessfuly !")
	// } else {
	// 	log.Println("Eror Email did not send !")
	// }
	fmt.Println("SENDING MAIL (TODO)")
	spew.Dump(emailInfo)
	return nil
}
