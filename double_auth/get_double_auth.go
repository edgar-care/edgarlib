package double_auth

import (
	"errors"
	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
	"net/http"
)

type GetDoubleAuthByIdResponse struct {
	DoubleAuth model.DoubleAuth
	Code       int
	Err        error
}

func GetDoubleAuthById(ownerId string) GetDoubleAuthByIdResponse {

	patient, errPatient := graphql.GetPatientById(ownerId)
	if errPatient == nil {
		if patient.DoubleAuthMethodsID == nil || *patient.DoubleAuthMethodsID == "" {
			return GetDoubleAuthByIdResponse{model.DoubleAuth{Methods: []string{}}, http.StatusOK, nil}
		}
		device, err := graphql.GetDoubleAuthById(*patient.DoubleAuthMethodsID)
		if err != nil {
			return GetDoubleAuthByIdResponse{model.DoubleAuth{}, http.StatusBadRequest, errors.New("id does not correspond to a double auth")}
		}
		return GetDoubleAuthByIdResponse{device, http.StatusOK, nil}
	}

	doctor, errDoctor := graphql.GetDoctorById(ownerId)
	if errDoctor == nil {
		if doctor.DoubleAuthMethodsID == nil || *doctor.DoubleAuthMethodsID == "" {
			return GetDoubleAuthByIdResponse{model.DoubleAuth{Methods: []string{}}, http.StatusOK, nil}
		}
		device, err := graphql.GetDoubleAuthById(*doctor.DoubleAuthMethodsID)
		if err != nil {
			return GetDoubleAuthByIdResponse{model.DoubleAuth{}, http.StatusBadRequest, errors.New("id does not correspond to a double auth")}
		}
		return GetDoubleAuthByIdResponse{device, http.StatusOK, nil}
	}

	return GetDoubleAuthByIdResponse{model.DoubleAuth{}, http.StatusBadRequest, errors.New("id does not correspond to a patient or doctor")}
}
