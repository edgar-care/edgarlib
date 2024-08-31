package auth

import (
	"errors"
	"fmt"

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
	patient, err := graphql.GetPatientByEmail(email)
	if err != nil {
		return MissingPasswordResponse{400, errors.New("no patient corresponds to this email")}
	}
	patient_uuid := uuid.New()
	expire := 600
	_, err = redis.SetKey(patient_uuid.String(), patient.ID, &expire)
	edgarlib.CheckError(err)

	err = edgarmail.SendEmail(edgarmail.Email{
		To:       email,
		Subject:  "Réinitialisation de votre mot de passe",
		Body:     fmt.Sprintf("Pour réinitialiser votre mot de passe, cliquez ici (app.edgar-sante.fr/reset-password?uuid=%s)", patient_uuid.String()),
		Template: "basic_with_button",
		TemplateInfos: map[string]interface{}{
			"Body":        "Pour réinitialiser votre mot de passe, cliquez ici",
			"ButtonUrl":   fmt.Sprintf("app.edgar-sante.fr/reset-password?uuid=%s", patient_uuid.String()),
			"ButtonTitle": "Cliquez ici",
		},
	})
	edgarlib.CheckError(err)

	return MissingPasswordResponse{200, nil}
}
