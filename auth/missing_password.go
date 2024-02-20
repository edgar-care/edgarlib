package auth

import (
	"context"
	"errors"
	"fmt"

	"github.com/edgar-care/edgarlib"
	edgarmail "github.com/edgar-care/edgarlib/email"
	"github.com/edgar-care/edgarlib/graphql"
	"github.com/edgar-care/edgarlib/redis"
	"github.com/google/uuid"
)

type MissingPasswordResponse struct {
	Code int
	Err  error
}

func MissingPassword(email string) MissingPasswordResponse {
	gqlClient := graphql.CreateClient()

	_, err := graphql.GetPatientByEmail(context.Background(), gqlClient, email)
	if err != nil {
		return MissingPasswordResponse{400, errors.New("no patient corresponds to this email")}
	}
	patient_uuid := uuid.New()
	expire := 600
	_, err = redis.SetKey(patient_uuid.String(), email, &expire)
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
