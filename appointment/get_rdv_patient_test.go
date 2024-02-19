package appointment

import (
	"context"
	"github.com/edgar-care/edgarlib/graphql"
	"github.com/joho/godotenv"
	"log"
	"testing"
)

func TestGetRdvPatient(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}
	gqlClient := graphql.CreateClient()
	patient, err := graphql.CreatePatient(context.Background(), gqlClient, "test_get_rdv_patient@edgar-sante.fr", "password")
	if err != nil {
		t.Errorf("Error creating patient: %v", err)
	}
	appointment, err := graphql.CreateRdv(context.Background(), gqlClient, patient.CreatePatient.Id, "doctorId", 0, 10, "WAITING_FOR_REVIEW", "")
	if err != nil {
		t.Errorf("Error creating appointment: %v", err)
	}

	response := GetRdvPatient(appointment.CreateRdv.Id, patient.CreatePatient.Id)

	if response.Err != nil {
		t.Errorf("Error getting rdv patient: %v", response.Err)
	}

	if response.Code != 200 {
		t.Errorf("Expected code 200, got %d", response.Code)
	}

	if response.Rdv.ID != appointment.CreateRdv.Id {
		t.Errorf("Expected appointment id %s, got appointment id %s", appointment.CreateRdv.Id, response.Rdv.ID)
	}
}

func TestGetRdvPatientInvalid(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}

	response := GetRdvPatient("invalidAppointmentGetPatient", "patient")

	if response.Err == nil {
		t.Error("Unexpected null error")
	}

	if response.Code != 400 {
		t.Errorf("Expected code 400, got %d", response.Code)
	}
}

func TestGetRdvPatientUnauthorized(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}
	gqlClient := graphql.CreateClient()
	appointment, err := graphql.CreateRdv(context.Background(), gqlClient, "patientId", "doctorId", 0, 10, "WAITING_FOR_REVIEW", "")
	if err != nil {
		t.Errorf("Error creating appointment: %v", err)
	}
	response := GetRdvPatient(appointment.CreateRdv.Id, "invalid")

	if response.Err == nil {
		t.Error("Unexpected null error")
	}

	if response.Code != 403 {
		t.Errorf("Expected code 403, got %d", response.Code)
	}
}
