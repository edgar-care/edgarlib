package auth

//func TestTokenCheckValidToken(t *testing.T) {
//	if err := godotenv.Load(".env.test"); err != nil {
//		log.Fatalf("Error loading .env.test file: %v", err)
//	}
//
//	token := "valid-token"
//	_, err := redis.SetKey("user-tokens", token, nil)
//	if err != nil {
//		t.Fatalf("Failed to set token in Redis: %v", err)
//	}
//
//	statusCode, message := TokenCheck(token)
//
//	if statusCode != http.StatusUnauthorized {
//		t.Errorf("Expected status code %d but got %d", http.StatusUnauthorized, statusCode)
//	}
//
//	if message != "Token is invalid" {
//		t.Errorf("Expected message 'Token is invalid' but got '%s'", message)
//	}
//}
//
//func TestTokenCheckInvalidToken(t *testing.T) {
//	if err := godotenv.Load(".env.test"); err != nil {
//		log.Fatalf("Error loading .env.test file: %v", err)
//	}
//
//	token := "invalid-token"
//
//	statusCode, message := TokenCheck(token)
//
//	if statusCode != http.StatusOK {
//		t.Errorf("Expected status code %d but got %d", http.StatusOK, statusCode)
//	}
//
//	if message != "Token is valid" {
//		t.Errorf("Expected message 'Token is valid' but got '%s'", message)
//	}
//}

//func TestTokenCheckError(t *testing.T) {
//	if err := godotenv.Load(".env.test"); err != nil {
//		log.Fatalf("Error loading .env.test file: %v", err)
//	}
//	originalCheckTokenPresence := CheckTokenPresence
//	CheckTokenPresence = func(token string) (bool, error) {
//		return false, fmt.Errorf("mock error")
//	}
//	defer func() { CheckTokenPresence = originalCheckTokenPresence }()
//
//	statusCode, message := TokenCheck("error-token")
//
//	if statusCode != http.StatusInternalServerError {
//		t.Errorf("Expected status code %d but got %d", http.StatusInternalServerError, statusCode)
//	}
//
//	if message != "Error checking token: mock error" {
//		t.Errorf("Expected message 'Error checking token: mock error' but got '%s'", message)
//	}
//}
