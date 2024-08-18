package double_auth

import (
	"errors"

	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
)

type UpdateDoubleAuthInput struct {
	Methods         string `json:"method_2fa"`
	Secret          string `json:"secret"`
	URL             string `json:"url"`
	TrustedDeviceId string `json:"trusted_device_id"`
}

type UpdateDoubleAuthResponse struct {
	DoubleAuth model.DoubleAuth
	Patient    model.Patient
	Code       int
	Err        error
}

func UpdateDoubleAuth(input UpdateDoubleAuthInput, patientId string) UpdateDoubleAuthResponse {
	patient, err := graphql.GetPatientById(patientId)
	if err != nil {
		return UpdateDoubleAuthResponse{DoubleAuth: model.DoubleAuth{}, Patient: model.Patient{}, Code: 400, Err: errors.New("id does not correspond to a patient")}
	}

	if patient.DoubleAuthMethodsID == nil || *patient.DoubleAuthMethodsID == "" {
		return UpdateDoubleAuthResponse{DoubleAuth: model.DoubleAuth{}, Patient: model.Patient{}, Code: 404, Err: errors.New("double auth not found on patient")}
	}

	updateAuth, err := graphql.GetDoubleAuthById(*patient.DoubleAuthMethodsID)
	if err != nil {
		return UpdateDoubleAuthResponse{DoubleAuth: model.DoubleAuth{}, Patient: model.Patient{}, Code: 400, Err: errors.New("id does not correspond to a double auth")}
	}

	auth, err := graphql.UpdateDoubleAuth(*patient.DoubleAuthMethodsID, model.UpdateDoubleAuthInput{
		Methods:       append(updateAuth.Methods, input.Methods),
		Secret:        &input.Secret,
		URL:           &input.URL,
		TrustDeviceID: &input.TrustedDeviceId,
	})
	if err != nil {
		return UpdateDoubleAuthResponse{DoubleAuth: model.DoubleAuth{}, Patient: model.Patient{}, Code: 400, Err: errors.New("unable  (check if you share all information)")}
	}

	return UpdateDoubleAuthResponse{
		DoubleAuth: auth,
		Code:       200,
		Err:        nil,
	}
}
