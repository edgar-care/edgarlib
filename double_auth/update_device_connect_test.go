package double_auth

import (
	"github.com/joho/godotenv"
	"log"
	"testing"

	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
)

func TestUpdateDeviceConnect_IP_Success(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}

	input := CreateDeviceConnectInput{
		DeviceType: "TestDevice",
		Browser:    "TestBrowser",
		Ip:         "192.168.0.1",
		City:       "testcity",
		Country:    "testcountry",
		Date:       1627880400,
	}
	patient, err := graphql.CreatePatient(model.CreatePatientInput{
		Email:    "test_patient_update_device_connect_success@edgar-sante.fr",
		Password: "password",
		Status:   true,
	})
	if err != nil {
		t.Errorf("Error while creating patient: %v", err)
	}

	test := CreateDeviceConnect(input, patient.ID)

	response := UpdateDeviceConnect(UpdateDeviceConnectInput{
		DeviceType: "Laptop",
		Browser:    "Chrome",
		Ip:         test.DeviceConnect.IPAddress,
		City:       "TestCity",
		Country:    "TestCountry",
		Date:       1627880400,
	}, test.DeviceConnect.IPAddress)

	if response.Code != 200 {
		t.Errorf("Expected response code 200, got %v", response.Code)
	}
	if response.Err != nil {
		t.Errorf("Unexpected error: %v", response.Err)
	}
}

func DeviceConnectByIP_NotFound(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}

	input := CreateDeviceConnectInput{
		DeviceType: "TestDevice",
		Browser:    "TestBrowser",
		Ip:         "192.168.0.1",
		City:       "testcity",
		Country:    "testcountry",
		Date:       1627880400,
	}
	patient, err := graphql.CreatePatient(model.CreatePatientInput{
		Email:    "test_patient_update_device_connect_failed@edgar-sante.fr",
		Password: "password",
		Status:   true,
	})
	if err != nil {
		t.Errorf("Error while creating patient: %v", err)
	}

	_ = CreateDeviceConnect(input, patient.ID)

	response := UpdateDeviceConnect(UpdateDeviceConnectInput{
		DeviceType: "Laptop",
		Browser:    "Chrome",
		Ip:         "ipAddress",
		City:       "TestCity",
		Country:    "TestCountry",
		Date:       1627880400,
	}, "ipAddress")

	if response.Code != 404 {
		t.Errorf("Expected response code 404, got %v", response.Code)
	}
	if response.Err == nil {
		t.Errorf("Expected an error in response")
	}
}
