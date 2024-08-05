package auth

import (
	"errors"

	"github.com/edgar-care/edgarlib/graphql"
	"github.com/edgar-care/edgarlib/graphql/model"
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
	var doctor model.Doctor
	var admin model.Admin
	var patient model.Patient
	var token string
	var err error

	if t == "d" {
		doctor, err = graphql.GetDoctorByEmail(input.Email)
	} else if t == "a" {
		admin, err = graphql.GetAdminByEmail(input.Email)
	} else {
		patient, err = graphql.GetPatientByEmail(input.Email)
	}
	if err != nil {
		resp.Code = 400
		resp.Err = errors.New("could not find user: " + err.Error())
		return resp
	}

	if !(t == "d" && CheckPassword(input.Password, doctor.Password)) &&
		!(t == "a" && CheckPassword(input.Password, admin.Password)) &&
		!(t == "p" && CheckPassword(input.Password, patient.Password)) {
		resp.Code = 401
		resp.Err = errors.New("username and password mismatch")
		return resp
	}

	if t == "d" {
		token, err = CreateToken(map[string]interface{}{
			"doctor": doctor,
		})
	} else if t == "a" {
		token, err = CreateToken(map[string]interface{}{
			"admin": admin,
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
