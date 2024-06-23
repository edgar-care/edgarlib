package auth

import (
	"encoding/json"
	"github.com/go-chi/jwtauth/v5"
	"github.com/joho/godotenv"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAuthenticateRequest(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}

	email := "test_checker_login@example.com"
	password := "password"
	response := RegisterAndLoginPatient(email, password, "12345")
	if response.Err != nil {
		t.Error("Error trying to create account")
	}

	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Set("Authorization", "Bearer "+response.Token)

	token := jwt.New()
	token.Set("patient", map[string]interface{}{"id": "test_patient_id"})

	ctx := jwtauth.NewContext(req.Context(), token, nil)

	req = req.WithContext(ctx)

	rw := httptest.NewRecorder()

	patientID, isAuthenticated := AuthenticateRequest(rw, req)

	if patientID == "" {
		t.Error("Expected non-empty patientID")
	}

	if !isAuthenticated {
		t.Error("Expected isAuthenticated to be true")
	}
}

func TestAuthenticateRequestEmptyPatientID(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}

	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Set("Authorization", "Bearer invalid_token")

	rw := httptest.NewRecorder()

	patientID, isAuthenticated := AuthenticateRequest(rw, req)

	if patientID != "" {
		t.Error("Expected empty patientID")
	}

	if isAuthenticated {
		t.Error("Expected isAuthenticated to be false")
	}

	if rw.Code != http.StatusUnauthorized {
		t.Errorf("Expected status code %d, got %d", http.StatusUnauthorized, rw.Code)
	}

	var responseBody map[string]string
	if err := json.Unmarshal(rw.Body.Bytes(), &responseBody); err != nil {
		t.Errorf("Error parsing response body: %v", err)
	}

	expectedMessage := "Not authenticated"
	if responseBody["message"] != expectedMessage {
		t.Errorf("Expected message %q, got %q", expectedMessage, responseBody["message"])
	}
}
