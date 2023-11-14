package auth

import (
	"github.com/edgar-care/edgarlib/graphql"
)

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(input LoginInput, t string) (string, error) {
	var doctor interface{}
	var admin interface{}
	var patient interface{}
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
		return "Could not find user: " + err.Error(), err
	}

	if !(t == "d" && CheckPassword(input.Password, doctor.(graphql.Doctor).Password)) &&
		!(t == "a" && CheckPassword(input.Password, admin.(graphql.Admin).Password)) &&
		!(t == "p" && CheckPassword(input.Password, patient.(graphql.Patient).Password)) {
		return "Username and password mismatch.", err
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
			"patient": patient,
		})
	}
	return token, err
}
