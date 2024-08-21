package double_auth

import (
	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
	"testing"
)

func TestDeleteDeviceConnect_Success(t *testing.T) {
	input := CreateDeviceConnectInput{
		DeviceName: "TestDevice",
		Ip:         "192.168.0.1",
		City:       "testCity",
		Country:    "testCountry",
		Date:       1627880400,
	}

	patient, err := graphql.CreatePatient(model.CreatePatientInput{
		Email:    "test_delete_device_connect_success@example.com",
		Password: "password",
		Status:   true,
	})
	if err != nil {
		t.Fatalf("failed to create patient: %s", err)
	}

	device := CreateDeviceConnect(input, patient.ID)
	response := DeleteDeviceConnect(device.DeviceConnect.ID, patient.ID)

	if response.Err != nil {
		t.Errorf("Expected no error, got: %v", response.Err)
	}
	if response.Code != 200 {
		t.Errorf("Expected status code 200, got: %d", response.Code)
	}
	if !response.Deleted {
		t.Errorf("Expected device to be deleted, but it wasn't")
	}

	updatedPatient, _ := graphql.GetPatientById(patient.ID)
	for _, devices := range updatedPatient.DeviceConnect {
		if *devices == device.DeviceConnect.ID {
			t.Errorf("Expected device %s to be removed from patient, but it still exists", device.DeviceConnect.ID)
		}
	}
}

func TestDeleteDeviceConnect_DoctorSuccess(t *testing.T) {
	input := CreateDeviceConnectInput{
		DeviceName: "TestDevice",
		Ip:         "192.168.0.1",
		City:       "testCity",
		Country:    "testCountry",
		Date:       1627880400,
	}

	doctor, err := graphql.CreateDoctor(model.CreateDoctorInput{
		Email:    "test_delete_device_connect_doctor_success@example.com",
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
	response := DeleteDeviceConnect(device.DeviceConnect.ID, doctor.ID)

	if response.Err != nil {
		t.Errorf("Expected no error, got: %v", response.Err)
	}
	if response.Code != 200 {
		t.Errorf("Expected status code 200, got: %d", response.Code)
	}
	if !response.Deleted {
		t.Errorf("Expected device to be deleted, but it wasn't")
	}

	updatedDoctor, _ := graphql.GetDoctorById(doctor.ID)
	for _, devices := range updatedDoctor.DeviceConnect {
		if *devices == device.DeviceConnect.ID {
			t.Errorf("Expected device %s to be removed from doctor, but it still exists", device.DeviceConnect.ID)
		}
	}
}

func TestDeleteDeviceConnect_InvalidDeviceId(t *testing.T) {
	invalidDeviceId := ""
	patientId := "valid_patient_id"

	response := DeleteDeviceConnect(invalidDeviceId, patientId)

	if response.Err == nil {
		t.Errorf("Expected an error, got none")
	}
	if response.Code != 400 {
		t.Errorf("Expected status code 400, got: %d", response.Code)
	}
}

func TestDeleteDeviceConnect_DeviceNotFound(t *testing.T) {
	deviceId := "invalid_device_id"
	patientId := "valid_patient_id"

	response := DeleteDeviceConnect(deviceId, patientId)

	if response.Err == nil {
		t.Errorf("Expected an error, got none")
	}
	if response.Code != 400 {
		t.Errorf("Expected status code 400, got: %d", response.Code)
	}
}

func TestDeleteDeviceConnect_PatientNotFound(t *testing.T) {
	deviceId := "valid_device_id"
	invalidPatientId := "invalid_patient_id"

	response := DeleteDeviceConnect(deviceId, invalidPatientId)

	if response.Err == nil {
		t.Errorf("Expected an error, got none")
	}
	if response.Code != 400 {
		t.Errorf("Expected status code 400, got: %d", response.Code)
	}
}

func TestDeleteDeviceConnect_DoctorNotFound(t *testing.T) {
	deviceId := "valid_device_id"
	invalidDoctorId := "invalid_doctor_id"

	response := DeleteDeviceConnect(deviceId, invalidDoctorId)

	if response.Err == nil {
		t.Errorf("Expected an error, got none")
	}
	if response.Code != 400 {
		t.Errorf("Expected status code 400, got: %d", response.Code)
	}
}
