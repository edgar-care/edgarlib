package auth

import (
	"github.com/edgar-care/edgarlib/v2/auth/utils"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"testing"
)

func TestRegisterAndLoginDoctor_Success(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}

	email := "test@example.com"
	password := "password"
	name := "last"
	firstname := "first"
	address := AddressInput{"123 Street", "City", "Country", "City"}
	nameDevice := "testDevice"

	response := RegisterAndLoginDoctor(email, password, name, firstname, address, nameDevice)

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

func TestRegisterAndLoginDoctor_Error(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}

	email := "test@example.com"
	password := "password"
	name := "last"
	firstname := "first"
	address := AddressInput{"123 Street", "City", "Country", "City"}
	nameDevice := "testDevice"

	response := RegisterAndLoginDoctor(email, password, name, firstname, address, nameDevice)

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

func TestRegisterAndLoginPatient_Success(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}

	email := "test@example.com"
	password := "password"

	response := RegisterAndLoginPatient(email, password, "123")

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

func TestRegisterAndLoginPatient_Error(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}

	email := "test@example.com"
	password := "password"

	response := RegisterAndLoginPatient(email, password, "1345")

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

func TestRegisterAndLoginAdmin_Success(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}

	email := "test@example.com"
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
	response := RegisterAndLoginAdmin(email, password, firstName, lastName, token)

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

func TestRegisterAndLoginAdmin_EmptyToken(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}

	email := "test@example.com"
	password := "password"
	firstName := "first"
	lastName := "last"
	token := ""
	response := RegisterAndLoginAdmin(email, password, firstName, lastName, token)

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

func TestRegisterAndLoginAdmin_Error(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}

	email := "test@example.com"
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
	response := RegisterAndLoginAdmin(email, password, firstName, lastName, token)
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
