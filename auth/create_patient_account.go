package auth

import (
	"fmt"

	"github.com/edgar-care/edgarlib/v2"
	"github.com/edgar-care/edgarlib/v2/auth/utils"
	edgarmail "github.com/edgar-care/edgarlib/v2/email"
	"github.com/edgar-care/edgarlib/v2/redis"
	"github.com/google/uuid"
)

type CreatePatientAccountResponse struct {
	Id   string
	Code int
	Err  error
}

func CreatePatientAccount(email string) CreatePatientAccountResponse {
	password := utils.GeneratePassword(10)

	patient, err := RegisterPatient(email, password)
	if err != nil {
		return CreatePatientAccountResponse{"", 400, err}
	}
	patient_uuid := uuid.New()
	expire := 43200
	_, err = redis.SetKey(patient_uuid.String(), email, &expire)
	edgarlib.CheckError(err)

	err = edgarmail.SendEmail(edgarmail.Email{
		To:       email,
		Subject:  "Création de votre compte - edgar-sante.fr",
		Body:     fmt.Sprintf("Votre compte à bien été créé, cliquez ici pour mettre à jour votre mot de passe (app.edgar-sante.fr/reset-password?uuid=%s)", patient_uuid.String()),
		Template: "basic_with_button",
		TemplateInfos: map[string]interface{}{
			"Body":        "Votre compte à bien été créé, cliquez ici pour mettre à jour votre mot de passe",
			"ButtonUrl":   fmt.Sprintf("app.edgar-sante.fr/reset-password?uuid=%s", patient_uuid.String()),
			"ButtonTitle": "Cliquez ici",
		},
	})
	edgarlib.CheckError(err)

	return CreatePatientAccountResponse{Id: patient.ID, Code: 200, Err: nil}
}
