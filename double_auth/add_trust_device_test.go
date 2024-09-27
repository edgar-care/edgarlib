package double_auth

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
	"github.com/joho/godotenv"
	"log"
	"testing"
)

func TestAddTrustDevice_Patient_Success(t *testing.T) {
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
		Email:    "test_patient_create_device_trust_success@edgar-sante.fr",
		Password: "password",
		Status:   true,
	})
	if err != nil {
		t.Errorf("Error while creating patient: %v", err)
	}

	device := CreateDeviceConnect(input, patient.ID)

	if device.Err != nil {
		t.Errorf("Expected no error, got: %v", device.Err)
	}
	if device.Code != 201 {
		t.Errorf("Expected status code 201, got: %d", device.Code)
	}

	input2 := CreateDoubleMobileInput{
		Methods:     "MOBILE",
		TrustDevice: device.DeviceConnect.ID,
	}

	_ = CreateDoubleAuthMobile(input2, patient.ID)

	input3 := CreateDeviceConnectInput{
		DeviceType: "Windows",
		Browser:    "Chrome",
		Ip:         "192.168.0.1",
		City:       "sdfsdf",
		Country:    "dfdgdfgdfg",
		Date:       1627880400,
	}
	device2 := CreateDeviceConnect(input3, patient.ID)

	response := AddTrustDevice(device2.DeviceConnect.ID, patient.ID)
	if response.Err != nil {
		t.Errorf("Expected no error, got: %v", response.Err)
	}
	if response.Code != 200 {
		t.Errorf("Expected status code 200, got: %d", response.Code)
	}
	if response.Patient == nil {
		t.Errorf("Expected a patient, got nil")
	}
	if response.Doctor != nil {
		t.Errorf("Expected no doctor, got: %v", response.Doctor)
	}

	_, err = graphql.GetDeviceConnectById(device2.DeviceConnect.ID)
}

func TestAddTrustDevice_Doctor_Success(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}
	input := CreateDeviceConnectInput{
		DeviceType: "TestDevice",
		Browser:    "TestBrowser",
		Ip:         "192.168.0.1",
		City:       "testCity",
		Country:    "testcountry",
		Date:       1627880400, // Timestamp pour une date fixe
	}

	doctor, err := graphql.CreateDoctor(model.CreateDoctorInput{
		Email:     "test_doctor_create_device_trust@edgar-sante.fr",
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
		t.Errorf("Error while creating patient: %v", err)
	}

	device := CreateDeviceConnect(input, doctor.ID)

	if device.Err != nil {
		t.Errorf("Expected no error, got: %v", device.Err)
	}
	if device.Code != 201 {
		t.Errorf("Expected status code 201, got: %d", device.Code)
	}

	_ = CreateDoubleAuthAppTier(doctor.ID)

	response := AddTrustDevice(device.DeviceConnect.ID, doctor.ID)
	// Vérifier les résultats
	if response.Err != nil {
		t.Errorf("Expected no error, got: %v", response.Err)
	}
	if response.Code != 200 {
		t.Errorf("Expected status code 200, got: %d", response.Code)
	}
	if response.Doctor == nil {
		t.Errorf("Expected a doctor, got nil")
	}
	if response.Patient != nil {
		t.Errorf("Expected no patient, got: %v", response.Patient)
	}
}

func TestAddTrustDevice_InvalidUser(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}
	idDevice := "test_device_id"
	idUser := "invalid_user_id"

	// Appeler la fonction
	response := AddTrustDevice(idDevice, idUser)

	if response.Err == nil {
		t.Errorf("Expected error but got none")
	}

}

func TestAddTrustDevice_Patient_SuccessDouble(t *testing.T) {
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

	//patient, err := graphql.CreatePatient(model.CreatePatientInput{
	//	Email:    "test_patient_create_device_trust_success@edgar-sante.fr",
	//	Password: "password",
	//	Status:   true,
	//})
	//if err != nil {
	//	t.Errorf("Error while creating patient: %v", err)
	//}

	doctor, err := graphql.CreateDoctor(model.CreateDoctorInput{
		Email:     "test_doctor_update_device_trust_add@edgar-sante.fr",
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
		t.Errorf("Error while creating patient: %v", err)
	}

	device := CreateDeviceConnect(input, doctor.ID)

	if device.Err != nil {
		t.Errorf("Expected no error, got: %v", device.Err)
	}
	if device.Code != 201 {
		t.Errorf("Expected status code 201, got: %d", device.Code)
	}

	input2 := CreateDoubleMobileInput{
		Methods:     "MOBILE",
		TrustDevice: device.DeviceConnect.ID,
	}

	rr := CreateDoubleAuthMobile(input2, doctor.ID)
	spew.Dump(rr)

	input3 := CreateDeviceConnectInput{
		DeviceType: "Windows",
		Browser:    "Chrome",
		Ip:         "192.168.0.1",
		City:       "sdfsdf",
		Country:    "dfdgdfgdfg",
		Date:       1627880400,
	}
	device2 := CreateDeviceConnect(input3, doctor.ID)

	response := AddTrustDevice(device2.DeviceConnect.ID, doctor.ID)
	spew.Dump(response)
	if response.Err != nil {
		t.Errorf("Expected no error, got: %v", response.Err)
	}
	if response.Code != 200 {
		t.Errorf("Expected status code 200, got: %d", response.Code)
	}
	if response.Doctor == nil {
		t.Errorf("Expected a doctor, got nil")
	}
	if response.Patient != nil {
		t.Errorf("Expected no doctor, got: %v", response.Doctor)
	}

	test, err := graphql.GetDoubleAuthById(rr.DoubleAuth.ID)
	spew.Dump(test)

	//ttt, err := graphql.GetDoubleAuthById(*doctor.DoubleAuthMethodsID)
	//spew.Dump(ttt)
}
