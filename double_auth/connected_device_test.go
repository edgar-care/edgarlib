package double_auth

import (
	"testing"
)

func TestCreateDeviceConnect_Success(t *testing.T) {
	// Configuration des valeurs de test
	input := CreateDeviceConnectInput{
		DeviceName: "TestDevice",
		Ip:         "192.168.0.1",
		Latitude:   48.8566,
		Longitude:  2.3522,
		Date:       1627880400, // Timestamp pour une date fixe
	}
	patientId := "valid_patient_id"

	// Appel de la fonction
	response := CreateDeviceConnect(input, patientId)

	// Vérification des résultats
	if response.Err != nil {
		t.Errorf("Expected no error, got: %v", response.Err)
	}
	if response.Code != 201 {
		t.Errorf("Expected status code 201, got: %d", response.Code)
	}
	if response.DeviceConnect.ID == "" {
		t.Errorf("Expected a valid DeviceConnect ID, got an empty string")
	}
	if response.Patient.ID != patientId {
		t.Errorf("Expected Patient ID %s, got: %s", patientId, response.Patient.ID)
	}
}

func TestCreateDeviceConnect_InvalidPatientId(t *testing.T) {
	// Configuration des valeurs de test
	input := CreateDeviceConnectInput{
		DeviceName: "TestDevice",
		Ip:         "192.168.0.1",
		Latitude:   48.8566,
		Longitude:  2.3522,
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
	if response.Patient.ID != "" {
		t.Errorf("Expected no Patient, but got ID: %s", response.Patient.ID)
	}
}

func TestCreateDeviceConnect_FailedDeviceCreation(t *testing.T) {
	// Configuration des valeurs de test pour simuler une création de device échouée
	input := CreateDeviceConnectInput{
		DeviceName: "", // Champ obligatoire manquant ou invalide
		Ip:         "192.168.0.1",
		Latitude:   48.8566,
		Longitude:  2.3522,
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
	if response.Patient.ID != "" {
		t.Errorf("Expected no Patient to be returned, but got ID: %s", response.Patient.ID)
	}
}
