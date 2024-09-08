package auth

import (
	"github.com/edgar-care/edgarlib/v2/auth/utils"
	"github.com/edgar-care/edgarlib/v2/double_auth"
	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
	"github.com/go-chi/jwtauth/v5"
	"github.com/joho/godotenv"
	"log"
	"net/http/httptest"
	"testing"
	"time"
)

//
//import (
//	"github.com/edgar-care/edgarlib/v2/auth/utils"
//	"github.com/edgar-care/edgarlib/v2/graphql/model"
//	"github.com/go-chi/jwtauth/v5"
//	"github.com/joho/godotenv"
//	"github.com/lestrrat-go/jwx/v2/jwt"
//	"log"
//	"net/http/httptest"
//	"testing"
//)
//
////func TestGetAuthenticatedUser(t *testing.T) {
////	req := httptest.NewRequest("GET", "/", nil)
////
////	token := jwt.New()
////	token.Set("patient", "test_patient_email")
////	token.Set("id", "test_patient_id")
////
////	ctx := jwtauth.NewContext(req.Context(), token, nil)
////
////	req = req.WithContext(ctx)
////
////	patientID, _ := GetAuthenticatedAccount(token)
////
////	expectedPatientID := "test_patient_id"
////	if patientID != expectedPatientID {
////		t.Errorf("Expected patient ID: %s, got: %s", expectedPatientID, patientID)
////	}
////}
////
////func TestGetAuthenticatedUserError(t *testing.T) {
////	req := httptest.NewRequest("GET", "/", nil)
////
////	token := jwt.New()
////	token.Set("patient", map[string]interface{}{"id": "test_patient_id"})
////
////	ctx := jwtauth.NewContext(req.Context(), token, nil)
////
////	req = req.WithContext(ctx)
////
////	patientID := GetAuthenticatedUser(nil, req)
////
////	expectedPatientID := "invalid_id"
////	if patientID == expectedPatientID {
////		t.Errorf("Expected error but didn't get one")
////	}
////}
////
////func TestGetAuthenticatedUserEmpty(t *testing.T) {
////	req := httptest.NewRequest("GET", "/", nil)
////
////	token := jwt.New()
////
////	ctx := jwtauth.NewContext(req.Context(), token, nil)
////
////	req = req.WithContext(ctx)
////
////	patientID := GetAuthenticatedUser(nil, req)
////
////	if patientID != "" {
////		t.Errorf("Expected empty ID but got: %s", patientID)
////	}
////}
////
////func TestGetAuthenticatedUserNoId(t *testing.T) {
////	req := httptest.NewRequest("GET", "/", nil)
////
////	token := jwt.New()
////	token.Set("patient", map[string]interface{}{})
////
////	ctx := jwtauth.NewContext(req.Context(), token, nil)
////
////	req = req.WithContext(ctx)
////
////	patientID := GetAuthenticatedUser(nil, req)
////
////	if patientID != "" {
////		t.Errorf("Expected empty id but got: %s", patientID)
////	}
////}
//
//func TestNewTokenAuth(t *testing.T) {
//	if err := godotenv.Load(".env.test"); err != nil {
//		log.Fatalf("Error loading .env.test file: %v", err)
//	}
//
//	tokenAuth := NewTokenAuth()
//
//	if tokenAuth == nil {
//		t.Error("Expected token authentication object, got nil")
//	}
//
//}
//
//func TestVerifyToken(t *testing.T) {
//	if err := godotenv.Load(".env.test"); err != nil {
//		log.Fatalf("Error loading .env.test file: %v", err)
//	}
//
//	token, err := utils.CreateToken(map[string]interface{}{
//		"patient": model.Patient{
//			ID:       "id",
//			Email:    "test@example.com",
//			Password: "password",
//		},
//	})
//
//	if err != nil {
//		t.Error("Error creating token")
//	}
//
//	valid := VerifyToken(token)
//
//	if !valid {
//		t.Error("Expected token verification to succeed, got false")
//	}
//
//}
//
////func TestAuthMiddlewareWithValidToken(t *testing.T) {
////	if err := godotenv.Load(".env.test"); err != nil {
////		log.Fatalf("Error loading .env.test file: %v", err)
////	}
////
////	email := "test_bearer@example.com"
////	password := "password"
////
////	response := RegisterAndLoginPatient(email, password, "1256")
////	if response.Err != nil {
////		t.Error("Error trying to create account")
////	}
////	req := httptest.NewRequest("GET", "/", nil)
////	req.Header.Set("Authorization", "Bearer "+response.Token)
////
////	token := jwt.New()
////	token.Set("patient", "test_patient_email")
////	token.Set("id", "test_patient_id")
////
////	ctx := jwtauth.NewContext(req.Context(), token, nil)
////
////	req = req.WithContext(ctx)
////
////	rw := httptest.NewRecorder()
////
////	authenticatedUser := AuthMiddleware(rw, req)
////
////	if authenticatedUser == "" {
////		t.Error("Expected authenticated user, got empty string")
////	}
////}
////
////func TestAuthMiddlewareWithoutToken(t *testing.T) {
////	req := httptest.NewRequest("GET", "/", nil)
////
////	rw := httptest.NewRecorder()
////
////	authenticatedUser := AuthMiddleware(rw, req)
////
////	if authenticatedUser != "" {
////		t.Error("Expected empty string, got authenticated user")
////	}
////}
////
////func TestAuthMiddlewareWithInvalidToken(t *testing.T) {
////	req := httptest.NewRequest("GET", "/", nil)
////	req.Header.Set("Authorization", "Bearer invalid_token")
////
////	rw := httptest.NewRecorder()
////
////	authenticatedUser := AuthMiddleware(rw, req)
////
////	if authenticatedUser != "" {
////		t.Error("Expected empty string, got authenticated user")
////	}
////}
//
//func TestHashPassword(t *testing.T) {
//	if err := godotenv.Load(".env.test"); err != nil {
//		log.Fatalf("Error loading .env.test file: %v", err)
//	}
//	password := "password123"
//
//	hashedPassword := HashPassword(password)
//
//	if hashedPassword == "" {
//		t.Error("Hashed password is empty")
//	}
//
//	if hashedPassword == password {
//		t.Error("Hashed password is the same as the original password")
//	}
//}
//
////func TestGetAuthenticatedDoctor(t *testing.T) {
////	req := httptest.NewRequest("GET", "/", nil)
////
////	token := jwt.New()
////	token.Set("doctor", "test_doctor_email")
////	token.Set("id", "test_doctor_id")
////
////	ctx := jwtauth.NewContext(req.Context(), token, nil)
////
////	req = req.WithContext(ctx)
////
////	doctorID := GetAuthenticatedDoctor(nil, req)
////
////	expectedDoctorID := "test_doctor_id"
////	if doctorID != expectedDoctorID {
////		t.Errorf("Expected patient ID: %s, got: %s", expectedDoctorID, doctorID)
////	}
////}
//
////func TestGetAuthenticatedDoctorError(t *testing.T) {
////	req := httptest.NewRequest("GET", "/", nil)
////
////	token := jwt.New()
////	token.Set("doctor", map[string]interface{}{"id": "test_doctor_id"})
////
////	ctx := jwtauth.NewContext(req.Context(), token, nil)
////
////	req = req.WithContext(ctx)
////
////	doctorID := GetAuthenticatedDoctor(nil, req)
////
////	expectedDoctorID := "invalid_id"
////	if doctorID == expectedDoctorID {
////		t.Errorf("Expected error but didn't get one")
////	}
////}
//
//func TestAuthMiddlewareDoctorWithValidToken(t *testing.T) {
//	if err := godotenv.Load(".env.test"); err != nil {
//		log.Fatalf("Error loading .env.test file: %v", err)
//	}
//
//	email := "test_doctor_auth_middelware@example.com"
//	password := "password"
//	input := AddressInput{
//		Street:  "12",
//		ZipCode: "1234",
//		Country: "France",
//		City:    "City",
//	}
//
//	response := RegisterAndLoginDoctor(email, password, "test_doctor_middle", "auth", input, "1256")
//	if response.Err != nil {
//		t.Error("Error trying to create account")
//	}
//	req := httptest.NewRequest("GET", "/", nil)
//	req.Header.Set("Authorization", "Bearer "+response.Token)
//
//	token := jwt.New()
//	token.Set("doctor", "test_doctor_email")
//	token.Set("id", "test_doctor_id")
//
//	ctx := jwtauth.NewContext(req.Context(), token, nil)
//
//	req = req.WithContext(ctx)
//
//	rw := httptest.NewRecorder()
//
//	authenticatedUser := AuthMiddlewareDoctor(rw, req)
//	if authenticatedUser == "" {
//		t.Error("Expected authenticated user, got empty string")
//	}
//
//}

func TestBlackListDevice(t *testing.T) {

	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}

	patient, err := graphql.CreatePatient(model.CreatePatientInput{
		Email:    "test_blacklist_middleware@gmail.com",
		Password: "testtest",
		Status:   true,
	})
	if err != nil {
		t.Errorf("Error while creating patient: %v", err)
	}

	operationTime := time.Now().Unix() // Unix timestamp

	input := double_auth.CreateDeviceConnectInput{
		DeviceType: "Windows",
		Browser:    "Chrome",
		Ip:         "127.0.0.1",
		City:       "Lyon",
		Country:    "France",
		Date:       int(operationTime),
	}
	test := double_auth.CreateDeviceConnect(input, patient.ID)
	if test.Err != nil {
		t.Errorf("Error while creating device connect: %v", test.Err)

	}

	tokenString, err := utils.CreateToken(map[string]interface{}{
		"patient":     patient.Email,
		"id":          patient.ID,
		"name_device": test.DeviceConnect.ID,
	})

	if err != nil {
		t.Fatalf("Error creating token: %v", err)
	}

	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Set("Authorization", "Bearer "+tokenString)

	token, _, err := jwtauth.FromContext(req.Context())
	if err != nil {
		t.Fatalf("Error getting token from context: %v", err)
	}

	ctx := jwtauth.NewContext(req.Context(), token, nil)

	req = req.WithContext(ctx)

	input2 := double_auth.CreateDeviceConnectInput{
		DeviceType: "Linux",
		Browser:    "Firefox",
		Ip:         "127.0.0.1",
		City:       "Marseille",
		Country:    "France",
		Date:       int(operationTime),
	}
	_ = double_auth.CreateDeviceConnect(input2, patient.ID)
	if test.Err != nil {
		t.Errorf("Error while creating device connect: %v", test.Err)

	}

	get := double_auth.GetDeviceConnect(patient.ID)
	if get.Err != nil {
		t.Errorf("Error getting device connect: %v", get.Err)
	}

	_ = double_auth.DeleteDeviceConnect(test.DeviceConnect.ID, patient.ID)

	_ = BlackListDevice(tokenString, patient.ID)
	//if authenticatedUser.Code == 401 {
	//	t.Errorf("Expected status code: 401 but got: %d", authenticatedUser.Code)
	//}

}
