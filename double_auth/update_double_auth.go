package double_auth

import (
	"context"
	"errors"

	"github.com/edgar-care/edgarlib/graphql"
	"github.com/edgar-care/edgarlib/graphql/server/model"
)

type UpdateDoubleAuthInput struct {
	Methods         string `json:"2fa_method"`
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

func UpdateAddDoubleAuth(method string, patientId string) UpdateDoubleAuthResponse {
	gqlClient := graphql.CreateClient()

	patient, err := graphql.GetPatientById(context.Background(), gqlClient, patientId)
	if err != nil {
		return UpdateDoubleAuthResponse{DoubleAuth: model.DoubleAuth{}, Patient: model.Patient{}, Code: 400, Err: errors.New("id does not correspond to a patient")}
	}

	updateAuth, err := graphql.GetDoubleAuthById(context.Background(), gqlClient, patient.GetPatientById.Double_auth_methods_id)
	if err != nil {
		return UpdateDoubleAuthResponse{DoubleAuth: model.DoubleAuth{}, Patient: model.Patient{}, Code: 400, Err: errors.New("id does not correspond to a double auth")}
	}

	auth, err := graphql.UpdateDoubleAuth(context.Background(), gqlClient, patient.GetPatientById.Double_auth_methods_id, append(updateAuth.GetDoubleAuthById.Methods, method), updateAuth.GetDoubleAuthById.Secret, updateAuth.GetDoubleAuthById.Url, updateAuth.GetDoubleAuthById.Trust_device_id)
	if err != nil {
		return UpdateDoubleAuthResponse{DoubleAuth: model.DoubleAuth{}, Patient: model.Patient{}, Code: 400, Err: errors.New("unable  (check if you share all information)")}
	}

	return UpdateDoubleAuthResponse{
		DoubleAuth: model.DoubleAuth{
			ID:            auth.UpdateDoubleAuth.Id,
			Methods:       auth.UpdateDoubleAuth.Methods,
			Secret:        auth.UpdateDoubleAuth.Secret,
			URL:           auth.UpdateDoubleAuth.Url,
			TrustDeviceID: auth.UpdateDoubleAuth.Trust_device_id,
		},
		Code: 200,
		Err:  nil,
	}
}

func UpdateDoubleAuth(input UpdateDoubleAuthInput, patientId string) UpdateDoubleAuthResponse {
	gqlClient := graphql.CreateClient()

	patient, err := graphql.GetPatientById(context.Background(), gqlClient, patientId)
	if err != nil {
		return UpdateDoubleAuthResponse{DoubleAuth: model.DoubleAuth{}, Patient: model.Patient{}, Code: 400, Err: errors.New("id does not correspond to a patient")}
	}

	updateAuth, err := graphql.GetDoubleAuthById(context.Background(), gqlClient, patient.GetPatientById.Double_auth_methods_id)
	if err != nil {
		return UpdateDoubleAuthResponse{DoubleAuth: model.DoubleAuth{}, Patient: model.Patient{}, Code: 400, Err: errors.New("id does not correspond to a double auth")}
	}

	auth, err := graphql.UpdateDoubleAuth(context.Background(), gqlClient, patient.GetPatientById.Double_auth_methods_id, updateAuth.GetDoubleAuthById.Methods, input.Secret, input.URL, input.TrustedDeviceId)
	if err != nil {
		return UpdateDoubleAuthResponse{DoubleAuth: model.DoubleAuth{}, Patient: model.Patient{}, Code: 400, Err: errors.New("unable  (check if you share all information)")}
	}

	return UpdateDoubleAuthResponse{
		DoubleAuth: model.DoubleAuth{
			ID:            auth.UpdateDoubleAuth.Id,
			Methods:       auth.UpdateDoubleAuth.Methods,
			Secret:        auth.UpdateDoubleAuth.Secret,
			URL:           auth.UpdateDoubleAuth.Url,
			TrustDeviceID: auth.UpdateDoubleAuth.Trust_device_id,
		},
		Code: 200,
		Err:  nil,
	}
}
