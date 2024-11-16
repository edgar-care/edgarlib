package medical_folder

import (
	"errors"
	"fmt"
	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
)

type CreateNewMedicalInfoInput struct {
	Name                   string                            `json:"name"`
	Firstname              string                            `json:"firstname"`
	Birthdate              int                               `json:"birthdate"`
	Sex                    string                            `json:"sex"`
	Weight                 int                               `json:"weight"`
	Height                 int                               `json:"height"`
	PrimaryDoctorID        string                            `json:"primary_doctor_id,omitempty"`
	MedicalAntecedents     []CreateNewMedicalAntecedentInput `json:"medical_antecedents"`
	FamilyMembersMedInfoId []string                          `json:"family_members_med_info_id"`
}

type CreateNewMedicalAntecedentInput struct {
	Name       string             `json:"name"`
	Symptoms   []string           `json:"symptoms"`
	Treatments []CreateTreatInput `json:"treatments"`
}

type CreateTreatInput struct {
	StartDate int                          `json:"start_date"`
	EndDate   int                          `json:"end_date"`
	Medicines []CreateAntecedentsMedicines `json:"medicines"`
}

type CreateAntecedentsMedicines struct {
	MedicineID string                    `json:"medicine_id"`
	Comment    string                    `json:"comment"`
	Period     []*CreateAntecedentPeriod `json:"period"`
}

type CreateAntecedentPeriod struct {
	Quantity       int    `json:"quantity"`
	Frequency      int    `json:"frequency"`
	FrequencyRatio int    `json:"frequency_ratio"`
	FrequencyUnit  string `json:"frequency_unit"`
	PeriodLength   int    `json:"period_length"`
	PeriodUnit     string `json:"period_unit"`
}

type CreateNewMedicalInfoResponse struct {
	MedicalInfo        model.MedicalInfo
	MedicalAntecedents []model.MedicalAntecedents
	Code               int
	Err                error
}

type AddMedicalAntecedentResponse struct {
	MedicalInfo        model.MedicalInfo
	MedicalAntecedents []model.MedicalAntecedents
	Code               int
	Err                error
}

func NewMedicalFolder(input CreateNewMedicalInfoInput, patientID string) CreateNewMedicalInfoResponse {
	control, err := graphql.GetPatientById(patientID)
	if err != nil {
		return CreateNewMedicalInfoResponse{Code: 400, Err: fmt.Errorf("unable to find patient by ID: %v", err)}
	}

	if control.MedicalInfoID != nil && *control.MedicalInfoID != "" {
		return CreateNewMedicalInfoResponse{Code: 400, Err: errors.New("medical folder has already been created")}
	}

	if len(input.MedicalAntecedents) == 0 {
		medical, err := graphql.CreateMedicalFolder(model.CreateMedicalFolderInput{
			Name:                   input.Name,
			Firstname:              input.Firstname,
			Birthdate:              input.Birthdate,
			Sex:                    input.Sex,
			Height:                 input.Height,
			Weight:                 input.Weight,
			PrimaryDoctorID:        input.PrimaryDoctorID,
			OnboardingStatus:       "DONE",
			FamilyMembersMedInfoID: input.FamilyMembersMedInfoId,
		})
		if err != nil {
			return CreateNewMedicalInfoResponse{Code: 400, Err: fmt.Errorf("unable to create medical folder: %v", err)}
		}

		_, err = graphql.UpdatePatient(patientID, model.UpdatePatientInput{
			MedicalInfoID: &medical.ID,
		})
		if err != nil {
			return CreateNewMedicalInfoResponse{Code: 400, Err: fmt.Errorf("unable to update patient: %v", err)}
		}

		return CreateNewMedicalInfoResponse{
			MedicalInfo: medical,
			Code:        201,
			Err:         nil,
		}
	}

	var createdAntecedents []model.MedicalAntecedents
	for _, medicalAntecedent := range input.MedicalAntecedents {
		treatments := make([]*model.CreateAntecedentTreatmentInput, 0)
		for _, treatment := range medicalAntecedent.Treatments {
			medicines := make([]*model.CreateAntecedentsMedicinesInput, 0)
			for _, medicine := range treatment.Medicines {
				periods := make([]*model.CreateAntecedentPeriodInput, 0)
				for _, period := range medicine.Period {
					periods = append(periods, &model.CreateAntecedentPeriodInput{
						Quantity:       period.Quantity,
						Frequency:      period.Frequency,
						FrequencyRatio: period.FrequencyRatio,
						FrequencyUnit:  model.TimeUnitEnum(period.FrequencyUnit),
						PeriodLength:   &period.PeriodLength,
						PeriodUnit:     (*model.TimeUnitEnum)(&period.PeriodUnit),
					})
				}
				medicines = append(medicines, &model.CreateAntecedentsMedicinesInput{
					MedicineID: medicine.MedicineID,
					Comment:    &medicine.Comment,
					Period:     periods,
				})
			}
			treatments = append(treatments, &model.CreateAntecedentTreatmentInput{
				CreatedBy: patientID,
				StartDate: treatment.StartDate,
				EndDate:   &treatment.EndDate,
				Medicines: medicines,
			})
		}

		createdAntecedent, err := graphql.CreateMedicalAntecedents(model.CreateMedicalAntecedentsInput{
			Name:       medicalAntecedent.Name,
			Symptoms:   medicalAntecedent.Symptoms,
			Treatments: treatments,
		})
		if err != nil {
			return CreateNewMedicalInfoResponse{Code: 400, Err: fmt.Errorf("unable to create medical antecedent: %v", err)}
		}
		createdAntecedents = append(createdAntecedents, createdAntecedent)
	}

	medical, err := graphql.CreateMedicalFolder(model.CreateMedicalFolderInput{
		Name:                   input.Name,
		Firstname:              input.Firstname,
		Birthdate:              input.Birthdate,
		Sex:                    input.Sex,
		Height:                 input.Height,
		Weight:                 input.Weight,
		PrimaryDoctorID:        input.PrimaryDoctorID,
		OnboardingStatus:       "DONE",
		FamilyMembersMedInfoID: input.FamilyMembersMedInfoId,
		AntecedentDiseaseIds:   convertToIDSlice(createdAntecedents),
	})
	if err != nil {
		return CreateNewMedicalInfoResponse{Code: 400, Err: fmt.Errorf("unable to create medical folder with antecedents: %v", err)}
	}

	_, err = graphql.UpdatePatient(patientID, model.UpdatePatientInput{
		MedicalInfoID: &medical.ID,
	})
	if err != nil {
		return CreateNewMedicalInfoResponse{Code: 400, Err: fmt.Errorf("unable to update patient: %v", err)}
	}

	return CreateNewMedicalInfoResponse{
		MedicalInfo:        medical,
		MedicalAntecedents: createdAntecedents,
		Code:               201,
		Err:                nil,
	}
}

func convertToIDSlice(antecedents []model.MedicalAntecedents) []string {
	var ids []string
	for _, antecedent := range antecedents {
		ids = append(ids, antecedent.ID)
	}
	return ids
}

func AddMedicalAntecedent(input CreateNewMedicalAntecedentInput, userID string) AddMedicalAntecedentResponse {
	control, err := graphql.GetPatientById(userID)
	if err != nil {
		return AddMedicalAntecedentResponse{Code: 400, Err: fmt.Errorf("unable to find patient by ID: %v", err)}
	}

	if control.MedicalInfoID == nil || *control.MedicalInfoID == "" {
		return AddMedicalAntecedentResponse{Code: 400, Err: errors.New("medical folder has not been created")}
	}

	checkMedicalFolder, err := graphql.GetMedicalFolderByID(*control.MedicalInfoID)
	if err != nil {
		return AddMedicalAntecedentResponse{Code: 400, Err: fmt.Errorf("unable to retrieve medical folder: %v", err)}
	}

	treatments := make([]*model.CreateAntecedentTreatmentInput, 0)
	for _, treatment := range input.Treatments {
		medicines := make([]*model.CreateAntecedentsMedicinesInput, 0)
		for _, medicine := range treatment.Medicines {
			if len(medicine.Period) == 0 {
				return AddMedicalAntecedentResponse{Code: 400, Err: errors.New("At least one period is required for medicine")}
			}

			periods := make([]*model.CreateAntecedentPeriodInput, 0)
			for _, period := range medicine.Period {
				if period.Quantity <= 0 || period.Frequency <= 0 {
					return AddMedicalAntecedentResponse{Code: 400, Err: errors.New("Quantity, Frequency, and PeriodLength must be greater than 0")}
				}
				periods = append(periods, &model.CreateAntecedentPeriodInput{
					Quantity:       period.Quantity,
					Frequency:      period.Frequency,
					FrequencyRatio: period.FrequencyRatio,
					FrequencyUnit:  model.TimeUnitEnum(period.FrequencyUnit),
					PeriodLength:   &period.PeriodLength,
					PeriodUnit:     (*model.TimeUnitEnum)(&period.PeriodUnit),
				})
			}
			medicines = append(medicines, &model.CreateAntecedentsMedicinesInput{
				MedicineID: medicine.MedicineID,
				Comment:    &medicine.Comment,
				Period:     periods,
			})
		}
		treatments = append(treatments, &model.CreateAntecedentTreatmentInput{
			CreatedBy: userID,
			StartDate: treatment.StartDate,
			EndDate:   &treatment.EndDate,
			Medicines: medicines,
		})
	}

	createMedicalAntecedent, err := graphql.CreateMedicalAntecedents(model.CreateMedicalAntecedentsInput{
		Name:       input.Name,
		Symptoms:   input.Symptoms,
		Treatments: treatments,
	})

	if err != nil {
		return AddMedicalAntecedentResponse{Code: 400, Err: fmt.Errorf("unable to create medical antecedent: %v", err)}
	}

	updatedMedicalFolder, err := graphql.UpdateMedicalFolder(*control.MedicalInfoID, model.UpdateMedicalFolderInput{
		AntecedentDiseaseIds: append(checkMedicalFolder.AntecedentDiseaseIds, createMedicalAntecedent.ID),
	})
	if err != nil {
		return AddMedicalAntecedentResponse{Code: 400, Err: fmt.Errorf("unable to update medical folder: %v", err)}
	}

	return AddMedicalAntecedentResponse{
		MedicalInfo:        updatedMedicalFolder,
		MedicalAntecedents: []model.MedicalAntecedents{createMedicalAntecedent},
		Code:               201,
	}
}
