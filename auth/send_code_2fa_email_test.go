package auth

import (
	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
	"github.com/edgar-care/edgarlib/v2/redis"
	"testing"
)

func TestEmail2faAuth(t *testing.T) {
	email := "testuser_log_2fa_email_password@example.com"

	_, err := graphql.CreatePatient(model.CreatePatientInput{
		Email:    email,
		Password: "initial_password",
		Status:   true,
	})
	if err != nil {
		t.Errorf("Failed to create patient: %v", err)
	}

	response := Email2faAuth(email)

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

func TestEmail2faAuth_InvalidEmail(t *testing.T) {
	email := "test_missing_password_nonexistentuser@example.com"

	response := Email2faAuth(email)

	if response.Code != 400 {
		t.Errorf("Expected response code 400, got %v", response.Code)
	}
	if response.Err == nil {
		t.Errorf("Expected an error in response, got nil")
	} else if response.Err.Error() != "email does not correspond to a valid patient or doctor" {
		t.Errorf("Expected 'email does not correspond to a valid patient or doctor' error, got %v", response.Err.Error())
	}
}
