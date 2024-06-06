package double_auth

import (
	"github.com/edgar-care/edgarlib/graphql"
	"testing"
)

func TestDeleteDeviceConnect_Success(t *testing.T) {
	// Configurer les valeurs de test
	deviceId := "valid_device_id"
	patientId := "valid_patient_id"

	// Appel de la fonction
	response := DeleteDeviceConnect(deviceId, patientId)

	// Vérifier les résultats
	if response.Err != nil {
		t.Errorf("Expected no error, got: %v", response.Err)
	}
	if response.Code != 200 {
		t.Errorf("Expected status code 200, got: %d", response.Code)
	}
	if !response.Deleted {
		t.Errorf("Expected device to be deleted, but it wasn't")
	}

	// Vérifier que l'appareil a bien été retiré du patient
	patient, _ := graphql.GetPatientById(patientId)
	for _, device := range patient.DeviceConnect {
		if *device == deviceId {
			t.Errorf("Expected device %s to be removed from patient, but it still exists", deviceId)
		}
	}
}

func TestDeleteDeviceConnect_InvalidDeviceId(t *testing.T) {
	// Configurer les valeurs de test
	invalidDeviceId := ""
	patientId := "valid_patient_id"

	// Appel de la fonction
	response := DeleteDeviceConnect(invalidDeviceId, patientId)

	// Vérifier les résultats
	if response.Err == nil {
		t.Errorf("Expected an error, got none")
	}
	if response.Code != 400 {
		t.Errorf("Expected status code 400, got: %d", response.Code)
	}
	if response.Err.Error() != "slot id is required" {
		t.Errorf("Expected error message 'slot id is required', got: %v", response.Err)
	}
}

func TestDeleteDeviceConnect_DeviceNotFound(t *testing.T) {
	// Configurer les valeurs de test
	deviceId := "invalid_device_id"
	patientId := "valid_patient_id"

	// Appel de la fonction
	response := DeleteDeviceConnect(deviceId, patientId)

	// Vérifier les résultats
	if response.Err == nil {
		t.Errorf("Expected an error, got none")
	}
	if response.Code != 400 {
		t.Errorf("Expected status code 400, got: %d", response.Code)
	}
	if response.Err.Error() != "id does not correspond to a slot" {
		t.Errorf("Expected error message 'id does not correspond to a slot', got: %v", response.Err)
	}
}

func TestDeleteDeviceConnect_PatientNotFound(t *testing.T) {
	// Configurer les valeurs de test
	deviceId := "valid_device_id"
	invalidPatientId := "invalid_patient_id"

	// Appel de la fonction
	response := DeleteDeviceConnect(deviceId, invalidPatientId)

	// Vérifier les résultats
	if response.Err == nil {
		t.Errorf("Expected an error, got none")
	}
	if response.Code != 400 {
		t.Errorf("Expected status code 400, got: %d", response.Code)
	}
	if response.Err.Error() != "id does not correspond to a doctor" {
		t.Errorf("Expected error message 'id does not correspond to a doctor', got: %v", response.Err)
	}
}

func TestDeleteDeviceConnect_ErrorUpdatingPatient(t *testing.T) {
	// Configurer les valeurs de test
	deviceId := "valid_device_id"
	patientId := "valid_patient_id_with_update_error"

	// Appel de la fonction
	response := DeleteDeviceConnect(deviceId, patientId)

	// Vérifier les résultats
	if response.Err == nil {
		t.Errorf("Expected an error, got none")
	}
	if response.Code != 500 {
		t.Errorf("Expected status code 500, got: %d", response.Code)
	}
	if response.Err.Error() != "error updating patient" {
		t.Errorf("Expected error message 'error updating patient', got: %v", response.Err)
	}
}
