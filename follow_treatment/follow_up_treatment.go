package follow_treatment

import (
	"context"
	"errors"
	"github.com/edgar-care/edgarlib/graphql"
	"github.com/edgar-care/edgarlib/graphql/server/model"
)

type CreateNewFollowUpInput struct {
	TreatmentId string   `json:"treatment_id"`
	Date        int      `json:"disease_id"`
	Period      []string `json:"period"`
}

type CreateTreatmentFollowUpResponse struct {
	TreatmentFollowUp model.TreatmentsFollowUp
	Code              int
	Err               error
}

func CreateTreatmentFollowUp(input CreateNewFollowUpInput, patientID string) CreateTreatmentFollowUpResponse {
	gqlClient := graphql.CreateClient()

	periods := make([]graphql.Period, len(input.Period))
	for i, p := range input.Period {
		periods[i] = graphql.Period(p)
	}

	followUp, err := graphql.CreateTreatmentsFollowUp(context.Background(), gqlClient, input.TreatmentId, input.Date, periods)
	if err != nil {
		return CreateTreatmentFollowUpResponse{TreatmentFollowUp: model.TreatmentsFollowUp{}, Code: 400, Err: errors.New("update failed" + err.Error())}
	}

	patient, err := graphql.GetPatientById(context.Background(), gqlClient, patientID)
	if err != nil {
		return CreateTreatmentFollowUpResponse{TreatmentFollowUp: model.TreatmentsFollowUp{}, Code: 400, Err: errors.New("id does not correspond to a patient")}
	}

	_, err = graphql.UpdatePatient(context.Background(), gqlClient, patientID, patient.GetPatientById.Email, patient.GetPatientById.Password, patient.GetPatientById.Medical_info_id, patient.GetPatientById.Rendez_vous_ids, patient.GetPatientById.Document_ids, append(patient.GetPatientById.Treatment_follow_up_ids, followUp.CreateTreatmentsFollowUp.Id))
	if err != nil {
		return CreateTreatmentFollowUpResponse{TreatmentFollowUp: model.TreatmentsFollowUp{}, Code: 500, Err: errors.New("unable to update patient")}
	}

	treatmentFollowPeriods := make([]model.Period, len(followUp.CreateTreatmentsFollowUp.Period))
	for i, p := range followUp.CreateTreatmentsFollowUp.Period {
		treatmentFollowPeriods[i] = model.Period(p)
	}

	return CreateTreatmentFollowUpResponse{
		TreatmentFollowUp: model.TreatmentsFollowUp{
			ID:          followUp.CreateTreatmentsFollowUp.Id,
			TreatmentID: followUp.CreateTreatmentsFollowUp.Treatment_id,
			Date:        followUp.CreateTreatmentsFollowUp.Date,
			Period:      treatmentFollowPeriods,
		},
		Code: 201,
		Err:  nil,
	}

}
