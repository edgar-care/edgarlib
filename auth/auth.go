package auth

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"github.com/edgar-care/edgarlib/v2/auth/utils"
	"github.com/edgar-care/edgarlib/v2/double_auth"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/go-chi/jwtauth/v5"
	"golang.org/x/crypto/bcrypt"
)

type AuthenticatedAccountResponse struct {
	ID   string
	Code int
	Err  error
}

type BlackListDeviceResponse struct {
	Code int
	Err  error
}

func NewTokenAuth() *jwtauth.JWTAuth {
	tokenAuth := jwtauth.New("HS256", []byte(os.Getenv("JWT_SECRET")), nil)
	return tokenAuth
}

func CreateToken(claims map[string]interface{}) (string, error) {
	_, token, err := NewTokenAuth().Encode(claims)
	return token, err
}

func VerifyToken(tokenString string) bool {
	token, err := jwtauth.VerifyToken(NewTokenAuth(), tokenString)
	if err != nil || token == nil {
		return false
	}
	return true
}

func HashPassword(password string) string {
	salt, _ := strconv.Atoi(os.Getenv("SALT"))
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), salt)
	return string(bytes)
}

func CheckPassword(password string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func AuthMiddlewarePatient(w http.ResponseWriter, r *http.Request) AuthenticatedAccountResponse {
	reqToken := r.Header.Get("Authorization")
	if reqToken == "" {
		return AuthenticatedAccountResponse{ID: "", Code: 401, Err: errors.New("authorization token missing")}
	}
	splitToken := strings.Split(reqToken, "Bearer ")
	reqToken = splitToken[1]

	if VerifyToken(reqToken) == false {
		return AuthenticatedAccountResponse{ID: "", Code: 401, Err: errors.New("invalid token")}
	}
	accountID, typeAccount := GetAuthenticatedAccount(reqToken)
	if typeAccount != "patient" {
		return AuthenticatedAccountResponse{ID: "", Code: 403, Err: errors.New("Not authorized, this account is not a patient")}
	}

	check_account := CheckAccountEnable(accountID)
	if check_account.Code == 409 {
		return AuthenticatedAccountResponse{ID: "", Code: 409, Err: errors.New("Not authorized, this account is disable")}

	}

	check_device := BlackListDevice(reqToken, accountID)
	if check_device.Code == 401 {
		return AuthenticatedAccountResponse{ID: "", Code: 401, Err: errors.New("Not authorized, this device is not connected")}
	}

	utils.DeviceConnectMiddleware(w, r, accountID)

	return AuthenticatedAccountResponse{accountID, 200, nil}
}

func GetBearerToken(req *http.Request) string {
	authHeader := req.Header.Get("Authorization")
	if authHeader == "" {
		return ""
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return ""
	}

	return parts[1]
}

func GetAccountType(token string) (string, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return "", errors.New("invalid token format")
	}
	payload, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return "", err
	}
	var claims map[string]interface{}
	if err := json.Unmarshal(payload, &claims); err != nil {
		return "", err
	}
	if _, valid := claims["patient"].(string); valid {
		return "patient", nil
	}
	if _, valid := claims["doctor"].(string); valid {
		return "doctor", nil
	}

	return "", errors.New("no account type found")
}

func AuthMiddlewareDoctor(w http.ResponseWriter, r *http.Request) AuthenticatedAccountResponse {
	reqToken := r.Header.Get("Authorization")
	if reqToken == "" {
		return AuthenticatedAccountResponse{ID: "", Code: 401, Err: errors.New("authorization token missing")}
	}
	splitToken := strings.Split(reqToken, "Bearer ")
	reqToken = splitToken[1]

	if VerifyToken(reqToken) == false {
		return AuthenticatedAccountResponse{ID: "", Code: 401, Err: errors.New("invalid token")}
	}
	accountID, typeAccount := GetAuthenticatedAccount(reqToken)
	if typeAccount != "doctor" {
		return AuthenticatedAccountResponse{ID: "", Code: 403, Err: errors.New("Not authorized, this account is not a doctor")}
	}

	check_account := CheckAccountEnable(accountID)
	if check_account.Code == 409 {
		return AuthenticatedAccountResponse{ID: "", Code: 409, Err: errors.New("Not authorized, this account is disable")}
	}
	check_device := BlackListDevice(reqToken, accountID)
	if check_device.Code == 401 {
		return AuthenticatedAccountResponse{ID: "", Code: 401, Err: errors.New("Not authorized, this device is not connected")}
	}

	utils.DeviceConnectMiddleware(w, r, accountID)

	return AuthenticatedAccountResponse{accountID, 200, nil}
}

func AuthMiddlewareAccount(w http.ResponseWriter, r *http.Request) AuthenticatedAccountResponse {
	reqToken := r.Header.Get("Authorization")
	if reqToken == "" {
		return AuthenticatedAccountResponse{ID: "", Code: 401, Err: errors.New("authorization token missing")}
	}
	splitToken := strings.Split(reqToken, "Bearer ")
	reqToken = splitToken[1]

	if VerifyToken(reqToken) == false {
		return AuthenticatedAccountResponse{ID: "", Code: 401, Err: errors.New("invalid token")}
	}
	accountID, _ := GetAuthenticatedAccount(reqToken)
	check_account := CheckAccountEnable(accountID)
	if check_account.Code == 409 {
		return AuthenticatedAccountResponse{ID: "", Code: 409, Err: errors.New("Not authorized, this account is disable")}
	}
	check_device := BlackListDevice(reqToken, accountID)
	if check_device.Code == 401 {
		return AuthenticatedAccountResponse{ID: "", Code: 401, Err: errors.New("Not authorized, this device is not connected")}
	}

	utils.DeviceConnectMiddleware(w, r, accountID)

	return AuthenticatedAccountResponse{accountID, 200, nil}
}

func GetAuthenticatedAccount(authToken string) (string, string) {
	parts := strings.Split(authToken, ".")
	if len(parts) != 3 {
		return "", ""
	}

	decodedBytes, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		CheckError(err)
		return "", ""
	}

	var jsonMap map[string]interface{}
	if err := json.Unmarshal(decodedBytes, &jsonMap); err != nil {
		CheckError(err)
		return "", ""
	}

	// Vérifie si l'utilisateur est un patient
	if patientID, ok := jsonMap["id"].(string); ok {
		if _, ok := jsonMap["patient"].(string); ok {
			return patientID, "patient"
		}
	}

	// Vérifie si l'utilisateur est un docteur
	if doctorID, ok := jsonMap["id"].(string); ok {
		if _, ok := jsonMap["doctor"].(string); ok {
			return doctorID, "doctor"
		}
	}

	return "", ""
}

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func BlackListDevice(token string, ownerID string) BlackListDeviceResponse {
	// Split the token into parts
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return BlackListDeviceResponse{Code: 401, Err: errors.New("Invalid token format")}
	}

	// Decode the token part that contains the payload
	decodedBytes, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return BlackListDeviceResponse{Code: 401, Err: errors.New("Error decoding token")}
	}

	// Unmarshal the payload into a map
	var jsonMap map[string]interface{}
	if err := json.Unmarshal(decodedBytes, &jsonMap); err != nil {
		return BlackListDeviceResponse{Code: 401, Err: errors.New("Error unmarshalling token")}
	}

	// Extract the device ID from the token
	deviceID, ok := jsonMap["name_device"].(string)
	if !ok {
		return BlackListDeviceResponse{Code: 401, Err: errors.New("Device name not found in token")}
	}

	// Get the list of connected devices for the owner (patient or doctor)
	allDeviceAccount := double_auth.GetDeviceConnect(ownerID, 0, 0)
	if allDeviceAccount.Err != nil {
		return BlackListDeviceResponse{Code: 401, Err: allDeviceAccount.Err}
	}

	for _, deviceConnected := range allDeviceAccount.DevicesConnect {
		if deviceConnected.ID == deviceID {
			return BlackListDeviceResponse{Code: 200, Err: nil}
		}
	}

	return BlackListDeviceResponse{Code: 401, Err: errors.New("Device not found in the list of connected devices")}

}
