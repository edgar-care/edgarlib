package double_auth

import (
	"github.com/edgar-care/edgarlib/v2/auth"
	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
	"testing"
)

func TestUpdateDoubleAuth_Success(t *testing.T) {
	patient, err := auth.RegisterPatient("test_update_double_auth_success@example.com", "password")
	if err != nil {
		t.Errorf("could not create patient: %v", err)
	}
	input := UpdateDoubleAuthInput{
		Methods:         "SMS",
		Secret:          "new_secret",
		URL:             "https://new.url",
		TrustedDeviceId: "new_device_id",
	}

	doubleAuth, err := graphql.CreateDoubleAuth(model.CreateDoubleAuthInput{
		Methods:       []string{"EMAIL"},
		Secret:        "old_secret",
		URL:           "https://old.url",
		TrustDeviceID: "old_device_id",
	})

	if err != nil {
		t.Errorf("could not create double-auth: %v", err)
	}

	_, err = graphql.UpdatePatient(patient.ID, model.UpdatePatientInput{DoubleAuthMethodsID: &doubleAuth.ID})

	if err != nil {
		t.Errorf("could not update patient: %v", err)
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

func TestUpdateDoubleAuth_DoctorSuccess(t *testing.T) {
	doctor, err := auth.RegisterDoctor("test_update_double_auth_doctor_success@example.com", "password", "poulet", "marvin", auth.AddressInput{
		Street:  "",
		ZipCode: "",
		Country: "",
		City:    "",
	})
	if err != nil {
		t.Errorf("could not create doctor: %v", err)
	}
	input := UpdateDoubleAuthInput{
		Methods:         "SMS",
		Secret:          "new_secret",
		URL:             "https://new.url",
		TrustedDeviceId: "new_device_id",
	}

	doubleAuth, err := graphql.CreateDoubleAuth(model.CreateDoubleAuthInput{
		Methods:       []string{"EMAIL"},
		Secret:        "old_secret",
		URL:           "https://old.url",
		TrustDeviceID: "old_device_id",
	})

	if err != nil {
		t.Errorf("could not create double-auth: %v", err)
	}

	_, err = graphql.UpdateDoctor(doctor.ID, model.UpdateDoctorInput{DoubleAuthMethodsID: &doubleAuth.ID})

	if err != nil {
		t.Errorf("could not update doctor: %v", err)
	}
	response := UpdateDoubleAuth(input, doctor.ID)

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
		Methods:         "SMS",
		Secret:          "secret",
		URL:             "https://url",
		TrustedDeviceId: "device_id",
	}

	response := UpdateDoubleAuth(input, patientID)

	if response.Err == nil {
		t.Errorf("Expected error but got none")
	}

	if response.Code != 400 {
		t.Errorf("Expected status code 400, but got: %d", response.Code)
	}
}

func TestUpdateDoubleAuth_DoctorNotFound(t *testing.T) {
	doctorID := "invalid_doctor_id"
	input := UpdateDoubleAuthInput{
		Methods:         "SMS",
		Secret:          "secret",
		URL:             "https://url",
		TrustedDeviceId: "device_id",
	}

	response := UpdateDoubleAuth(input, doctorID)

	if response.Err == nil {
		t.Errorf("Expected error but got none")
	}

	if response.Code != 400 {
		t.Errorf("Expected status code 400, but got: %d", response.Code)
	}
}

func TestUpdateDoubleAuth_NoDoubleAuth(t *testing.T) {
	patient, err := auth.RegisterPatient("test_update_double_auth_no_double_auth@example.com", "password")
	if err != nil {
		t.Errorf("could not create patient: %v", err)
	}

	input := UpdateDoubleAuthInput{
		Methods:         "SMS",
		Secret:          "secret",
		URL:             "https://url",
		TrustedDeviceId: "device_id",
	}

	response := UpdateDoubleAuth(input, patient.ID)

	if response.Err == nil {
		t.Errorf("Expected error but got none")
	}

	if response.Code != 404 {
		t.Errorf("Expected status code 404, but got: %d", response.Code)
	}
}

func TestUpdateDoubleAuth_NoDoubleAuthForDoctor(t *testing.T) {
	doctor, err := auth.RegisterDoctor("test_update_double_auth_no_double_auth_doctor@example.com", "password", "checko", "jean", auth.AddressInput{
		Street:  "",
		ZipCode: "",
		Country: "",
		City:    "",
	})
	if err != nil {
		t.Errorf("could not create doctor: %v", err)
	}

	input := UpdateDoubleAuthInput{
		Methods:         "SMS",
		Secret:          "secret",
		URL:             "https://url",
		TrustedDeviceId: "device_id",
	}

	response := UpdateDoubleAuth(input, doctor.ID)

	if response.Err == nil {
		t.Errorf("Expected error but got none")
	}

	if response.Code != 404 {
		t.Errorf("Expected status code 404, but got: %d", response.Code)
	}
}
