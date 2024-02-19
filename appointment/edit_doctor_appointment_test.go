package appointment

import (
	"context"
	"github.com/joho/godotenv"
	"log"
	"testing"

	"github.com/edgar-care/edgarlib/graphql"
)

func TestUpdateDoctorAppointment(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}
	gqlClient := graphql.CreateClient()

	patient, err := graphql.CreatePatient(context.Background(), gqlClient, "update_doctor_appointment@edgar-sante.fr", "password")
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

	response := UpdateDoctorAppointment(newAppointment.CreateRdv.Id, appointment.CreateRdv.Id)

	if response.Err != nil {
		t.Errorf("Unexpected error: %v", response.Err)
	}

	if response.Code != 200 {
		t.Errorf("Expected code 200, got %d", response.Code)
	}
}

func TestUpdateDoctorAppointmentEmptyId(t *testing.T) {
	response := UpdateDoctorAppointment("", "")
	if response.Code != 400 || response.Err == nil {
		t.Errorf("Expected error and code 400 but got code %d and err: %s", response.Code, response.Err.Error())
	}
}

func TestUpdateDoctorAppointmentInvalidId(t *testing.T) {
	response := UpdateDoctorAppointment("invalidId", "invalidId")
	if response.Code != 400 || response.Err == nil {
		t.Errorf("Expected error and code 400 but got code %d and err: %s", response.Code, response.Err.Error())
	}
}

func TestUpdateDoctorAppointmentAlreadyBooked(t *testing.T) {
	gqlClient := graphql.CreateClient()
	appointment, err := graphql.CreateRdv(context.Background(), gqlClient, "booked", "doctorId", 0, 100, "WAITING_FOR_REVIEW", "")
	if err != nil {
		t.Errorf("Error while creating appointment: %v", err)
	}
	response := UpdateDoctorAppointment(appointment.CreateRdv.Id, "test")
	if response.Code != 400 || response.Err == nil {
		t.Errorf("Expected error and code 400 but got code %d and err: %s", response.Code, response.Err.Error())
	}
}

func TestUpdateDoctorAppointmentInvalidSecondId(t *testing.T) {
	gqlClient := graphql.CreateClient()
	appointment, err := graphql.CreateRdv(context.Background(), gqlClient, "", "doctorId", 0, 100, "WAITING_FOR_REVIEW", "")
	if err != nil {
		t.Errorf("Error while creating appointment: %v", err)
	}
	response := UpdateDoctorAppointment(appointment.CreateRdv.Id, "invalidId")
	if response.Code != 400 || response.Err == nil {
		t.Errorf("Expected error and code 400 but got code %d and err: %s", response.Code, response.Err.Error())
	}
}

func TestUpdateDoctorAppointmentInvalidPatientId(t *testing.T) {
	gqlClient := graphql.CreateClient()
	appointment, err := graphql.CreateRdv(context.Background(), gqlClient, "invalidPatientId", "doctorId", 0, 100, "WAITING_FOR_REVIEW", "")
	if err != nil {
		t.Errorf("Error while creating appointment: %v", err)
	}
	newAppointment, err := graphql.CreateRdv(context.Background(), gqlClient, "", "newDoctor", 0, 10, "WAITING_FOR_REVIEW", "")
	if err != nil {
		t.Errorf("Error while creating appointment: %v", err)
	}
	response := UpdateDoctorAppointment(newAppointment.CreateRdv.Id, appointment.CreateRdv.Id)
	if response.Code != 400 || response.Err == nil {
		t.Errorf("Expected error and code 400 but got code %d and err: %s", response.Code, response.Err.Error())
	}
}
