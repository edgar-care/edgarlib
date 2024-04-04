package auth

import (
	"context"
	"github.com/edgar-care/edgarlib/graphql"
	"github.com/joho/godotenv"
	"log"
	"testing"
)

func TestGetDoctorById(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}
	gqlClient := graphql.CreateClient()
	doctor, err := graphql.CreateDoctor(context.Background(), gqlClient, "test_get_doctor@edgar-sante.fr", "password", "name", "first", graphql.AddressInput{"", "", "", ""})
	if err != nil {
		t.Errorf("Error while creating doctor: %v", err)
	}

	response := GetDoctorById(doctor.CreateDoctor.Id)

	if response.Err != nil {
		t.Errorf("Unexpected error: %v", response.Err)
	}

	if response.Code != 200 {
		t.Errorf("Expected code 200, got %d", response.Code)
	}

	if response.Doctor.ID != doctor.CreateDoctor.Id {
		t.Errorf("Expected appointment ID %s, got %s", doctor.CreateDoctor.Id, response.Doctor.ID)
	}
}

func TestGetDoctorByIdInvalid(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}
	response := GetDoctorById("invalid")

	if response.Err == nil {
		t.Error("Unexpected null error")
	}

	if response.Code != 400 {
		t.Errorf("Expected code 400, got %d", response.Code)
	}
}

func TestGetDoctors(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}
	response := GetDoctors()
	if response.Err != nil {
		t.Errorf("Unexpected error: %v", response.Err)
	}

	if response.Code != 200 {
		t.Errorf("Expected code 200, got %d", response.Code)
	}

	if len(response.Doctors) == 0 {
		t.Errorf("Expected list of doctors but got none")
	}
}
