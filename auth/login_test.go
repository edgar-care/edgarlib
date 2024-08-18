package auth

import (
	"github.com/edgar-care/edgarlib/v2/auth/utils"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"testing"
)

func TestLoginDoctor_Success(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}

	email := "test_login@example.com"
	password := "password"
	name := "last"
	firstname := "first"
	address := AddressInput{"123 Street", "City", "Country", "City"}

	doctor, err := RegisterDoctor(email, password, name, firstname, address)

	if err != nil {
		t.Error("Error trying to create account")
	}

	response := Login(LoginInput{
		Email:    doctor.Email,
		Password: password,
	}, "d", "")

	if response.Token == "" {
		t.Error("Expected token to be non-empty, but got an empty token")
	}

	if response.Code != http.StatusOK {
		t.Errorf("Expected status code: %d, but got: %d", http.StatusOK, response.Code)
	}

	if response.Err != nil {
		t.Errorf("Expected no error, but got: %v", response.Err)
	}
}

func TestLoginDoctor_MismatchError(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}

	response := Login(LoginInput{
		Email:    "test_login@example.com",
		Password: "invalid",
	}, "d", "")

	if response.Token != "" {
		t.Error("Expected token to be empty, but got an non-empty token")
	}

	if response.Code != http.StatusUnauthorized {
		t.Errorf("Expected status code: %d, but got: %d", http.StatusUnauthorized, response.Code)
	}

	if response.Err == nil {
		t.Error("Expected error, but got no error")
	}
}

func TestLoginDoctor_Error(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}

	response := Login(LoginInput{
		Email:    "inexistant@email.com",
		Password: "invalid",
	}, "d", "")

	if response.Token != "" {
		t.Error("Expected token to be empty, but got an non-empty token")
	}

	if response.Code != http.StatusBadRequest {
		t.Errorf("Expected status code: %d, but got: %d", http.StatusBadRequest, response.Code)
	}

	if response.Err == nil {
		t.Error("Expected error, but got no error")
	}
}

func TestLoginPatient_Success(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}

	email := "test_login@example.com"
	password := "password"

	patient, err := RegisterPatient(email, password)

	if err != nil {
		t.Error("Error trying to create account")
	}

	response := Login(LoginInput{
		Email:    patient.Email,
		Password: password,
	}, "p", "")

	if response.Token == "" {
		t.Error("Expected token to be non-empty, but got an empty token")
	}

	if response.Code != http.StatusOK {
		t.Errorf("Expected status code: %d, but got: %d", http.StatusOK, response.Code)
	}

	if response.Err != nil {
		t.Errorf("Expected no error, but got: %v", response.Err)
	}
}

func TestLoginPatient_Error(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}

	response := Login(LoginInput{
		Email:    "inexistant@email.com",
		Password: "invalid",
	}, "p", "")

	if response.Token != "" {
		t.Error("Expected token to be empty, but got an non-empty token")
	}

	if response.Code != http.StatusBadRequest {
		t.Errorf("Expected status code: %d, but got: %d", http.StatusBadRequest, response.Code)
	}

	if response.Err == nil {
		t.Error("Expected error, but got no error")
	}
}

func TestLoginAdmin_Success(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}

	email := "test_login@example.com"
	password := "password"
	firstName := "first"
	lastName := "last"
	token, err := utils.CreateToken(map[string]interface{}{
		"admin": model.Admin{
			ID:       "id",
			Email:    email,
			Password: password,
			Name:     firstName,
			LastName: lastName,
		},
	})
	if err != nil {
		t.Errorf("Expected no error, but got: %v", err)
	}
	admin, err := RegisterAdmin(email, password, firstName, lastName, token)

	if err != nil {
		t.Error("Error trying to create account")
	}

	response := Login(LoginInput{
		Email:    admin.Email,
		Password: password,
	}, "a", "")

	if response.Token == "" {
		t.Error("Expected token to be non-empty, but got an empty token")
	}

	if response.Code != http.StatusOK {
		t.Errorf("Expected status code: %d, but got: %d", http.StatusOK, response.Code)
	}

	if response.Err != nil {
		t.Errorf("Expected no error, but got: %v", response.Err)
	}
}

func TestLoginAdmin_Error(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}

	response := Login(LoginInput{
		Email:    "inexistant@email.com",
		Password: "invalid",
	}, "a", "")

	if response.Token != "" {
		t.Errorf("Expected token to be empty, but got %s:", response.Token)
	}

	if response.Code != http.StatusBadRequest {
		t.Errorf("Expected status code: %d, but got: %d", http.StatusBadRequest, response.Code)
	}

	if response.Err == nil {
		t.Error("Expected error, but got no error")
	}
}
