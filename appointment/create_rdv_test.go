package appointment

import (
	"context"
	"log"
	"testing"

	"github.com/edgar-care/edgarlib/graphql"
	"github.com/joho/godotenv"
)

func TestCreateRdv(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}
	gqlClient := graphql.CreateClient()
	doctor, err := graphql.CreateDoctor(context.Background(), gqlClient, "test_doctor_create_appointment@edgar-sante.fr", "password", "name", "first", graphql.AddressInput{"", "", "", ""})
	if err != nil {
		t.Errorf("Error while creating doctor: %v", err)
	}

	patient, err := graphql.CreatePatient(context.Background(), gqlClient, "test_patient_create_appointment@edgar-sante.fr", "password")
	if err != nil {
		t.Errorf("Error while creating patient: %v", err)
	}

	response := CreateRdv(patient.CreatePatient.Id, doctor.CreateDoctor.Id, 0, 10, "")

	if response.Code != 200 {
		t.Errorf("Expected code 200 but got %d", response.Code)
	}
	if response.Err != nil {
		t.Errorf("Expected no error but got: %s", response.Err.Error())
	}

	appointment, err := graphql.GetRdvById(context.Background(), gqlClient, response.Rdv.ID)

	if err != nil {
		t.Errorf("Error while retrieving created appointment: %s", err.Error())
	}
	if appointment.GetRdvById.Doctor_id != doctor.CreateDoctor.Id {
		t.Error("Appointment wasn't created with the correct doctor id")
	}
}

func TestCreateRdvInvalidPatient(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}
	gqlClient := graphql.CreateClient()
	doctor, err := graphql.CreateDoctor(context.Background(), gqlClient, "test_doctor_create_appointment_invalid_patient@edgar-sante.fr", "password", "name", "first", graphql.AddressInput{"", "", "", ""})
	if err != nil {
		t.Errorf("Error while creating doctor: %v", err)
	}

	response := CreateRdv("invalid", doctor.CreateDoctor.Id, 0, 10, "")

	if response.Code != 400 {
		t.Errorf("Expected code 400 but got %d", response.Code)
	}
	if response.Err == nil {
		t.Error("Expected error but got none")
	}
}

func TestCreateRdvInvalidId(t *testing.T) {
	response := CreateRdv("patientId", "invalid", 0, 10, "")
	if response.Code != 400 || response.Err == nil {
		t.Errorf("Expected error and code 400 but got code %d and err: %s", response.Code, response.Err.Error())
	}
}
