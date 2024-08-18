package appointment

import (
	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
	"github.com/joho/godotenv"
	"log"
	"testing"
)

func TestGetRdvPatient(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}
	patient, err := graphql.CreatePatient(model.CreatePatientInput{
		Email:    "test_get_rdv_patient@edgar-sante.fr",
		Password: "password",
	})
	if err != nil {
		t.Errorf("Error creating patient: %v", err)
	}
	appointment, err := graphql.CreateRdv(model.CreateRdvInput{
		IDPatient:         patient.ID,
		DoctorID:          "doctorId",
		StartDate:         0,
		EndDate:           10,
		AppointmentStatus: "WAITING_FOR_REVIEW",
		SessionID:         "",
	})
	if err != nil {
		t.Errorf("Error creating appointment: %v", err)
	}

	response := GetRdvPatient(appointment.ID, patient.ID)

	if response.Err != nil {
		t.Errorf("Error getting rdv patient: %v", response.Err)
	}

	if response.Code != 200 {
		t.Errorf("Expected code 200, got %d", response.Code)
	}

	if response.Rdv.ID != appointment.ID {
		t.Errorf("Expected appointment id %s, got appointment id %s", appointment.ID, response.Rdv.ID)
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
	appointment, err := graphql.CreateRdv(model.CreateRdvInput{
		IDPatient:         "patientId",
		DoctorID:          "doctorId",
		StartDate:         0,
		EndDate:           10,
		AppointmentStatus: "WAITING_FOR_REVIEW",
		SessionID:         "",
	})
	if err != nil {
		t.Errorf("Error creating appointment: %v", err)
	}
	response := GetRdvPatient(appointment.ID, "invalid")

	if response.Err == nil {
		t.Error("Unexpected null error")
	}

	if response.Code != 403 {
		t.Errorf("Expected code 403, got %d", response.Code)
	}
}
