package follow_treatment

import (
	"errors"
	"github.com/edgar-care/edgarlib/graphql"
	"github.com/edgar-care/edgarlib/graphql/model"
)

type GetTreatmentFollowUpByIdResponse struct {
	TreatmentFollowUp model.TreatmentsFollowUp
	Code              int
	Err               error
}

type GetTreatmentFollowUpsResponse struct {
	TreatmentFollowUps []model.TreatmentsFollowUp
	Code               int
	Err                error
}

func GetTreatmentFollowUpById(id string) GetTreatmentFollowUpByIdResponse {
	followUp, err := graphql.GetTreatmentsFollowUpById(id)
	if err != nil {
		return GetTreatmentFollowUpByIdResponse{model.TreatmentsFollowUp{}, 400, errors.New("id does not correspond to a medicament")}
	}

	treatmentFollowPeriods := make([]model.Period, len(followUp.Period))
	for i, p := range followUp.Period {
		treatmentFollowPeriods[i] = model.Period(p)
	}

	return GetTreatmentFollowUpByIdResponse{followUp, 200, nil}
}

func GetTreatmentFollowUp(patientId string) GetTreatmentFollowUpsResponse {
	followUps, err := graphql.GetTreatmentsFollowUps(patientId, nil)
	if err != nil {
		return GetTreatmentFollowUpsResponse{[]model.TreatmentsFollowUp{}, 400, errors.New("invalid input: " + err.Error())}
	}
	return GetTreatmentFollowUpsResponse{followUps, 200, nil}
}
