package dashboard

import (
	"errors"
	"github.com/edgar-care/edgarlib/v2/medical_folder"

	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
)

type GetPatientByIdResponse struct {
	PatientInfo PatientWithMedicalInfo `json:"patient_info"`
	Code        int
	Err         error
}

type PatientWithMedicalInfo struct {
	ID                string                     `json:"id"`
	Email             string                     `json:"email"`
	MedicalInfo       model.MedicalInfo          `json:"medical_info"`
	Antedisease       []model.MedicalAntecedents `json:"antedisease"`
	RendezVousIds     []*string                  `json:"rendez_vous_ids"`
	DocumentsIds      []*string                  `json:"documents_ids"`
	TreatmentFollowUp []*string                  `json:"treatment_follow_up"`
}

type GetPatientsResponse struct {
	PatientsInfo []PatientWithMedicalInfo `json:"patients_info"`
	Code         int
	Err          error
}

func GetPatientById(id string, doctorid string) GetPatientByIdResponse {
	doctor, err := graphql.GetDoctorById(doctorid)
	if err != nil {
		return GetPatientByIdResponse{Code: 400, Err: errors.New("id does not correspond to a doctor")}
	}
	idFound := false
	for _, item := range doctor.PatientIds {
		if *item == id {
			idFound = true
			break
		}
	}
	if !idFound {
		return GetPatientByIdResponse{Code: 403, Err: errors.New("unauthorized to access to this account")}
	}

	patient, err := graphql.GetPatientById(id)
	if err != nil {
		return GetPatientByIdResponse{Code: 404, Err: errors.New("id does not correspond to a patient")}
	}

	var patients PatientWithMedicalInfo
	medicalInfo := medical_folder.GetMedicalFolder(id)
	if medicalInfo.Err != nil {
		return GetPatientByIdResponse{Code: 404, Err: errors.New("error while retrieving medical info by id")}
	}

	patients = PatientWithMedicalInfo{
		ID:                patient.ID,
		Email:             patient.Email,
		RendezVousIds:     patient.RendezVousIds,
		DocumentsIds:      patient.DocumentIds,
		MedicalInfo:       medicalInfo.MedicalInfo,
		TreatmentFollowUp: patient.TreatmentFollowUpIds,
		Antedisease:       medicalInfo.MedicalAntecedents,
	}

	return GetPatientByIdResponse{
		PatientInfo: patients,
		Code:        200,
		Err:         nil,
	}
}

func GetPatients(doctorId string) GetPatientsResponse {
	patientDoctor, err := graphql.GetPatientsFromDoctorById(doctorId, nil)
	if err != nil {
		return GetPatientsResponse{Code: 400, Err: errors.New("id does not correspond to a doctor")}
	}

	var patients []PatientWithMedicalInfo
	for _, patient := range patientDoctor {
		medicalInfo := medical_folder.GetMedicalFolder(patient.ID)
		if medicalInfo.Err != nil {
			return GetPatientsResponse{Code: 400, Err: errors.New("error while retrieving medical info by id")}
		}
		patients = append(patients, PatientWithMedicalInfo{
			ID:                patient.ID,
			Email:             patient.Email,
			RendezVousIds:     patient.RendezVousIds,
			DocumentsIds:      patient.DocumentIds,
			MedicalInfo:       medicalInfo.MedicalInfo,
			TreatmentFollowUp: patient.TreatmentFollowUpIds,
			Antedisease:       medicalInfo.MedicalAntecedents,
		})
	}
	return GetPatientsResponse{
		PatientsInfo: patients,
		Code:         200,
		Err:          nil,
	}
}
