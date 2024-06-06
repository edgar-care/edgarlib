package appointment

import (
	"github.com/edgar-care/edgarlib/graphql"
	"github.com/edgar-care/edgarlib/graphql/model"
	"github.com/joho/godotenv"
	"log"
	"testing"
)

func TestValidateRdv(t *testing.T) {
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
	response := ValidateRdv(appointment.ID, ReviewInput{
		Reason:     "testing",
		Validation: true,
	})

	if response.Err != nil {
		t.Errorf("Error getting rdv patient: %v", response.Err)
	}

	if response.Code != 200 {
		t.Errorf("Expected code 200, got %d", response.Code)
	}
	if response.Rdv.AppointmentStatus != "ACCEPTED_DUE_TO_REVIEW" {
		t.Errorf("Expected status ACCEPTED_DUE_TO_REVIEW, got %s", response.Rdv.AppointmentStatus)
	}
}

func TestValidateRdvRefuse(t *testing.T) {
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
	response := ValidateRdv(appointment.ID, ReviewInput{
		Reason:     "testing",
		Validation: false,
	})

	if response.Err != nil {
		t.Errorf("Error getting rdv patient: %v", response.Err)
	}

	if response.Code != 200 {
		t.Errorf("Expected code 200, got %d", response.Code)
	}
	if response.Rdv.AppointmentStatus != "CANCELED_DUE_TO_REVIEW" {
		t.Errorf("Expected status CANCELED_DUE_TO_REVIEW, got %s", response.Rdv.AppointmentStatus)
	}
}

func TestValideRdvEmpty(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}

	response := ValidateRdv("", ReviewInput{
		Reason:     "",
		Validation: false,
	})

	if response.Err == nil {
		t.Error("Unexpected null error")
	}

	if response.Code != 400 {
		t.Errorf("Expected code 400, got %d", response.Code)
	}
}

func TestValideRdvInvalid(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}

	response := ValidateRdv("invalidRdvToValidate", ReviewInput{
		Reason:     "",
		Validation: false,
	})

	if response.Err == nil {
		t.Error("Unexpected null error")
	}

	if response.Code != 400 {
		t.Errorf("Expected code 400, got %d", response.Code)
	}
}
