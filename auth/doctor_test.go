package auth

import (
	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
	"github.com/joho/godotenv"
	"log"
	"testing"
)

func TestGetDoctorById(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}
	doctor, err := graphql.CreateDoctor(model.CreateDoctorInput{
		Email:     "test_get_doctor@edgar-sante.fr",
		Password:  "password",
		Name:      "name",
		Firstname: "first",
		Address: &model.AddressInput{
			Street:  "street",
			City:    "city",
			ZipCode: "zip",
			Country: "country",
		},
	})
	if err != nil {
		t.Errorf("Error while creating doctor: %v", err)
	}

	response := GetDoctorById(doctor.ID)

	if response.Err != nil {
		t.Errorf("Unexpected error: %v", response.Err)
	}

	if response.Code != 200 {
		t.Errorf("Expected code 200, got %d", response.Code)
	}

	if response.Doctor.ID != doctor.ID {
		t.Errorf("Expected appointment ID %s, got %s", doctor.ID, response.Doctor.ID)
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
	response := GetDoctors(0, 0)
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
