package double_auth

import (
	"errors"
	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
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

func CreateDoubleAuthMobile(input CreateDoubleMobileInput, patientId string) CreateDoubleMobileResponse {
	check, err := graphql.GetPatientById(patientId)
	if err != nil {
		return CreateDoubleMobileResponse{DoubleAuth: model.DoubleAuth{}, Patient: model.Patient{}, Code: http.StatusBadRequest, Err: errors.New("unable to check double auth")}
	}

	if check.DoubleAuthMethodsID == nil || *check.DoubleAuthMethodsID == "" {
		auth, err := graphql.CreateDoubleAuth(model.CreateDoubleAuthInput{
			Methods:       []string{input.Methods},
			TrustDeviceID: input.TrustDevice,
		})
		if err != nil {
			return CreateDoubleMobileResponse{DoubleAuth: model.DoubleAuth{}, Patient: model.Patient{}, Code: http.StatusBadRequest, Err: errors.New("unable to create double auth")}
		}

		_, err = graphql.UpdatePatient(patientId, model.UpdatePatientInput{DoubleAuthMethodsID: &auth.ID, TrustDevices: []*string{&input.TrustDevice}})
		if err != nil {
			return CreateDoubleMobileResponse{DoubleAuth: model.DoubleAuth{}, Patient: model.Patient{}, Code: http.StatusBadRequest, Err: errors.New("update failed: " + err.Error())}
		}

		return CreateDoubleMobileResponse{
			DoubleAuth: auth,
			Code:       http.StatusCreated,
			Err:        nil,
		}
	} else {
		if input.Methods != "MOBILE" {
			return CreateDoubleMobileResponse{DoubleAuth: model.DoubleAuth{}, Patient: model.Patient{}, Code: http.StatusBadRequest, Err: errors.New("only method_2fa MOBILE is supported")}
		}

		if !isTrustedDevice(check.TrustDevices, input.TrustDevice) {
			return CreateDoubleMobileResponse{DoubleAuth: model.DoubleAuth{}, Patient: model.Patient{}, Code: http.StatusBadRequest, Err: errors.New("trusted_device_id is not valid")}
		}

		doubleAuth, err := graphql.GetDoubleAuthById(*check.DoubleAuthMethodsID)
		if err != nil {
			return CreateDoubleMobileResponse{DoubleAuth: model.DoubleAuth{}, Patient: model.Patient{}, Code: http.StatusBadRequest, Err: errors.New("get double_auth failed: " + err.Error())}
		}

		if sliceContains(doubleAuth.Methods, input.Methods) {
			return CreateDoubleMobileResponse{
				DoubleAuth: doubleAuth,
				Code:       http.StatusOK,
				Err:        nil,
			}
		}

		newMethods := append(doubleAuth.Methods, input.Methods)
		updated, err := graphql.UpdateDoubleAuth(doubleAuth.ID, model.UpdateDoubleAuthInput{Methods: newMethods, TrustDeviceID: &input.TrustDevice})
		if err != nil {
			return CreateDoubleMobileResponse{DoubleAuth: model.DoubleAuth{}, Patient: model.Patient{}, Code: http.StatusBadRequest, Err: errors.New("update double_auth failed: " + err.Error())}
		}

		return CreateDoubleMobileResponse{
			DoubleAuth: updated,
			Code:       http.StatusOK,
			Err:        nil,
		}
	}
}

func isTrustedDevice(trustDevices []*string, trustedDeviceID string) bool {
	for _, device := range trustDevices {
		if device == &trustedDeviceID {
			return true
		}
	}
	return false
}
