package appointment

import (
	"github.com/edgar-care/edgarlib/v2/graphql/model"
	"github.com/joho/godotenv"
	"log"
	"testing"

	"github.com/edgar-care/edgarlib/v2/graphql"
)

func TestUpdateDoctorAppointment(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}

	patient, err := graphql.CreatePatient(model.CreatePatientInput{
		Email:    "update_doctor_appointment@edgar-sante.fr",
		Password: "password",
	})
	if err != nil {
		t.Errorf("Error while creating patient: %v", err)
	}
	appointment, err := graphql.CreateRdv(model.CreateRdvInput{
		IDPatient:         patient.ID,
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
		EndDate:           10,
		AppointmentStatus: "WAITING_FOR_REVIEW",
		SessionID:         "",
	})
	if err != nil {
		t.Errorf("Error while creating appointment: %v", err)
	}

	response := UpdateDoctorAppointment(newAppointment.ID, appointment.ID)

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
	appointment, err := graphql.CreateRdv(model.CreateRdvInput{"booked", "doctorId", 0, 100, "WAITING_FOR_REVIEW", ""})
	if err != nil {
		t.Errorf("Error while creating appointment: %v", err)
	}
	response := UpdateDoctorAppointment(appointment.ID, "test")
	if response.Code != 400 || response.Err == nil {
		t.Errorf("Expected error and code 400 but got code %d and err: %s", response.Code, response.Err.Error())
	}
}

func TestUpdateDoctorAppointmentInvalidSecondId(t *testing.T) {
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
	response := UpdateDoctorAppointment(appointment.ID, "invalidId")
	if response.Code != 400 || response.Err == nil {
		t.Errorf("Expected error and code 400 but got code %d and err: %s", response.Code, response.Err.Error())
	}
}

func TestUpdateDoctorAppointmentInvalidPatientId(t *testing.T) {
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
		EndDate:           10,
		AppointmentStatus: "WAITING_FOR_REVIEW",
		SessionID:         "",
	})
	if err != nil {
		t.Errorf("Error while creating appointment: %v", err)
	}
	response := UpdateDoctorAppointment(newAppointment.ID, appointment.ID)
	if response.Code != 400 || response.Err == nil {
		t.Errorf("Expected error and code 400 but got code %d and err: %s", response.Code, response.Err.Error())
	}
}
