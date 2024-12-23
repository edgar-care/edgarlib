package treatment

import (
	"errors"
	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
)

type CreateTreatInput struct {
	MedicalantecedentID string                       `json:"medical_antecedent_id"`
	StartDate           int                          `json:"start_date"`
	EndDate             int                          `json:"end_date"`
	Medicines           []CreateAntecedentsMedicines `json:"medicines"`
}

type CreateAntecedentsMedicines struct {
	MedicineID string                   `json:"medicine_id"`
	Comment    string                   `json:"comment"`
	Period     []CreateAntecedentPeriod `json:"period"`
}

type CreateAntecedentPeriod struct {
	Quantity       int     `json:"quantity"`
	Frequency      int     `json:"frequency"`
	FrequencyRatio int     `json:"frequency_ratio"`
	FrequencyUnit  string  `json:"frequency_unit"`
	PeriodLength   *int    `json:"period_length"`
	PeriodUnit     *string `json:"period_unit"`
}

type CreateTreatmentResponse struct {
	Treatment []model.AntecedentTreatment
	Code      int
	Err       error
}

func ConvertPeriods(periods []CreateAntecedentPeriod) []model.AntecedentPeriod {
	convertedPeriods := make([]model.AntecedentPeriod, len(periods))
	for i, p := range periods {
		freqUnit := model.TimeUnitEnum(p.FrequencyUnit)
		var periodUnit *model.TimeUnitEnum
		if p.PeriodUnit != nil {
			unit := model.TimeUnitEnum(*p.PeriodUnit)
			periodUnit = &unit
		}
		convertedPeriods[i] = model.AntecedentPeriod{
			Quantity:       p.Quantity,
			Frequency:      p.Frequency,
			FrequencyRatio: p.FrequencyRatio,
			FrequencyUnit:  freqUnit,
			PeriodLength:   p.PeriodLength,
			PeriodUnit:     periodUnit,
		}
	}
	return convertedPeriods
}

func convertToCreateAntecedentPeriodInputSlice(periods []model.AntecedentPeriod) []*model.CreateAntecedentPeriodInput {
	pointerSlice := make([]*model.CreateAntecedentPeriodInput, len(periods))
	for i, p := range periods {
		pointerSlice[i] = &model.CreateAntecedentPeriodInput{
			Quantity:       p.Quantity,
			Frequency:      p.Frequency,
			FrequencyRatio: p.FrequencyRatio,
			FrequencyUnit:  p.FrequencyUnit,
			PeriodLength:   p.PeriodLength,
			PeriodUnit:     p.PeriodUnit,
		}
	}
	return pointerSlice
}

func CreateTreatment(input CreateTreatInput, patientID string) CreateTreatmentResponse {
	var res []model.AntecedentTreatment

	patient, err := graphql.GetPatientById(patientID)
	if err != nil {
		return CreateTreatmentResponse{Code: 400, Err: errors.New("unable to find patient by ID: " + err.Error())}
	}
	if patient.MedicalInfoID == nil || *patient.MedicalInfoID == "" {
		return CreateTreatmentResponse{Code: 400, Err: errors.New("medical folder has not been created")}
	}

	medicalFolder, err := graphql.GetMedicalFolderByID(*patient.MedicalInfoID)
	if err != nil {
		return CreateTreatmentResponse{Code: 400, Err: errors.New("unable to get medical folder: " + err.Error())}
	}

	found := false
	for _, id := range medicalFolder.AntecedentDiseaseIds {
		if id == input.MedicalantecedentID {
			found = true
			break
		}
	}
	if !found {
		return CreateTreatmentResponse{Code: 404, Err: errors.New("antecedent ID not found in medical folder")}
	}

	var medicines []*model.CreateAntecedentsMedicinesInput
	for _, medicine := range input.Medicines {
		comment := medicine.Comment
		periods := ConvertPeriods(medicine.Period)
		medicines = append(medicines, &model.CreateAntecedentsMedicinesInput{
			MedicineID: medicine.MedicineID,
			Comment:    &comment,
			Period:     convertToCreateAntecedentPeriodInputSlice(periods),
		})
	}

	treatment, err := graphql.CreateAntecdentTreatment(input.MedicalantecedentID, model.CreateAntecedentTreatmentInput{
		CreatedBy: patientID,
		StartDate: input.StartDate,
		EndDate:   &input.EndDate,
		Medicines: medicines,
	})
	if err != nil {
		return CreateTreatmentResponse{Code: 400, Err: errors.New("unable to create antecedent treatment: " + err.Error())}
	}
	res = append(res, treatment)

	return CreateTreatmentResponse{
		Treatment: res,
		Code:      201,
		Err:       nil,
	}
}
