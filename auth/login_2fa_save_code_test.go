package auth

import (
	"github.com/edgar-care/edgarlib/v2/auth/utils"
	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
	"github.com/joho/godotenv"
	"log"
	"net/http/httptest"
	"testing"
)

func TestLogin2faSaveCode(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}

	email := "test_login_2fa_save_code@example.com"
	password := "passwordtest"
	haspassword := HashPassword(password)

	patient, err := graphql.CreatePatient(model.CreatePatientInput{
		Email:    email,
		Password: haspassword,
		Status:   true,
	})
	if err != nil {
		t.Errorf("Error while creating patient: %v", err)
	}

	req := httptest.NewRequest("GET", "/", nil)

	token, err := utils.CreateToken(map[string]interface{}{
		"patient":     patient.Email,
		"id":          patient.ID,
		"name_device": "nameDevice",
	})
	req.Header.Set("Authorization", "Bearer "+token)

	saveCode := CreateBackupCodes(patient.ID, req)
	UseSaveCode := saveCode.SaveCode.Code[0]

	input := Login2faSaveCodeInput{
		Email:      email,
		Password:   password,
		BackupCode: UseSaveCode,
	}

	response := Login2faSaveCode(input, "test_device")
	if response.Token == "" {
		t.Error("Expected token to be non-empty, but got an empty token")
	}
	if response.Code != 200 {
		t.Errorf("Expected status code: 200, but got: %d", response.Code)
	}
	if response.Err != nil {
		t.Errorf("Expected no error, but got: %v", response)
	}

}

func TestLogin2faSaveCodeDoctor(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}

	email := "test_login_2fa_save_code_doctor@example.com"
	password := "password"
	haspassword := HashPassword(password)

	doctor, err := graphql.CreateDoctor(model.CreateDoctorInput{
		Email:     email,
		Password:  haspassword,
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

	req := httptest.NewRequest("GET", "/", nil)

	token, err := utils.CreateToken(map[string]interface{}{
		"doctor":      doctor.Email,
		"id":          doctor.ID,
		"name_device": "nameDevice",
	})
	req.Header.Set("Authorization", "Bearer "+token)

	saveCode := CreateBackupCodes(doctor.ID, req)
	UseSaveCode := saveCode.SaveCode.Code[0]

	input := Login2faSaveCodeInput{
		Email:      email,
		Password:   password,
		BackupCode: UseSaveCode,
	}

	response := Login2faSaveCode(input, "test_device")
	if response.Token == "" {
		t.Error("Expected token to be non-empty, but got an empty token")
	}
	if response.Code != 200 {
		t.Errorf("Expected status code: 200, but got: %d", response.Code)
	}
	if response.Err != nil {
		t.Errorf("Expected no error, but got: %v", response)
	}

}

func TestLogin2faSaveCode_InvalidCode(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}

	email := "test_login_2fa_save_code_doctor_invalid_code@example.com"
	password := "password"

	doctor, err := graphql.CreateDoctor(model.CreateDoctorInput{
		Email:     email,
		Password:  password,
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

	req := httptest.NewRequest("GET", "/", nil)

	token, err := utils.CreateToken(map[string]interface{}{
		"doctor":      doctor.Email,
		"id":          doctor.ID,
		"name_device": "nameDevice",
	})
	req.Header.Set("Authorization", "Bearer "+token)

	input := Login2faSaveCodeInput{
		Email:      email,
		Password:   password,
		BackupCode: "Code",
	}

	response := Login2faSaveCode(input, "test_device")
	if token == "" {
		t.Error("Expected token to be empty, but got a non-empty token")
	}
	if response.Code != 400 {
		t.Errorf("Expected status code: 400, but got: %d", response.Code)
	}

}
