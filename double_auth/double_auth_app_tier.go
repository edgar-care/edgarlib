package double_auth

import (
	"errors"
	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
	"github.com/pquerna/otp/totp"
	"net/http"
)

type CreateDoubleAuthTierResponse struct {
	TotpInfo TotpInfo
	Code     int
	Err      error
}

type ActivateDoubleAuthTierResponse struct {
	DoubleAuth model.DoubleAuth
	Code       int
	Err        error
}

type TotpInfo struct {
	Secret string
	Url    string
}

type GetSecretResponse struct {
	Secret string
	Code   int
	Err    error
}

func ActivateDoubleAuthTier(ownerId string) ActivateDoubleAuthTierResponse {
	var doubleAuthMethodsID *string

	checkPatient, errPatient := graphql.GetPatientById(ownerId)
	if errPatient == nil {
		doubleAuthMethodsID = checkPatient.DoubleAuthMethodsID
	} else {
		checkDoctor, errDoctor := graphql.GetDoctorById(ownerId)
		if errDoctor == nil {
			doubleAuthMethodsID = checkDoctor.DoubleAuthMethodsID
		} else {
			return ActivateDoubleAuthTierResponse{
				DoubleAuth: model.DoubleAuth{},
				Code:       http.StatusBadRequest,
				Err:        errors.New("unable to check double auth: owner ID does not correspond to a valid patient or doctor"),
			}
		}
	}
	doubleAuth, err := graphql.GetDoubleAuthById(*doubleAuthMethodsID)
	if err != nil {
		return ActivateDoubleAuthTierResponse{
			DoubleAuth: model.DoubleAuth{},
			Code:       http.StatusBadRequest,
			Err:        errors.New("get double_auth failed: " + err.Error()),
		}
	}

	if sliceContains(doubleAuth.Methods, "AUTHENTIFICATOR") {
		return ActivateDoubleAuthTierResponse{
			DoubleAuth: doubleAuth,
			Code:       http.StatusBadRequest,
			Err:        nil,
		}
	}

	newMethods := append(doubleAuth.Methods, "AUTHENTIFICATOR")
	updatedDoubleAuth, err := graphql.UpdateDoubleAuth(doubleAuth.ID, model.UpdateDoubleAuthInput{
		Methods:       newMethods,
		TrustDeviceID: doubleAuth.TrustDeviceID,
	})
	if err != nil {
		return ActivateDoubleAuthTierResponse{
			DoubleAuth: model.DoubleAuth{},
			Code:       http.StatusBadRequest,
			Err:        errors.New("update double_auth failed: " + err.Error()),
		}
	}

	return ActivateDoubleAuthTierResponse{
		DoubleAuth: updatedDoubleAuth,
		Code:       http.StatusOK,
		Err:        nil,
	}
}

func CreateDoubleAuthAppTier(ownerId string) CreateDoubleAuthTierResponse {
	var doubleAuthMethodsID *string
	var userEmail string
	var isPatient bool

	checkPatient, errPatient := graphql.GetPatientById(ownerId)
	if errPatient == nil {
		userEmail = checkPatient.Email
		doubleAuthMethodsID = checkPatient.DoubleAuthMethodsID
		isPatient = true
	} else {
		checkDoctor, errDoctor := graphql.GetDoctorById(ownerId)
		if errDoctor == nil {
			userEmail = checkDoctor.Email
			doubleAuthMethodsID = checkDoctor.DoubleAuthMethodsID
			isPatient = false
		} else {
			return CreateDoubleAuthTierResponse{
				TotpInfo: TotpInfo{},
				Code:     http.StatusBadRequest,
				Err:      errors.New("unable to check double auth: owner ID does not correspond to a valid patient or doctor"),
			}
		}
	}

	key, err := totp.Generate(totp.GenerateOpts{
		Issuer:      "edgar-sante.fr",
		AccountName: userEmail,
		SecretSize:  15,
	})
	if err != nil {
		return CreateDoubleAuthTierResponse{
			TotpInfo: TotpInfo{},
			Code:     http.StatusBadRequest,
			Err:      errors.New("unable to generate secret"),
		}
	}

	secret := key.Secret()

	if doubleAuthMethodsID == nil || *doubleAuthMethodsID == "" {
		auth, err := graphql.CreateDoubleAuth(model.CreateDoubleAuthInput{
			Methods: []string{},
			Code:    secret,
		})
		if err != nil {
			return CreateDoubleAuthTierResponse{
				TotpInfo: TotpInfo{},
				Code:     http.StatusBadRequest,
				Err:      errors.New("unable to create double auth"),
			}
		}

		if isPatient {
			_, err = graphql.UpdatePatient(ownerId, model.UpdatePatientInput{DoubleAuthMethodsID: &auth.ID})
			if err != nil {
				return CreateDoubleAuthTierResponse{
					TotpInfo: TotpInfo{},
					Code:     400,
					Err:      errors.New("failed to update patient: " + err.Error()),
				}
			}
		} else {
			_, err := graphql.UpdateDoctor(ownerId, model.UpdateDoctorInput{DoubleAuthMethodsID: &auth.ID})
			if err != nil {
				return CreateDoubleAuthTierResponse{
					TotpInfo: TotpInfo{},
					Code:     400,
					Err:      errors.New("failed to update doctor: " + err.Error()),
				}
			}
		}
		if err != nil {
			return CreateDoubleAuthTierResponse{
				TotpInfo: TotpInfo{},
				Code:     http.StatusBadRequest,
				Err:      errors.New("update failed: " + err.Error()),
			}
		}

		return CreateDoubleAuthTierResponse{
			TotpInfo: TotpInfo{Secret: secret, Url: key.URL()},
			Code:     http.StatusCreated,
			Err:      nil,
		}
	}

	doubleAuth, err := graphql.GetDoubleAuthById(*doubleAuthMethodsID)
	if err != nil {
		return CreateDoubleAuthTierResponse{
			TotpInfo: TotpInfo{},
			Code:     http.StatusBadRequest,
			Err:      errors.New("get double_auth failed: " + err.Error()),
		}
	}

	if sliceContains(doubleAuth.Methods, "AUTHENTIFICATOR") {
		return CreateDoubleAuthTierResponse{
			TotpInfo: TotpInfo{},
			Code:     http.StatusBadRequest,
			Err:      nil,
		}
	}

	_, err = graphql.UpdateDoubleAuth(doubleAuth.ID, model.UpdateDoubleAuthInput{
		Methods:       doubleAuth.Methods,
		Code:          &secret,
		TrustDeviceID: doubleAuth.TrustDeviceID,
	})
	if err != nil {
		return CreateDoubleAuthTierResponse{
			TotpInfo: TotpInfo{},
			Code:     http.StatusBadRequest,
			Err:      errors.New("update double_auth failed: " + err.Error()),
		}
	}

	return CreateDoubleAuthTierResponse{
		TotpInfo: TotpInfo{Secret: secret, Url: key.URL()},
		Code:     http.StatusOK,
		Err:      nil,
	}
}

func GetSecretThirdParty(ownerID string) GetSecretResponse {

	var authMethodsId string

	patient, errPatient := graphql.GetPatientById(ownerID)
	if errPatient != nil {
		doctor, errDoctor := graphql.GetDoctorById(ownerID)
		if errDoctor != nil {
			return GetSecretResponse{
				Secret: "",
				Code:   http.StatusBadRequest,
				Err:    errors.New("Id does not correspond to a valid patient or doctor"),
			}
		} else {
			if doctor.DoubleAuthMethodsID == nil {
				return GetSecretResponse{
					Secret: "",
					Code:   http.StatusBadRequest,
					Err:    errors.New("Unable to check double auth: auth method id is missing"),
				}
			}
			authMethodsId = *doctor.DoubleAuthMethodsID
		}
	} else {
		if patient.DoubleAuthMethodsID == nil {
			return GetSecretResponse{
				Secret: "",
				Code:   http.StatusBadRequest,
				Err:    errors.New("Unable to check double auth: auth method id is missing"),
			}
		}
		authMethodsId = *patient.DoubleAuthMethodsID
	}

	getDoubleAuth, err := graphql.GetDoubleAuthById(authMethodsId)
	if err != nil {
		return GetSecretResponse{
			Secret: "",
			Code:   http.StatusBadRequest,
			Err:    errors.New("Unable to check double auth: owner ID does not correspond to a valid patient or doctor"),
		}
	}
	if getDoubleAuth.Code == "" {
		return GetSecretResponse{
			Secret: "",
			Code:   http.StatusBadRequest,
			Err:    errors.New("Error: value secret empty"),
		}
	}

	return GetSecretResponse{
		Secret: getDoubleAuth.Code,
		Code:   http.StatusOK,
		Err:    nil,
	}

}
