package chat

import (
	"context"
	"github.com/edgar-care/edgarlib/graphql"
	"log"
	"testing"

	"github.com/joho/godotenv"
)

func TestGetChat(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}
	gqlClient := graphql.CreateClient()

	patient, err := graphql.CreatePatient(context.Background(), gqlClient, "test_getchat@edgar-sante.fr", "password")
	if err != nil {
		t.Errorf("Error creating patient: %v", err)
	}

	doctor, err := graphql.CreateDoctor(context.Background(), gqlClient, "test_doctor_getchat@edgar-sante.fr", "password", "name", "first", graphql.AddressInput{"", "", "", ""})
	if err != nil {
		t.Errorf("Error while creating doctor: %v", err)
	}

	input := ContentInput{
		Message:      "test add",
		RecipientIds: []string{patient.CreatePatient.Id, doctor.CreateDoctor.Id},
	}

	response := CreateChat(patient.CreatePatient.Id, input)

	if response.Code != 201 {
		t.Errorf("Expected code 200 but got %d", response.Code)
	}
	if response.Err != nil {
		t.Errorf("Expected no error but got: %s", response.Err.Error())
	}

	getchat := GetChat(patient.CreatePatient.Id)

	if getchat.Err != nil {
		t.Errorf("Unexpected error: %v", getchat.Err)
	}

	if getchat.Code != 200 {
		t.Errorf("Expected code 200, got %d", getchat.Code)
	}

}

func TestGetChatInvalid(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}

	response := GetChat("")

	if response.Err == nil {
		t.Error("Unexpected null error")
	}

	if response.Code != 400 {
		t.Errorf("Expected code 400, got %d", response.Code)
	}
}
