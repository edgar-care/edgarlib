package treatment

import (
	"errors"
	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
)

type UpdateTreatmentInput struct {
	ID        string                       `json:"id"`
	CreatedBy string                       `json:"created_by"`
	StartDate int                          `json:"start_date"`
	EndDate   int                          `json:"end_date"`
	Medicines []UpdateAntecedentsMedicines `json:"medicines"`
}

type UpdateAntecedentsMedicines struct {
	Period []UpdateAntecedentPeriod `json:"period"`
}

type UpdateAntecedentPeriod struct {
	Quantity       int    `json:"quantity"`
	Frequency      int    `json:"frequency"`
	FrequencyRatio int    `json:"frequency_ratio"`
	FrequencyUnit  string `json:"frequency_unit"`
	PeriodLength   int    `json:"period_length"`
	PeriodUnit     string `json:"period_unit"`
	Comment        string `json:"comment"`
}

type UpdateTreatmentResponse struct {
	Treatment []model.AntecedentTreatment
	Code      int
	Err       error
}

func ConvertUpdatePeriods(periods []UpdateAntecedentPeriod) []model.AntecedentPeriod {
	convertedPeriods := make([]model.AntecedentPeriod, len(periods))
	for i, p := range periods {
		freqUnit := model.TimeUnitEnum(p.FrequencyUnit)
		periodUnit := model.TimeUnitEnum(p.PeriodUnit)
		convertedPeriods[i] = model.AntecedentPeriod{
			Quantity:       p.Quantity,
			Frequency:      p.Frequency,
			FrequencyRatio: p.FrequencyRatio,
			FrequencyUnit:  freqUnit,
			PeriodLength:   &p.PeriodLength,
			PeriodUnit:     &periodUnit,
			Comment:        &p.Comment,
		}
	}
	return convertedPeriods
}

func UpdateTreatment(input UpdateTreatmentInput, patientID string, antecedentID string) UpdateTreatmentResponse {
	var res []model.AntecedentTreatment

	control, err := graphql.GetPatientById(patientID)
	if err != nil {
		return UpdateTreatmentResponse{Code: 400, Err: errors.New("unable to find patient by ID: " + err.Error())}
	}
	if control.MedicalInfoID == nil || *control.MedicalInfoID == "" {
		return UpdateTreatmentResponse{Code: 400, Err: errors.New("medical folder has not been created")}
	}

	for _, medicine := range input.Medicines {
		periods := ConvertUpdatePeriods(medicine.Period)

		treatment, err := graphql.UpdateAntecedentTreatment(input.ID, antecedentID, model.UpdateAntecedentTreatmentInput{
			CreatedBy: &input.CreatedBy,
			StartDate: &input.StartDate,
			EndDate:   &input.EndDate,
			Medicines: []*model.UpdateAntecedentsMedicinesInput{
				{
					Period: convertToPointerSliceInput(periods),
				},
			},
		})
		if err != nil {
			return UpdateTreatmentResponse{Code: 400, Err: errors.New("unable to update antecedent treatment: " + err.Error())}
		}
		res = append(res, treatment)
	}

	return UpdateTreatmentResponse{
		Treatment: res,
		Code:      200,
		Err:       nil,
	}
}

func convertToPointerSliceInput(periods []model.AntecedentPeriod) []*model.UpdateAntecedentPeriodInput {
	pointerSlice := make([]*model.UpdateAntecedentPeriodInput, len(periods))
	for i := range periods {
		pointerSlice[i] = &model.UpdateAntecedentPeriodInput{
			Quantity:       &periods[i].Quantity,
			Frequency:      &periods[i].Frequency,
			FrequencyRatio: &periods[i].FrequencyRatio,
			FrequencyUnit:  &periods[i].FrequencyUnit,
			PeriodLength:   periods[i].PeriodLength,
			PeriodUnit:     periods[i].PeriodUnit,
			Comment:        periods[i].Comment,
		}
	}
	return pointerSlice
}
