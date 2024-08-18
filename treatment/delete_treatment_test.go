package treatment

import (
	"testing"

	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
	"github.com/google/uuid"
)

func TestDeleteTreatmentWithValidInput(t *testing.T) {
	treatmentID := uuid.New().String()

	anteDisease, err := graphql.CreateAnteDisease(model.CreateAnteDiseaseInput{
		Name:          "Hypertension",
		TreatmentIds:  []string{treatmentID},
		SurgeryIds:    []string{},
		Symptoms:      []string{},
		StillRelevant: true,
	})
	if err != nil {
		t.Errorf("Failed to create AnteDisease: %v", err)
	}

	treatment, err := graphql.CreateTreatment(model.CreateTreatmentInput{
		MedicineID: uuid.New().String(),
		Quantity:   1,
		Period:     []model.Period{"MORNING"},
		Day:        []model.Day{"MONDAY"},
	})
	if err != nil {
		t.Errorf("Failed to create treatment: %v", err)
	}

	response := DeleteTreatment(treatment.ID)

	if response.Code != 200 {
		t.Errorf("Expected response code 200, got %v", response.Code)
	}
	if response.Err != nil {
		t.Errorf("Unexpected error: %v", response.Err)
	}
	if !response.Deleted {
		t.Errorf("Expected treatment to be deleted, but it was not")
	}

	updatedAnteDisease, err := graphql.GetAnteDiseaseByID(anteDisease.ID)
	if err != nil {
		t.Errorf("Failed to retrieve AnteDisease: %v", err)
	}

	if contains(updatedAnteDisease.TreatmentIds, treatment.ID) {
		t.Errorf("Treatment ID was not removed from AnteDisease")
	}
}

func TestDeleteTreatmentWithEmptyTreatmentID(t *testing.T) {
	treatmentID := ""

	response := DeleteTreatment(treatmentID)

	if response.Code != 400 {
		t.Errorf("Expected response code 400, got %v", response.Code)
	}
	if response.Err == nil {
		t.Errorf("Expected an error in response")
	} else if response.Err.Error() != "treatment id is required" {
		t.Errorf("Expected 'treatment id is required' error, got %v", response.Err.Error())
	}
	if response.Deleted {
		t.Errorf("Expected treatment not to be deleted")
	}
}
