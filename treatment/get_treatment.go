package treatment

import (
	"context"
	"errors"
	"github.com/edgar-care/edgarlib/graphql"
	"github.com/edgar-care/edgarlib/graphql/server/model"
)

type GetTreatmentByIdResponse struct {
	Treatment model.Treatment
	Code      int
	Err       error
}

type GetTreatmentsResponse struct {
	Treatments []model.Treatment
	Code       int
	Err        error
}

func GetTreatmentById(id string) GetTreatmentByIdResponse {
	gqlClient := graphql.CreateClient()
	var res model.Treatment

	treatment, err := graphql.GetTreatmentByID(context.Background(), gqlClient, id)
	if err != nil {
		return GetTreatmentByIdResponse{model.Treatment{}, 400, errors.New("id does not correspond to a Treatment")}
	}

	treatmentPeriods := make([]model.Period, len(treatment.GetTreatmentByID.Period))
	for i, p := range treatment.GetTreatmentByID.Period {
		treatmentPeriods[i] = model.Period(p)
	}

	treatmentDays := make([]model.Day, len(treatment.GetTreatmentByID.Day))
	for i, d := range treatment.GetTreatmentByID.Day {
		treatmentDays[i] = model.Day(d)
	}

	res = model.Treatment{
		ID:         treatment.GetTreatmentByID.Id,
		Period:     treatmentPeriods,
		Day:        treatmentDays,
		Quantity:   treatment.GetTreatmentByID.Quantity,
		MedicineID: treatment.GetTreatmentByID.Medicine_id,
	}
	return GetTreatmentByIdResponse{res, 200, nil}
}

func GetTreatments() GetTreatmentsResponse {
	gqlClient := graphql.CreateClient()
	var res []model.Treatment

	treatments, err := graphql.GetTreatments(context.Background(), gqlClient)
	if err != nil {
		return GetTreatmentsResponse{[]model.Treatment{}, 400, errors.New("invalid input: " + err.Error())}
	}

	for _, Treatment := range treatments.GetTreatments {
		treatmentPeriods := make([]model.Period, len(treatments.GetTreatments.Period))
		for i, p := range treatments.GetTreatments.Period {
			treatmentPeriods[i] = model.Period(p)
		}

		treatmentDays := make([]model.Day, len(treatments.GetTreatments.Day))
		for i, d := range treatments.GetTreatments.Day {
			treatmentDays[i] = model.Day(d)
		}
		res = append(res, model.Treatment{
			ID:         Treatment.Id,
			Period:     treatmentPeriods,
			Day:        treatmentDays,
			Quantity:   Treatment.Quantity,
			MedicineID: Treatment.Medicine_id,
		})
	}
	return GetTreatmentsResponse{res, 200, nil}
}
