package auth

import (
	"errors"
	"fmt"
	"os"

	"github.com/edgar-care/edgarlib/v2"
	edgarmail "github.com/edgar-care/edgarlib/v2/email"
	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/edgar-care/edgarlib/v2/redis"
	"github.com/google/uuid"
)

type MissingPasswordResponse struct {
	Code int
	Err  error
}

func MissingPassword(email string) MissingPasswordResponse {
	var userID string

	patient, err := graphql.GetPatientByEmail(email)
	if err != nil {
		doctor, err := graphql.GetDoctorByEmail(email)
		if err != nil {
			return MissingPasswordResponse{400, errors.New("no account corresponds to this email")}
		}
		userID = doctor.ID
	}
	userID = patient.ID
	patient_uuid := uuid.New()
	expire := 600
	_, err = redis.SetKey(patient_uuid.String(), userID, &expire)
	edgarlib.CheckError(err)

	link := fmt.Sprintf("app.edgar-sante.fr/reset-password?uuid=%s", patient_uuid.String())
	if os.Getenv("ENV") == "demo" {
		link = fmt.Sprintf("demo.app.edgar-sante.fr/reset-password?uuid=%s", patient_uuid.String())
	}

	err = edgarmail.SendEmail(edgarmail.Email{
		To:       email,
		Subject:  "Réinitialisation de votre mot de passe",
		Body:     fmt.Sprintf("Pour réinitialiser votre mot de passe, cliquez ici (app.edgar-sante.fr/reset-password?uuid=%s)", patient_uuid.String()),
		Template: "basic_with_button",
		TemplateInfos: map[string]interface{}{
			"Body":        "Pour réinitialiser votre mot de passe",
			"ButtonUrl":   link,
			"ButtonTitle": "Merci de suivre ce lien:",
		},
	})
	edgarlib.CheckError(err)

	return MissingPasswordResponse{200, nil}
}
