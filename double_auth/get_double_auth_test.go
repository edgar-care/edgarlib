package double_auth

import (
	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
	"testing"
)

func TestGetDoubleAuthById_Succes(t *testing.T) {
	patient, err := graphql.CreatePatient(model.CreatePatientInput{
		Email:    "test_get_double_auth_success@example.com",
		Password: "password",
		Status:   true,
	})
	if err != nil {
		t.Fatalf("failed to create patient: %s", err)
	}
	input := CreateDoubleAuthInput{Methods: "EMAIL"}
	_ = CreateDoubleAuthEmail(input, patient.ID)

	response := GetDoubleAuthById(patient.ID)
	if response.Err != nil {
		t.Errorf("Expected no error, got: %v", response.Err)
	}
	if response.Code != 200 {
		t.Errorf("Expected status code 200, got: %d", response.Code)
	}
	if response.DoubleAuth.ID == "" {
		t.Errorf("Expected a valid DeviceConnect ID, got an empty string")
	}
}

func TestGetDoubleAuthByIdDoctor_Succes(t *testing.T) {
	doctor, err := graphql.CreateDoctor(model.CreateDoctorInput{
		Email:     "test_get_double_auth_doctor_succes@example.com",
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

	input := CreateDoubleAuthInput{Methods: "EMAIL"}
	_ = CreateDoubleAuthEmail(input, doctor.ID)

	response := GetDoubleAuthById(doctor.ID)
	if response.Err != nil {
		t.Errorf("Expected no error, got: %v", response.Err)
	}
	if response.Code != 200 {
		t.Errorf("Expected status code 200, got: %d", response.Code)
	}
	if response.DoubleAuth.ID == "" {
		t.Errorf("Expected a valid DeviceConnect ID, got an empty string")
	}
}

func TestGetDoubleAuthByIdInvalid(t *testing.T) {
	invalidDeviceId := "invalid_device_id"

	response := GetDoubleAuthById(invalidDeviceId)

	if response.Err == nil {
		t.Errorf("Expected an error, got none")
	}
	if response.Code != 400 {
		t.Errorf("Expected status code 400, got: %d", response.Code)
	}
}
