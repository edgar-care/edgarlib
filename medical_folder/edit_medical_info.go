package medical_folder

import (
	"context"
	"errors"

	"github.com/edgar-care/edgarlib/graphql"
	"github.com/edgar-care/edgarlib/graphql/server/model"
)

type UpdateMedicalFolderResponse struct {
	MedicalInfo model.MedicalInfo
	Code        int
	Err         error
}

func UpdateMedicalFolder(newMedicalInfo CreateMedicalInfoInput, medicalInfoID string) UpdateMedicalFolderResponse {
	gqlClient := graphql.CreateClient()
	if medicalInfoID == "" {
		return UpdateMedicalFolderResponse{MedicalInfo: model.MedicalInfo{}, Code: 400, Err: errors.New("medical info ID is required")}
	}

	medicalAntecedents := make([]graphql.MedicalAntecedentsInput, len(newMedicalInfo.MedicalAntecedents))
	for i, antecedent := range newMedicalInfo.MedicalAntecedents {
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

	updatedMedicalFolder, err := graphql.UpdateMedicalFolder(context.Background(), gqlClient, medicalInfoID, newMedicalInfo.Name, newMedicalInfo.Firstname, newMedicalInfo.Birthdate, newMedicalInfo.Sex, newMedicalInfo.Weight, newMedicalInfo.Height, newMedicalInfo.PrimaryDoctorID, medicalAntecedents, newMedicalInfo.OnboardingStatus)
	if err != nil {
		return UpdateMedicalFolderResponse{MedicalInfo: model.MedicalInfo{}, Code: 500, Err: errors.New("unable to update medical folder")}
	}

	medicalAntecedentsResp := make([]*model.MedicalAntecedents, len(updatedMedicalFolder.UpdateMedicalFolder.Medical_antecedents))
	for i, antecedent := range updatedMedicalFolder.UpdateMedicalFolder.Medical_antecedents {
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

	return UpdateMedicalFolderResponse{
		MedicalInfo: model.MedicalInfo{
			ID:                 updatedMedicalFolder.UpdateMedicalFolder.Id,
			Name:               updatedMedicalFolder.UpdateMedicalFolder.Name,
			Firstname:          updatedMedicalFolder.UpdateMedicalFolder.Firstname,
			Birthdate:          updatedMedicalFolder.UpdateMedicalFolder.Birthdate,
			Sex:                model.Sex(updatedMedicalFolder.UpdateMedicalFolder.Sex),
			Weight:             updatedMedicalFolder.UpdateMedicalFolder.Weight,
			Height:             updatedMedicalFolder.UpdateMedicalFolder.Height,
			PrimaryDoctorID:    updatedMedicalFolder.UpdateMedicalFolder.Primary_doctor_id,
			MedicalAntecedents: medicalAntecedentsResp,
			OnboardingStatus:   model.OnboardingStatus(updatedMedicalFolder.UpdateMedicalFolder.Onboarding_status),
		},
		Code: 200,
		Err:  nil,
	}
}

func UpdateMedicalFolderFromDoctor(newMedicalInfo CreateMedicalInfoInput, PatientID string) UpdateMedicalFolderResponse {
	gqlClient := graphql.CreateClient()
	if PatientID == "" {
		return UpdateMedicalFolderResponse{MedicalInfo: model.MedicalInfo{}, Code: 400, Err: errors.New("medical info ID is required")}
	}
	patient, err := graphql.GetPatientById(context.Background(), gqlClient, PatientID)
	if err != nil {
		return UpdateMedicalFolderResponse{MedicalInfo: model.MedicalInfo{}, Code: 400, Err: errors.New("id does not correspond to a patient")}
	}

	medicalAntecedents := make([]graphql.MedicalAntecedentsInput, len(newMedicalInfo.MedicalAntecedents))
	for i, antecedent := range newMedicalInfo.MedicalAntecedents {
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

	updatedMedicalFolder, err := graphql.UpdateMedicalFolder(context.Background(), gqlClient, patient.GetPatientById.Medical_info_id, newMedicalInfo.Name, newMedicalInfo.Firstname, newMedicalInfo.Birthdate, newMedicalInfo.Sex, newMedicalInfo.Weight, newMedicalInfo.Height, newMedicalInfo.PrimaryDoctorID, medicalAntecedents, newMedicalInfo.OnboardingStatus)
	if err != nil {
		return UpdateMedicalFolderResponse{MedicalInfo: model.MedicalInfo{}, Code: 500, Err: errors.New("unable to update medical folder")}
	}

	medicalAntecedentsResp := make([]*model.MedicalAntecedents, len(updatedMedicalFolder.UpdateMedicalFolder.Medical_antecedents))
	for i, antecedent := range updatedMedicalFolder.UpdateMedicalFolder.Medical_antecedents {
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

	return UpdateMedicalFolderResponse{
		MedicalInfo: model.MedicalInfo{
			ID:                 updatedMedicalFolder.UpdateMedicalFolder.Id,
			Name:               updatedMedicalFolder.UpdateMedicalFolder.Name,
			Firstname:          updatedMedicalFolder.UpdateMedicalFolder.Firstname,
			Birthdate:          updatedMedicalFolder.UpdateMedicalFolder.Birthdate,
			Sex:                model.Sex(updatedMedicalFolder.UpdateMedicalFolder.Sex),
			Weight:             updatedMedicalFolder.UpdateMedicalFolder.Weight,
			Height:             updatedMedicalFolder.UpdateMedicalFolder.Height,
			PrimaryDoctorID:    updatedMedicalFolder.UpdateMedicalFolder.Primary_doctor_id,
			MedicalAntecedents: medicalAntecedentsResp,
			OnboardingStatus:   model.OnboardingStatus(updatedMedicalFolder.UpdateMedicalFolder.Onboarding_status),
		},
		Code: 200,
		Err:  nil,
	}
}
