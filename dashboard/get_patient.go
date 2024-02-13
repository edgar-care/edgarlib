package dashboard

import (
	"context"
	"errors"
	"github.com/edgar-care/edgarlib/graphql"
	"github.com/edgar-care/edgarlib/graphql/server/model"
)

type GetPatientByIdResponse struct {
	Patient model.Patient
	Info    model.Info
	Health  model.Health
	Code    int
	Err     error
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

	if patient.GetPatientById.Onboarding_info_id == "" || patient.GetPatientById.Onboarding_health_id == "" {
		return GetPatientByIdResponse{Code: 400, Err: errors.New("onboarding not started")}
	}

	info, err := graphql.GetInfoById(context.Background(), gqlClient, patient.GetPatientById.Onboarding_info_id)
	if err != nil {
		return GetPatientByIdResponse{Code: 400, Err: errors.New("id does not correspond to an info table")}
	}

	health, err := graphql.GetHealthById(context.Background(), gqlClient, patient.GetPatientById.Onboarding_health_id)
	if err != nil {
		return GetPatientByIdResponse{Code: 400, Err: errors.New("id does not correspond to a health table")}
	}

	return GetPatientByIdResponse{
		Patient: model.Patient{
			ID:                 patient.GetPatientById.Id,
			Email:              patient.GetPatientById.Email,
			Password:           patient.GetPatientById.Password,
			OnboardingInfoID:   &patient.GetPatientById.Onboarding_info_id,
			OnboardingHealthID: &patient.GetPatientById.Onboarding_health_id,
			RendezVousIds:      graphql.ConvertStringSliceToPointerSlice(patient.GetPatientById.Rendez_vous_ids),
			DocumentIds:        graphql.ConvertStringSliceToPointerSlice(patient.GetPatientById.Document_ids),
		},
		Info: model.Info{
			ID:        info.GetInfoById.Id,
			Name:      info.GetInfoById.Name,
			Birthdate: info.GetInfoById.Birthdate,
			Height:    info.GetInfoById.Height,
			Weight:    info.GetInfoById.Weight,
			Sex:       model.Sex(info.GetInfoById.Sex),
			Surname:   info.GetInfoById.Surname,
		},
		Health: model.Health{
			ID:                    health.GetHealthById.Id,
			PatientsAllergies:     health.GetHealthById.Patients_allergies,
			PatientsIllness:       health.GetHealthById.Patients_illness,
			PatientsTreatments:    health.GetHealthById.Patients_treatments,
			PatientsPrimaryDoctor: health.GetHealthById.Patients_primary_doctor,
		},
		Code: 200,
		Err:  nil,
	}
}

func GetPatients(doctorId string) {
	//gqlClient := graphql.CreateClient()
	//doctor, err := graphql.GetDoctorById(context.Background(), gqlClient, doctorId)
}
