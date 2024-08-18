package follow_treatment

import (
	"errors"
	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
)

type CreateNewFollowUpInput struct {
	TreatmentId string   `json:"treatment_id"`
	Date        int      `json:"date"`
	Period      []string `json:"period"`
}

type CreateTreatmentFollowUpResponse struct {
	TreatmentFollowUp model.TreatmentsFollowUp
	Code              int
	Err               error
}

func CreateTreatmentFollowUp(input CreateNewFollowUpInput, patientID string) CreateTreatmentFollowUpResponse {
	periods := make([]model.Period, len(input.Period))
	for i, p := range input.Period {
		periods[i] = model.Period(p)
	}

	followUp, err := graphql.CreateTreatmentsFollowUp(model.CreateTreatmentsFollowUpInput{TreatmentID: input.TreatmentId, Date: input.Date, Period: periods})
	if err != nil {
		return CreateTreatmentFollowUpResponse{TreatmentFollowUp: model.TreatmentsFollowUp{}, Code: 400, Err: errors.New("update failed" + err.Error())}
	}

	patient, err := graphql.GetPatientById(patientID)
	if err != nil {
		return CreateTreatmentFollowUpResponse{TreatmentFollowUp: model.TreatmentsFollowUp{}, Code: 400, Err: errors.New("id does not correspond to a patient")}
	}

	_, err = graphql.UpdatePatient(patientID, model.UpdatePatientInput{
		TreatmentFollowUpIds: append(patient.TreatmentFollowUpIds, &followUp.ID),
	})
	if err != nil {
		return CreateTreatmentFollowUpResponse{TreatmentFollowUp: model.TreatmentsFollowUp{}, Code: 500, Err: errors.New("unable to update patient")}
	}

	return CreateTreatmentFollowUpResponse{
		TreatmentFollowUp: followUp,
		Code:              201,
		Err:               nil,
	}
}
