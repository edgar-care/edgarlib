package black_list

import (
	"github.com/joho/godotenv"
	"log"
	"testing"
)

func TestGetBlackListById(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}

	createBlackList := CreateBlackList("test_token_create_blacklist")
	if createBlackList.Code != 201 {
		t.Errorf("Expected code 201 but got %d", createBlackList.Code)
	}
	if createBlackList.Err != nil {
		t.Errorf("Expected no error but got: %s", createBlackList.Err.Error())
	}
	response := GetBlackListById(createBlackList.BlackList.ID)
	if response.Code != 200 {
		t.Errorf("Expected code 200 but got %d", response.Code)
	}
	if response.Err != nil {
		t.Errorf("Expected no error but got: %s", response.Err.Error())
	}
}

func TestGetBlackListInvalidId(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}

	response := GetBlackListById("test_invalid_id_blacklist")
	if response.Code != 400 || response.Err == nil {
		t.Errorf("Expected error and code 400 but got code %d and err: %s", response.Code, response.Err.Error())
	}
}

func TestGetBlackList(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}

	createBlackList := CreateBlackList("test_token_create_blacklist_all")
	if createBlackList.Code != 201 {
		t.Errorf("Expected code 201 but got %d", createBlackList.Code)
	}
	if createBlackList.Err != nil {
		t.Errorf("Expected no error but got: %s", createBlackList.Err.Error())
	}
	response := GetBlackList()
	if response.Code != 200 {
		t.Errorf("Expected code 200 but got %d", response.Code)
	}
	if response.Err != nil {
		t.Errorf("Expected no error but got: %s", response.Err.Error())
	}
}
