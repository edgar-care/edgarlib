package double_auth

import (
	"errors"
	"net/http"

	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
)

type CreateDoubleAuthInput struct {
	Methods string `json:"method_2fa"`
}

type CreateDoubleAuthResponse struct {
	DoubleAuth model.DoubleAuth
	Code       int
	Err        error
}

func CreateDoubleAuthEmail(input CreateDoubleAuthInput, ownerId string) CreateDoubleAuthResponse {
	var doubleAuthMethodsID *string
	var isPatient bool

	if input.Methods != "EMAIL" {
		return CreateDoubleAuthResponse{
			DoubleAuth: model.DoubleAuth{},
			Code:       http.StatusBadRequest,
			Err:        errors.New("only method_2fa EMAIL is supported"),
		}
	}

	checkPatient, errPatient := graphql.GetPatientById(ownerId)
	if errPatient == nil {
		doubleAuthMethodsID = checkPatient.DoubleAuthMethodsID
		isPatient = true
	} else {

		checkDoctor, errDoctor := graphql.GetDoctorById(ownerId)
		if errDoctor == nil {
			doubleAuthMethodsID = checkDoctor.DoubleAuthMethodsID
			isPatient = false
		} else {

			return CreateDoubleAuthResponse{
				DoubleAuth: model.DoubleAuth{},
				Code:       http.StatusBadRequest,
				Err:        errors.New("unable to check double auth: owner ID does not correspond to a valid patient or doctor"),
			}
		}
	}

	if doubleAuthMethodsID == nil || *doubleAuthMethodsID == "" {
		auth, err := graphql.CreateDoubleAuth(model.CreateDoubleAuthInput{
			Methods: []string{"EMAIL"},
		})
		if err != nil {
			return CreateDoubleAuthResponse{
				DoubleAuth: model.DoubleAuth{},
				Code:       http.StatusBadRequest,
				Err:        errors.New("unable to create double auth"),
			}
		}

		if isPatient {
			_, err = graphql.UpdatePatient(ownerId, model.UpdatePatientInput{DoubleAuthMethodsID: &auth.ID})
			if err != nil {
				return CreateDoubleAuthResponse{DoubleAuth: model.DoubleAuth{}, Code: 400, Err: errors.New("update patient failed: " + err.Error())}
			}
		} else {
			_, err = graphql.UpdateDoctor(ownerId, model.UpdateDoctorInput{DoubleAuthMethodsID: &auth.ID})
			if err != nil {
				return CreateDoubleAuthResponse{DoubleAuth: model.DoubleAuth{}, Code: 400, Err: errors.New("update doctor failed: " + err.Error())}
			}
		}
		if err != nil {
			return CreateDoubleAuthResponse{
				DoubleAuth: model.DoubleAuth{},
				Code:       http.StatusBadRequest,
				Err:        errors.New("update failed: " + err.Error()),
			}
		}

		return CreateDoubleAuthResponse{
			DoubleAuth: auth,
			Code:       http.StatusCreated,
			Err:        nil,
		}
	}

	if input.Methods != "EMAIL" {
		return CreateDoubleAuthResponse{
			DoubleAuth: model.DoubleAuth{},
			Code:       http.StatusBadRequest,
			Err:        errors.New("only method_2fa EMAIL is supported"),
		}
	}

	doubleAuth, err := graphql.GetDoubleAuthById(*doubleAuthMethodsID)
	if err != nil {
		return CreateDoubleAuthResponse{
			DoubleAuth: model.DoubleAuth{},
			Code:       http.StatusBadRequest,
			Err:        errors.New("get double_auth failed: " + err.Error()),
		}
	}

	if sliceContains(doubleAuth.Methods, input.Methods) {
		return CreateDoubleAuthResponse{
			DoubleAuth: doubleAuth,
			Code:       http.StatusOK,
			Err:        nil,
		}
	}

	newMethods := append(doubleAuth.Methods, input.Methods)
	updatedDoubleAuth, err := graphql.UpdateDoubleAuth(doubleAuth.ID, model.UpdateDoubleAuthInput{Methods: newMethods, TrustDeviceID: doubleAuth.TrustDeviceID})
	if err != nil {
		return CreateDoubleAuthResponse{
			DoubleAuth: model.DoubleAuth{},
			Code:       http.StatusBadRequest,
			Err:        errors.New("update double_auth failed: " + err.Error()),
		}
	}

	return CreateDoubleAuthResponse{
		DoubleAuth: updatedDoubleAuth,
		Code:       http.StatusOK,
		Err:        nil,
	}
}

func sliceContains(slice []string, s string) bool {
	for _, value := range slice {
		if value == s {
			return true
		}
	}
	return false
}
