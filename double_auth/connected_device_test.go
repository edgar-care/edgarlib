package double_auth

import (
	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
	"github.com/joho/godotenv"
	"log"
	"testing"
)

func TestCreateDeviceConnect_Success(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}

	input := CreateDeviceConnectInput{
		DeviceName: "TestDevice",
		Ip:         "192.168.0.1",
		City:       "testcity",
		Country:    "testcountry",
		Date:       1627880400, // Timestamp pour une date fixe
	}
	patient, err := graphql.CreatePatient(model.CreatePatientInput{
		Email:    "test_patient_create_device_connect_success@edgar-sante.fr",
		Password: "password",
		Status:   true,
	})
	if err != nil {
		t.Errorf("Error while creating patient: %v", err)
	}

	response := CreateDeviceConnect(input, patient.ID)

	if response.Err != nil {
		t.Errorf("Expected no error, got: %v", response.Err)
	}
	if response.Code != 201 {
		t.Errorf("Expected status code 201, got: %d", response.Code)
	}
	if response.DeviceConnect.ID == "" {
		t.Errorf("Expected a valid DeviceConnect ID, got an empty string")
	}
}

func TestCreateDeviceConnect_InvalidPatientId(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}

	input := CreateDeviceConnectInput{
		DeviceName: "TestDevice",
		Ip:         "192.168.0.1",
		City:       "testCity",
		Country:    "testcountry",
		Date:       1627880400,
	}
	invalidPatientId := "invalid_patient_id"

	// Appel de la fonction
	response := CreateDeviceConnect(input, invalidPatientId)

	// Vérification des résultats
	if response.Err == nil {
		t.Errorf("Expected an error, got none")
	}
	if response.Code != 400 {
		t.Errorf("Expected status code 400, got: %d", response.Code)
	}
	if response.DeviceConnect.ID != "" {
		t.Errorf("Expected no DeviceConnect to be created, but got ID: %s", response.DeviceConnect.ID)
	}

}

func TestCreateDeviceConnect_FailedDeviceCreation(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}

	input := CreateDeviceConnectInput{
		DeviceName: "", // Champ obligatoire manquant ou invalide
		Ip:         "192.168.0.1",
		City:       "testcity",
		Country:    "testcountry",
		Date:       1627880400,
	}
	patientId := "valid_patient_id"

	// Appel de la fonction
	response := CreateDeviceConnect(input, patientId)

	// Vérification des résultats
	if response.Err == nil {
		t.Errorf("Expected an error, got none")
	}
	if response.Code != 400 {
		t.Errorf("Expected status code 400, got: %d", response.Code)
	}
	if response.DeviceConnect.ID != "" {
		t.Errorf("Expected no DeviceConnect to be created, but got ID: %s", response.DeviceConnect.ID)
	}
}

func TestCreateDeviceConnect_SuccessDoctor(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}

	input := CreateDeviceConnectInput{
		DeviceName: "TestDevice",
		Ip:         "192.168.0.1",
		City:       "testcity",
		Country:    "testcountry",
		Date:       1627880400,
	}
	doctor, err := graphql.CreateDoctor(model.CreateDoctorInput{
		Email:     "test_doctor_create_device@edgar-sante.fr",
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
		t.Errorf("Error while creating patient: %v", err)
	}

	response := CreateDeviceConnect(input, doctor.ID)

	if response.Err != nil {
		t.Errorf("Expected no error, got: %v", response.Err)
	}
	if response.Code != 201 {
		t.Errorf("Expected status code 201, got: %d", response.Code)
	}
	if response.DeviceConnect.ID == "" {
		t.Errorf("Expected a valid DeviceConnect ID, got an empty string")
	}
}
