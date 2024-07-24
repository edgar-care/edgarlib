package dashboard

import (
	"context"
	"github.com/edgar-care/edgarlib/graphql"
	"github.com/joho/godotenv"
	"log"
	"testing"
)

func TestDeletePatient(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}
	gqlClient := graphql.CreateClient()

	patient, err := graphql.CreatePatient(context.Background(), gqlClient, "test_delete_patient@edgar-sante.fr", "password")
	if err != nil {
		t.Errorf("Error while creating patient: %v", err)
	}
	doctor, err := graphql.CreateDoctor(context.Background(), gqlClient, "test_delete_patient@edgar-sante.fr", "password", "name", "first", graphql.AddressInput{"", "", "", ""})
	if err != nil {
		t.Errorf("Error while creating doctor: %v", err)
	}

	_, err = graphql.UpdateDoctor(context.Background(), gqlClient, doctor.CreateDoctor.Id, doctor.CreateDoctor.Email, doctor.CreateDoctor.Password, doctor.CreateDoctor.Name, doctor.CreateDoctor.Firstname, []string{}, []string{patient.CreatePatient.Id}, graphql.AddressInput{"", "", "", ""}, []string{}, []string{}, "", []string{})
	if err != nil {
		t.Errorf("Error while updating doctor: %v", err)
	}

	response := DeletePatient(patient.CreatePatient.Id, doctor.CreateDoctor.Id)

	if response.Err != nil {
		t.Errorf("Unexpected error: %v", response.Err)
	}

	if response.Code != 200 {
		t.Errorf("Expected code 200, got %d", response.Code)
	}

	if len(response.UpdatedDoctor.PatientIds) != 0 {
		t.Errorf("Expected empty patient id list but got %v", response.UpdatedDoctor.PatientIds)
	}
}

func TestDeletePatientEmpty(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}

	response := DeletePatient("", "doctorID")

	if response.Err == nil {
		t.Error("Unexpected null error")
	}

	if response.Code != 400 {
		t.Errorf("Expected code 400, got %d", response.Code)
	}
}

func TestDeletePatientInvalidPatient(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}

	response := DeletePatient("invalid", "doctorID")

	if response.Err == nil {
		t.Error("Unexpected null error")
	}

	if response.Code != 400 {
		t.Errorf("Expected code 400, got %d", response.Code)
	}
}

func TestDeletePatientInvalidDoctor(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}
	gqlClient := graphql.CreateClient()

	patient, err := graphql.CreatePatient(context.Background(), gqlClient, "test_delete_patient_invalid_doc@edgar-sante.fr", "password")
	if err != nil {
		t.Errorf("Error while creating patient: %v", err)
	}
	response := DeletePatient(patient.CreatePatient.Id, "invalid")

	if response.Err == nil {
		t.Error("Unexpected null error")
	}

	if response.Code != 400 {
		t.Errorf("Expected code 400, got %d", response.Code)
	}
}
