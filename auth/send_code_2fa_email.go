package auth

import (
	"errors"
	"fmt"
	edgarmail "github.com/edgar-care/edgarlib/v2/email"
	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/edgar-care/edgarlib/v2/redis"
	"math/rand"
)

type CodeEmailResponse struct {
	Code int
	Err  error
}

func Email2faAuth(email string) CodeEmailResponse {
	_, patientErr := graphql.GetPatientByEmail(email)
	if patientErr != nil {
		_, doctorErr := graphql.GetDoctorByEmail(email)
		if doctorErr != nil {
			return CodeEmailResponse{Code: 400, Err: errors.New("email does not correspond to a valid patient or doctor")}
		}
	}
	code := fmt.Sprintf("%06d", rand.Intn(1000000))

	expire := 600
	_, err := redis.SetKey(email, code, &expire)
	if err != nil {
		return CodeEmailResponse{500, err}
	}

	err = edgarmail.SendEmail(edgarmail.Email{
		To:       email,
		Subject:  "Votre code d'authentification à deux facteurs",
		Body:     fmt.Sprintf("Votre code d'authentification à deux facteurs est: %s", code),
		Template: "basic_with_button",
		TemplateInfos: map[string]interface{}{
			"Body":        fmt.Sprintf("Votre code d'authentification à deux facteurs est: %s", code),
			"ButtonUrl":   "",
			"ButtonTitle": "",
		},
	})
	if err != nil {
		return CodeEmailResponse{500, err}
	}

	return CodeEmailResponse{200, nil}
}
