package auth

import (
	"errors"
	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
	"net/http"
)

type CheckAccountStatusResponse struct {
	Status bool
	Code   int
	Err    error
}

type ModifyAccountStatusResponse struct {
	Patient model.Patient
	Doctor  model.Doctor
	Code    int
	Err     error
}

func CheckAccountEnable(id string) CheckAccountStatusResponse {
	patient, err := graphql.GetPatientById(id)
	if err == nil {
		if patient.Status == false {
			return CheckAccountStatusResponse{Status: false, Code: 409, Err: errors.New("patient account is disabled")}
		}
		return CheckAccountStatusResponse{Status: true, Code: http.StatusOK, Err: nil}
	}

	doctor, err := graphql.GetDoctorById(id)
	if err == nil {
		if doctor.Status == false {
			return CheckAccountStatusResponse{Status: false, Code: 409, Err: errors.New("doctor account is disabled")}
		}
		return CheckAccountStatusResponse{Status: true, Code: http.StatusOK, Err: nil}
	}

	return CheckAccountStatusResponse{Code: http.StatusBadRequest, Err: errors.New("id does not correspond to a patient or a doctor")}
}

func ModifyStatusAccount(id string, status bool) ModifyAccountStatusResponse {
	_, err := graphql.GetPatientById(id)
	if err == nil {
		patient, err := graphql.UpdatePatient(id, model.UpdatePatientInput{Status: &status})
		if err != nil {
			return ModifyAccountStatusResponse{Patient: model.Patient{}, Code: http.StatusInternalServerError, Err: err}
		}
		return ModifyAccountStatusResponse{Patient: patient, Code: http.StatusOK, Err: nil}
	}

	_, err = graphql.GetDoctorById(id)
	if err == nil {
		doctor, err := graphql.UpdateDoctor(id, model.UpdateDoctorInput{Status: &status})
		if err != nil {
			return ModifyAccountStatusResponse{Doctor: model.Doctor{}, Code: http.StatusInternalServerError, Err: err}
		}

		return ModifyAccountStatusResponse{Doctor: doctor, Code: http.StatusOK, Err: nil}
	}
	return ModifyAccountStatusResponse{Code: http.StatusBadRequest, Err: errors.New("id does not correspond to a patient or a doctor")}
}
