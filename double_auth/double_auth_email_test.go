package double_auth

import (
	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
	"testing"
)

func TestCreateDoubleAuthEmail_Succes(t *testing.T) {
	patient, err := graphql.CreatePatient(model.CreatePatientInput{
		Email:    "test_add_double_auth_email_success_create@example.com",
		Password: "password",
		Status:   true,
	})
	if err != nil {
		t.Fatalf("Failed to create patient: %s", err)
	}

	input := CreateDoubleAuthInput{
		Methods: "EMAIL",
	}

	response := CreateDoubleAuthEmail(input, patient.ID)

	if response.Err != nil {
		t.Errorf("Expected no error, got: %v", response.Err)
	}
	if response.Code != 201 {
		t.Errorf("Expected status code 200, got: %d", response.Code)
	}
}

func TestCreateDoubleAuthEmailDoctor_Succes(t *testing.T) {
	doctor, err := graphql.CreateDoctor(model.CreateDoctorInput{
		Email:     "test_create_double_auth_email_succes_create@example.com",
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

	input := CreateDoubleAuthInput{
		Methods: "EMAIL",
	}

	response := CreateDoubleAuthEmail(input, doctor.ID)

	if response.Err != nil {
		t.Errorf("Expected no error, got: %v", response.Err)
	}
	if response.Code != 201 {
		t.Errorf("Expected status code 200, got: %d", response.Code)
	}
}

func TestCreateDoubleAuthEmail_invalid(t *testing.T) {

	input := CreateDoubleAuthInput{
		Methods: "EMAIL",
	}

	response := CreateDoubleAuthEmail(input, "test_invalid_id")

	if response.Err == nil {
		t.Errorf("Expected an error, got none")
	}
	if response.Code != 400 {
		t.Errorf("Expected status code 400, got: %d", response.Code)
	}
}

func TestCreateDoubleAuthEmail_InvalidMethods(t *testing.T) {

	patient, err := graphql.CreatePatient(model.CreatePatientInput{
		Email:    "test_add_double_auth_email_invalid_method@example.com",
		Password: "password",
		Status:   true,
	})
	if err != nil {
		t.Fatalf("Failed to create patient: %s", err)
	}

	input := CreateDoubleAuthInput{
		Methods: "TEST",
	}

	response := CreateDoubleAuthEmail(input, patient.ID)

	if response.Err == nil {
		t.Errorf("Expected an error, got none")
	}
	if response.Code != 400 {
		t.Errorf("Expected status code 400, got: %d", response.Code)
	}
}

func TestCreateDoubleAuthEmail_SuccesAddAUth(t *testing.T) {
	patient, err := graphql.CreatePatient(model.CreatePatientInput{
		Email:    "test_add_double_auth_email_success_method_add@example.com",
		Password: "password",
		Status:   true,
	})
	if err != nil {
		t.Fatalf("Failed to create patient: %s", err)
	}

	input := CreateDeviceConnectInput{
		DeviceType: "TestDevice",
		Browser:    "TestBrowser",
		Ip:         "192.168.0.1",
		City:       "testCity",
		Country:    "testCountry",
		Date:       1627880400,
	}

	device := CreateDeviceConnect(input, patient.ID)
	_ = AddTrustDevice(device.DeviceConnect.ID, patient.ID)

	mobile := CreateDoubleMobileInput{Methods: "MOBILE", TrustDevice: device.DeviceConnect.ID}
	_ = CreateDoubleAuthMobile(mobile, patient.ID)

	email := CreateDoubleAuthInput{Methods: "EMAIL"}
	response := CreateDoubleAuthEmail(email, patient.ID)

	if response.Err != nil {
		t.Errorf("Expected no error, got: %v", response.Err)
	}
	if response.Code != 200 {
		t.Errorf("Expected status code 200, got: %d", response.Code)
	}
}

func TestCreateDoubleAuthEmail_InvalidAddAUthMethod(t *testing.T) {
	patient, err := graphql.CreatePatient(model.CreatePatientInput{
		Email:    "test_add_double_auth_email_invalid_methods@example.com",
		Password: "password",
		Status:   true,
	})
	if err != nil {
		t.Fatalf("Failed to create patient: %s", err)
	}

	input := CreateDeviceConnectInput{
		DeviceType: "TestDevice",
		Browser:    "TestBrowser",
		Ip:         "192.168.0.1",
		City:       "testCity",
		Country:    "testCountry",
		Date:       1627880400,
	}

	device := CreateDeviceConnect(input, patient.ID)
	_ = AddTrustDevice(device.DeviceConnect.ID, patient.ID)

	mobile := CreateDoubleMobileInput{Methods: "MOBILE", TrustDevice: device.DeviceConnect.ID}
	_ = CreateDoubleAuthMobile(mobile, patient.ID)

	email := CreateDoubleAuthInput{Methods: "TEST"}
	response := CreateDoubleAuthEmail(email, patient.ID)

	if response.Err == nil {
		t.Errorf("Expected an error, got none")
	}
	if response.Code != 400 {
		t.Errorf("Expected status code 400, got: %d", response.Code)
	}
}
