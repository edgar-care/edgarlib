package auth

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/go-chi/jwtauth/v5"
	"golang.org/x/crypto/bcrypt"
)

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

func GetAuthenticatedUser(w http.ResponseWriter, r *http.Request) string {
	_, claims, _ := jwtauth.FromContext(r.Context())

	patientClaim, ok := claims["patient"].(map[string]interface{})
	if !ok || patientClaim == nil {
		return ""
	}

	id, ok := patientClaim["id"].(string)
	if !ok {
		return ""
	}

	return id
}

func AuthMiddleware(w http.ResponseWriter, r *http.Request) string {
	reqToken := r.Header.Get("Authorization")
	if reqToken == "" {
		return ""
	}
	splitToken := strings.Split(reqToken, "Bearer ")
	reqToken = splitToken[1]

	if VerifyToken(reqToken) == false {
		return ""
	}
	return GetAuthenticatedUser(w, r)
}

//func AuthMiddlewareGetAccountType(r *http.Request) string {
//	reqToken := r.Header.Get("Authorization")
//	if reqToken == "" {
//		return ""
//	}
//	splitToken := strings.Split(reqToken, "Bearer ")
//	reqToken = splitToken[1]
//
//	if VerifyToken(reqToken) == false {
//		return ""
//	}
//	return GetAccountType(r)
//}
//
//func GetAccountType(r *http.Request) string {
//	_, claims, _ := jwtauth.FromContext(r.Context())
//
//	_, valid := claims["patient"].(map[string]interface{})
//	if valid {
//		return "patient"
//	}
//	_, valid = claims["doctor"].(map[string]interface{})
//	if valid {
//		return "doctor"
//	}
//	return ""
//}

func GetBearerToken(req *http.Request) string {
	authHeader := req.Header.Get("Authorization")
	if authHeader == "" {
		return ""
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return ""
	}
	//checkTokenCode, _ := CheckAccountEnable(parts[1])
	//if checkTokenCode == http.StatusUnauthorized || checkTokenCode == http.StatusInternalServerError {
	//	return ""
	//}

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
