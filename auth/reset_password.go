package auth

import (
	"context"
	"errors"
	"strings"

	"github.com/edgar-care/edgarlib/auth/utils"
	"github.com/edgar-care/edgarlib/graphql"
	"github.com/edgar-care/edgarlib/redis"
)

type ResetPasswordResponse struct {
	Code int
	Err  error
}

func ResetPassword(email string, password string, uuid string) ResetPasswordResponse {
	gqlClient := graphql.CreateClient()

	if uuid == "" {
		return ResetPasswordResponse{403, errors.New("uuid has to be provided")}
	}
	value, err := redis.GetKey(uuid)
	value = strings.Replace(value, "\n", "", -1)
	if value == "" || err != nil {
		return ResetPasswordResponse{403, errors.New("uuid is expired")}
	}

	patient, err := graphql.GetPatientByEmail(context.Background(), gqlClient, email)
	if err != nil {
		return ResetPasswordResponse{403, errors.New("no patient correspond to this email")}
	}

	_, err = graphql.UpdatePatient(context.Background(), gqlClient, patient.GetPatientByEmail.Id, patient.GetPatientByEmail.Email, utils.HashPassword(password), patient.GetPatientByEmail.Medical_info_id, patient.GetPatientByEmail.Rendez_vous_ids, patient.GetPatientByEmail.Document_ids, patient.GetPatientByEmail.Treatment_follow_up_ids)
	if err != nil {
		return ResetPasswordResponse{400, err}
	}
	return ResetPasswordResponse{200, nil}
}
