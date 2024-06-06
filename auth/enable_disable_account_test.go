package auth

import (
	"github.com/edgar-care/edgarlib/graphql"
	"github.com/edgar-care/edgarlib/graphql/model"
	"github.com/joho/godotenv"
	"log"
	"testing"
)

func TestModifyStatusAccountPatient(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}

	patient, err := graphql.CreatePatient(model.CreatePatientInput{
		Email:    "test_patient_disable_patient_code@edgar-sante.fr",
		Password: "password",
		Status:   true,
	})
	if err != nil {
		t.Errorf("Error while creating patient: %v", err)
	}

	response := ModifyStatusAccount(patient.ID, false)

	if response.Code != 200 {
		t.Errorf("Expected code 200 but got %d", response.Code)
	}
	if response.Err != nil {
		t.Errorf("Expected no error but got: %s", response.Err.Error())
	}

}

func TestModifyStatusAccountDoctor(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}

	doctor, err := graphql.CreateDoctor(model.CreateDoctorInput{
		Email:     "test_doctor_eanble_doctor@edgar-sante.fr",
		Password:  "password",
		Name:      "name",
		Firstname: "first",
		Address: &model.AddressInput{
			Street:  "",
			ZipCode: "",
			Country: "",
			City:    "",
		},
		Status: false,
	})
	if err != nil {
		t.Errorf("Error while creating patient: %v", err)
	}

	response := ModifyStatusAccount(doctor.ID, true)

	if response.Code != 200 {
		t.Errorf("Expected code 200 but got %d", response.Code)
	}
	if response.Err != nil {
		t.Errorf("Expected no error but got: %s", response.Err.Error())
	}

}

func TestCheckAccountEnablePatientGood(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}

	patient, err := graphql.CreatePatient(model.CreatePatientInput{
		Email:    "test_patient_check_enable_code@edgar-sante.fr",
		Password: "password",
		Status:   true,
	})
	if err != nil {
		t.Errorf("Error while creating patient: %v", err)
	}

	response := CheckAccountEnable(patient.ID)
	if response.Code != 200 {
		t.Errorf("Expected code 200 but got %d", response.Code)
	}

	if response.Err != nil {
		t.Errorf("Expected no error but got: %s", response.Err.Error())
	}

}

func TestCheckAccountEnablePatientNotGood(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}

	patient, err := graphql.CreatePatient(model.CreatePatientInput{
		Email:    "test_patient_check_disable_code@edgar-sante.fr",
		Password: "password",
		Status:   false,
	})
	if err != nil {
		t.Errorf("Error while creating patient: %v", err)
	}

	response := CheckAccountEnable(patient.ID)
	if response.Code != 409 {
		t.Errorf("Expected code 409 but got %d", response.Code)
	}

	if response.Err == nil {
		t.Errorf("Expected error but got none")
	}

}

func TestCheckAccountEnableDoctorNotGood(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}

	doctor, err := graphql.CreateDoctor(model.CreateDoctorInput{
		Email:     "test_doctor_enable_doctor_not_good@edgar-sante.fr",
		Password:  "password",
		Name:      "name",
		Firstname: "first",
		Address: &model.AddressInput{
			Street:  "",
			ZipCode: "",
			Country: "",
			City:    "",
		},
		Status: false,
	})
	if err != nil {
		t.Errorf("Error while creating patient: %v", err)
	}

	response := CheckAccountEnable(doctor.ID)
	if response.Code != 409 {
		t.Errorf("Expected code 200 but got %d", response.Code)
	}

	if response.Err == nil {
		t.Errorf("Expected error but got none")
	}

}

func TestCheckAccountEnableDoctorGood(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}

	doctor, err := graphql.CreateDoctor(model.CreateDoctorInput{
		Email:     "test_doctor_enable_doctor_good@edgar-sante.fr",
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

	response := CheckAccountEnable(doctor.ID)
	if response.Code != 200 {
		t.Errorf("Expected code 200 but got %d", response.Code)
	}

	if response.Err != nil {
		t.Errorf("Expected no error but got: %s", response.Err.Error())
	}

}

func TestModifyStatusAccountInvalidId(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}

	response := ModifyStatusAccount("id", false)

	if response.Code != 400 || response.Err == nil {
		t.Errorf("Expected error and code 400 but got code %d and err: %s", response.Code, response.Err.Error())
	}

}
