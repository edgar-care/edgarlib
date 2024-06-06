package double_auth

import (
	"github.com/edgar-care/edgarlib/auth"
	"testing"
)

func TestUpdateDoubleAuth_Success(t *testing.T) {
	patient, err := auth.RegisterPatient("test_update_double_auth_success@example.com", "password")
	if err != nil {
		t.Errorf("could not create patient: %v", err)
	}
	input := UpdateDoubleAuthInput{
		Secret:          "new_secret",
		URL:             "https://new.url",
		TrustedDeviceId: "new_device_id",
	}

	response := UpdateDoubleAuth(input, patient.ID)

	if response.Err != nil {
		t.Errorf("Expected no error, but got: %v", response.Err)
	}

	if response.Code != 200 {
		t.Errorf("Expected status code 200, but got: %d", response.Code)
	}

	if response.DoubleAuth.Secret != input.Secret {
		t.Errorf("Expected secret %s, but got: %s", input.Secret, response.DoubleAuth.Secret)
	}

	if response.DoubleAuth.URL != input.URL {
		t.Errorf("Expected URL %s, but got: %s", input.URL, response.DoubleAuth.URL)
	}

	if response.DoubleAuth.TrustDeviceID != input.TrustedDeviceId {
		t.Errorf("Expected trusted device ID %s, but got: %s", input.TrustedDeviceId, response.DoubleAuth.TrustDeviceID)
	}
}

func TestUpdateDoubleAuth_PatientNotFound(t *testing.T) {
	patientID := "invalid_patient_id"
	input := UpdateDoubleAuthInput{
		Secret:          "secret",
		URL:             "https://url",
		TrustedDeviceId: "device_id",
	}

	response := UpdateDoubleAuth(input, patientID)

	if response.Err == nil || response.Err.Error() != "id does not correspond to a patient" {
		t.Errorf("Expected error 'id does not correspond to a patient', but got: %v", response.Err)
	}

	if response.Code != 400 {
		t.Errorf("Expected status code 400, but got: %d", response.Code)
	}
}

func TestUpdateDoubleAuth_NoDoubleAuth(t *testing.T) {
	patient, err := auth.RegisterPatient("test_update_double_auth_success@example.com", "password")
	if err != nil {
		t.Errorf("could not create patient: %v", err)
	}

	input := UpdateDoubleAuthInput{
		Secret:          "secret",
		URL:             "https://url",
		TrustedDeviceId: "device_id",
	}

	response := UpdateDoubleAuth(input, patient.ID)

	if response.Err == nil || response.Err.Error() != "double auth not found on patient" {
		t.Errorf("Expected error 'double auth not found on patient', but got: %v", response.Err)
	}

	if response.Code != 404 {
		t.Errorf("Expected status code 404, but got: %d", response.Code)
	}
}
