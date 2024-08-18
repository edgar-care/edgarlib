package follow_treatment

import (
	"github.com/edgar-care/edgarlib/v2/graphql/model"
	"log"
	"testing"

	"github.com/joho/godotenv"

	"github.com/edgar-care/edgarlib/v2/graphql"
)

func TestGetTreatmentFollowUp(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}

	patient, err := graphql.CreatePatient(model.CreatePatientInput{Email: "test_patient_get_treatment_follow_up@edgar-sante.fr", Password: "password"})
	if err != nil {
		t.Errorf("Error while creating patient: %v", err)
	}

	input := CreateNewFollowUpInput{
		Period: []string{"MORNING"},
	}

	periods := make([]model.Period, len(input.Period))
	for i, p := range input.Period {
		periods[i] = model.Period(p)
	}

	_, err = graphql.CreateTreatmentsFollowUp(model.CreateTreatmentsFollowUpInput{
		TreatmentID: "treament_id",
		Date:        123456,
		Period:      periods,
	})
	if err != nil {
		t.Errorf("Error while creating follow up treatment: %v", err)
	}

	response := GetTreatmentFollowUp(patient.ID)

	if response.Err != nil {
		t.Errorf("Unexpected error: %v", response.Err)
	}

	if response.Code != 200 {
		t.Errorf("Expected code 200, got %d", response.Code)
	}

	//if len(response.TreatmentFollowUps) == 0 {
	//	t.Errorf("Expected non-empty Treatment follow up slice, got empty slice")
	//}
	//if response.TreatmentFollowUps[0].ID != treatment_follow.CreateTreatmentsFollowUp.Id {
	//	t.Errorf("Expected first Treatment follow up slice to have it's ID=%s but go ID=%s", response.TreatmentFollowUps[0].ID, treatment_follow.CreateTreatmentsFollowUp.Id)
	//}
}

func TestGetTreatmentFollowUpInvalid(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}

	response := GetTreatmentFollowUp("111111111111111111111111")

	if response.Err == nil {
		t.Error("Unexpected null error")
	}

	if response.Code != 400 {
		t.Errorf("Expected code 400, got %d", response.Code)
	}
}

func TestGetTreatmentFollowUpById(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}

	_, err := graphql.CreatePatient(model.CreatePatientInput{
		Email:               "test_patient_appointment2@edgar-sante.fr",
		Password:            "password",
		Status:              false,
		DeviceConnect:       nil,
		DoubleAuthMethodsID: nil,
		TrustDevices:        nil,
	})
	if err != nil {
		t.Errorf("Error while creating patient: %v", err)
	}

	input := CreateNewFollowUpInput{
		Period: []string{"MORNING"},
	}

	periods := make([]model.Period, len(input.Period))
	for i, p := range input.Period {
		periods[i] = model.Period(p)
	}

	follow, err := graphql.CreateTreatmentsFollowUp(model.CreateTreatmentsFollowUpInput{
		TreatmentID: "treament_id",
		Date:        123456,
		Period:      periods,
	})
	if err != nil {
		t.Errorf("Error while creating follow up treatment: %v", err)
	}

	response := GetTreatmentFollowUpById(follow.ID)

	if response.Err != nil {
		t.Errorf("Unexpected error: %v", response.Err)
	}

	if response.Code != 200 {
		t.Errorf("Expected code 200, got %d", response.Code)
	}
}

func TestGetTreatmentFollowUpByIdInvalid(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}

	response := GetTreatmentFollowUpById("111111111111111111111111")

	if response.Err == nil {
		t.Error("Unexpected null error")
	}

	if response.Code != 400 {
		t.Errorf("Expected code 400, got %d", response.Code)
	}
}
