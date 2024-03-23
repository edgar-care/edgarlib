package treatment

import (
	"context"
	"errors"
	"github.com/edgar-care/edgarlib/graphql"
)

type DeleteTreatmentResponse struct {
	Deleted bool
	Code    int
	Err     error
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

func DeleteTreatment(treatmentId string, patientID string) DeleteTreatmentResponse {
	if treatmentId == "" {
		return DeleteTreatmentResponse{Deleted: false, Code: 400, Err: errors.New("slot id is required")}
	}
	gqlClient := graphql.CreateClient()

	control, err := graphql.GetPatientById(context.Background(), gqlClient, patientID)
	if err != nil {
		return DeleteTreatmentResponse{Code: 400, Err: errors.New("unable to find patient by ID: " + err.Error())}
	}
	if control.GetPatientById.Medical_info_id != "" {
		return DeleteTreatmentResponse{Code: 400, Err: errors.New("medical folder has already been created")}
	}

	deleted, err := graphql.DeleteTreatment(context.Background(), gqlClient, treatmentId)
	if err != nil {
		return DeleteTreatmentResponse{Deleted: false, Code: 500, Err: errors.New("error while deleting slot: " + err.Error())}
	}

	getAnte, err := graphql.GetAnteDiseaseByID(context.Background(), gqlClient, treatmentId)
	if err != nil {
		return DeleteTreatmentResponse{Deleted: false, Code: 500, Err: errors.New("error updating patient: " + err.Error())}
	}

	_, err = graphql.UpdateAnteDisease(context.Background(), gqlClient, treatmentId, getAnte.GetAnteDiseaseByID.Name, getAnte.GetAnteDiseaseByID.Chronicity, getAnte.GetAnteDiseaseByID.Surgery_ids, []string{""}, remElement(getAnte.GetAnteDiseaseByID.Treatment_ids, treatmentId), getAnte.GetAnteDiseaseByID.Still_relevant)
	if err != nil {
		return DeleteTreatmentResponse{Deleted: false, Code: 500, Err: errors.New("error updating patient: " + err.Error())}
	}

	return DeleteTreatmentResponse{
		Deleted: deleted.DeleteTreatment,
		Code:    200,
		Err:     nil}
}
