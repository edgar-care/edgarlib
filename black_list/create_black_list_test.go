package black_list

import (
	"github.com/joho/godotenv"
	"log"
	"testing"
)

func TestCreateBlackList(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}

	response := CreateBlackList("test_token_create_blacklist")
	if response.Code != 201 {
		t.Errorf("Expected code 200 but got %d", response.Code)
	}
	if response.Err != nil {
		t.Errorf("Expected no error but got: %s", response.Err.Error())
	}
}
