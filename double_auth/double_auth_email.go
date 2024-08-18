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
	Patient    model.Patient
	Code       int
	Err        error
}

func CreateDoubleAuthEmail(input CreateDoubleAuthInput, patientId string) CreateDoubleAuthResponse {
	check, err := graphql.GetPatientById(patientId)
	if err != nil {
		return CreateDoubleAuthResponse{DoubleAuth: model.DoubleAuth{}, Patient: model.Patient{}, Code: http.StatusBadRequest, Err: errors.New("unable to check double auth")}
	}

	if check.DoubleAuthMethodsID == nil || *check.DoubleAuthMethodsID == "" {
		auth, err := graphql.CreateDoubleAuth(model.CreateDoubleAuthInput{})
		if err != nil {
			return CreateDoubleAuthResponse{DoubleAuth: model.DoubleAuth{}, Patient: model.Patient{}, Code: http.StatusBadRequest, Err: errors.New("unable to create double auth")}
		}

		_, err = graphql.UpdatePatient(patientId, model.UpdatePatientInput{DoubleAuthMethodsID: &auth.ID})
		if err != nil {
			return CreateDoubleAuthResponse{DoubleAuth: model.DoubleAuth{}, Patient: model.Patient{}, Code: http.StatusBadRequest, Err: errors.New("update failed: " + err.Error())}
		}

		return CreateDoubleAuthResponse{
			DoubleAuth: auth,
			Code:       http.StatusCreated,
			Err:        nil,
		}
	} else {
		if input.Methods != "EMAIL" {
			return CreateDoubleAuthResponse{DoubleAuth: model.DoubleAuth{}, Patient: model.Patient{}, Code: http.StatusBadRequest, Err: errors.New("only method_2fa EMAIL is supported")}
		}
		doubleAuth, err := graphql.GetDoubleAuthById(*check.DoubleAuthMethodsID)
		if err != nil {
			return CreateDoubleAuthResponse{DoubleAuth: model.DoubleAuth{}, Patient: model.Patient{}, Code: http.StatusBadRequest, Err: errors.New("get double_auth failed: " + err.Error())}
		}

		if sliceContains(doubleAuth.Methods, input.Methods) {
			return CreateDoubleAuthResponse{
				DoubleAuth: doubleAuth,
				Code:       http.StatusOK,
				Err:        nil,
			}
		}

		newMethods := append(doubleAuth.Methods, input.Methods)
		updatedDoubleAuth, err := graphql.UpdateDoubleAuth(doubleAuth.ID, model.UpdateDoubleAuthInput{Methods: newMethods})
		if err != nil {
			return CreateDoubleAuthResponse{DoubleAuth: model.DoubleAuth{}, Patient: model.Patient{}, Code: http.StatusBadRequest, Err: errors.New("update double_auth failed: " + err.Error())}
		}

		return CreateDoubleAuthResponse{
			DoubleAuth: updatedDoubleAuth,
			Code:       http.StatusOK,
			Err:        nil,
		}
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
