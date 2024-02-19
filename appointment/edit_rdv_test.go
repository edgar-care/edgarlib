package appointment

import (
	"context"
	"github.com/joho/godotenv"
	"log"
	"testing"

	"github.com/edgar-care/edgarlib/graphql"
)

func TestEditRdv(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}
	gqlClient := graphql.CreateClient()

	patient, err := graphql.CreatePatient(context.Background(), gqlClient, "edit_rdv@edgar-sante.fr", "password")
	if err != nil {
		t.Errorf("Error while creating patient: %v", err)
	}
	appointment, err := graphql.CreateRdv(context.Background(), gqlClient, patient.CreatePatient.Id, "doctorId", 0, 100, "WAITING_FOR_REVIEW", "")
	if err != nil {
		t.Errorf("Error while creating appointment: %v", err)
	}
	newAppointment, err := graphql.CreateRdv(context.Background(), gqlClient, "", "newDoctor", 0, 10, "WAITING_FOR_REVIEW", "")
	if err != nil {
		t.Errorf("Error while creating appointment: %v", err)
	}

	response := EditRdv(newAppointment.CreateRdv.Id, appointment.CreateRdv.Id, patient.CreatePatient.Id)

	if response.Err != nil {
		t.Errorf("Unexpected error: %v", response.Err)
	}

	if response.Code != 200 {
		t.Errorf("Expected code 200, got %d", response.Code)
	}
}

func TestEditRdvEmptyId(t *testing.T) {
	response := EditRdv("", "", "")
	if response.Code != 400 || response.Err == nil {
		t.Errorf("Expected error and code 400 but got code %d and err: %s", response.Code, response.Err.Error())
	}
}

func TestEditRdvInvalidId(t *testing.T) {
	response := EditRdv("invalidId", "invalidId", "")
	if response.Code != 400 || response.Err == nil {
		t.Errorf("Expected error and code 400 but got code %d and err: %s", response.Code, response.Err.Error())
	}
}

func TestEditRdvAlreadyBooked(t *testing.T) {
	gqlClient := graphql.CreateClient()
	appointment, err := graphql.CreateRdv(context.Background(), gqlClient, "booked", "doctorId", 0, 100, "WAITING_FOR_REVIEW", "")
	if err != nil {
		t.Errorf("Error while creating appointment: %v", err)
	}
	response := EditRdv(appointment.CreateRdv.Id, "test", "")
	if response.Code != 400 || response.Err == nil {
		t.Errorf("Expected error and code 400 but got code %d and err: %s", response.Code, response.Err.Error())
	}
}

func TestEditRdvInvalidSecondId(t *testing.T) {
	gqlClient := graphql.CreateClient()
	patient, err := graphql.CreatePatient(context.Background(), gqlClient, "edit_rdv_invalid@edgar-sante.fr", "password")
	if err != nil {
		t.Errorf("Error while creating patient: %v", err)
	}
	appointment, err := graphql.CreateRdv(context.Background(), gqlClient, "", "doctorId", 0, 100, "WAITING_FOR_REVIEW", "")
	if err != nil {
		t.Errorf("Error while creating appointment: %v", err)
	}
	response := EditRdv(appointment.CreateRdv.Id, "invalidId", patient.CreatePatient.Id)
	if response.Code != 400 || response.Err == nil {
		t.Errorf("Expected error and code 400 but got code %d and err: %s", response.Code, response.Err.Error())
	}
}

func TestEditRdvInvalidPatientId(t *testing.T) {
	gqlClient := graphql.CreateClient()
	appointment, err := graphql.CreateRdv(context.Background(), gqlClient, "invalidPatientId", "doctorId", 0, 100, "WAITING_FOR_REVIEW", "")
	if err != nil {
		t.Errorf("Error while creating appointment: %v", err)
	}
	newAppointment, err := graphql.CreateRdv(context.Background(), gqlClient, "", "newDoctor", 0, 10, "WAITING_FOR_REVIEW", "")
	if err != nil {
		t.Errorf("Error while creating appointment: %v", err)
	}
	response := EditRdv(newAppointment.CreateRdv.Id, appointment.CreateRdv.Id, "invalidPatientId")
	if response.Code != 400 || response.Err == nil {
		t.Errorf("Expected error and code 400 but got code %d and err: %s", response.Code, response.Err.Error())
	}
}
