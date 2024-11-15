package treatment

import (
	"errors"
	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
)

type DeleteTreatmentResponse struct {
	Deleted bool
	Code    int
	Err     error
}

func remElement(slice []string, element string) []string {
	var result []string
	for _, v := range slice {
		if v != element {
			result = append(result, v)
		}
	}
	return result
}

func DeleteTreatment(treatmentID string) DeleteTreatmentResponse {
	if treatmentID == "" {
		return DeleteTreatmentResponse{Deleted: false, Code: 400, Err: errors.New("treatment id is required")}
	}

	deleted, err := graphql.DeleteAntecdentTreatment(treatmentID)
	if err != nil {
		return DeleteTreatmentResponse{Deleted: false, Code: 500, Err: errors.New("error while deleting treatment: " + err.Error())}
	}

	medicalAntecedent, err := graphql.GetMedicalAntecedents(nil)
	if err != nil {
		return DeleteTreatmentResponse{Deleted: false, Code: 500, Err: errors.New("error getting antediseases: " + err.Error())}
	}

	var affectedAntediseases []string

	for _, ad := range medicalAntecedent {
		for _, treatment := range ad.Treatments {
			if treatment.ID == treatmentID {
				affectedAntediseases = append(affectedAntediseases, ad.ID)
				break
			}
		}
	}

	for _, adID := range affectedAntediseases {
		ad, err := graphql.GetMedicalAntecedentsById(adID)
		if err != nil {
			return DeleteTreatmentResponse{Deleted: false, Code: 500, Err: errors.New("error getting antedisease by ID: " + err.Error())}
		}

		updatedTreatments := make([]*model.AntecedentTreatment, 0)
		for _, treatment := range ad.Treatments {
			if treatment.ID != treatmentID {
				updatedTreatments = append(updatedTreatments, treatment)
			}
		}

		updatedTreatmentsInput := make([]*model.UpdateTreatmentMedicalInfoInput, len(updatedTreatments))
		for i, treatment := range updatedTreatments {
			medicinesInput := make([]*model.UpdateTreatmentMedicineMedicalInput, len(treatment.Medicines))
			for j, medicine := range treatment.Medicines {
				periodsInput := make([]*model.UpdateTreatmentPeriodMedicalInput, len(medicine.Period))
				for k, period := range medicine.Period {
					periodsInput[k] = &model.UpdateTreatmentPeriodMedicalInput{
						Quantity:       &period.Quantity,
						Frequency:      &period.Frequency,
						FrequencyRatio: &period.FrequencyRatio,
						FrequencyUnit:  &period.FrequencyUnit,
						PeriodLength:   period.PeriodLength,
						PeriodUnit:     period.PeriodUnit,
					}
				}
				medicinesInput[j] = &model.UpdateTreatmentMedicineMedicalInput{
					MedicineID: &medicine.MedicineID,
					Comment:    medicine.Comment,
					Period:     periodsInput,
				}
			}
			updatedTreatmentsInput[i] = &model.UpdateTreatmentMedicalInfoInput{
				ID:        &treatment.ID,
				CreatedBy: &treatment.CreatedBy,
				StartDate: &treatment.StartDate,
				EndDate:   treatment.EndDate,
				Medicines: medicinesInput,
			}
		}

		_, err = graphql.UpdateTreatmentsMedicalAntecedents(adID, model.UpdateTreatmentMedicalAntecedentsInput{
			Treatments: updatedTreatmentsInput,
		})
	}

	return DeleteTreatmentResponse{
		Deleted: deleted,
		Code:    200,
		Err:     nil,
	}
}

func contains(slice []string, element string) bool {
	for _, e := range slice {
		if e == element {
			return true
		}
	}
	return false
}
