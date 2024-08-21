package double_auth

import (
	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
	"testing"
)

func TestCreateDoubleAuthMobilePatient_Succes(t *testing.T) {
	patient, err := graphql.CreatePatient(model.CreatePatientInput{
		Email:    "test_add_double_auth_mobile_success@example.com",
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
	_ = AddTrustDevice(device.DeviceConnect.ID, patient.ID)

	mobile := CreateDoubleMobileInput{Methods: "MOBILE", TrustDevice: device.DeviceConnect.ID}
	response := CreateDoubleAuthMobile(mobile, patient.ID)

	if response.Err != nil {
		t.Errorf("Expected no error, got: %v", response.Err)
	}
	if response.Code != 201 {
		t.Errorf("Expected status code 200, got: %d", response.Code)
	}
}

func TestCreateDoubleAuthMobileDoctor_Success(t *testing.T) {
	doctor, err := graphql.CreateDoctor(model.CreateDoctorInput{
		Email:     "test_create_double_auth_mobile_succes@example.com",
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
	_ = AddTrustDevice(device.DeviceConnect.ID, doctor.ID)

	mobile := CreateDoubleMobileInput{Methods: "MOBILE", TrustDevice: device.DeviceConnect.ID}
	response := CreateDoubleAuthMobile(mobile, doctor.ID)

	if response.Err != nil {
		t.Errorf("Expected no error, got: %v", response.Err)
	}
	if response.Code != 201 {
		t.Errorf("Expected status code 200, got: %d", response.Code)
	}
}

func TestCreateDoubleAuthMobile_Invalid(t *testing.T) {
	doctor, err := graphql.CreateDoctor(model.CreateDoctorInput{
		Email:     "test_create_double_auth_mobile_invalid@example.com",
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

	mobile := CreateDoubleMobileInput{Methods: "MOBILE", TrustDevice: "test_invalid_id"}
	response := CreateDoubleAuthMobile(mobile, doctor.ID)

	if response.Err == nil {
		t.Errorf("Expected an error, got none")
	}
	if response.Code != 400 {
		t.Errorf("Expected status code 400, got: %d", response.Code)
	}
}

func TestCreateDoubleAuthMobile_InvalidAccount(t *testing.T) {

	patient, err := graphql.CreatePatient(model.CreatePatientInput{
		Email:    "test_add_double_auth_mobile_invalid_id@example.com",
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
	_ = AddTrustDevice(device.DeviceConnect.ID, patient.ID)

	mobile := CreateDoubleMobileInput{Methods: "MOBILE", TrustDevice: device.DeviceConnect.ID}
	response := CreateDoubleAuthMobile(mobile, "test_invalid_id")

	if response.Err == nil {
		t.Errorf("Expected an error, got none")
	}
	if response.Code != 400 {
		t.Errorf("Expected status code 400, got: %d", response.Code)
	}
}

func TestCreateDoubleAuthMobile_InvalidMethod(t *testing.T) {

	patient, err := graphql.CreatePatient(model.CreatePatientInput{
		Email:    "test_add_double_auth_mobile_invalid_method@example.com",
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
	_ = AddTrustDevice(device.DeviceConnect.ID, patient.ID)

	mobile := CreateDoubleMobileInput{Methods: "TEST", TrustDevice: device.DeviceConnect.ID}
	response := CreateDoubleAuthMobile(mobile, patient.ID)

	if response.Err == nil {
		t.Errorf("Expected an error, got none")
	}
	if response.Code != 400 {
		t.Errorf("Expected status code 400, got: %d", response.Code)
	}
}

func TestCreateDoubleAuthMobileADDPatient_Succes(t *testing.T) {
	patient, err := graphql.CreatePatient(model.CreatePatientInput{
		Email:    "test_add_double_auth_mobile_success_method_add@example.com",
		Password: "password",
		Status:   true,
	})
	if err != nil {
		t.Fatalf("Failed to create patient: %s", err)
	}

	email := CreateDoubleAuthInput{Methods: "EMAIL"}
	_ = CreateDoubleAuthEmail(email, patient.ID)

	input := CreateDeviceConnectInput{
		DeviceName: "TestDevice",
		Ip:         "192.168.0.1",
		City:       "testCity",
		Country:    "testCountry",
		Date:       1627880400,
	}

	device := CreateDeviceConnect(input, patient.ID)
	_ = AddTrustDevice(device.DeviceConnect.ID, patient.ID)

	mobile := CreateDoubleMobileInput{Methods: "MOBILE", TrustDevice: device.DeviceConnect.ID}
	response := CreateDoubleAuthMobile(mobile, patient.ID)

	if response.Err != nil {
		t.Errorf("Expected no error, got: %v", response.Err)
	}
	if response.Code != 200 {
		t.Errorf("Expected status code 200, got: %d", response.Code)
	}
}

func TestCreateDoubleAuthMobileADDPatient_InvalidMehtod(t *testing.T) {
	patient, err := graphql.CreatePatient(model.CreatePatientInput{
		Email:    "test_add_double_auth_mobile_invalid_methods@example.com",
		Password: "password",
		Status:   true,
	})
	if err != nil {
		t.Fatalf("Failed to create patient: %s", err)
	}

	email := CreateDoubleAuthInput{Methods: "EMAIL"}
	_ = CreateDoubleAuthEmail(email, patient.ID)

	input := CreateDeviceConnectInput{
		DeviceName: "TestDevice",
		Ip:         "192.168.0.1",
		City:       "testCity",
		Country:    "testCountry",
		Date:       1627880400,
	}

	device := CreateDeviceConnect(input, patient.ID)
	_ = AddTrustDevice(device.DeviceConnect.ID, patient.ID)

	mobile := CreateDoubleMobileInput{Methods: "TEST", TrustDevice: device.DeviceConnect.ID}
	response := CreateDoubleAuthMobile(mobile, patient.ID)

	if response.Err == nil {
		t.Errorf("Expected an error, got none")
	}
	if response.Code != 400 {
		t.Errorf("Expected status code 400, got: %d", response.Code)
	}
}

func TestCreateDoubleAuthMobileADDPatient_InvalidID(t *testing.T) {
	patient, err := graphql.CreatePatient(model.CreatePatientInput{
		Email:    "test_add_double_auth_mobile_invalid_ID_addd@example.com",
		Password: "password",
		Status:   true,
	})
	if err != nil {
		t.Fatalf("Failed to create patient: %s", err)
	}

	email := CreateDoubleAuthInput{Methods: "EMAIL"}
	_ = CreateDoubleAuthEmail(email, patient.ID)

	input := CreateDeviceConnectInput{
		DeviceName: "TestDevice",
		Ip:         "192.168.0.1",
		City:       "testCity",
		Country:    "testCountry",
		Date:       1627880400,
	}

	device := CreateDeviceConnect(input, patient.ID)
	_ = AddTrustDevice(device.DeviceConnect.ID, patient.ID)

	mobile := CreateDoubleMobileInput{Methods: "MOBILE", TrustDevice: "test_invalid_id"}
	response := CreateDoubleAuthMobile(mobile, patient.ID)

	if response.Err == nil {
		t.Errorf("Expected an error, got none")
	}
	if response.Code != 400 {
		t.Errorf("Expected status code 400, got: %d", response.Code)
	}
}
