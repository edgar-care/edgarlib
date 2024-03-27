package dashboard

import (
	"context"
	"errors"
	"github.com/edgar-care/edgarlib/medical_folder"

	"github.com/edgar-care/edgarlib/graphql"
	"github.com/edgar-care/edgarlib/graphql/server/model"
)

type GetPatientByIdResponse struct {
	PatientInfo PatientWithMedicalInfo `json:"patient_info"`
	Code        int
	Err         error
}

type PatientWithMedicalInfo struct {
	ID                string            `json:"id"`
	Email             string            `json:"email"`
	MedicalInfo       model.MedicalInfo `json:"medical_info"`
	RendezVousIds     []string          `json:"rendez_vous_ids"`
	DocumentsIds      []string          `json:"documents_ids"`
	TreatmentFollowUp []string          `json:"treatment_follow_up"`
}

type GetPatientsResponse struct {
	PatientsInfo []PatientWithMedicalInfo `json:"patients_info"`
	Code         int
	Err          error
}

func GetPatientById(id string, doctorid string) GetPatientByIdResponse {
	gqlClient := graphql.CreateClient()

	doctor, err := graphql.GetDoctorById(context.Background(), gqlClient, doctorid)
	if err != nil {
		return GetPatientByIdResponse{Code: 400, Err: errors.New("id does not correspond to a doctor")}
	}
	idFound := false
	for _, item := range doctor.GetDoctorById.Patient_ids {
		if item == id {
			idFound = true
			break
		}
	}
	if !idFound {
		return GetPatientByIdResponse{Code: 401, Err: errors.New("unauthorized to access to this account")}
	}

	patient, err := graphql.GetPatientById(context.Background(), gqlClient, id)
	if err != nil {
		return GetPatientByIdResponse{Code: 404, Err: errors.New("id does not correspond to a patient")}
	}

	var patients PatientWithMedicalInfo
	medicalInfo := medical_folder.GetMedicalInfo(id)
	if medicalInfo.Err != nil {
		return GetPatientByIdResponse{Code: 404, Err: errors.New("error while retrieving medical info by id")}
	}

	patients = PatientWithMedicalInfo{
		ID:                patient.GetPatientById.Id,
		Email:             patient.GetPatientById.Email,
		RendezVousIds:     patient.GetPatientById.Rendez_vous_ids,
		DocumentsIds:      patient.GetPatientById.Document_ids,
		MedicalInfo:       medicalInfo.MedicalInfo,
		TreatmentFollowUp: patient.GetPatientById.Treatment_follow_up_ids,
	}

	return GetPatientByIdResponse{
		PatientInfo: patients,
		Code:        200,
		Err:         nil,
	}
}

func GetPatients(doctorId string) GetPatientsResponse {
	gqlClient := graphql.CreateClient()

	patientDoctor, err := graphql.GetPatientsFromDoctorById(context.Background(), gqlClient, doctorId)
	if err != nil {
		return GetPatientsResponse{Code: 400, Err: errors.New("id does not correspond to a doctor")}
	}

	var patients []PatientWithMedicalInfo
	for _, patient := range patientDoctor.GetPatientsFromDoctorById {
		patientId := patient.Medical_info_id
		medicalInfo := medical_folder.GetMedicalInfo(patientId)
		if medicalInfo.Err != nil {
			return GetPatientsResponse{Code: 401, Err: errors.New("error while retrieving medical info by id")}
		}
		patients = append(patients, PatientWithMedicalInfo{
			ID:                patient.Id,
			Email:             patient.Email,
			RendezVousIds:     patient.Rendez_vous_ids,
			DocumentsIds:      patient.Document_ids,
			MedicalInfo:       medicalInfo.MedicalInfo,
			TreatmentFollowUp: patient.Treatment_follow_up_ids,
		})
	}
	return GetPatientsResponse{
		PatientsInfo: patients,
		Code:         200,
		Err:          nil,
	}
}
