package double_auth

import (
	"context"
	"errors"
	"github.com/edgar-care/edgarlib/graphql"
	"github.com/edgar-care/edgarlib/graphql/server/model"
)

type GetDoubleAuthByIdResponse struct {
	DoubleAuth model.DoubleAuth
	Code       int
	Err        error
}

func GetDoubleAuthById(id string) GetDoubleAuthByIdResponse {
	gqlClient := graphql.CreateClient()
	var res model.DoubleAuth

	patient, err := graphql.GetPatientById(context.Background(), gqlClient, id)
	if err != nil {
		return GetDoubleAuthByIdResponse{model.DoubleAuth{}, 400, errors.New("id does not correspond to a patient")}
	}

	device, err := graphql.GetDoubleAuthById(context.Background(), gqlClient, patient.GetPatientById.Double_auth_methods_id)
	if err != nil {
		return GetDoubleAuthByIdResponse{model.DoubleAuth{}, 400, errors.New("id does not correspond to a double auth")}
	}
	res = model.DoubleAuth{
		ID:            device.GetDoubleAuthById.Id,
		Methods:       device.GetDoubleAuthById.Methods,
		Secret:        device.GetDoubleAuthById.Secret,
		URL:           device.GetDoubleAuthById.Url,
		TrustDeviceID: device.GetDoubleAuthById.Trust_device_id,
	}
	return GetDoubleAuthByIdResponse{res, 200, nil}
}
