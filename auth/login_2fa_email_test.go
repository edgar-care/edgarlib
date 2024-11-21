package auth

import (
	"github.com/edgar-care/edgarlib/v2/double_auth"
	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
	"github.com/edgar-care/edgarlib/v2/redis"
	"github.com/joho/godotenv"
	"log"
	"testing"
)

func TestLogin2faEmailPatient(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}

	email := "test_login_2fa_email@example.com"
	password := "testtestpassword"

	haspassword := HashPassword(password)

	patient, err := graphql.CreatePatient(model.CreatePatientInput{
		Email:    email,
		Password: haspassword,
		Status:   true,
	})
	if err != nil {
		t.Errorf("Error while creating patient: %v", err)
	}

	inputDoubleAuth := double_auth.CreateDoubleAuthInput{
		Methods: "EMAIL",
	}

	_ = double_auth.CreateDoubleAuthEmail(inputDoubleAuth, patient.ID)

	testSendEmail := Email2faAuth(email)
	if testSendEmail.Code == 500 {
		t.Errorf("Expected no error, but got: %v", testSendEmail)

	}

	getToken2fa, err := redis.GetKey(email)
	if err != nil {
		t.Errorf("Failed to get key from Redis")
	}

	input := Login2faEmailInput{
		Email:    email,
		Password: password,
		Token2fa: getToken2fa,
	}

	response := Login2faEmail(input, "test_device", patient.ID)
	if response.Token == "" {
		t.Error("Expected token to be non-empty, but got an empty token")
	}
	if response.Code != 200 {
		t.Errorf("Expected status code: 200, but got: %d", response.Code)
	}
	if response.Err != nil {
		t.Errorf("Expected no error, but got: %v", response.Err)
	}

}

func TestLogin2faEmailDoctor(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}

	email := "test_login_2fa_email_doctor@example.com"
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

	inputDoubleAuth := double_auth.CreateDoubleAuthInput{
		Methods: "EMAIL",
	}

	_ = double_auth.CreateDoubleAuthEmail(inputDoubleAuth, doctor.ID)

	testSendEmail := Email2faAuth(email)
	if testSendEmail.Code == 500 {
		t.Errorf("Expected no error, but got: %v", testSendEmail)

	}

	getToken2fa, err := redis.GetKey(email)
	if err != nil {
		t.Errorf("Failed to get key from Redis")
	}

	input := Login2faEmailInput{
		Email:    email,
		Password: password,
		Token2fa: getToken2fa,
	}

	response := Login2faEmail(input, "test_device", doctor.ID)
	if response.Token == "" {
		t.Error("Expected token to be non-empty, but got an empty token")
	}
	if response.Code != 200 {
		t.Errorf("Expected status code: 200, but got: %d", response.Code)
	}
	if response.Err != nil {
		t.Errorf("Expected no error, but got: %v", response.Err)
	}

}

func TestLogin2faEmail_InvalidDoubleAuth(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}

	email := "test_login_2fa_email_doctor_invalid_double_auth@example.com"
	password := "password"

	_, err := graphql.CreateDoctor(model.CreateDoctorInput{
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

	testSendEmail := Email2faAuth(email)
	if testSendEmail.Code == 500 {
		t.Errorf("Expected no error, but got: %v", testSendEmail)

	}

	getToken2fa, err := redis.GetKey(email)
	if err != nil {
		t.Errorf("Failed to get key from Redis")
	}

	input := Login2faEmailInput{
		Email:    email,
		Password: password,
		Token2fa: getToken2fa,
	}

	response := Login2faEmail(input, "test_device", "invalid_id")
	if response.Err == nil {
		t.Error("Expected an error, but got no error")
	}
	if response.Code != 400 {
		t.Errorf("Expected status code: 400, but got: %d", response.Code)
	}

}

func TestLogin2faEmail_InvalidToken2fa(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}

	email := "test_login_2fa_email_invalid_Token2fa@example.com"
	password := "password"

	patient, err := graphql.CreatePatient(model.CreatePatientInput{
		Email:    email,
		Password: password,
		Status:   true,
	})
	if err != nil {
		t.Errorf("Error while creating patient: %v", err)
	}

	inputDoubleAuth := double_auth.CreateDoubleAuthInput{
		Methods: "EMAIL",
	}

	_ = double_auth.CreateDoubleAuthEmail(inputDoubleAuth, patient.ID)

	testSendEmail := Email2faAuth(email)
	if testSendEmail.Code == 500 {
		t.Errorf("Expected no error, but got: %v", testSendEmail)

	}

	input := Login2faEmailInput{
		Email:    email,
		Password: password,
		Token2fa: "12",
	}

	response := Login2faEmail(input, "test_device", patient.ID)
	if response.Code != 400 {
		t.Error("Expected token to be different, but got the same token")
	}
	if response.Err == nil {
		t.Error("Expected an error, but got no error")
	}

}
