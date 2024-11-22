package treatment

import (
	"errors"
	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
	"github.com/edgar-care/edgarlib/v2/medical_folder"
)

type UpdateTreatmentInput struct {
	StartDate int                          `json:"start_date"`
	EndDate   int                          `json:"end_date"`
	Medicines []UpdateAntecedentsMedicines `json:"medicines"`
}

type UpdateAntecedentsMedicines struct {
	MedicineID string                   `json:"medicine_id"`
	Comment    string                   `json:"comment"`
	Period     []UpdateAntecedentPeriod `json:"period"`
}

type UpdateAntecedentPeriod struct {
	Quantity       int    `json:"quantity"`
	Frequency      int    `json:"frequency"`
	FrequencyRatio int    `json:"frequency_ratio"`
	FrequencyUnit  string `json:"frequency_unit"`
	PeriodLength   int    `json:"period_length"`
	PeriodUnit     string `json:"period_unit"`
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
		}
	}
	return convertedPeriods
}

func UpdateTreatment(input UpdateTreatmentInput, patientID string, treatmentID string) UpdateTreatmentResponse {
	var res []model.AntecedentTreatment

	control, err := graphql.GetPatientById(patientID)
	if err != nil {
		return UpdateTreatmentResponse{Code: 400, Err: errors.New("unable to find patient by ID: " + err.Error())}
	}
	if control.MedicalInfoID == nil || *control.MedicalInfoID == "" {
		return UpdateTreatmentResponse{Code: 400, Err: errors.New("medical folder has not been created")}
	}

	medicalAntecedent := medical_folder.GetMedicalAntecedents(patientID)
	if medicalAntecedent.Err != nil {
		return UpdateTreatmentResponse{Code: 400, Err: errors.New("unable to find medical folder by ID: " + err.Error())}
	}

	var antecedentID string
	for _, antecedent := range medicalAntecedent.MedicalAntecedents {
		for _, treatment := range antecedent.Treatments {
			if treatment.ID == treatmentID {
				antecedentID = antecedent.ID
				break
			}
		}
		if antecedentID != "" {
			break
		}
	}

	if antecedentID == "" {
		return UpdateTreatmentResponse{Code: 404, Err: errors.New("treatment not found in any antecedent")}
	}

	var medicines []*model.UpdateAntecedentsMedicinesInput
	for _, medicine := range input.Medicines {
		medicineID := medicine.MedicineID
		comment := medicine.Comment
		periods := ConvertUpdatePeriods(medicine.Period)

		medicines = append(medicines, &model.UpdateAntecedentsMedicinesInput{
			MedicineID: &medicineID,
			Comment:    &comment,
			Period:     convertToPointerSliceInput(periods),
		})
	}

	treatment, err := graphql.UpdateAntecedentTreatment(treatmentID, antecedentID, model.UpdateAntecedentTreatmentInput{
		StartDate: &input.StartDate,
		EndDate:   &input.EndDate,
		Medicines: medicines,
	})
	if err != nil {
		return UpdateTreatmentResponse{Code: 400, Err: errors.New("unable to update antecedent treatment: " + err.Error())}
	}
	res = append(res, treatment)

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
		}
	}
	return pointerSlice
}
