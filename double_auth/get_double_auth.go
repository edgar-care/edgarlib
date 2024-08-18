package double_auth

import (
	"errors"
	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
)

type GetDoubleAuthByIdResponse struct {
	DoubleAuth model.DoubleAuth
	Code       int
	Err        error
}

func GetDoubleAuthById(id string) GetDoubleAuthByIdResponse {
	patient, err := graphql.GetPatientById(id)
	if err != nil {
		return GetDoubleAuthByIdResponse{model.DoubleAuth{}, 400, errors.New("id does not correspond to a patient")}
	}

	if patient.DoubleAuthMethodsID == nil {
		return GetDoubleAuthByIdResponse{model.DoubleAuth{}, 404, errors.New("double auth not found on patient")}
	}
	device, err := graphql.GetDoubleAuthById(*patient.DoubleAuthMethodsID)
	if err != nil {
		return GetDoubleAuthByIdResponse{model.DoubleAuth{}, 400, errors.New("id does not correspond to a double auth")}
	}

	return GetDoubleAuthByIdResponse{device, 200, nil}
}
