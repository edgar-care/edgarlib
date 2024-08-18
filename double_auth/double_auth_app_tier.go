package double_auth

import (
	"errors"
	"net/http"

	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
)

type CreateDoubleAuthTierInput struct {
	Methods string `json:"method_2fa"`
	//Code    string `json:"code"`
}

type CreateDoubleAuthTierResponse struct {
	DoubleAuth model.DoubleAuth
	Patient    model.Patient
	Code       int
	Err        error
}

func CreateDoubleAuthAppTier(input CreateDoubleAuthTierInput, url string, patientId string) CreateDoubleAuthResponse {

	check, err := graphql.GetPatientById(patientId)
	if err != nil {
		return CreateDoubleAuthResponse{DoubleAuth: model.DoubleAuth{}, Patient: model.Patient{}, Code: http.StatusBadRequest, Err: errors.New("unable to check double auth")}
	}

	if check.DoubleAuthMethodsID == nil || *check.DoubleAuthMethodsID == "" {
		auth, err := graphql.CreateDoubleAuth(model.CreateDoubleAuthInput{
			Methods: []string{input.Methods},
			URL:     url,
		})
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
		if input.Methods != "AUTHENTIFICATOR" {
			return CreateDoubleAuthResponse{DoubleAuth: model.DoubleAuth{}, Patient: model.Patient{}, Code: http.StatusBadRequest, Err: errors.New("only method_2fa AUTHENTIFICATOR is supported")}
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
		updatedDoubleAuth, err := graphql.UpdateDoubleAuth(doubleAuth.ID, model.UpdateDoubleAuthInput{Methods: newMethods, URL: &url})
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
