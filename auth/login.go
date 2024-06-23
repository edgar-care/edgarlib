package auth

import (
	"context"
	"errors"
	"github.com/edgar-care/edgarlib/graphql"
)

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string
	Code  int
	Err   error
}

func Login(input LoginInput, t string, ip string) LoginResponse {
	var resp LoginResponse
	var doctor *graphql.GetDoctorByEmailResponse
	var admin *graphql.GetAdminByEmailResponse
	var patient *graphql.GetPatientByEmailResponse
	var token string
	var err error
	gqlClient := graphql.CreateClient()

	if t == "d" {
		doctor, err = graphql.GetDoctorByEmail(context.Background(), gqlClient, input.Email)
	} else if t == "a" {
		admin, err = graphql.GetAdminByEmail(context.Background(), gqlClient, input.Email)
	} else {
		patient, err = graphql.GetPatientByEmail(context.Background(), gqlClient, input.Email)
	}
	if err != nil {
		resp.Code = 400
		resp.Err = errors.New("could not find user: " + err.Error())
		return resp
	}

	if !(t == "d" && CheckPassword(input.Password, doctor.GetDoctorByEmail.Password)) &&
		!(t == "a" && CheckPassword(input.Password, admin.GetAdminByEmail.Password)) &&
		!(t == "p" && CheckPassword(input.Password, patient.GetPatientByEmail.Password)) {
		resp.Code = 401
		resp.Err = errors.New("username and password mismatch")
		return resp
	}

	if t == "d" {
		token, err = CreateToken(map[string]interface{}{
			"doctor": doctor.GetDoctorByEmail,
		})
	} else if t == "a" {
		token, err = CreateToken(map[string]interface{}{
			"admin": admin.GetAdminByEmail,
		})
	} else {
		token, err = CreateToken(map[string]interface{}{
			"patient":   patient.GetPatientByEmail.Email,
			"id":        patient.GetPatientByEmail.Id,
			"ip_device": ip,
		})
	}
	resp.Token = token
	resp.Code = 200
	resp.Err = nil
	return resp
}
