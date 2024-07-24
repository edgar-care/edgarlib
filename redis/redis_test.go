package redis

import "testing"

func TestAddTokenToList(t *testing.T) {
	token := "test-token"

	resp, err := AddTokenToList(token)
	if err != nil {
		t.Errorf("Expected no error but got: %s", err.Error())
	}
	if resp != "1" {
		t.Errorf("Expected response '1' but got %s", resp)
	}
}

func TestRemoveTokenFromList(t *testing.T) {
	token := "test-token"

	_, err := AddTokenToList(token)
	if err != nil {
		t.Errorf("Error adding token: %s", err.Error())
	}

	resp, err := RemoveTokenFromList(token)
	if err != nil {
		t.Errorf("Expected no error but got: %s", err.Error())
	}
	if resp != "1" {
		t.Errorf("Expected response '1' but got %s", resp)
	}
}
