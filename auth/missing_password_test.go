package auth

import (
	"testing"

	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
	"github.com/edgar-care/edgarlib/v2/redis"
)

func TestMissingPassword(t *testing.T) {
	email := "testuser_missing_password@example.com"

	patient, err := graphql.CreatePatient(model.CreatePatientInput{
		Email:    email,
		Password: "initial_password",
		Status:   true,
	})
	if err != nil {
		t.Errorf("Failed to create patient: %v", err)
	}

	response := MissingPassword(email)

	if response.Code != 200 {
		t.Errorf("Expected response code 200, got %v", response.Code)
	}
	if response.Err != nil {
		t.Errorf("Unexpected error: %v", response.Err)
	}

	_, err = redis.GetKey(patient.ID)
	if err != nil {
		t.Errorf("Failed to get key from Redis: %v", err)
	}

}

func TestMissingPasswordWithNonExistentEmail(t *testing.T) {
	email := "test_missing_password_nonexistentuser@example.com"

	response := MissingPassword(email)

	if response.Code != 400 {
		t.Errorf("Expected response code 400, got %v", response.Code)
	}
	if response.Err == nil {
		t.Errorf("Expected an error in response, got nil")
	} else if response.Err.Error() != "no account corresponds to this email" {
		t.Errorf("Expected 'no account corresponds to this email' error, got %v", response.Err.Error())
	}
}

func TestMissingPasswordDoctor(t *testing.T) {
	email := "testdoctor_missing_password@example.com"

	doctor, err := graphql.CreateDoctor(model.CreateDoctorInput{
		Email:     email,
		Password:  "haspassword",
		Name:      "name",
		Firstname: "first",
		Address: &model.AddressInput{
			Street:  "",
			ZipCode: "",
			Country: "",
			City:    "",
		},
		Status: true,
	})
	response := MissingPassword(email)

	if response.Code != 200 {
		t.Errorf("Expected response code 200, got %v", response.Code)
	}
	if response.Err != nil {
		t.Errorf("Unexpected error: %v", response.Err)
	}

	_, err = redis.GetKey(doctor.ID)
	if err != nil {
		t.Errorf("Failed to get key from Redis: %v", err)
	}

}
