package auth

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"github.com/edgar-care/edgarlib/v2/auth/utils"
	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
	"github.com/go-chi/jwtauth/v5"
	"github.com/joho/godotenv"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"log"
	"net/http/httptest"
	"testing"
)

func TestCreateBackupCodes(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}

	patient, err := graphql.CreatePatient(model.CreatePatientInput{
		Email:    "test_patient_create_save_code@edgar-sante.fr",
		Password: "password",
		Status:   false,
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

	response := CreateBackupCodes(patient.ID, req)
	if response.Code != 201 {
		t.Errorf("Expected code 200 but got %d", response.Code)
	}
	if response.Err != nil {
		t.Errorf("Expected no error but got: %s", response.Err.Error())
	}

}

func TestCreateBackupCodesInvalidId(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)

	token := jwt.New()
	token.Set("patient", map[string]interface{}{"id": "patientId"})

	ctx := jwtauth.NewContext(req.Context(), token, nil)

	req = req.WithContext(ctx)
	response := CreateBackupCodes("patientId", req)
	if response.Code != 400 || response.Err == nil {
		t.Errorf("Expected error and code 400 but got code %d and err: %s", response.Code, response.Err.Error())
	}
}

func TestGenerateRandomCode(t *testing.T) {
	// Test successful code generation
	length := 8
	code, err := generateRandomCode(length)
	if err != nil {
		t.Errorf("Expected no error but got: %s", err.Error())
	}
	if len(code) != length {
		t.Errorf("Expected code of length %d but got %d", length, len(code))
	}
}

func TestGenerateBackupCodes(t *testing.T) {
	// Test successful backup code generation
	codes, err := generateBackupCodes()
	if err != nil {
		t.Errorf("Expected no error but got: %s", err.Error())
	}
	const expectedNumCodes = 10
	if len(codes) != expectedNumCodes {
		t.Errorf("Expected %d codes but got %d", expectedNumCodes, len(codes))
	}
	for _, code := range codes {
		const expectedCodeLength = 8
		if len(code) != expectedCodeLength {
			t.Errorf("Expected code length of %d but got %d", expectedCodeLength, len(code))
		}
	}
}

func TestHashCode(t *testing.T) {
	// Test hashing of a code
	code := "testcode"
	hashed := hashCode(code)
	expectedHash := sha256.Sum256([]byte(code))
	expectedHashString := hex.EncodeToString(expectedHash[:])
	if hashed != expectedHashString {
		t.Errorf("Expected hash %s but got %s", expectedHashString, hashed)
	}
}

func TestGenerateRandomCode_ErrorHandling(t *testing.T) {
	// Test if the function correctly handles an error case
	originalRandReader := rand.Reader
	defer func() { rand.Reader = originalRandReader }()
	rand.Reader = &errorReader{}

	_, err := generateRandomCode(8)
	if err == nil {
		t.Errorf("Expected an error but got none")
	}
}

// errorReader is a helper struct to simulate an error during random code generation
type errorReader struct{}

func (r *errorReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("simulated error")
}
