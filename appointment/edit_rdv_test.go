package appointment

import (
	"github.com/edgar-care/edgarlib/v2/graphql/model"
	"github.com/joho/godotenv"
	"log"
	"testing"

	"github.com/edgar-care/edgarlib/v2/graphql"
)

func TestEditRdv(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}

	patient, err := graphql.CreatePatient(model.CreatePatientInput{
		Email:    "edit_rdv@edgar-sante.fr",
		Password: "password",
		Status:   true,
	})
	if err != nil {
		t.Errorf("Error while creating patient: %v", err)
	}
	appointment, err := graphql.CreateRdv(model.CreateRdvInput{
		IDPatient:         patient.ID,
		DoctorID:          "doctorID",
		StartDate:         0,
		EndDate:           100,
		AppointmentStatus: "WAITING_FOR_REVIEW",
		SessionID:         "",
	})

	if err != nil {
		t.Errorf("Error while creating appointment: %v", err)
	}
	newAppointment, err := graphql.CreateRdv(model.CreateRdvInput{
		IDPatient:         "",
		DoctorID:          "newDoctor",
		StartDate:         0,
		EndDate:           10,
		AppointmentStatus: "WAITING_FOR_REVIEW",
		SessionID:         "",
	})
	if err != nil {
		t.Errorf("Error while creating appointment: %v", err)
	}

	response := EditRdv(newAppointment.ID, appointment.ID, patient.ID)

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

	appointment, err := graphql.CreateRdv(model.CreateRdvInput{
		IDPatient:         "booked",
		DoctorID:          "doctorId",
		StartDate:         0,
		EndDate:           100,
		AppointmentStatus: "WAITING_FOR_REVIEW",
		SessionID:         "",
	})
	if err != nil {
		t.Errorf("Error while creating appointment: %v", err)
	}
	response := EditRdv(appointment.ID, "test", "")
	if response.Code != 400 || response.Err == nil {
		t.Errorf("Expected error and code 400 but got code %d and err: %s", response.Code, response.Err.Error())
	}
}

func TestEditRdvInvalidSecondId(t *testing.T) {

	patient, err := graphql.CreatePatient(model.CreatePatientInput{
		Email:    "edit_rdv_invalid@edgar-sante.fr",
		Password: "password",
		Status:   true,
	})

	if err != nil {
		t.Errorf("Error while creating patient: %v", err)
	}
	appointment, err := graphql.CreateRdv(model.CreateRdvInput{
		IDPatient:         "",
		DoctorID:          "doctorId",
		StartDate:         0,
		EndDate:           100,
		AppointmentStatus: "WAITING_FOR_REVIEW",
		SessionID:         "",
	})
	if err != nil {
		t.Errorf("Error while creating appointment: %v", err)
	}
	response := EditRdv(appointment.ID, "invalidId", patient.ID)
	if response.Code != 400 || response.Err == nil {
		t.Errorf("Expected error and code 400 but got code %d and err: %s", response.Code, response.Err.Error())
	}
}

func TestEditRdvInvalidPatientId(t *testing.T) {

	appointment, err := graphql.CreateRdv(model.CreateRdvInput{
		IDPatient:         "invalidPatientId",
		DoctorID:          "doctorId",
		StartDate:         0,
		EndDate:           100,
		AppointmentStatus: "WAITING_FOR_REVIEW",
		SessionID:         "",
	})
	if err != nil {
		t.Errorf("Error while creating appointment: %v", err)
	}
	newAppointment, err := graphql.CreateRdv(model.CreateRdvInput{
		IDPatient:         "",
		DoctorID:          "newDoctor",
		StartDate:         0,
		EndDate:           100,
		AppointmentStatus: "WAITING_FOR_REVIEW",
		SessionID:         "",
	})
	if err != nil {
		t.Errorf("Error while creating appointment: %v", err)
	}
	response := EditRdv(newAppointment.ID, appointment.ID, "invalidPatientId")
	if response.Code != 400 || response.Err == nil {
		t.Errorf("Expected error and code 400 but got code %d and err: %s", response.Code, response.Err.Error())
	}
}
