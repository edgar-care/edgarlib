package auth

import (
	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/joho/godotenv"
	"log"
	"testing"
)

func TestCreatePatientAccount(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}

	response := CreatePatientAccount("test_create_patient_account@edgar-sante.fr")

	if response.Code != 200 {
		t.Errorf("Expected code 200 but got %d", response.Code)
	}
	if response.Err != nil {
		t.Errorf("Expected no error but got: %s", response.Err.Error())
	}

	_, err := graphql.GetPatientById(response.Id)
	if err != nil {
		t.Errorf("Expected no error but got: %s", err.Error())
	}
}

func TestCreatePatientAccountInvalidEmail(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}

	response := CreatePatientAccount("test_create_patient_account@edgar-sante.fr")

	if response.Code != 400 || response.Err == nil {
		t.Errorf("Expected error and code 400 but got code %d and err: %s", response.Code, response.Err.Error())
	}
}
