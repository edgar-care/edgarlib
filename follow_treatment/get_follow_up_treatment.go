package follow_treatment

import (
	"context"
	"errors"
	"github.com/edgar-care/edgarlib/graphql"
	"github.com/edgar-care/edgarlib/graphql/server/model"
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
	gqlClient := graphql.CreateClient()
	var res model.TreatmentsFollowUp

	followUp, err := graphql.GetTreatmentsFollowUpByID(context.Background(), gqlClient, id)
	if err != nil {
		return GetTreatmentFollowUpByIdResponse{model.TreatmentsFollowUp{}, 400, errors.New("id does not correspond to a medicament")}
	}

	treatmentFollowPeriods := make([]model.Period, len(followUp.GetTreatmentsFollowUpById.Period))
	for i, p := range followUp.GetTreatmentsFollowUpById.Period {
		treatmentFollowPeriods[i] = model.Period(p)
	}

	res = model.TreatmentsFollowUp{
		ID:          followUp.GetTreatmentsFollowUpById.Id,
		TreatmentID: followUp.GetTreatmentsFollowUpById.Treatment_id,
		Date:        followUp.GetTreatmentsFollowUpById.Date,
		Period:      treatmentFollowPeriods,
	}
	return GetTreatmentFollowUpByIdResponse{res, 200, nil}
}

func GetTreatmentFollowUp(patientId string) GetTreatmentFollowUpsResponse {
	gqlClient := graphql.CreateClient()
	var res []model.TreatmentsFollowUp

	followUps, err := graphql.GetTreatmentsFollowUp(context.Background(), gqlClient, patientId)
	if err != nil {
		return GetTreatmentFollowUpsResponse{[]model.TreatmentsFollowUp{}, 400, errors.New("invalid input: " + err.Error())}
	}

	for _, follow := range followUps.GetTreatmentsFollowUps {
		var followPeriod []model.Period
		for _, p := range follow.Period {
			followPeriod = append(followPeriod, model.Period(p))
		}

		res = append(res, model.TreatmentsFollowUp{
			ID:          follow.Id,
			TreatmentID: follow.Treatment_id,
			Date:        follow.Date,
			Period:      followPeriod,
		})
	}
	return GetTreatmentFollowUpsResponse{res, 200, nil}
}
