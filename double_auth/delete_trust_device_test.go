package double_auth

import (
	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
	"testing"
)

func TestRemoveTrustDevice_Success(t *testing.T) {
	input := CreateDeviceConnectInput{
		DeviceType: "TestDevice",
		Browser:    "TestBrowser",
		Ip:         "192.168.0.1",
		City:       "testCity",
		Country:    "testCountry",
		Date:       1627880400,
	}

	patient, err := graphql.CreatePatient(model.CreatePatientInput{
		Email:    "test_rem_device_trust_success@example.com",
		Password: "password",
		Status:   true,
	})
	if err != nil {
		t.Fatalf("failed to create patient: %s", err)
	}

	device := CreateDeviceConnect(input, patient.ID)

	mobile := CreateDoubleMobileInput{
		Methods:     "MOBILE",
		TrustDevice: device.DeviceConnect.ID,
	}
	_ = CreateDoubleAuthMobile(mobile, patient.ID)
	response := RemoveTrustDevice(device.DeviceConnect.ID, patient.ID)

	if response.Err != nil {
		t.Errorf("Expected no error, got: %v", response.Err)
	}
	if response.Code != 200 {
		t.Errorf("Expected status code 200, got: %d", response.Code)
	}
	if response.Code != 200 {
		t.Errorf("Expected device to be deleted, but it wasn't")
	}

	updatedPatient, _ := graphql.GetPatientById(patient.ID)
	for _, devices := range updatedPatient.TrustDevices {
		if *devices == device.DeviceConnect.ID {
			t.Errorf("Expected device %s to be removed from patient, but it still exists", device.DeviceConnect.ID)
		}
	}
}

func TestRemoveTrustDevice_DoctorSuccess(t *testing.T) {
	input := CreateDeviceConnectInput{
		DeviceType: "TestDevice",
		Browser:    "TestBrowser",
		Ip:         "192.168.0.1",
		City:       "testCity",
		Country:    "testCountry",
		Date:       1627880400,
	}

	doctor, err := graphql.CreateDoctor(model.CreateDoctorInput{
		Email:    "test_rem_device_trust_doctor_success@example.com",
		Password: "password",
		Status:   true,
		Address: &model.AddressInput{
			Street:  "",
			ZipCode: "",
			Country: "",
			City:    "",
		},
	})
	if err != nil {
		t.Fatalf("failed to create doctor: %s", err)
	}

	device := CreateDeviceConnect(input, doctor.ID)

	mobile := CreateDoubleMobileInput{
		Methods:     "MOBILE",
		TrustDevice: device.DeviceConnect.ID,
	}
	_ = CreateDoubleAuthMobile(mobile, doctor.ID)
	response := RemoveTrustDevice(device.DeviceConnect.ID, doctor.ID)

	if response.Err != nil {
		t.Errorf("Expected no error, got: %v", response.Err)
	}
	if response.Code != 200 {
		t.Errorf("Expected status code 200, got: %d", response.Code)
	}
	if response.Code != 200 {
		t.Errorf("Expected device to be deleted, but it wasn't")
	}

	updatedDoctor, _ := graphql.GetDoctorById(doctor.ID)
	for _, devices := range updatedDoctor.TrustDevices {
		if *devices == device.DeviceConnect.ID {
			t.Errorf("Expected device %s to be removed from doctor, but it still exists", device.DeviceConnect.ID)
		}
	}
}

func TestRemoveTrustDevice_InvalidDeviceId(t *testing.T) {
	invalidDeviceId := ""
	patientId := "valid_patient_id"

	response := RemoveTrustDevice(invalidDeviceId, patientId)

	if response.Err == nil {
		t.Errorf("Expected an error, got none")
	}
	if response.Code != 400 {
		t.Errorf("Expected status code 400, got: %d", response.Code)
	}
}

func TestRemoveTrustDevice_DeviceNotFound(t *testing.T) {
	deviceId := "invalid_device_id"
	patientId := "valid_patient_id"

	response := RemoveTrustDevice(deviceId, patientId)

	if response.Err == nil {
		t.Errorf("Expected an error, got none")
	}
	if response.Code != 400 {
		t.Errorf("Expected status code 400, got: %d", response.Code)
	}
}

func TestRemoveTrustDevice_PatientNotFound(t *testing.T) {
	deviceId := "valid_device_id"
	invalidPatientId := "invalid_patient_id"

	response := RemoveTrustDevice(deviceId, invalidPatientId)

	if response.Err == nil {
		t.Errorf("Expected an error, got none")
	}
	if response.Code != 400 {
		t.Errorf("Expected status code 400, got: %d", response.Code)
	}
}

func TestRemoveTrustDevice_DoctorNotFound(t *testing.T) {
	deviceId := "valid_device_id"
	invalidDoctorId := "invalid_doctor_id"

	response := RemoveTrustDevice(deviceId, invalidDoctorId)

	if response.Err == nil {
		t.Errorf("Expected an error, got none")
	}
	if response.Code != 400 {
		t.Errorf("Expected status code 400, got: %d", response.Code)
	}
}

func TestRemoveTrustDevice_ErrorUpdatingPatient(t *testing.T) {
	deviceId := "valid_device_id"
	patientId := "valid_patient_id_with_update_error"

	response := RemoveTrustDevice(deviceId, patientId)

	if response.Err == nil {
		t.Errorf("Expected an error, got none")
	}
	if response.Code != 400 {
		t.Errorf("Expected status code 400, got: %d", response.Code)
	}
}

func TestRemoveTrustDevice_ErrorUpdatingDoctor(t *testing.T) {
	deviceId := "valid_device_id"
	doctorId := "valid_doctor_id_with_update_error"

	response := RemoveTrustDevice(deviceId, doctorId)

	if response.Err == nil {
		t.Errorf("Expected an error, got none")
	}
	if response.Code != 400 {
		t.Errorf("Expected status code 400, got: %d", response.Code)
	}
}
