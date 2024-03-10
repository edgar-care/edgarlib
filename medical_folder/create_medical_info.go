package medical_folder

import (
	"context"
	"errors"
	"github.com/edgar-care/edgarlib/graphql"
	"github.com/edgar-care/edgarlib/graphql/server/model"
)

type CreateMedicalInfoInput struct {
	Name               string                    `json:"name"`
	Firstname          string                    `json:"firstname"`
	Birthdate          int                       `json:"birthdate"`
	Sex                string                    `json:"sex"`
	Weight             int                       `json:"weight"`
	Height             int                       `json:"height"`
	PrimaryDoctorID    string                    `json:"primary_doctor_id"`
	MedicalAntecedents []CreateMedicalAntecedent `json:"medical_antecedents"`
	OnboardingStatus   string                    `json:"onboarding_status"`
}

type CreateMedicalAntecedent struct {
	Name          string           `json:"name"`
	Medicines     []CreateMedicine `json:"medicines"`
	StillRelevant bool             `json:"still_relevant"`
}

type CreateMedicine struct {
	Period   []string `json:"period"`
	Day      []string `json:"day"`
	Quantity int      `json:"quantity"`
}

type CreateMedicalInfoResponse struct {
	MedicalInfo model.MedicalInfo
	Patient     model.Patient
	Code        int
	Err         error
}

func CreateMedicalInfo(input CreateMedicalInfoInput, patientID string) CreateMedicalInfoResponse {
	gqlClient := graphql.CreateClient()

	medicalAntecedents := make([]graphql.MedicalAntecedentsInput, len(input.MedicalAntecedents))
	for i, antecedent := range input.MedicalAntecedents {
		medicines := make([]graphql.MedicinesInput, len(antecedent.Medicines))
		for j, med := range antecedent.Medicines {
			periods := make([]graphql.Period, len(med.Period))
			for k, p := range med.Period {
				periods[k] = graphql.Period(p)
			}
			days := make([]graphql.Day, len(med.Day))
			for k, d := range med.Day {
				days[k] = graphql.Day(d)
			}
			medicines[j] = graphql.MedicinesInput{
				Period:   periods,
				Day:      days,
				Quantity: med.Quantity,
			}
		}
		medicalAntecedents[i] = graphql.MedicalAntecedentsInput{
			Name:           antecedent.Name,
			Medicines:      medicines,
			Still_relevant: antecedent.StillRelevant,
		}
	}
	control, err := graphql.GetPatientById(context.Background(), gqlClient, patientID)
	if err != nil {
		return CreateMedicalInfoResponse{Code: 400, Err: errors.New("unable to find patient by ID: " + err.Error())}
	}

	if control.GetPatientById.Medical_info_id != "" {
		return CreateMedicalInfoResponse{Code: 400, Err: errors.New("medical folder has already been create")}
	}

	medical, err := graphql.CreateMedicalFolder(context.Background(), gqlClient, input.Name, input.Firstname, input.Birthdate, input.Sex, input.Weight, input.Height, input.PrimaryDoctorID, medicalAntecedents, input.OnboardingStatus)
	if err != nil {
		return CreateMedicalInfoResponse{Code: 400, Err: errors.New("unable to create medical folder: " + err.Error())}
	}

	patient, err := graphql.GetPatientById(context.Background(), gqlClient, patientID)
	if err != nil {
		return CreateMedicalInfoResponse{Code: 400, Err: errors.New("unable to find patient by ID: " + err.Error())}
	}

	_, err = graphql.UpdatePatient(context.Background(), gqlClient, patientID, patient.GetPatientById.Email, patient.GetPatientById.Password, medical.CreateMedicalFolder.Id, patient.GetPatientById.Rendez_vous_ids, patient.GetPatientById.Document_ids)
	if err != nil {
		return CreateMedicalInfoResponse{Code: 400, Err: errors.New("unable to update patient: " + err.Error())}
	}

	medicalAntecedentsResp := make([]*model.MedicalAntecedents, len(medical.CreateMedicalFolder.Medical_antecedents))
	for i, antecedent := range medical.CreateMedicalFolder.Medical_antecedents {
		medicines := make([]*model.Medicines, len(antecedent.Medicines))
		for j, med := range antecedent.Medicines {
			periods := make([]*model.Period, len(med.Period))
			for k, p := range med.Period {
				period := model.Period(p)
				periods[k] = &period
			}
			days := make([]*model.Day, len(med.Day))
			for k, d := range med.Day {
				day := model.Day(d)
				days[k] = &day
			}
			medicines[j] = &model.Medicines{
				Period:   periods,
				Day:      days,
				Quantity: med.Quantity,
			}
		}
		medicalAntecedentsResp[i] = &model.MedicalAntecedents{
			Name:          antecedent.Name,
			Medicines:     medicines,
			StillRelevant: antecedent.Still_relevant,
		}
	}

	return CreateMedicalInfoResponse{
		MedicalInfo: model.MedicalInfo{
			ID:                 medical.CreateMedicalFolder.Id,
			Name:               medical.CreateMedicalFolder.Name,
			Firstname:          medical.CreateMedicalFolder.Firstname,
			Birthdate:          medical.CreateMedicalFolder.Birthdate,
			Sex:                model.Sex(medical.CreateMedicalFolder.Sex),
			Weight:             medical.CreateMedicalFolder.Weight,
			Height:             medical.CreateMedicalFolder.Height,
			PrimaryDoctorID:    medical.CreateMedicalFolder.Primary_doctor_id,
			MedicalAntecedents: medicalAntecedentsResp,
			OnboardingStatus:   model.OnboardingStatus(medical.CreateMedicalFolder.Onboarding_status),
		},
		Code: 201,
		Err:  nil,
	}
}
