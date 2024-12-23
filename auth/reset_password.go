package auth

import (
	"errors"
	"github.com/edgar-care/edgarlib/v2/auth/utils"
	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
	"github.com/edgar-care/edgarlib/v2/redis"
)

type ResetPasswordResponse struct {
	Code int
	Err  error
}

func ResetPassword(password string, uuid string, accountype string) ResetPasswordResponse {
	if uuid == "" {
		return ResetPasswordResponse{400, errors.New("uuid has to be provided")}
	}
	value, err := redis.GetKey(uuid)
	if err != nil {
		return ResetPasswordResponse{500, err}
	}

	if value == "" {
		return ResetPasswordResponse{403, errors.New("uuid is expired")}
	}

	if accountype == "p" {
		patient, err := graphql.GetPatientById(value)
		password = utils.HashPassword(password)
		_, err = graphql.UpdatePatient(patient.ID, model.UpdatePatientInput{Password: &password})
		if err != nil {
			return ResetPasswordResponse{500, err}
		}
	} else if accountype == "d" {
		doctor, err := graphql.GetDoctorById(value)
		if err != nil {

			return ResetPasswordResponse{400, errors.New("no account correspond to this email")}
		}
		password = utils.HashPassword(password)
		_, err = graphql.UpdateDoctor(doctor.ID, model.UpdateDoctorInput{Password: &password})
		if err != nil {
			return ResetPasswordResponse{500, err}
		}
	} else {
		return ResetPasswordResponse{400, errors.New("no account type not found")}
	}

	return ResetPasswordResponse{200, nil}

}
