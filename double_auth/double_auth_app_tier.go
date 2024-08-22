package double_auth

import (
	"errors"
	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
	"net/http"
)

type CreateDoubleAuthTierResponse struct {
	DoubleAuth model.DoubleAuth
	Code       int
	Err        error
}

type DecrypteSecretResponse struct {
	Secret string
	Code   int
	Err    error
}

func CreateDoubleAuthAppTier(ownerId string, secret string) CreateDoubleAuthResponse {
	var doubleAuthMethodsID *string
	var isPatient bool

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
			Methods: []string{"AUTHENTIFICATOR"},
			Code:    secret,
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

	if sliceContains(doubleAuth.Methods, "AUTHENTIFICATOR") {
		return CreateDoubleAuthResponse{
			DoubleAuth: doubleAuth,
			Code:       http.StatusOK,
			Err:        nil,
		}
	}

	newMethods := append(doubleAuth.Methods, "AUTHENTIFICATOR")
	updatedDoubleAuth, err := graphql.UpdateDoubleAuth(doubleAuth.ID, model.UpdateDoubleAuthInput{
		Methods: newMethods,
		Code:    &secret,
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

func GetSecretThirdParty(ownerID string) DecrypteSecretResponse {

	var authMethodsId string

	patient, errPatient := graphql.GetPatientById(ownerID)
	if errPatient != nil {
		doctor, errDoctor := graphql.GetDoctorById(ownerID)
		if errDoctor != nil {
			return DecrypteSecretResponse{
				Secret: "",
				Code:   http.StatusBadRequest,
				Err:    errors.New("Id does not correspond to a valid patient or doctor"),
			}
		} else {
			if doctor.DoubleAuthMethodsID == nil {
				return DecrypteSecretResponse{
					Secret: "",
					Code:   http.StatusBadRequest,
					Err:    errors.New("Unable to check double auth: auth method id is missing"),
				}
			}
			authMethodsId = *doctor.DoubleAuthMethodsID
		}
	} else {
		if patient.DoubleAuthMethodsID == nil {
			return DecrypteSecretResponse{
				Secret: "",
				Code:   http.StatusBadRequest,
				Err:    errors.New("Unable to check double auth: auth method id is missing"),
			}
		}
		authMethodsId = *patient.DoubleAuthMethodsID
	}

	getDoubleAuth, err := graphql.GetDoubleAuthById(authMethodsId)
	if err != nil {
		return DecrypteSecretResponse{
			Secret: "",
			Code:   http.StatusBadRequest,
			Err:    errors.New("Unable to check double auth: owner ID does not correspond to a valid patient or doctor"),
		}
	}
	if getDoubleAuth.Code == "" {
		return DecrypteSecretResponse{
			Secret: "",
			Code:   http.StatusBadRequest,
			Err:    errors.New("Error: value secret empty"),
		}
	}

	return DecrypteSecretResponse{
		Secret: getDoubleAuth.Code,
		Code:   http.StatusOK,
		Err:    nil,
	}

}
