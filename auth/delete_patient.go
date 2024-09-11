package auth

import (
	"errors"
	"fmt"
	edgarmail "github.com/edgar-care/edgarlib/v2/email"
	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/edgar-care/edgarlib/v2/redis"
)

func sendDeleteConfirmationEmail(email string) error {
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

func sendReminderEmail(email string, daysLeft int) error {
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

func deleteDefinitlyAccount(ownerID, accountType string) error {
	if accountType == "patient" {
		_, err := graphql.DeletePatient(ownerID)
		if err != nil {
			return fmt.Errorf("failed to delete patient: %v", err)
		}
	} else {
		_, err := graphql.DeleteDoctor(ownerID)
		if err != nil {
			return fmt.Errorf("failed to delete doctor: %v", err)
		}
	}

	// Suppression de la clé Redis
	_, err := redis.DeleteKey(ownerID + "_delete_request")
	if err != nil {
		return fmt.Errorf("failed to delete redis key: %v", err)
	}
	return nil
}

func getEmailByOwnerID(ownerID string) (string, string, error) {
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
