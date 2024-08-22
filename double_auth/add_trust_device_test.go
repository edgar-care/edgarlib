package double_auth

import (
	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
	"github.com/joho/godotenv"
	"log"
	"testing"
)

func TestAddTrustDevice_Patient_Success(t *testing.T) {
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

	patient, err := graphql.CreatePatient(model.CreatePatientInput{
		Email:    "test_patient_create_device_trust_success@edgar-sante.fr",
		Password: "password",
		Status:   true,
	})
	if err != nil {
		t.Errorf("Error while creating patient: %v", err)
	}

	device := CreateDeviceConnect(input, patient.ID)

	if device.Err != nil {
		t.Errorf("Expected no error, got: %v", device.Err)
	}
	if device.Code != 201 {
		t.Errorf("Expected status code 201, got: %d", device.Code)
	}
	_ = CreateDoubleAuthAppTier(patient.ID, "1234")

	response := AddTrustDevice(device.DeviceConnect.ID, patient.ID)

	if response.Err != nil {
		t.Errorf("Expected no error, got: %v", response.Err)
	}
	if response.Code != 200 {
		t.Errorf("Expected status code 200, got: %d", response.Code)
	}
	if response.Patient == nil {
		t.Errorf("Expected a patient, got nil")
	}
	if response.Doctor != nil {
		t.Errorf("Expected no doctor, got: %v", response.Doctor)
	}
}

func TestAddTrustDevice_Doctor_Success(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}
	input := CreateDeviceConnectInput{
		DeviceName: "TestDevice",
		Ip:         "192.168.0.1",
		City:       "testCity",
		Country:    "testcountry",
		Date:       1627880400, // Timestamp pour une date fixe
	}

	doctor, err := graphql.CreateDoctor(model.CreateDoctorInput{
		Email:     "test_doctor_create_device_trust@edgar-sante.fr",
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

	device := CreateDeviceConnect(input, doctor.ID)

	if device.Err != nil {
		t.Errorf("Expected no error, got: %v", device.Err)
	}
	if device.Code != 201 {
		t.Errorf("Expected status code 201, got: %d", device.Code)
	}

	_ = CreateDoubleAuthAppTier(doctor.ID, "1234")

	response := AddTrustDevice(device.DeviceConnect.ID, doctor.ID)

	// Vérifier les résultats
	if response.Err != nil {
		t.Errorf("Expected no error, got: %v", response.Err)
	}
	if response.Code != 200 {
		t.Errorf("Expected status code 200, got: %d", response.Code)
	}
	if response.Doctor == nil {
		t.Errorf("Expected a doctor, got nil")
	}
	if response.Patient != nil {
		t.Errorf("Expected no patient, got: %v", response.Patient)
	}
}

func TestAddTrustDevice_InvalidUser(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}
	idDevice := "test_device_id"
	idUser := "invalid_user_id"

	// Appeler la fonction
	response := AddTrustDevice(idDevice, idUser)

	if response.Err == nil {
		t.Errorf("Expected error but got none")
	}

}
