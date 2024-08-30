package redis

import (
	"github.com/joho/godotenv"
	"log"
	"testing"
)

func TestAddTokenToList(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}

	token := "test-token"

	resp, err := AddTokenToList(token)
	if err != nil {
		t.Errorf("Expected no error but got: %s", err.Error())
	}
	if resp != "1\n" {
		t.Errorf("Expected response '1' but got %s", resp)
	}
}

func TestRemoveTokenFromList(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}

	token := "test-token"

	_, err := AddTokenToList(token)
	if err != nil {
		t.Errorf("Error adding token: %s", err.Error())
	}

	resp, err := RemoveTokenFromList(token)
	if err != nil {
		t.Errorf("Expected no error but got: %s", err.Error())
	}
	if resp != "2\n" {
		t.Errorf("Expected response '1' but got %s", resp)
	}
}

func TestSetKey(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}

	expire := 5
	resp, err := SetKey("test_key_set", "test_value_set", &expire)
	if err != nil {
		t.Errorf("Expected no error but got: %s", err.Error())
	}
	if resp != "OK\n" {
		t.Errorf("Expected response 'OK' but got %s", resp)
	}
}

func TestGetKey(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}

	expire := 5
	_, err := SetKey("test_key_get", "test_value_get", &expire)
	if err != nil {
		t.Errorf("Expected no error but got: %s", err.Error())
	}

	resp, err := GetKey("test_key_get")
	if err != nil {
		t.Errorf("Expected no error but got: %s", err.Error())
	}
	if resp != "test_value_get" {
		t.Errorf("Expected response 'OK' but got %s", resp)
	}
}

func TestDeleteKey(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}

	expire := 5
	_, err := SetKey("test_delete_key", "test_delete_value", &expire)
	if err != nil {
		t.Errorf("Expected no error but got: %s", err.Error())
	}

	resp, err := DeleteKey("test_key_delete")
	if err != nil {
		t.Errorf("Expected no error but got: %s", err.Error())
	}
	if resp != "0\n" {
		t.Errorf("Expected response 'OK' but got %s", resp)
	}
}
