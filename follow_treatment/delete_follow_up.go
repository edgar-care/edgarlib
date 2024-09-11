package follow_treatment

import (
	"errors"

	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
)

type DeleteFollowUpResponse struct {
	Deleted        bool
	UpdatedPatient model.Patient
	Code           int
	Err            error
}

func remElement(slice []*string, element string) []*string {
	var result []*string
	for _, v := range slice {
		if *v != element {
			result = append(result, v)
		}
	}
	return result
}

func Delete_follow_up(id string, patientId string) DeleteFollowUpResponse {
	if id == "" {
		return DeleteFollowUpResponse{Deleted: false, UpdatedPatient: model.Patient{}, Code: 400, Err: errors.New("treatment follow up id is required")}
	}

	_, err := graphql.GetTreatmentsFollowUpById(id)
	if err != nil {
		return DeleteFollowUpResponse{Deleted: false, UpdatedPatient: model.Patient{}, Code: 400, Err: errors.New("id does not correspond to a follow up treatment")}
	}

	deleted, err := graphql.DeleteTreatmentsFollowUp(id)
	if err != nil {
		return DeleteFollowUpResponse{Deleted: false, UpdatedPatient: model.Patient{}, Code: 500, Err: errors.New("error while deleting treatment follow up: " + err.Error())}
	}

	patient, err := graphql.GetPatientById(patientId)
	if err != nil {
		return DeleteFollowUpResponse{Deleted: false, UpdatedPatient: model.Patient{}, Code: 400, Err: errors.New("id does not correspond to a doctor")}
	}

	_, err = graphql.UpdatePatientFollowTreatment(patientId, model.UpdatePatientFollowTreatmentInput{TreatmentFollowUpIds: remElement(patient.TreatmentFollowUpIds, id)})
	if err != nil {
		return DeleteFollowUpResponse{Deleted: false, UpdatedPatient: model.Patient{}, Code: 500, Err: errors.New("unable to update patient")}
	}

	return DeleteFollowUpResponse{
		Deleted: deleted,
		Code:    200,
		Err:     nil}
}
