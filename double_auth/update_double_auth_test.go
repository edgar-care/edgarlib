package double_auth

import (
	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
	"testing"
)

func TestUpdateDoubleAuth_Success(t *testing.T) {
	patient, err := graphql.CreatePatient(model.CreatePatientInput{
		Email:    "test_patient_update_double_auth_success@edgar-sante.fr",
		Password: "password",
		Status:   true,
	})
	input := UpdateDoubleAuthInput{
		Methods:         "SMS",
		Secret:          "new_secret",
		TrustedDeviceId: "new_device_id",
	}

	doubleAuth, err := graphql.CreateDoubleAuth(model.CreateDoubleAuthInput{
		Methods:       []string{"EMAIL"},
		Secret:        "old_secret",
		TrustDeviceID: []string{"old_device_id"},
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

}

func TestUpdateDoubleAuth_DoctorSuccess(t *testing.T) {
	doctor, err := graphql.CreateDoctor(model.CreateDoctorInput{
		Email:     "test_doctor_update_doubleauth_success@example.com",
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

	input := UpdateDoubleAuthInput{
		Methods:         "SMS",
		Secret:          "new_secret",
		TrustedDeviceId: "new_device_id",
	}

	doubleAuth, err := graphql.CreateDoubleAuth(model.CreateDoubleAuthInput{
		Methods:       []string{"EMAIL"},
		Secret:        "old_secret",
		TrustDeviceID: []string{"old_device_id"},
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

}

func TestUpdateDoubleAuth_PatientNotFound(t *testing.T) {
	patientID := "invalid_patient_id"
	input := UpdateDoubleAuthInput{
		Methods:         "SMS",
		Secret:          "secret",
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
	patient, _ := graphql.CreatePatient(model.CreatePatientInput{
		Email:    "test_patient_update_not_doubleauth_wrong@edgar-sante.fr",
		Password: "password",
		Status:   true,
	})

	input := UpdateDoubleAuthInput{
		Methods:         "SMS",
		Secret:          "secret",
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
	doctor, err := graphql.CreateDoctor(model.CreateDoctorInput{
		Email:     "test_doctor_update_wrong_no_doubleuath@example.com",
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

	input := UpdateDoubleAuthInput{
		Methods:         "SMS",
		Secret:          "secret",
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
