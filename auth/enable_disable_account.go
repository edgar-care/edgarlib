package auth

import (
	"context"
	"errors"
	"github.com/edgar-care/edgarlib/graphql"
	"github.com/edgar-care/edgarlib/graphql/server/model"
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
	gqlClient := graphql.CreateClient()

	patient, err := graphql.GetPatientById(context.Background(), gqlClient, id)
	if err == nil {
		if patient.GetPatientById.Status == false {
			return CheckAccountStatusResponse{Status: false, Code: 409, Err: errors.New("patient account is disabled")}
		}
		return CheckAccountStatusResponse{Status: true, Code: http.StatusOK, Err: nil}
	}

	doctor, err := graphql.GetDoctorById(context.Background(), gqlClient, id)
	if err == nil {
		if doctor.GetDoctorById.Status == false {
			return CheckAccountStatusResponse{Status: false, Code: 409, Err: errors.New("doctor account is disabled")}
		}
		return CheckAccountStatusResponse{Status: true, Code: http.StatusOK, Err: nil}
	}

	return CheckAccountStatusResponse{Code: http.StatusBadRequest, Err: errors.New("id does not correspond to a patient or a doctor")}
}

func ModifyStatusAccount(id string, status bool) ModifyAccountStatusResponse {
	gqlClient := graphql.CreateClient()

	patient, err := graphql.GetPatientById(context.Background(), gqlClient, id)
	if err == nil {
		_, updateErr := graphql.UpdatePatient(context.Background(), gqlClient, id, patient.GetPatientById.Email, patient.GetPatientById.Password, patient.GetPatientById.Medical_info_id, patient.GetPatientById.Rendez_vous_ids, patient.GetPatientById.Document_ids, patient.GetPatientById.Treatment_follow_up_ids, patient.GetPatientById.Chat_ids, patient.GetPatientById.Device_connect, patient.GetPatientById.Double_auth_methods_id, status)
		if updateErr != nil {
			return ModifyAccountStatusResponse{Patient: model.Patient{}, Code: http.StatusInternalServerError, Err: updateErr}
		}
		patient.GetPatientById.Status = status
		return ModifyAccountStatusResponse{Patient: model.Patient{}, Code: http.StatusOK, Err: nil}
	}

	doctor, err := graphql.GetDoctorById(context.Background(), gqlClient, id)
	if err == nil {
		_, updateErr := graphql.UpdateDoctor(context.Background(), gqlClient, id, doctor.GetDoctorById.Email, doctor.GetDoctorById.Password, doctor.GetDoctorById.Name, doctor.GetDoctorById.Firstname, doctor.GetDoctorById.Rendez_vous_ids, doctor.GetDoctorById.Patient_ids, graphql.AddressInput{Street: doctor.GetDoctorById.Address.Street, Zip_code: doctor.GetDoctorById.Address.Zip_code, Country: doctor.GetDoctorById.Address.Country}, doctor.GetDoctorById.Chat_ids, status)
		if updateErr != nil {
			return ModifyAccountStatusResponse{Doctor: model.Doctor{}, Code: http.StatusInternalServerError, Err: updateErr}
		}
		doctor.GetDoctorById.Status = status
		return ModifyAccountStatusResponse{Doctor: model.Doctor{}, Code: http.StatusOK, Err: nil}
	}
	return ModifyAccountStatusResponse{Code: http.StatusBadRequest, Err: errors.New("id does not correspond to a patient or a doctor")}
}
