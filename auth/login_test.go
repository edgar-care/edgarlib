package auth

import (
	"testing"
)

func TestLogin(t *testing.T) {
	// Set up the GraphQL client

	// Test cases
	tests := []struct {
		name         string
		input        LoginInput
		tokenType    string
		expectedCode int
	}{
		{
			name: "Doctor login success",
			input: LoginInput{
				Email:    "doctor@example.com",
				Password: "password",
			},
			tokenType:    "d",
			expectedCode: 200,
		},
		{
			name: "Admin login success",
			input: LoginInput{
				Email:    "admin@example.com",
				Password: "password",
			},
			tokenType:    "a",
			expectedCode: 200,
		},
		// Add more test cases as needed
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			resp := Login(tc.input, tc.tokenType)

			if resp.Code != tc.expectedCode {
				t.Errorf("expected code: %d, got: %d", tc.expectedCode, resp.Code)
			}

			// Optionally, you can check for other conditions as well
			// such as the validity of the token, etc.
		})
	}
}
