package follow_treatment

import (
	"context"
	"github.com/edgar-care/edgarlib/graphql"
	"github.com/joho/godotenv"
	"log"
	"testing"
)

func TestDelete_follow_up(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}
	gqlClient := graphql.CreateClient()

	patient, err := graphql.CreatePatient(context.Background(), gqlClient, "delete_appointment@edgar-sante.fr", "password")
	if err != nil {
		t.Errorf("Error while creating patient: %v", err)
	}

	input := CreateNewFollowUpInput{
		Period: []string{"NIGHT"},
	}

	periods := make([]graphql.Period, len(input.Period))
	for i, p := range input.Period {
		periods[i] = graphql.Period(p)
	}

	follow_up, err := graphql.CreateTreatmentsFollowUp(context.Background(), gqlClient, "test_treatment_id", 123456, periods)
	if err != nil {
		t.Errorf("Error while creating follow up treatment: %v", err)
	}
	follow_upID := follow_up.CreateTreatmentsFollowUp.Id
	patientID := patient.CreatePatient.Id

	_, err = graphql.UpdatePatient(context.Background(), gqlClient, patientID, patient.CreatePatient.Email, patient.CreatePatient.Password, patient.CreatePatient.Medical_info_id, patient.CreatePatient.Rendez_vous_ids, patient.CreatePatient.Document_ids, append(patient.CreatePatient.Treatment_follow_up_ids, follow_upID))
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

	newPatient, err := graphql.GetPatientById(context.Background(), gqlClient, patientID)
	if err != nil {
		t.Errorf("Error getting updated patient: %v", err)
	}
	for _, v := range newPatient.GetPatientById.Treatment_follow_up_ids {
		if v == follow_upID {
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
	gqlClient := graphql.CreateClient()
	input := CreateNewFollowUpInput{
		Period: []string{"MORNING"},
	}

	periods := make([]graphql.Period, len(input.Period))
	for i, p := range input.Period {
		periods[i] = graphql.Period(p)
	}

	follow_up, err := graphql.CreateTreatmentsFollowUp(context.Background(), gqlClient, "test_treatment_id", 123456, periods)
	if err != nil {
		t.Errorf("Error while creating follow up treatment: %v", err)
	}

	response := Delete_follow_up(follow_up.CreateTreatmentsFollowUp.Id, "invalid")
	if response.Code != 400 || response.Err == nil {
		t.Errorf("Expected error and code 400 but got code %d and err: %s", response.Code, response.Err.Error())
	}
}
