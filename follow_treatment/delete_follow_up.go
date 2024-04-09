package follow_treatment

import (
	"context"
	"errors"

	"github.com/edgar-care/edgarlib/graphql"
	"github.com/edgar-care/edgarlib/graphql/server/model"
)

type DeleteFollowUpResponse struct {
	Deleted        bool
	UpdatedPatient model.Patient
	Code           int
	Err            error
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

func Delete_follow_up(id string, patientId string) DeleteFollowUpResponse {
	if id == "" {
		return DeleteFollowUpResponse{Deleted: false, UpdatedPatient: model.Patient{}, Code: 400, Err: errors.New("treatment follow up id is required")}
	}
	gqlClient := graphql.CreateClient()

	_, err := graphql.GetTreatmentsFollowUpByID(context.Background(), gqlClient, id)
	if err != nil {
		return DeleteFollowUpResponse{Deleted: false, UpdatedPatient: model.Patient{}, Code: 400, Err: errors.New("id does not correspond to a follow up treatment")}
	}

	deleted, err := graphql.DeleteTreatmentsFollowUp(context.Background(), gqlClient, id)
	if err != nil {
		return DeleteFollowUpResponse{Deleted: false, UpdatedPatient: model.Patient{}, Code: 500, Err: errors.New("error while deleting treatment follow up: " + err.Error())}
	}

	patient, err := graphql.GetPatientById(context.Background(), gqlClient, patientId)
	if err != nil {
		return DeleteFollowUpResponse{Deleted: false, UpdatedPatient: model.Patient{}, Code: 400, Err: errors.New("id does not correspond to a doctor")}
	}

	_, err = graphql.UpdatePatient(context.Background(), gqlClient, patientId, patient.GetPatientById.Email, patient.GetPatientById.Password, patient.GetPatientById.Medical_info_id, patient.GetPatientById.Rendez_vous_ids, patient.GetPatientById.Document_ids, remElement(patient.GetPatientById.Treatment_follow_up_ids, id), patient.GetPatientById.Chat_ids)
	if err != nil {
		return DeleteFollowUpResponse{Deleted: false, UpdatedPatient: model.Patient{}, Code: 500, Err: errors.New("unable to update patient")}
	}

	return DeleteFollowUpResponse{
		Deleted: deleted.DeleteTreatmentsFollowUp,
		Code:    200,
		Err:     nil}
}
