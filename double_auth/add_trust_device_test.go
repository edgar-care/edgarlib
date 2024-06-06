package double_auth

import (
	"testing"
)

func TestAddTrustDevice_Patient_Success(t *testing.T) {
	// Configurer des valeurs de test
	idDevice := "test_device_id"
	idUser := "test_patient_id"

	// Appeler la fonction
	response := AddTrustDevice(idDevice, idUser)

	// Vérifier les résultats
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
	// Configurer des valeurs de test
	idDevice := "test_device_id"
	idUser := "test_doctor_id"

	// Appeler la fonction
	response := AddTrustDevice(idDevice, idUser)

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
	// Configurer des valeurs de test
	idDevice := "test_device_id"
	idUser := "invalid_user_id"

	// Appeler la fonction
	response := AddTrustDevice(idDevice, idUser)

	// Vérifier les résultats
	if response.Err == nil {
		t.Errorf("Expected an error, got none")
	}
	if response.Code == 200 {
		t.Errorf("Expected non-200 status code, got: %d", response.Code)
	}
	if response.Patient != nil || response.Doctor != nil {
		t.Errorf("Expected no patient or doctor, got: %v, %v", response.Patient, response.Doctor)
	}
}
