package treatment

import (
	"errors"
	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
)

type DeleteTreatmentResponse struct {
	Deleted bool
	Code    int
	Err     error
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

func DeleteTreatment(treatmentID string) DeleteTreatmentResponse {
	if treatmentID == "" {
		return DeleteTreatmentResponse{Deleted: false, Code: 400, Err: errors.New("treatment id is required")}
	}

	deleted, err := graphql.DeleteTreatment(treatmentID)
	if err != nil {
		return DeleteTreatmentResponse{Deleted: false, Code: 500, Err: errors.New("error while deleting treatment: " + err.Error())}
	}

	antediseases, err := graphql.GetAnteDiseases(nil)
	if err != nil {
		return DeleteTreatmentResponse{Deleted: false, Code: 500, Err: errors.New("error getting antediseases: " + err.Error())}
	}

	var affectedAntediseases []string

	for _, ad := range antediseases {
		if contains(ad.TreatmentIds, treatmentID) {
			affectedAntediseases = append(affectedAntediseases, ad.ID)
		}
	}

	for _, adID := range affectedAntediseases {
		ad, err := graphql.GetAnteDiseaseByID(adID)
		if err != nil {
			return DeleteTreatmentResponse{Deleted: false, Code: 500, Err: errors.New("error getting antedisease by ID: " + err.Error())}
		}

		updatedTreatmentIDs := make([]*string, len(ad.TreatmentIds))
		for i, v := range ad.TreatmentIds {
			updatedTreatmentIDs[i] = &v
		}

		_, err = graphql.UpdatePatientAntediesae(adID, model.UpdatePatientAntediseaseInput{
			TreatmentIds: remElement(updatedTreatmentIDs, treatmentID),
		})
		if err != nil {
			return DeleteTreatmentResponse{Deleted: false, Code: 500, Err: errors.New("error updating antedisease: " + err.Error())}
		}
	}

	return DeleteTreatmentResponse{
		Deleted: deleted,
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
