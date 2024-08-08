package treatment

import (
	"errors"
	"github.com/edgar-care/edgarlib/graphql"
	"github.com/edgar-care/edgarlib/graphql/model"
)

type UpdateTreatmentInput struct {
	Treatments []TreatmentsInput `json:"treatments"`
}

type TreatmentsInput struct {
	ID         string   `json:"id"`
	MedicineId string   `json:"medicine_id"`
	Period     []string `json:"period"`
	Day        []string `json:"day"`
	Quantity   int      `json:"quantity"`
}

type UpdateTreatmentResponse struct {
	Treatment []model.Treatment
	Code      int
	Err       error
}

func UpdateTreatment(input UpdateTreatmentInput, patientID string) UpdateTreatmentResponse {
	var res []model.Treatment

	control, err := graphql.GetPatientById(patientID)
	if err != nil {
		return UpdateTreatmentResponse{Code: 400, Err: errors.New("unable to find patient by ID: " + err.Error())}
	}
	if control.MedicalInfoID == nil || *control.MedicalInfoID == "" {
		return UpdateTreatmentResponse{Code: 400, Err: errors.New("medical folder has not been created")}
	}

	for _, treat := range input.Treatments {
		periods, days := ConvertPeriodsAndDays(treat.Period, treat.Day)

		_, err = graphql.GetTreatmentByID(treat.ID)
		if err != nil {
			return UpdateTreatmentResponse{Code: 400, Err: errors.New("unable to get treatment by id: " + err.Error())}
		}

		treatment, err := graphql.UpdateTreatment(treat.ID, model.UpdateTreatmentInput{Period: periods, Day: days, Quantity: &treat.Quantity, MedicineID: &treat.MedicineId})
		if err != nil {
			return UpdateTreatmentResponse{Code: 400, Err: errors.New("unable to update treatment: " + err.Error())}
		}

		res = append(res, treatment)
	}

	return UpdateTreatmentResponse{
		Treatment: res,
		Code:      200,
		Err:       nil,
	}
}
