package auth

import (
	"errors"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
	"strings"

	"github.com/edgar-care/edgarlib/v2/auth/utils"
	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/edgar-care/edgarlib/v2/redis"
)

type ResetPasswordResponse struct {
	Code int
	Err  error
}

func ResetPassword(email string, password string, uuid string) ResetPasswordResponse {
	if uuid == "" {
		return ResetPasswordResponse{403, errors.New("uuid has to be provided")}
	}
	value, err := redis.GetKey(uuid)
	if err != nil {
		return ResetPasswordResponse{500, err}
	}

	value = strings.Replace(value, "\n", "", -1)
	if value == "" || err != nil {
		return ResetPasswordResponse{403, errors.New("uuid is expired")}
	}

	patient, err := graphql.GetPatientByEmail(email)
	if err != nil {
		return ResetPasswordResponse{403, errors.New("no patient correspond to this email")}
	}

	password = utils.HashPassword(password)
	_, err = graphql.UpdatePatient(patient.ID, model.UpdatePatientInput{Password: &password})
	if err != nil {
		return ResetPasswordResponse{400, err}
	}
	return ResetPasswordResponse{200, nil}
}
