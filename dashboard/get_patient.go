package dashboard

import (
	"context"
	"errors"

	"github.com/edgar-care/edgarlib/graphql"
	"github.com/edgar-care/edgarlib/graphql/server/model"
)

type GetPatientByIdResponse struct {
	Patient     model.Patient
	MedicalInfo model.MedicalInfo
	Code        int
	Err         error
}

type GetPatientsResponse struct {
	Patients []model.Patient
}

func GetPatientById(id string) GetPatientByIdResponse {
	gqlClient := graphql.CreateClient()
	patient, err := graphql.GetPatientById(context.Background(), gqlClient, id)
	if err != nil {
		return GetPatientByIdResponse{Code: 400, Err: errors.New("id does not correspond to a patient")}
	}

	if patient.GetPatientById.Medical_info_id == "" {
		return GetPatientByIdResponse{Code: 400, Err: errors.New("onboarding not started")}
	}

	medical_info, err := graphql.GetMedicalFolderByID(context.Background(), gqlClient, patient.GetPatientById.Medical_info_id)
	if err != nil {
		return GetPatientByIdResponse{Code: 400, Err: errors.New("id does not correspond to an medical table")}
	}

	return GetPatientByIdResponse{
		Patient: model.Patient{
			ID:            patient.GetPatientById.Id,
			Email:         patient.GetPatientById.Email,
			Password:      patient.GetPatientById.Password,
			MedicalInfoID: &patient.GetPatientById.Medical_info_id,
			RendezVousIds: graphql.ConvertStringSliceToPointerSlice(patient.GetPatientById.Rendez_vous_ids),
			DocumentIds:   graphql.ConvertStringSliceToPointerSlice(patient.GetPatientById.Document_ids),
		},
		MedicalInfo: model.MedicalInfo{
			ID:               medical_info.GetMedicalFolderById.Id,
			Name:             medical_info.GetMedicalFolderById.Name,
			Firstname:        medical_info.GetMedicalFolderById.Firstname,
			Birthdate:        medical_info.GetMedicalFolderById.Birthdate,
			Sex:              model.Sex(medical_info.GetMedicalFolderById.Sex),
			Weight:           medical_info.GetMedicalFolderById.Weight,
			Height:           medical_info.GetMedicalFolderById.Height,
			PrimaryDoctorID:  medical_info.GetMedicalFolderById.Primary_doctor_id,
			OnboardingStatus: model.OnboardingStatus(medical_info.GetMedicalFolderById.Onboarding_status),
		},
		Code: 200,
		Err:  nil,
	}
}

func GetPatients(doctorId string) {
	gqlClient := graphql.CreateClient()
	doctor, err := graphql.GetDoctorById(context.Background(), gqlClient, doctorId)
	if err != nil {
		return GetPatientByIdResponse{Code: 400, Err: errors.New("id does not correspond to a doctor")}
	}

	patient_doct, err := graphql.GetPatientById(context.Background(), gqlClient, doctor.GetDoctorById.Patient_ids[])
	if err != nil {
		return GetPatientByIdResponse{Code: 400, Err: errors.New("id does not correspond to a patient")}
	}

	return GetPatientByIdResponse{
		Patient: model.Patient{
			ID:            patient.GetPatientById.Id,
			Email:         patient.GetPatientById.Email,
			Password:      patient.GetPatientById.Password,
			MedicalInfoID: &patient.GetPatientById.Medical_info_id,
			RendezVousIds: graphql.ConvertStringSliceToPointerSlice(patient.GetPatientById.Rendez_vous_ids),
			DocumentIds:   graphql.ConvertStringSliceToPointerSlice(patient.GetPatientById.Document_ids),
		},
		MedicalInfo: model.MedicalInfo{
			ID:               medical_info.GetMedicalFolderById.Id,
			Name:             medical_info.GetMedicalFolderById.Name,
			Firstname:        medical_info.GetMedicalFolderById.Firstname,
			Birthdate:        medical_info.GetMedicalFolderById.Birthdate,
			Sex:              model.Sex(medical_info.GetMedicalFolderById.Sex),
			Weight:           medical_info.GetMedicalFolderById.Weight,
			Height:           medical_info.GetMedicalFolderById.Height,
			PrimaryDoctorID:  medical_info.GetMedicalFolderById.Primary_doctor_id,
			OnboardingStatus: model.OnboardingStatus(medical_info.GetMedicalFolderById.Onboarding_status),
		},
		Code: 200,
		Err:  nil,
	}
}
