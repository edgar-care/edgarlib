package auth

import (
	"errors"
	"fmt"
	"github.com/edgar-care/edgarlib/v2"
	"os"

	edgarmail "github.com/edgar-care/edgarlib/v2/email"
	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/edgar-care/edgarlib/v2/redis"
	"github.com/google/uuid"
)

type MissingPasswordResponse struct {
	Code int
	Err  error
}

func MissingPassword(email string, accountype string) MissingPasswordResponse {
	var userID string
	var link string
	if accountype == "d" {
		doctor, err := graphql.GetDoctorByEmail(email)
		if err != nil {
			return MissingPasswordResponse{400, errors.New("no account corresponds to this email")}
		}
		userID = doctor.ID
	} else if accountype == "p" {
		patient, err := graphql.GetPatientByEmail(email)
		if err != nil {
			return MissingPasswordResponse{400, errors.New("no account corresponds to this email")}
		}
		userID = patient.ID
	} else {
		return MissingPasswordResponse{400, errors.New("no account type not found")}
	}

	account_uuid := uuid.New()
	expire := 600
	_, err := redis.SetKey(account_uuid.String(), userID, &expire)
	edgarlib.CheckError(err)

	link = fmt.Sprintf("app.edgar-sante.fr/%s/reset-password?uuid=%s", accountype, account_uuid.String())
	if os.Getenv("ENV") == "demo" {
		link = fmt.Sprintf("demo.app.edgar-sante.fr/%s/reset-password?uuid=%s", accountype, account_uuid.String())
	}

	err = edgarmail.SendEmail(edgarmail.Email{
		To:       email,
		Subject:  "Réinitialisation de votre mot de passe",
		Body:     fmt.Sprintf("Pour réinitialiser votre mot de passe, cliquez ici (app.edgar-sante.fr/%s/reset-password?uuid=%s)", accountype, account_uuid.String()),
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
