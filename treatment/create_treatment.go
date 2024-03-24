package treatment

import (
	"context"
	"errors"
	"github.com/edgar-care/edgarlib/graphql"
	"github.com/edgar-care/edgarlib/graphql/server/model"
)

type CreateNewTreatmentInput struct {
	Name          string             `json:"name"`
	DiseaseId     string             `json:"disease_id"`
	Treatments    []CreateTreatInput `json:"treatments"`
	StillRelevant bool               `json:"still_relevant"`
}

type CreateTreatInput struct {
	MedicineId string   `json:"medicine_id"`
	Period     []string `json:"period"`
	Day        []string `json:"day"`
	Quantity   int      `json:"quantity"`
}

type CreateTreatmentResponse struct {
	AnteDisease model.AnteDisease
	Treatment   []model.Treatment
	Code        int
	Err         error
}

func CreateTreatment(input CreateNewTreatmentInput, patientID string) CreateTreatmentResponse {
	gqlClient := graphql.CreateClient()
	var res []model.Treatment

	control, err := graphql.GetPatientById(context.Background(), gqlClient, patientID)
	if err != nil {
		return CreateTreatmentResponse{Code: 400, Err: errors.New("unable to find patient by ID: " + err.Error())}
	}
	if control.GetPatientById.Medical_info_id != "" {
		return CreateTreatmentResponse{Code: 400, Err: errors.New("medical folder has already been created")}
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

		treatment, err := graphql.CreateTreatment(context.Background(), gqlClient, periods, days, treat.Quantity, treat.MedicineId)
		if err != nil {
			return CreateTreatmentResponse{Code: 400, Err: errors.New("unable to create treatment: " + err.Error())}
		}

		treatmentPeriods := make([]model.Period, len(treatment.CreateTreatment.Period))
		for i, p := range treatment.CreateTreatment.Period {
			treatmentPeriods[i] = model.Period(p)
		}

		treatmentDays := make([]model.Day, len(treatment.CreateTreatment.Day))
		for i, d := range treatment.CreateTreatment.Day {
			treatmentDays[i] = model.Day(d)
		}

		res = append(res, model.Treatment{
			ID:         treatment.CreateTreatment.Id,
			Period:     treatmentPeriods,
			Day:        treatmentDays,
			Quantity:   treatment.CreateTreatment.Quantity,
			MedicineID: treatment.CreateTreatment.Medicine_id,
		})

	}

	anteDisease, err := graphql.CreateAnteDisease(context.Background(), gqlClient, input.Name, 0, []string{""}, []string{""}, []string{""}, input.StillRelevant)
	if err != nil {
		return CreateTreatmentResponse{Code: 400, Err: errors.New("unable to create ante disease: " + err.Error())}
	}

	var treatmentIds []string
	for _, treatment := range res {
		treatmentIds = append(treatmentIds, treatment.ID)
	}

	return CreateTreatmentResponse{
		AnteDisease: model.AnteDisease{
			Name:          anteDisease.CreateAnteDisease.Name,
			StillRelevant: anteDisease.CreateAnteDisease.Still_relevant,
			TreatmentIds:  treatmentIds,
		},
		Treatment: res,
		Code:      201,
		Err:       nil,
	}
}
