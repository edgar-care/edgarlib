package double_auth

import (
	"context"
	"errors"
	"net/http"

	"github.com/edgar-care/edgarlib/graphql"
	"github.com/edgar-care/edgarlib/graphql/server/model"
)

type CreateDoubleAuthTierInput struct {
	Methods string `json:"method_2fa"`
	Code    string `json:"code"`
}

type CreateDoubleAuthTierResponse struct {
	DoubleAuth model.DoubleAuth
	Patient    model.Patient
	Code       int
	Err        error
}

func CreateDoubleAuthAppTier(input CreateDoubleAuthTierInput, url string, patientId string) CreateDoubleAuthResponse {
	gqlClient := graphql.CreateClient()

	// Check if patient exists and retrieve their current double_auth_methods_id
	check, err := graphql.GetPatientById(context.Background(), gqlClient, patientId)
	if err != nil {
		return CreateDoubleAuthResponse{DoubleAuth: model.DoubleAuth{}, Patient: model.Patient{}, Code: http.StatusBadRequest, Err: errors.New("unable to check double auth")}
	}

	// If patient does not have a double_auth_methods_id, create new double auth
	if check.GetPatientById.Double_auth_methods_id == "" {
		auth, err := graphql.CreateDoubleAuth(context.Background(), gqlClient, []string{input.Methods}, input.Code, url, "")
		if err != nil {
			return CreateDoubleAuthResponse{DoubleAuth: model.DoubleAuth{}, Patient: model.Patient{}, Code: http.StatusBadRequest, Err: errors.New("unable to create double auth")}
		}

		// Update patient with new double_auth_methods_id
		_, err = graphql.UpdatePatient(context.Background(), gqlClient, patientId, check.GetPatientById.Email, check.GetPatientById.Password, check.GetPatientById.Medical_info_id, check.GetPatientById.Rendez_vous_ids, check.GetPatientById.Document_ids, check.GetPatientById.Treatment_follow_up_ids, check.GetPatientById.Chat_ids, check.GetPatientById.Device_connect, auth.CreateDoubleAuth.Id, check.GetPatientById.Trust_devices, check.GetPatientById.Status)
		if err != nil {
			return CreateDoubleAuthResponse{DoubleAuth: model.DoubleAuth{}, Patient: model.Patient{}, Code: http.StatusBadRequest, Err: errors.New("update failed: " + err.Error())}
		}

		return CreateDoubleAuthResponse{
			DoubleAuth: model.DoubleAuth{
				ID:            auth.CreateDoubleAuth.Id,
				Methods:       auth.CreateDoubleAuth.Methods,
				Secret:        auth.CreateDoubleAuth.Secret,
				URL:           auth.CreateDoubleAuth.Url,
				TrustDeviceID: auth.CreateDoubleAuth.Trust_device_id,
			},
			Code: http.StatusCreated,
			Err:  nil,
		}
	} else {
		// Patient already has a double_auth_methods_id, check if email method exists
		doubleAuth, err := graphql.GetDoubleAuthById(context.Background(), gqlClient, check.GetPatientById.Double_auth_methods_id)
		if err != nil {
			return CreateDoubleAuthResponse{DoubleAuth: model.DoubleAuth{}, Patient: model.Patient{}, Code: http.StatusBadRequest, Err: errors.New("get double_auth failed: " + err.Error())}
		}

		// Check if the email method already exists in the current methods
		if sliceContains(doubleAuth.GetDoubleAuthById.Methods, input.Methods) {
			return CreateDoubleAuthResponse{
				DoubleAuth: model.DoubleAuth{
					ID:      doubleAuth.GetDoubleAuthById.Id,
					Methods: doubleAuth.GetDoubleAuthById.Methods,
				},
				Code: http.StatusOK,
				Err:  nil,
			}
		}

		// If not already exists, update double_auth with new method
		newMethods := append(doubleAuth.GetDoubleAuthById.Methods, input.Methods)
		_, err = graphql.UpdateDoubleAuth(context.Background(), gqlClient, doubleAuth.GetDoubleAuthById.Id, newMethods, doubleAuth.GetDoubleAuthById.Secret, doubleAuth.GetDoubleAuthById.Url, doubleAuth.GetDoubleAuthById.Trust_device_id)
		if err != nil {
			return CreateDoubleAuthResponse{DoubleAuth: model.DoubleAuth{}, Patient: model.Patient{}, Code: http.StatusBadRequest, Err: errors.New("update double_auth failed: " + err.Error())}
		}

		return CreateDoubleAuthResponse{
			DoubleAuth: model.DoubleAuth{
				ID:      doubleAuth.GetDoubleAuthById.Id,
				Methods: newMethods,
			},
			Code: http.StatusOK,
			Err:  nil,
		}
	}
}
