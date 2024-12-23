package auth

import (
	"errors"
	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/pquerna/otp/totp"
	"net/http"
)

type Login2faThirdPartyInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Token2fa string `json:"token_2fa"`
}

type Login2faThirdPartyResponse struct {
	Token string
	Code  int
	Err   error
}

func Login2faThirdParty(input Login2faThirdPartyInput, nameDevice string, accountID string) Login2faThirdPartyResponse {
	var token string

	patient, patientErr := graphql.GetPatientById(accountID)
	if patientErr == nil {

		doubleAuth, err := graphql.GetDoubleAuthById(*patient.DoubleAuthMethodsID)
		if err != nil {
			return Login2faThirdPartyResponse{
				Token: "",
				Code:  http.StatusBadRequest,
				Err:   errors.New("get double_auth failed: " + err.Error()),
			}
		}

		valid := totp.Validate(input.Token2fa, doubleAuth.Code)
		if !valid {
			return Login2faThirdPartyResponse{Token: token, Code: 400, Err: errors.New("token 2fa is invalid")}
		}

		check := CheckPassword(input.Password, patient.Password)
		if !check {
			return Login2faThirdPartyResponse{Token: "", Code: 401, Err: errors.New("invalid password")}
		}
		// Create a token
		token, _ := CreateToken(map[string]interface{}{
			"patient":     patient.Email,
			"id":          patient.ID,
			"name_device": nameDevice,
		})

		return Login2faThirdPartyResponse{Token: token, Code: 200, Err: nil}
	}

	doctor, doctorErr := graphql.GetDoctorById(accountID)
	if doctorErr == nil {

		doubleAuth, err := graphql.GetDoubleAuthById(*doctor.DoubleAuthMethodsID)
		if err != nil {
			return Login2faThirdPartyResponse{
				Token: "",
				Code:  http.StatusBadRequest,
				Err:   errors.New("get double_auth failed: " + err.Error()),
			}
		}

		valid := totp.Validate(input.Token2fa, doubleAuth.Code)
		if !valid {
			return Login2faThirdPartyResponse{Token: token, Code: 400, Err: errors.New("token 2fa is invalid")}
		}

		check := CheckPassword(input.Password, doctor.Password)
		if !check {
			return Login2faThirdPartyResponse{Token: "", Code: 401, Err: errors.New("invalid password")}
		}

		token, _ = CreateToken(map[string]interface{}{
			"doctor":      doctor.Email,
			"id":          doctor.ID,
			"name_device": nameDevice,
		})

		return Login2faThirdPartyResponse{Token: token, Code: 200, Err: nil}
	}

	return Login2faThirdPartyResponse{Token: "", Code: 400, Err: errors.New("email does not correspond to a valid patient or doctor")}
}
