package double_auth

import (
	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
	"testing"
)

func TestGetTrustDeviceConnectById_Success(t *testing.T) {
	input := CreateDeviceConnectInput{
		DeviceName: "TestDevice",
		Ip:         "192.168.0.1",
		City:       "testCity",
		Country:    "testCountry",
		Date:       1627880400,
	}

	patient, err := graphql.CreatePatient(model.CreatePatientInput{
		Email:    "test_get_device_trust_success@example.com",
		Password: "password",
		Status:   true,
	})
	if err != nil {
		t.Fatalf("failed to create patient: %s", err)
	}

	device := CreateDeviceConnect(input, patient.ID)

	tier := CreateDoubleAuthTierInput{Methods: "AUTHENTIFICATOR"}
	_ = CreateDoubleAuthAppTier(tier, "url", patient.ID)

	_ = AddTrustDevice(device.DeviceConnect.ID, patient.ID)

	response := GetTrustDeviceConnectById(device.DeviceConnect.ID)

	if response.Err != nil {
		t.Errorf("Expected no error, got: %v", response.Err)
	}
	if response.Code != 200 {
		t.Errorf("Expected status code 200, got: %d", response.Code)
	}
	if response.DeviceConnect.ID != device.DeviceConnect.ID {
		t.Errorf("Expected device ID %s, got: %s", device.DeviceConnect.ID, response.DeviceConnect.ID)
	}
}

func TestGetTrustDeviceConnectById_NotFound(t *testing.T) {
	invalidDeviceId := "invalid_device_id"

	response := GetTrustDeviceConnectById(invalidDeviceId)

	if response.Err == nil {
		t.Errorf("Expected an error, got none")
	}
	if response.Code != 400 {
		t.Errorf("Expected status code 400, got: %d", response.Code)
	}
}

func TestGetTrustDeviceConnect_PatientSuccess(t *testing.T) {
	patient, err := graphql.CreatePatient(model.CreatePatientInput{
		Email:    "test_get_trust_device_success@example.com",
		Password: "password",
		Status:   true,
	})
	if err != nil {
		t.Fatalf("Failed to create patient: %s", err)
	}

	input := CreateDeviceConnectInput{
		DeviceName: "TestDevice",
		Ip:         "192.168.0.1",
		City:       "testCity",
		Country:    "testCountry",
		Date:       1627880400,
	}

	device := CreateDeviceConnect(input, patient.ID)

	tier := CreateDoubleAuthTierInput{Methods: "AUTHENTIFICATOR"}
	_ = CreateDoubleAuthAppTier(tier, "url", patient.ID)

	_ = AddTrustDevice(device.DeviceConnect.ID, patient.ID)

	response := GetTrustDeviceConnect(patient.ID)

	if response.Err != nil {
		t.Errorf("Expected no error, got: %v", response.Err)
	}
	if response.Code != 200 {
		t.Errorf("Expected status code 200, got: %d", response.Code)
	}
	if len(response.DevicesConnect) == 0 || response.DevicesConnect[0].ID != device.DeviceConnect.ID {
		t.Errorf("Expected device ID %s, got: %v", device.DeviceConnect.ID, response.DevicesConnect)
	}
}

func TestGetTrustDeviceConnect_DoctorSuccess(t *testing.T) {
	doctor, err := graphql.CreateDoctor(model.CreateDoctorInput{
		Email:     "test_get_trust_device_doctor@example.com",
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

	input := CreateDeviceConnectInput{
		DeviceName: "TestDevice",
		Ip:         "192.168.0.1",
		City:       "testCity",
		Country:    "testCountry",
		Date:       1627880400,
	}

	device := CreateDeviceConnect(input, doctor.ID)

	tier := CreateDoubleAuthTierInput{Methods: "AUTHENTIFICATOR"}
	_ = CreateDoubleAuthAppTier(tier, "url", doctor.ID)

	_ = AddTrustDevice(device.DeviceConnect.ID, doctor.ID)

	response := GetTrustDeviceConnect(doctor.ID)

	if response.Err != nil {
		t.Errorf("Expected no error, got: %v", response.Err)
	}
	if response.Code != 200 {
		t.Errorf("Expected status code 200, got: %d", response.Code)
	}
	if len(response.DevicesConnect) == 0 || response.DevicesConnect[0].ID != device.DeviceConnect.ID {
		t.Errorf("Expected device ID %s, got: %v", device.DeviceConnect.ID, response.DevicesConnect)
	}
}

func TestGetTrustDeviceConnect_InvalidUserId(t *testing.T) {
	invalidUserId := "invalid_user_id"

	response := GetTrustDeviceConnect(invalidUserId)

	if response.Err == nil {
		t.Errorf("Expected an error, got none")
	}
	if response.Code != 400 {
		t.Errorf("Expected status code 400, got: %d", response.Code)
	}
	if len(response.DevicesConnect) != 0 {
		t.Errorf("Expected no devices, got: %v", response.DevicesConnect)
	}
}
