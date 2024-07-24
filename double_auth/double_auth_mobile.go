package double_auth

import (
	"context"
	"errors"
	"github.com/edgar-care/edgarlib/graphql"
	"github.com/edgar-care/edgarlib/graphql/server/model"
	"net/http"
)

type CreateDoubleMobileInput struct {
	Methods     string `json:"method_2fa"`
	TrustDevice string `json:"trusted_device_id"`
}

type CreateDoubleMobileResponse struct {
	DoubleAuth model.DoubleAuth
	Patient    model.Patient
	Code       int
	Err        error
}

func CreateDoubleAuthMobile(input CreateDoubleMobileInput, url string, patientId string) CreateDoubleMobileResponse {
	gqlClient := graphql.CreateClient()

	check, err := graphql.GetPatientById(context.Background(), gqlClient, patientId)
	if err != nil {
		return CreateDoubleMobileResponse{DoubleAuth: model.DoubleAuth{}, Patient: model.Patient{}, Code: http.StatusBadRequest, Err: errors.New("unable to check double auth")}
	}

	if check.GetPatientById.Double_auth_methods_id == "" {
		auth, err := graphql.CreateDoubleAuth(context.Background(), gqlClient, []string{input.Methods}, "", url, input.TrustDevice)
		if err != nil {
			return CreateDoubleMobileResponse{DoubleAuth: model.DoubleAuth{}, Patient: model.Patient{}, Code: http.StatusBadRequest, Err: errors.New("unable to create double auth")}
		}

		_, err = graphql.UpdatePatient(context.Background(), gqlClient, patientId, check.GetPatientById.Email, check.GetPatientById.Password, check.GetPatientById.Medical_info_id, check.GetPatientById.Rendez_vous_ids, check.GetPatientById.Document_ids, check.GetPatientById.Treatment_follow_up_ids, check.GetPatientById.Chat_ids, check.GetPatientById.Device_connect, auth.CreateDoubleAuth.Id, check.GetPatientById.Trust_devices, check.GetPatientById.Status)
		if err != nil {
			return CreateDoubleMobileResponse{DoubleAuth: model.DoubleAuth{}, Patient: model.Patient{}, Code: http.StatusBadRequest, Err: errors.New("update failed: " + err.Error())}
		}

		return CreateDoubleMobileResponse{
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
		if input.Methods != "MOBILE" {
			return CreateDoubleMobileResponse{DoubleAuth: model.DoubleAuth{}, Patient: model.Patient{}, Code: http.StatusBadRequest, Err: errors.New("only method_2fa MOBILE is supported")}
		}

		if !isTrustedDevice(check.GetPatientById.Trust_devices, input.TrustDevice) {
			return CreateDoubleMobileResponse{DoubleAuth: model.DoubleAuth{}, Patient: model.Patient{}, Code: http.StatusBadRequest, Err: errors.New("trusted_device_id is not valid")}
		}

		doubleAuth, err := graphql.GetDoubleAuthById(context.Background(), gqlClient, check.GetPatientById.Double_auth_methods_id)
		if err != nil {
			return CreateDoubleMobileResponse{DoubleAuth: model.DoubleAuth{}, Patient: model.Patient{}, Code: http.StatusBadRequest, Err: errors.New("get double_auth failed: " + err.Error())}
		}

		if sliceContains(doubleAuth.GetDoubleAuthById.Methods, input.Methods) {
			return CreateDoubleMobileResponse{
				DoubleAuth: model.DoubleAuth{
					ID:      doubleAuth.GetDoubleAuthById.Id,
					Methods: doubleAuth.GetDoubleAuthById.Methods,
				},
				Code: http.StatusOK,
				Err:  nil,
			}
		}

		newMethods := append(doubleAuth.GetDoubleAuthById.Methods, input.Methods)
		_, err = graphql.UpdateDoubleAuth(context.Background(), gqlClient, doubleAuth.GetDoubleAuthById.Id, newMethods, doubleAuth.GetDoubleAuthById.Secret, doubleAuth.GetDoubleAuthById.Url, doubleAuth.GetDoubleAuthById.Trust_device_id)
		if err != nil {
			return CreateDoubleMobileResponse{DoubleAuth: model.DoubleAuth{}, Patient: model.Patient{}, Code: http.StatusBadRequest, Err: errors.New("update double_auth failed: " + err.Error())}
		}

		return CreateDoubleMobileResponse{
			DoubleAuth: model.DoubleAuth{
				ID:      doubleAuth.GetDoubleAuthById.Id,
				Methods: newMethods,
			},
			Code: http.StatusOK,
			Err:  nil,
		}
	}
}

func isTrustedDevice(trustDevices []string, trustedDeviceID string) bool {
	for _, device := range trustDevices {
		if device == trustedDeviceID {
			return true
		}
	}
	return false
}
