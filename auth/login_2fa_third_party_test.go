package auth

import (
	"github.com/edgar-care/edgarlib/v2/double_auth"
	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
	"github.com/joho/godotenv"
	"github.com/pquerna/otp/totp"
	"log"
	"testing"
	"time"
)

func TestLogin2faThirdParty(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}

	email := "test_login_2fa_third_party_success@example.com"
	password := "passwordtesttest"
	haspassword := HashPassword(password)

	patient, err := graphql.CreatePatient(model.CreatePatientInput{
		Email:    email,
		Password: haspassword,
		Status:   true,
	})
	if err != nil {
		t.Errorf("Error while creating patient: %v", err)
	}

	test := double_auth.CreateDoubleAuthAppTier(patient.ID)

	token, err := totp.GenerateCode(test.TotpInfo.Secret, time.Now().UTC())
	if err != nil {
		t.Errorf("Error while generating token: %v", err)
	}

	input := Login2faThirdPartyInput{
		Email:    email,
		Password: password,
		Token2fa: token,
	}

	response := Login2faThirdParty(input, "test_device", patient.ID)
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
