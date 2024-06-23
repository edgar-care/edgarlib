package double_auth

import (
	"context"
	"errors"
	"net/http"

	"github.com/edgar-care/edgarlib/graphql"
	"github.com/edgar-care/edgarlib/graphql/server/model"
)

type CreateDoubleAuthTierInput struct {
	Methods string `json:"2fa_method"`
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

	check, err := graphql.GetPatientById(context.Background(), gqlClient, patientId)
	if err != nil {
		return CreateDoubleAuthResponse{DoubleAuth: model.DoubleAuth{}, Patient: model.Patient{}, Code: http.StatusBadRequest, Err: errors.New("unable to check double auth")}
	}

	if check.GetPatientById.Double_auth_methods_id == "" {
		auth, err := graphql.CreateDoubleAuth(context.Background(), gqlClient, []string{input.Methods}, input.Code, url, "")
		if err != nil {
			return CreateDoubleAuthResponse{DoubleAuth: model.DoubleAuth{}, Patient: model.Patient{}, Code: http.StatusBadRequest, Err: errors.New("unable to create double auth")}
		}

		patient, err := graphql.GetPatientById(context.Background(), gqlClient, patientId)
		if err != nil {
			return CreateDoubleAuthResponse{DoubleAuth: model.DoubleAuth{}, Patient: model.Patient{}, Code: http.StatusBadRequest, Err: errors.New("id does not correspond to a patient")}
		}

		_, err = graphql.UpdatePatient(context.Background(), gqlClient, patientId, patient.GetPatientById.Email, patient.GetPatientById.Password, patient.GetPatientById.Medical_info_id, patient.GetPatientById.Rendez_vous_ids, patient.GetPatientById.Document_ids, patient.GetPatientById.Treatment_follow_up_ids, patient.GetPatientById.Chat_ids, patient.GetPatientById.Device_connect, auth.CreateDoubleAuth.Id)
		if err != nil {
			return CreateDoubleAuthResponse{DoubleAuth: model.DoubleAuth{}, Patient: model.Patient{}, Code: http.StatusBadRequest, Err: errors.New("update failed" + err.Error())}
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
		update := UpdateAddDoubleAuth(input.Methods, patientId)
		if update.Err != nil {
			return CreateDoubleAuthResponse{DoubleAuth: model.DoubleAuth{}, Patient: model.Patient{}, Code: http.StatusBadRequest, Err: errors.New("update failed" + update.Err.Error())}
		}
		return CreateDoubleAuthResponse{
			DoubleAuth: model.DoubleAuth{
				ID:      update.DoubleAuth.ID,
				Methods: update.DoubleAuth.Methods,
				Secret:  update.DoubleAuth.Secret,
				URL:     update.DoubleAuth.URL,
			},
			Code: http.StatusOK,
			Err:  nil,
		}
	}
}
