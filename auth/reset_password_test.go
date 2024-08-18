package auth

import (
	"testing"

	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
	"github.com/edgar-care/edgarlib/v2/redis"
	"github.com/google/uuid"
)

func TestResetPassword(t *testing.T) {
	email := "testuser_reset_password@example.com"
	password := "new_password"
	patientUUID := uuid.New()

	_, err := redis.SetKey(patientUUID.String(), email, nil)
	if err != nil {
		t.Errorf("Failed to set patient key: %v", err)
	}

	_, err = graphql.CreatePatient(model.CreatePatientInput{
		Email:    email,
		Password: "first_password",
		Status:   true,
	})
	if err != nil {
		t.Errorf("Failed to create patient: %v", err)
	}

	response := ResetPassword(email, password, patientUUID.String())

	if response.Code != 200 {
		t.Errorf("Expected response code 200, got %v", response.Code)
	}
	if response.Err != nil {
		t.Errorf("Unexpected error: %v", response.Err)
	}

	updatedPatient, err := graphql.GetPatientByEmail(email)
	if err != nil {
		t.Errorf("Unable to retrieve patient: %v", err)
	}

	if updatedPatient.Password == "first_password" {
		t.Errorf("password was not updated")
	}
}

func TestResetPasswordWithInvalidUUID(t *testing.T) {
	email := "testuser@example.com"
	password := "newpassword123"
	invalidUUID := "invalid-uuid"

	response := ResetPassword(email, password, invalidUUID)

	if response.Code != 403 {
		t.Errorf("Expected response code 403, got %v", response.Code)
	}
	if response.Err == nil {
		t.Errorf("Expected an error in response")
	} else if response.Err.Error() != "uuid is expired" {
		t.Errorf("Expected 'uuid is expired' error, got %v", response.Err.Error())
	}
}

func TestResetPasswordWithInvalidEmail(t *testing.T) {
	email := "invaliduser@example.com"
	password := "newpassword123"

	patientUUID := uuid.New()

	uuid, err := redis.SetKey(patientUUID.String(), email, nil)
	if err != nil {
		t.Errorf("Failed to set patient key: %v", err)
	}

	response := ResetPassword(email, password, uuid)

	if response.Code != 403 {
		t.Errorf("Expected response code 403, got %v", response.Code)
	}
	if response.Err == nil {
		t.Errorf("Expected an error in response")
	}
}

func TestResetPasswordWithoutUUID(t *testing.T) {
	email := "testuser@example.com"
	password := "newpassword123"
	uuid := ""

	response := ResetPassword(email, password, uuid)

	if response.Code != 403 {
		t.Errorf("Expected response code 403, got %v", response.Code)
	}
	if response.Err == nil {
		t.Errorf("Expected an error in response")
	} else if response.Err.Error() != "uuid has to be provided" {
		t.Errorf("Expected 'uuid has to be provided' error, got %v", response.Err.Error())
	}
}
