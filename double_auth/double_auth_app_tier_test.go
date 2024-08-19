package double_auth

import (
	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
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

	tier := CreateDoubleAuthTierInput{Methods: "AUTHENTIFICATOR", Code: "21234"}
	response := CreateDoubleAuthAppTier(tier, "url", patient.ID)

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

	tier := CreateDoubleAuthTierInput{Methods: "AUTHENTIFICATOR", Code: "21234"}
	response := CreateDoubleAuthAppTier(tier, "url", doctor.ID)

	if response.Err != nil {
		t.Errorf("Expected no error, got: %v", response.Err)
	}
	if response.Code != 201 {
		t.Errorf("Expected status code 200, got: %d", response.Code)
	}
}

func TestCreateDoubleAuthAppTier_invalidAccount(t *testing.T) {

	tier := CreateDoubleAuthTierInput{Methods: "AUTHENTIFICATOR", Code: "21234"}
	response := CreateDoubleAuthAppTier(tier, "url", "test_invalid_id")

	if response.Err == nil {
		t.Errorf("Expected an error, got none")
	}
	if response.Code != 400 {
		t.Errorf("Expected status code 400, got: %d", response.Code)
	}
}

func TestCreateDoubleAuthAppTier_invalidMethods(t *testing.T) {

	patient, err := graphql.CreatePatient(model.CreatePatientInput{
		Email:    "test_add_double_auth_apptier_invalid_method@example.com",
		Password: "password",
		Status:   true,
	})
	if err != nil {
		t.Fatalf("Failed to create patient: %s", err)
	}

	tier := CreateDoubleAuthTierInput{Methods: "TEST", Code: "21234"}
	response := CreateDoubleAuthAppTier(tier, "url", patient.ID)

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

	tier := CreateDoubleAuthTierInput{Methods: "AUTHENTIFICATOR", Code: "21234"}
	response := CreateDoubleAuthAppTier(tier, "url", patient.ID)

	if response.Err != nil {
		t.Errorf("Expected no error, got: %v", response.Err)
	}
	if response.Code != 200 {
		t.Errorf("Expected status code 200, got: %d", response.Code)
	}
}

func TestCreateDoubleAuthAppTier_ADDInvalidMethods(t *testing.T) {
	patient, err := graphql.CreatePatient(model.CreatePatientInput{
		Email:    "test_add_double_auth_apptier_invalid_methods@example.com",
		Password: "password",
		Status:   true,
	})
	if err != nil {
		t.Fatalf("Failed to create patient: %s", err)
	}

	email := CreateDoubleAuthInput{Methods: "EMAIL"}
	_ = CreateDoubleAuthEmail(email, patient.ID)

	tier := CreateDoubleAuthTierInput{Methods: "TEST", Code: "21234"}
	response := CreateDoubleAuthAppTier(tier, "url", patient.ID)

	if response.Err == nil {
		t.Errorf("Expected an error, got none")
	}
	if response.Code != 400 {
		t.Errorf("Expected status code 400, got: %d", response.Code)
	}
}
