package double_auth

import (
	"errors"
	"net/http"

	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
)

type CreateDoubleAuthTierInput struct {
	Methods string `json:"method_2fa"`
	Code    string `json:"code"`
}

type CreateDoubleAuthTierResponse struct {
	DoubleAuth model.DoubleAuth
	Code       int
	Err        error
}

func CreateDoubleAuthAppTier(input CreateDoubleAuthTierInput, url string, ownerId string) CreateDoubleAuthResponse {
	var doubleAuthMethodsID *string
	var isPatient bool

	if input.Methods != "AUTHENTIFICATOR" {
		return CreateDoubleAuthResponse{
			DoubleAuth: model.DoubleAuth{},
			Code:       http.StatusBadRequest,
			Err:        errors.New("only method_2fa AUTHENTIFICATOR is supported"),
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
			Methods: []string{input.Methods},
			URL:     url,
			Secret:  input.Code,
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
				return CreateDoubleAuthResponse{
					DoubleAuth: model.DoubleAuth{},
					Code:       400,
					Err:        errors.New("failed to update patient: " + err.Error()),
				}
			}
		} else {
			_, err := graphql.UpdateDoctor(ownerId, model.UpdateDoctorInput{DoubleAuthMethodsID: &auth.ID})
			if err != nil {
				return CreateDoubleAuthResponse{
					DoubleAuth: model.DoubleAuth{},
					Code:       400,
					Err:        errors.New("failed to update doctor: " + err.Error()),
				}
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
	updatedDoubleAuth, err := graphql.UpdateDoubleAuth(doubleAuth.ID, model.UpdateDoubleAuthInput{
		Methods: newMethods,
		URL:     &url,
		Secret:  &input.Code,
	})
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
