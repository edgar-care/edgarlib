package double_auth

import (
	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
	"github.com/joho/godotenv"
	"log"
	"testing"
)

func TestCreateDoubleAuthAppTier_Success(t *testing.T) {
	patient, err := graphql.CreatePatient(model.CreatePatientInput{
		Email:    "test_add_double_auth_tierapp_success@example.com",
		Password: "password",
		Status:   true,
	})
	if err != nil {
		t.Fatalf("Failed to create patient: %s", err)
	}

	response := CreateDoubleAuthAppTier(patient.ID, "12345")

	if response.Err != nil {
		t.Errorf("Expected no error, got: %v", response.Err)
	}
	if response.Code != 201 {
		t.Errorf("Expected status code 200, got: %d", response.Code)
	}
}

func TestCreateDoubleAuthAppTierDoctor_Success(t *testing.T) {
	doctor, err := graphql.CreateDoctor(model.CreateDoctorInput{
		Email:     "test_create_double_tierapp_mobile_succes@example.com",
		Password:  "password",
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
	if err != nil {
		t.Fatalf("Failed to create doctor: %s", err)
	}

	response := CreateDoubleAuthAppTier(doctor.ID, "1234")

	if response.Err != nil {
		t.Errorf("Expected no error, got: %v", response.Err)
	}
	if response.Code != 201 {
		t.Errorf("Expected status code 200, got: %d", response.Code)
	}
}

func TestCreateDoubleAuthAppTier_invalidAccount(t *testing.T) {

	response := CreateDoubleAuthAppTier("test_invalid_id", "1233")

	if response.Err == nil {
		t.Errorf("Expected an error, got none")
	}
	if response.Code != 400 {
		t.Errorf("Expected status code 400, got: %d", response.Code)
	}
}

func TestCreateDoubleAuthAppTier_ADDSuccess(t *testing.T) {
	patient, err := graphql.CreatePatient(model.CreatePatientInput{
		Email:    "test_add_double_auth_apptier_success_method_add@example.com",
		Password: "password",
		Status:   true,
	})
	if err != nil {
		t.Fatalf("Failed to create patient: %s", err)
	}

	email := CreateDoubleAuthInput{Methods: "EMAIL"}
	_ = CreateDoubleAuthEmail(email, patient.ID)

	response := CreateDoubleAuthAppTier(patient.ID, "1234")

	if response.Err != nil {
		t.Errorf("Expected no error, got: %v", response.Err)
	}
	if response.Code != 200 {
		t.Errorf("Expected status code 200, got: %d", response.Code)
	}
}

func TestGetSecretThirdParty(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}

	patient, err := graphql.CreatePatient(model.CreatePatientInput{
		Email:    "test_add_double_auth_apptier_get_secret_success@example.com",
		Password: "password",
		Status:   true,
	})
	if err != nil {
		t.Fatalf("Failed to create patient: %s", err)
	}

	ThirdParty := CreateDoubleAuthAppTier(patient.ID, "12345")
	if ThirdParty.Err != nil {
		t.Errorf("Expected no error, got: %v", ThirdParty.Err)
	}
	response := GetSecretThirdParty(patient.ID)

	if response.Err != nil {
		t.Errorf("Expected no error, got: %v", response.Err)
	}
	if response.Code != 200 {
		t.Errorf("Expected status code 200, got: %d", response.Code)
	}

}

func TestGetSecretThirdPartyInvalidID(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}

	response := GetSecretThirdParty("test_invalid_id")
	if response.Err == nil {
		t.Errorf("Expected an error, got none")
	}
	if response.Code != 400 {
		t.Errorf("Expected status code 400, got: %d", response.Code)
	}

}

func TestGetSecretThirdPartyInvalidDoubleAuth(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}

	patient, err := graphql.CreatePatient(model.CreatePatientInput{
		Email:    "test_add_double_auth_apptier_get_secret_failed@example.com",
		Password: "password",
		Status:   true,
	})
	if err != nil {
		t.Fatalf("Failed to create patient: %s", err)
	}

	response := GetSecretThirdParty(patient.ID)
	if response.Err == nil {
		t.Errorf("Expected an error, got none")
	}
	if response.Code != 400 {
		t.Errorf("Expected status code 400, got: %d", response.Code)
	}

}
