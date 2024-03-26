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

func DeleteTreatment(treatmentID string, patientID string) DeleteTreatmentResponse {
	if treatmentID == "" {
		return DeleteTreatmentResponse{Deleted: false, Code: 400, Err: errors.New("treatment id is required")}
	}
	gqlClient := graphql.CreateClient()

	deleted, err := graphql.DeleteTreatment(context.Background(), gqlClient, treatmentID)
	if err != nil {
		return DeleteTreatmentResponse{Deleted: false, Code: 500, Err: errors.New("error while deleting treatment: " + err.Error())}
	}

	antediseases, err := graphql.GetAnteDiseases(context.Background(), gqlClient)
	if err != nil {
		return DeleteTreatmentResponse{Deleted: false, Code: 500, Err: errors.New("error getting antediseases: " + err.Error())}
	}

	var affectedAntediseases []string

	for _, ad := range antediseases.GetAnteDiseases {
		if contains(ad.Treatment_ids, treatmentID) {
			affectedAntediseases = append(affectedAntediseases, ad.Id)
		}
	}

	for _, adID := range affectedAntediseases {
		ad, err := graphql.GetAnteDiseaseByID(context.Background(), gqlClient, adID)
		if err != nil {
			return DeleteTreatmentResponse{Deleted: false, Code: 500, Err: errors.New("error getting antedisease by ID: " + err.Error())}
		}

		updatedTreatmentIDs := remElement(ad.GetAnteDiseaseByID.Treatment_ids, treatmentID)

		_, err = graphql.UpdateAnteDisease(context.Background(), gqlClient, adID, ad.GetAnteDiseaseByID.Name, ad.GetAnteDiseaseByID.Chronicity, ad.GetAnteDiseaseByID.Surgery_ids, ad.GetAnteDiseaseByID.Symptoms, updatedTreatmentIDs, ad.GetAnteDiseaseByID.Still_relevant)
		if err != nil {
			return DeleteTreatmentResponse{Deleted: false, Code: 500, Err: errors.New("error updating antedisease: " + err.Error())}
		}
	}

	return DeleteTreatmentResponse{
		Deleted: deleted.DeleteTreatment,
		Code:    200,
		Err:     nil,
	}
}

func contains(slice []string, element string) bool {
	for _, e := range slice {
		if e == element {
			return true
		}
	}
	return false
}
