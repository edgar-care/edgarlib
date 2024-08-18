package auth

import (
	"testing"

	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
	"github.com/edgar-care/edgarlib/v2/redis"
)

func TestMissingPassword(t *testing.T) {
	email := "testuser_missing_password@example.com"

	_, err := graphql.CreatePatient(model.CreatePatientInput{
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

	_, err = redis.GetKey(email)
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
	} else if response.Err.Error() != "no patient corresponds to this email" {
		t.Errorf("Expected 'no patient corresponds to this email' error, got %v", response.Err.Error())
	}
}
