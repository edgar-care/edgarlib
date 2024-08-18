package follow_treatment

import (
	"github.com/edgar-care/edgarlib/v2/graphql/model"
	"log"
	"testing"

	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/joho/godotenv"
)

func TestDelete_follow_up(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}

	patient, err := graphql.CreatePatient(model.CreatePatientInput{Email: "delete_follow_up@edgar-sante.fr", Password: "password"})
	if err != nil {
		t.Errorf("Error while creating patient: %v", err)
	}

	input := CreateNewFollowUpInput{
		Period: []string{"NIGHT"},
	}

	periods := make([]model.Period, len(input.Period))
	for i, p := range input.Period {
		periods[i] = model.Period(p)
	}

	follow_up, err := graphql.CreateTreatmentsFollowUp(model.CreateTreatmentsFollowUpInput{TreatmentID: "test_treatment_id", Date: 123456, Period: periods})
	if err != nil {
		t.Errorf("Error while creating follow up treatment: %v", err)
	}
	follow_upID := follow_up.ID
	patientID := patient.ID

	_, err = graphql.UpdatePatient(patientID, model.UpdatePatientInput{TreatmentFollowUpIds: append(patient.TreatmentFollowUpIds, &follow_upID)})
	if err != nil {
		t.Errorf("Error while updating patient: %v", err)
	}
	response := Delete_follow_up(follow_upID, patientID)

	if response.Err != nil {
		t.Errorf("Unexpected error: %v", response.Err)
	}

	if response.Code != 200 {
		t.Errorf("Expected code 200, got %d", response.Code)
	}

	newPatient, err := graphql.GetPatientById(patientID)
	if err != nil {
		t.Errorf("Error getting updated patient: %v", err)
	}
	for _, v := range newPatient.TreatmentFollowUpIds {
		if v == &follow_upID {
			t.Error("follow up treatment's id has not been deleted on patient")
		}
	}
}

func TestDelete_follow_upEmptyId(t *testing.T) {
	response := Delete_follow_up("", "test")
	if response.Code != 400 || response.Err == nil {
		t.Errorf("Expected error and code 400 but got code %d and err: %s", response.Code, response.Err.Error())
	}
}

func TestDelete_follow_upInvalidId(t *testing.T) {
	response := Delete_follow_up("invalid", "test")
	if response.Code != 400 || response.Err == nil {
		t.Errorf("Expected error and code 400 but got code %d and err: %s", response.Code, response.Err.Error())
	}
}

func TestDelete_follow_upInvalidTreatment(t *testing.T) {
	input := CreateNewFollowUpInput{
		Period: []string{"MORNING"},
	}

	periods := make([]model.Period, len(input.Period))
	for i, p := range input.Period {
		periods[i] = model.Period(p)
	}

	follow_up, err := graphql.CreateTreatmentsFollowUp(model.CreateTreatmentsFollowUpInput{TreatmentID: "test_treatment_id", Date: 123456, Period: periods})
	if err != nil {
		t.Errorf("Error while creating follow up treatment: %v", err)
	}

	response := Delete_follow_up(follow_up.ID, "invalid")
	if response.Code != 400 || response.Err == nil {
		t.Errorf("Expected error and code 400 but got code %d and err: %s", response.Code, response.Err.Error())
	}
}
