package treatment

import (
	"context"
	"errors"
	"github.com/edgar-care/edgarlib/graphql"
	"github.com/edgar-care/edgarlib/graphql/server/model"
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
	gqlClient := graphql.CreateClient()
	var res []model.Treatment

	control, err := graphql.GetPatientById(context.Background(), gqlClient, patientID)
	if err != nil {
		return UpdateTreatmentResponse{Code: 400, Err: errors.New("unable to find patient by ID: " + err.Error())}
	}
	if control.GetPatientById.Medical_info_id == "" {
		return UpdateTreatmentResponse{Code: 400, Err: errors.New("medical folder has not been created")}
	}

	for _, treat := range input.Treatments {
		periods := make([]graphql.Period, len(treat.Period))
		for i, p := range treat.Period {
			periods[i] = graphql.Period(p)
		}

		days := make([]graphql.Day, len(treat.Day))
		for i, d := range treat.Day {
			days[i] = graphql.Day(d)
		}

		_, err = graphql.GetTreatmentByID(context.Background(), gqlClient, treat.ID)
		if err != nil {
			return UpdateTreatmentResponse{Code: 400, Err: errors.New("unable to get treatment by id: " + err.Error())}
		}

		treatment, err := graphql.UpdateTreatment(context.Background(), gqlClient, treat.ID, periods, days, treat.Quantity, treat.MedicineId)
		if err != nil {
			return UpdateTreatmentResponse{Code: 400, Err: errors.New("unable to update treatment: " + err.Error())}
		}

		// Convert periods and days to model types
		treatmentPeriods := make([]model.Period, len(treatment.UpdateTreatment.Period))
		for i, p := range treatment.UpdateTreatment.Period {
			treatmentPeriods[i] = model.Period(p)
		}

		treatmentDays := make([]model.Day, len(treatment.UpdateTreatment.Day))
		for i, d := range treatment.UpdateTreatment.Day {
			treatmentDays[i] = model.Day(d)
		}

		res = append(res, model.Treatment{
			ID:         treatment.UpdateTreatment.Id,
			Period:     treatmentPeriods,
			Day:        treatmentDays,
			Quantity:   treatment.UpdateTreatment.Quantity,
			MedicineID: treatment.UpdateTreatment.Medicine_id,
		})
	}

	return UpdateTreatmentResponse{
		Treatment: res,
		Code:      200,
		Err:       nil,
	}
}
