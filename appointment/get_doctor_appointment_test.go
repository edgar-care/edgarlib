package appointment

import (
	"github.com/edgar-care/edgarlib/graphql/model"
	"github.com/joho/godotenv"
	"log"
	"testing"

	"github.com/edgar-care/edgarlib/graphql"
)

func TestGetDoctorAppointment(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}

	doctorID := "doctorId"

	appointment, err := graphql.CreateRdv(model.CreateRdvInput{
		IDPatient:         "",
		DoctorID:          "doctorId",
		StartDate:         0,
		EndDate:           10,
		AppointmentStatus: "WAITING_FOR_REVIEW",
		SessionID:         "",
	})
	if err != nil {
		t.Errorf("Error creating appointment: %v", err)
	}

	response := GetDoctorAppointment(appointment.ID, doctorID)

	if response.Err != nil {
		t.Errorf("Unexpected error: %v", response.Err)
	}

	if response.Code != 200 {
		t.Errorf("Expected code 200, got %d", response.Code)
	}

	if response.Appointment.ID != appointment.ID {
		t.Errorf("Expected appointment ID %s, got %s", appointment.ID, response.Appointment.ID)
	}
}

func TestGetDoctorAppointmentInvalid(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}

	response := GetDoctorAppointment("invalid", "doctorID")

	if response.Err == nil {
		t.Error("Unexpected null error")
	}

	if response.Code != 400 {
		t.Errorf("Expected code 400, got %d", response.Code)
	}
}

func TestGetDoctorAppointmentUnauthorized(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}

	appointment, err := graphql.CreateRdv(model.CreateRdvInput{
		IDPatient:         "",
		DoctorID:          "doctorId",
		StartDate:         0,
		EndDate:           10,
		AppointmentStatus: "WAITING_FOR_REVIEW",
		SessionID:         "",
	})
	if err != nil {
		t.Errorf("Error creating appointment: %v", err)
	}

	response := GetDoctorAppointment(appointment.ID, "invalid")

	if response.Err == nil {
		t.Error("Unexpected null error")
	}

	if response.Code != 403 {
		t.Errorf("Expected code 400, got %d", response.Code)
	}
}

func TestGetAllDoctorAppointment(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}

	doctor, err := graphql.CreateDoctor(model.CreateDoctorInput{
		Email:     "test_get_doctor_appointment@edgar-sante.fr",
		Password:  "password",
		Name:      "name",
		Firstname: "first",
		Address:   &model.AddressInput{"", "", "", ""},
		Status:    true,
	})
	if err != nil {
		t.Errorf("Error while creating doctor: %v", err)
	}
	doctorID := doctor.ID
	_, err = graphql.CreateRdv(model.CreateRdvInput{
		IDPatient:         "",
		DoctorID:          doctorID,
		StartDate:         0,
		EndDate:           10,
		AppointmentStatus: "WAITING_FOR_REVIEW",
		SessionID:         "",
	})
	if err != nil {
		t.Errorf("Error while creating appointment: %v", err)
	}

	response := GetAllDoctorAppointment(doctorID)

	if response.Err != nil {
		t.Errorf("Unexpected error: %v", response.Err)
	}

	if response.Code != 201 {
		t.Errorf("Expected code 201, got %d", response.Code)
	}

	if len(response.Slots) == 0 {
		t.Errorf("Expected non-empty slots list, got empty list")
	}
}

func TestGetAllDoctorAppointmentInvalid(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}

	response := GetAllDoctorAppointment("invalid")

	if response.Err == nil {
		t.Error("Unexpected null error")
	}

	if response.Code != 400 {
		t.Errorf("Expected code 400, got %d", response.Code)
	}
}
