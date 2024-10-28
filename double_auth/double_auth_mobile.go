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

func CreateDoubleAuthMobile(input CreateDoubleMobileInput, ownerId string) CreateDoubleMobileResponse {
	var doubleAuthMethodsID *string
	var trustDevices []*string
	var isPatient bool

	_, err := graphql.GetDeviceConnectById(input.TrustDevice)
	if err != nil {
		return CreateDoubleMobileResponse{
			DoubleAuth: model.DoubleAuth{},
			Code:       400,
			Err:        errors.New("failed, id invalid: " + err.Error()),
		}
	}

	if input.Methods != "MOBILE" {
		return CreateDoubleMobileResponse{
			DoubleAuth: model.DoubleAuth{},
			Patient:    model.Patient{},
			Code:       http.StatusBadRequest,
			Err:        errors.New("only method_2fa MOBILE is supported"),
		}
	}

	checkPatient, errPatient := graphql.GetPatientById(ownerId)
	if errPatient == nil {
		doubleAuthMethodsID = checkPatient.DoubleAuthMethodsID
		trustDevices = checkPatient.TrustDevices
		isPatient = true
	} else {

		checkDoctor, errDoctor := graphql.GetDoctorById(ownerId)
		if errDoctor == nil {
			doubleAuthMethodsID = checkDoctor.DoubleAuthMethodsID
			trustDevices = checkDoctor.TrustDevices
			isPatient = false
		} else {

			return CreateDoubleMobileResponse{
				DoubleAuth: model.DoubleAuth{},
				Patient:    model.Patient{},
				Code:       http.StatusBadRequest,
				Err:        errors.New("unable to check double auth: owner ID does not correspond to a valid patient or doctor"),
			}
		}
	}

	if doubleAuthMethodsID == nil || *doubleAuthMethodsID == "" {
		auth, err := graphql.CreateDoubleAuth(model.CreateDoubleAuthInput{
			Methods:       []string{input.Methods},
			TrustDeviceID: []string{input.TrustDevice},
		})
		if err != nil {
			return CreateDoubleMobileResponse{
				DoubleAuth: model.DoubleAuth{},
				Patient:    model.Patient{},
				Code:       http.StatusBadRequest,
				Err:        errors.New("unable to create double auth"),
			}
		}

		if isPatient {
			_, err = graphql.UpdatePatient(ownerId, model.UpdatePatientInput{
				DoubleAuthMethodsID: &auth.ID,
				TrustDevices:        []*string{&input.TrustDevice},
			})
		} else {
			_, err = graphql.UpdateDoctor(ownerId, model.UpdateDoctorInput{
				DoubleAuthMethodsID: &auth.ID,
				TrustDevices:        []*string{&input.TrustDevice},
			})
		}
		if err != nil {
			return CreateDoubleMobileResponse{
				DoubleAuth: model.DoubleAuth{},
				Patient:    model.Patient{},
				Code:       http.StatusBadRequest,
				Err:        errors.New("update failed: " + err.Error()),
			}
		}

		status := true
		_, err = graphql.UpdateDeviceConnect(input.TrustDevice, model.UpdateDeviceConnectInput{
			TrustDevice: &status,
		})

		return CreateDoubleMobileResponse{
			DoubleAuth: auth,
			Code:       http.StatusCreated,
			Err:        nil,
		}
	} else {
		if input.Methods != "MOBILE" {
			return CreateDoubleMobileResponse{
				DoubleAuth: model.DoubleAuth{},
				Patient:    model.Patient{},
				Code:       http.StatusBadRequest,
				Err:        errors.New("only method_2fa MOBILE is supported"),
			}
		}

		if isTrustedDevice(trustDevices, input.TrustDevice) {
			return CreateDoubleMobileResponse{
				DoubleAuth: model.DoubleAuth{},
				Patient:    model.Patient{},
				Code:       http.StatusBadRequest,
				Err:        errors.New("trusted_device_id is not valid"),
			}
		}

		doubleAuth, err := graphql.GetDoubleAuthById(*doubleAuthMethodsID)
		if err != nil {
			return CreateDoubleMobileResponse{
				DoubleAuth: model.DoubleAuth{},
				Patient:    model.Patient{},
				Code:       http.StatusBadRequest,
				Err:        errors.New("get double_auth failed: " + err.Error()),
			}
		}

		if sliceContains(doubleAuth.Methods, input.Methods) {
			return CreateDoubleMobileResponse{
				DoubleAuth: doubleAuth,
				Code:       http.StatusOK,
				Err:        nil,
			}
		}

		newMethods := append(doubleAuth.Methods, input.Methods)
		updated, err := graphql.UpdateDoubleAuth(doubleAuth.ID, model.UpdateDoubleAuthInput{
			Methods:       newMethods,
			TrustDeviceID: append(doubleAuth.TrustDeviceID, input.TrustDevice),
		})
		if err != nil {
			return CreateDoubleMobileResponse{
				DoubleAuth: model.DoubleAuth{},
				Patient:    model.Patient{},
				Code:       http.StatusBadRequest,
				Err:        errors.New("update double_auth failed: " + err.Error()),
			}
		}

		status := true
		_, err = graphql.UpdateDeviceConnect(input.TrustDevice, model.UpdateDeviceConnectInput{
			TrustDevice: &status,
		})
		if err != nil {
			return CreateDoubleMobileResponse{DoubleAuth: model.DoubleAuth{}, Code: 400, Err: errors.New("update trust device failed: " + err.Error())}
		}

		if isPatient {
			if trustDevices == nil {
				trustDevices = []*string{&input.TrustDevice}
			} else {
				trustDevices = append(trustDevices, &input.TrustDevice)
			}
			_, err = graphql.UpdatePatientTrustDevice(ownerId, model.UpdatePatientTrustDeviceInput{
				TrustDevices: trustDevices,
			})
		} else {
			if trustDevices == nil {
				trustDevices = []*string{&input.TrustDevice}
			} else {
				trustDevices = append(trustDevices, &input.TrustDevice)
			}
			_, err = graphql.UpdateDoctorsTrustDevice(ownerId, model.UpdateDoctorsTrustDeviceInput{
				TrustDevices: trustDevices,
			})
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
		if device != nil && *device == trustedDeviceID {
			return true
		}
	}
	return false
}
