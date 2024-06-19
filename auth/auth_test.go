package auth

import (
	"github.com/edgar-care/edgarlib/auth/utils"
	"github.com/edgar-care/edgarlib/graphql/server/model"
	"github.com/go-chi/jwtauth/v5"
	"github.com/joho/godotenv"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"log"
	"net/http/httptest"
	"testing"
)

func TestGetAuthenticatedUser(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)

	token := jwt.New()
	token.Set("patient", map[string]interface{}{"id": "test_patient_id"})

	ctx := jwtauth.NewContext(req.Context(), token, nil)

	req = req.WithContext(ctx)

	patientID := GetAuthenticatedUser(nil, req)

	expectedPatientID := "test_patient_id"
	if patientID != expectedPatientID {
		t.Errorf("Expected patient ID: %s, got: %s", expectedPatientID, patientID)
	}
}

func TestGetAuthenticatedUserError(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)

	token := jwt.New()
	token.Set("patient", map[string]interface{}{"id": "test_patient_id"})

	ctx := jwtauth.NewContext(req.Context(), token, nil)

	req = req.WithContext(ctx)

	patientID := GetAuthenticatedUser(nil, req)

	expectedPatientID := "invalid_id"
	if patientID == expectedPatientID {
		t.Errorf("Expected error but didn't get one")
	}
}

func TestGetAuthenticatedUserEmpty(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)

	token := jwt.New()

	ctx := jwtauth.NewContext(req.Context(), token, nil)

	req = req.WithContext(ctx)

	patientID := GetAuthenticatedUser(nil, req)

	if patientID != "" {
		t.Errorf("Expected empty ID but got: %s", patientID)
	}
}

func TestGetAuthenticatedUserNoId(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)

	token := jwt.New()
	token.Set("patient", map[string]interface{}{})

	ctx := jwtauth.NewContext(req.Context(), token, nil)

	req = req.WithContext(ctx)

	patientID := GetAuthenticatedUser(nil, req)

	if patientID != "" {
		t.Errorf("Expected empty id but got: %s", patientID)
	}
}

func TestNewTokenAuth(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}

	tokenAuth := NewTokenAuth()

	if tokenAuth == nil {
		t.Error("Expected token authentication object, got nil")
	}

}

func TestVerifyToken(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}

	token, err := utils.CreateToken(map[string]interface{}{
		"patient": model.Patient{
			ID:       "id",
			Email:    "test@example.com",
			Password: "password",
		},
	})

	if err != nil {
		t.Error("Error creating token")
	}

	valid := VerifyToken(token)

	if !valid {
		t.Error("Expected token verification to succeed, got false")
	}

}

func TestAuthMiddlewareWithValidToken(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}

	email := "test_bearer@example.com"
	password := "password"

	response := RegisterAndLoginPatient(email, password, "1256")
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

	authenticatedUser := AuthMiddleware(rw, req)

	if authenticatedUser == "" {
		t.Error("Expected authenticated user, got empty string")
	}
}

func TestAuthMiddlewareWithoutToken(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)

	rw := httptest.NewRecorder()

	authenticatedUser := AuthMiddleware(rw, req)

	if authenticatedUser != "" {
		t.Error("Expected empty string, got authenticated user")
	}
}

func TestAuthMiddlewareWithInvalidToken(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Set("Authorization", "Bearer invalid_token")

	rw := httptest.NewRecorder()

	authenticatedUser := AuthMiddleware(rw, req)

	if authenticatedUser != "" {
		t.Error("Expected empty string, got authenticated user")
	}
}

func TestHashPassword(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}
	password := "password123"

	hashedPassword := HashPassword(password)

	if hashedPassword == "" {
		t.Error("Hashed password is empty")
	}

	if hashedPassword == password {
		t.Error("Hashed password is the same as the original password")
	}
}
