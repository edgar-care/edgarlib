package follow_treatment

import (
	"github.com/edgar-care/edgarlib/graphql/model"
	"log"
	"testing"

	"github.com/edgar-care/edgarlib/graphql"
	"github.com/joho/godotenv"
)

func TestCreateTreatmentFollowUp(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}

	patient, err := graphql.CreatePatient(model.CreatePatientInput{
		Email:    "test_patient_create_treatment_follow_up@edgar-sante.fr",
		Password: "password",
		Status:   false,
	})
	if err != nil {
		t.Errorf("Error while creating patient: %v", err)
	}

	input := CreateNewFollowUpInput{
		TreatmentId: "test_treatment_id",
		Date:        123456,
		Period:      []string{"NIGHT"},
	}

	response := CreateTreatmentFollowUp(input, patient.ID)

	if response.Code != 201 {
		t.Errorf("Expected code 200 but got %d", response.Code)
	}
	if response.Err != nil {
		t.Errorf("Expected no error but got: %s", response.Err.Error())
	}

	_, err = graphql.GetTreatmentsFollowUpById(response.TreatmentFollowUp.ID)

	if err != nil {
		t.Errorf("Error while retrieving created appointment: %s", err.Error())
	}

}

func TestCreateRdvInvalidId(t *testing.T) {

	input := CreateNewFollowUpInput{
		TreatmentId: "test_treatment_id",
		Date:        123456,
		Period:      []string{"NIGHT"},
	}

	response := CreateTreatmentFollowUp(input, "patientId")
	if response.Code != 400 || response.Err == nil {
		t.Errorf("Expected error and code 400 but got code %d and err: %s", response.Code, response.Err.Error())
	}
}
