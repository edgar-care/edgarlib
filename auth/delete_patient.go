package auth

import (
	"errors"
	"fmt"
	edgarmail "github.com/edgar-care/edgarlib/v2/email"
	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/edgar-care/edgarlib/v2/redis"
)

type InitDeleteAccountResponse struct {
	Validate bool
	Code     int
	Err      error
}

type DeleteDefinitlyAccountResponse struct {
	Confirm bool
	Code    int
	Err     error
}

func InitDeleteAccount(ownerID string) InitDeleteAccountResponse {

	email, _, err := GetEmailByOwnerID(ownerID)
	if err != nil {
		return InitDeleteAccountResponse{Validate: false, Code: 400, Err: err}
	}

	err = SendDeleteConfirmationEmail(email)
	if err != nil {
		return InitDeleteAccountResponse{Validate: false, Code: 400, Err: errors.New("failed to send email")}
	}

	time := 30 * 24 * 60 * 60

	_, err = redis.SetKey(ownerID+"_delete_request", "true", &time)
	if err != nil {
		return InitDeleteAccountResponse{Validate: false, Code: 500, Err: errors.New("failed to set key in Redis")}
	}

	return InitDeleteAccountResponse{Validate: true, Err: nil, Code: 200}

}

func SendDeleteConfirmationEmail(email string) error {
	return edgarmail.SendEmail(edgarmail.Email{
		To:       email,
		Subject:  "Confirmation de la demande de suppression de compte",
		Body:     "Votre demande de suppression de compte a été enregistrée. Votre compte sera supprimé dans 30 jours.",
		Template: "basic_with_button",
		TemplateInfos: map[string]interface{}{
			"Body":        "Votre demande de suppression de compte a été enregistrée. Votre compte sera supprimé dans 30 jours.",
			"ButtonUrl":   "",
			"ButtonTitle": "",
		},
	})
}

func SendReminderEmail(email string, daysLeft int) error {
	return edgarmail.SendEmail(edgarmail.Email{
		To:       email,
		Subject:  "Rappel : Demande de suppression de compte",
		Body:     fmt.Sprintf("Il reste %d jours avant la suppression définitive de votre compte.", daysLeft),
		Template: "basic_with_button",
		TemplateInfos: map[string]interface{}{
			"Body":        fmt.Sprintf("Votre compte sera supprimé dans %d jours.", daysLeft),
			"ButtonUrl":   "",
			"ButtonTitle": "",
		},
	})
}

func DeleteDefinitlyAccount(ownerID string) DeleteDefinitlyAccountResponse {
	_, accountType, err := GetEmailByOwnerID(ownerID)
	if err != nil {
		return DeleteDefinitlyAccountResponse{Confirm: false, Code: 400, Err: err}
	}

	if accountType == "patient" {
		_, err := graphql.DeletePatient(ownerID)
		if err != nil {
			return DeleteDefinitlyAccountResponse{Confirm: false, Code: 500, Err: errors.New("failed to delete patient")}
		}
	} else {
		_, err := graphql.DeleteDoctor(ownerID)
		if err != nil {
			return DeleteDefinitlyAccountResponse{Confirm: false, Code: 500, Err: errors.New("failed to delete doctor")}
		}
	}

	_, err = redis.DeleteKey(ownerID + "_delete_request")
	if err != nil {
		return DeleteDefinitlyAccountResponse{Confirm: false, Code: 500, Err: errors.New("failed to delete key in Redis")}
	}

	return DeleteDefinitlyAccountResponse{Confirm: true, Code: 200, Err: nil}
}

func GetEmailByOwnerID(ownerID string) (string, string, error) {
	patient, err := graphql.GetPatientById(ownerID)
	if err == nil {
		return patient.Email, "patient", nil
	}
	doctor, err := graphql.GetDoctorById(ownerID)
	if err == nil {
		return doctor.Email, "doctor", nil
	}
	return "", "", errors.New("owner ID does not correspond to a valid patient or doctor")
}
