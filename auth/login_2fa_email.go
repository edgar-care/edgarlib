package auth

import (
	"errors"
	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/edgar-care/edgarlib/v2/redis"
)

type Login2faEmailInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Token2fa string `json:"token_2fa"`
}

type Login2faEmailResponse struct {
	Token string
	Code  int
	Err   error
}

func Login2faEmail(input Login2faEmailInput, nameDevice string) Login2faEmailResponse {
	var token string
	var doubleAuthId *string

	verifyToken2fa, err := redis.GetKey(input.Email)
	if err != nil {
		return Login2faEmailResponse{Token: "", Code: 500, Err: err}
	}

	if verifyToken2fa != input.Token2fa {
		return Login2faEmailResponse{Token: "", Code: 401, Err: errors.New("invalid 2FA token")}
	}

	patient, patientErr := graphql.GetPatientByEmail(input.Email)
	if patientErr == nil {
		doubleAuthId = patient.DoubleAuthMethodsID
		if doubleAuthId == nil {
			return Login2faEmailResponse{Token: "", Code: 400, Err: errors.New("no 2FA method associated with")}
		}
		checkDoubleAuth, err := graphql.GetDoubleAuthById(*doubleAuthId)
		if err != nil {
			return Login2faEmailResponse{Token: "", Code: 500, Err: err}
		}
		if !isEmail2faValid(checkDoubleAuth.Methods) {
			return Login2faEmailResponse{Token: "", Code: 400, Err: errors.New("email is not a valid 2FA method")}
		}
		check := CheckPassword(input.Password, patient.Password)
		if !check {
			return Login2faEmailResponse{Token: "", Code: 401, Err: errors.New("invalid password")}
		}

		token, err = CreateToken(map[string]interface{}{
			"patient":     patient.Email,
			"id":          patient.ID,
			"name_device": nameDevice,
		})

		return Login2faEmailResponse{Token: token, Code: 200, Err: nil}
	}

	doctor, doctorErr := graphql.GetDoctorByEmail(input.Email)
	if doctorErr == nil {
		doubleAuthId = doctor.DoubleAuthMethodsID
		if doubleAuthId == nil {
			return Login2faEmailResponse{Token: "", Code: 400, Err: errors.New("no 2FA method associated with")}
		}
		checkDoubleAuth, err := graphql.GetDoubleAuthById(*doubleAuthId)
		if err != nil {
			return Login2faEmailResponse{Token: "", Code: 500, Err: err}
		}

		if !isEmail2faValid(checkDoubleAuth.Methods) {
			return Login2faEmailResponse{Token: "", Code: 400, Err: errors.New("email is not a valid 2FA method")}
		}

		check := CheckPassword(input.Password, doctor.Password)
		if !check {
			return Login2faEmailResponse{Token: "", Code: 401, Err: errors.New("invalid password")}
		}
		token, err = CreateToken(map[string]interface{}{
			"doctor":      doctor.Email,
			"id":          doctor.ID,
			"name_device": nameDevice,
		})

		return Login2faEmailResponse{Token: token, Code: 200, Err: nil}
	}

	return Login2faEmailResponse{Token: "", Code: 400, Err: errors.New("email does not correspond to a valid patient or doctor")}
}

func isEmail2faValid(methods []string) bool {
	for _, method := range methods {
		if method == "EMAIL" || method == "MOBILE" {
			return true
		}
	}
	return false
}
