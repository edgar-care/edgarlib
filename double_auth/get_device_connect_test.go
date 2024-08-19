package double_auth

import (
	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
	"github.com/joho/godotenv"
	"log"
	"testing"
)

func TestGetDeviceConnectById_Success(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}

	input := CreateDeviceConnectInput{
		DeviceName: "TestDevice",
		Ip:         "192.168.0.1",
		Latitude:   48.8566,
		Longitude:  2.3522,
		Date:       1627880400, // Timestamp pour une date fixe
	}
	patient, err := graphql.CreatePatient(model.CreatePatientInput{
		Email:    "test_patient_get_device_connect_unique_success@edgar-sante.fr",
		Password: "password",
		Status:   true,
	})
	if err != nil {
		t.Errorf("Error while creating patient: %v", err)
	}

	device := CreateDeviceConnect(input, patient.ID)

	response := GetDeviceConnectById(device.DeviceConnect.ID)
	if response.Err != nil {
		t.Errorf("Expected no error, got: %v", response.Err)
	}
	if response.Code != 200 {
		t.Errorf("Expected status code 200, got: %d", response.Code)
	}
	if response.DeviceConnect.ID == "" {
		t.Errorf("Expected a valid DeviceConnect ID, got an empty string")
	}
}

func TestGetDeviceConnectByIdDoctor_Success(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}

	input := CreateDeviceConnectInput{
		DeviceName: "TestDevice",
		Ip:         "192.168.0.1",
		Latitude:   48.8566,
		Longitude:  2.3522,
		Date:       1627880400, // Timestamp pour une date fixe
	}
	doctor, err := graphql.CreateDoctor(model.CreateDoctorInput{
		Email:     "test_get_device_connect_doctor_unique_succes@example.com",
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

	device := CreateDeviceConnect(input, doctor.ID)

	response := GetDeviceConnectById(device.DeviceConnect.ID)
	if response.Err != nil {
		t.Errorf("Expected no error, got: %v", response.Err)
	}
	if response.Code != 200 {
		t.Errorf("Expected status code 200, got: %d", response.Code)
	}
	if response.DeviceConnect.ID == "" {
		t.Errorf("Expected a valid DeviceConnect ID, got an empty string")
	}
}

func TestGetDeviceConnectById_Invalid(t *testing.T) {
	invalidDeviceId := "invalid_device_id"

	response := GetDeviceConnectById(invalidDeviceId)

	if response.Err == nil {
		t.Errorf("Expected an error, got none")
	}
	if response.Code != 400 {
		t.Errorf("Expected status code 400, got: %d", response.Code)
	}
}

func TestGetDeviceConnect(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}

	input := CreateDeviceConnectInput{
		DeviceName: "TestDevice",
		Ip:         "192.168.0.1",
		Latitude:   48.8566,
		Longitude:  2.3522,
		Date:       1627880400, // Timestamp pour une date fixe
	}
	patient, err := graphql.CreatePatient(model.CreatePatientInput{
		Email:    "test_patient_get_device_connect_all_success@edgar-sante.fr",
		Password: "password",
		Status:   true,
	})
	if err != nil {
		t.Errorf("Error while creating patient: %v", err)
	}

	_ = CreateDeviceConnect(input, patient.ID)

	response := GetDeviceConnect(patient.ID)
	if response.Err != nil {
		t.Errorf("Expected no error, got: %v", response.Err)
	}
	if response.Code != 200 {
		t.Errorf("Expected status code 200, got: %d", response.Code)
	}
}

func TestGetDeviceConnectDoctor_Success(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}

	input := CreateDeviceConnectInput{
		DeviceName: "TestDevice",
		Ip:         "192.168.0.1",
		Latitude:   48.8566,
		Longitude:  2.3522,
		Date:       1627880400, // Timestamp pour une date fixe
	}
	doctor, err := graphql.CreateDoctor(model.CreateDoctorInput{
		Email:     "test_get_device_connect_doctor_all_succes@example.com",
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

	_ = CreateDeviceConnect(input, doctor.ID)

	response := GetDeviceConnect(doctor.ID)
	if response.Err != nil {
		t.Errorf("Expected no error, got: %v", response.Err)
	}
	if response.Code != 200 {
		t.Errorf("Expected status code 200, got: %d", response.Code)
	}
}

func TestGetDeviceConnect_Invalid(t *testing.T) {
	invalidDeviceId := "invalid_device_id"

	response := GetDeviceConnect(invalidDeviceId)

	if response.Err == nil {
		t.Errorf("Expected an error, got none")
	}
	if response.Code != 400 {
		t.Errorf("Expected status code 400, got: %d", response.Code)
	}
}
