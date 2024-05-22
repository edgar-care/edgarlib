package chat

import (
	"context"
	"log"
	"testing"

	"github.com/edgar-care/edgarlib/graphql"
	"github.com/joho/godotenv"
)

func TestCreateChat(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}
	gqlClient := graphql.CreateClient()

	patient, err := graphql.CreatePatient(context.Background(), gqlClient, "test_create_chat@edgar-sante.fr", "password")
	if err != nil {
		t.Errorf("Error creating patient: %v", err)
	}

	doctor, err := graphql.CreateDoctor(context.Background(), gqlClient, "test_doctor_createchat@edgar-sante.fr", "password", "name", "first", graphql.AddressInput{"", "", "", ""})
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

	_, err = graphql.GetChatById(context.Background(), gqlClient, response.Chat.ID)

	if err != nil {
		t.Errorf("Error while retrieving created appointment: %s", err.Error())
	}

}

func TestCreateChatInvalidId(t *testing.T) {

	input := ContentInput{
		Message:      "test add",
		RecipientIds: []string{"hello", "world"},
	}

	response := CreateChat("11111111", input)
	if response.Code != 400 || response.Err == nil {
		t.Errorf("Expected error and code 400 but got code %d and err: %s", response.Code, response.Err.Error())
	}
}
