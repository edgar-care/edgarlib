package auth

import (
	"github.com/joho/godotenv"
	"log"
	"testing"
)

func TestUpdatePassword(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}

	patient, err := RegisterPatient("test_patient_update_password_succes", "password")
	if err != nil {
		t.Errorf("Error while creating patient: %v", err)
	}
	input := UpdatePasswordInput{
		OldPassword: "password",
		NewPassword: "newpassword",
	}

	response := UpdatePassword(input, patient.ID)
	if response.Code != 200 {
		t.Errorf("Expected code 200 but got %d", response.Code)
	}
	if response.Err != nil {
		t.Errorf("Expected no error but got: %s", response.Err.Error())
	}

}

func TestUpdatePassword_Wrong(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}

	patient, err := RegisterPatient("test_patient_update_password_wrong", "password")

	if err != nil {
		t.Errorf("Error while creating patient: %v", err)
	}
	input := UpdatePasswordInput{
		OldPassword: "testpassword",
		NewPassword: "newpassword",
	}

	response := UpdatePassword(input, patient.ID)
	if response.Code != 403 || response.Err == nil {
		t.Errorf("Expected error and code 400 but got code %d and err: %s", response.Code, response.Err.Error())
	}
}

func TestUpdatePasswordDoctor(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}

	doctor, err := RegisterDoctor("test_doctor_update_password_doctor_success", "password", "name", "first", AddressInput{"", "", "", ""})
	if err != nil {
		t.Errorf("Error while creating patient: %v", err)
	}

	input := UpdatePasswordInput{
		OldPassword: "password",
		NewPassword: "newpassword",
	}

	response := UpdatePassword(input, doctor.ID)
	if response.Code != 200 {
		t.Errorf("Expected code 200 but got %d", response.Code)
	}
	if response.Err != nil {
		t.Errorf("Expected no error but got: %s", response.Err.Error())
	}

}

func TestUpdatePasswordDoctor_Wrong(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}

	doctor, err := RegisterDoctor("test_doctor_update_password_doctor_wrong", "password", "name", "first", AddressInput{"", "", "", ""})

	if err != nil {
		t.Errorf("Error while creating patient: %v", err)
	}

	input := UpdatePasswordInput{
		OldPassword: "testpassword",
		NewPassword: "newpassword",
	}

	response := UpdatePassword(input, doctor.ID)
	if response.Code != 403 || response.Err == nil {
		t.Errorf("Expected error and code 400 but got code %d and err: %s", response.Code, response.Err.Error())
	}
}

func TestUpdatePassword_InvalidId(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}

	input := UpdatePasswordInput{
		OldPassword: "testpassword",
		NewPassword: "newpassword",
	}

	response := UpdatePassword(input, "invalidId")
	if response.Code != 400 || response.Err == nil {
		t.Errorf("Expected error and code 400 but got code %d and err: %s", response.Code, response.Err.Error())
	}
}
