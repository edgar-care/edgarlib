package chat

import (
	"github.com/edgar-care/edgarlib/graphql"
	"github.com/edgar-care/edgarlib/graphql/model"
	"github.com/joho/godotenv"
	"log"
	"testing"
)

func TestGetChat(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}

	patient, err := graphql.CreatePatient(model.CreatePatientInput{
		Email:    "test_patient_chat_get_chat@edgar-sante.fr",
		Password: "password",
		Status:   true,
	})
	if err != nil {
		t.Errorf("Error while creating patient: %v", err)
	}

	doctor, err := graphql.CreateDoctor(model.CreateDoctorInput{
		Email:     "test_doctor_chat_get_chat@edgar-sante.fr",
		Password:  "password",
		Name:      "name",
		Firstname: "first",
		Address: &model.AddressInput{
			Street:  "",
			ZipCode: "",
			Country: "",
			City:    "",
		},
		Status: true,
	})
	if err != nil {
		t.Errorf("Error while creating patient: %v", err)
	}

	content := ContentInput{
		Message:      "test_create_chat",
		RecipientIds: []string{patient.ID, doctor.ID},
	}
	getChat := CreateChat(patient.ID, content)

	if getChat.Code != 201 {
		t.Errorf("Expected code 200 but got %d", getChat.Code)
	}
	if getChat.Err != nil {
		t.Errorf("Expected no error but got: %s", getChat.Err.Error())
	}

	response := GetChat(patient.ID)

	if response.Err != nil {
		t.Errorf("Unexpected error: %v", response.Err)
	}

	if response.Code != 200 {
		t.Errorf("Expected code 200, got %d", response.Code)
	}
}

func TestGetChatInvalidPatient(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}

	response := GetChat("test_chat_get_invalid_patient@edgar-sante.fr")

	if response.Code != 400 || response.Err == nil {
		t.Errorf("Expected error and code 400 but got code %d and err: %s", response.Code, response.Err.Error())
	}
}

func TestGetChatNotCreate(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}

	patient, err := graphql.CreatePatient(model.CreatePatientInput{
		Email:    "test_patient_chat_get_chat_not_create@edgar-sante.fr",
		Password: "password",
		Status:   true,
	})
	if err != nil {
		t.Errorf("Error while creating patient: %v", err)
	}

	response := GetChat(patient.ID)

	if response.Code != 400 && response.Err != nil {
		t.Errorf("Expected error and code 400 but got code %d and err: %s", response.Code, response.Err.Error())
	}
}

func TestUpdateMessageRead(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}

	patient, err := graphql.CreatePatient(model.CreatePatientInput{
		Email:    "test_patient_chat_update_chat@edgar-sante.fr",
		Password: "password",
		Status:   true,
	})
	if err != nil {
		t.Errorf("Error while creating patient: %v", err)
	}

	doctor, err := graphql.CreateDoctor(model.CreateDoctorInput{
		Email:     "test_doctor_chat_update_chat@edgar-sante.fr",
		Password:  "password",
		Name:      "name",
		Firstname: "first",
		Address: &model.AddressInput{
			Street:  "",
			ZipCode: "",
			Country: "",
			City:    "",
		},
		Status: true,
	})
	if err != nil {
		t.Errorf("Error while creating patient: %v", err)
	}

	content := ContentInput{
		Message:      "test_create_chat",
		RecipientIds: []string{patient.ID, doctor.ID},
	}
	getChat := CreateChat(patient.ID, content)

	if getChat.Code != 201 {
		t.Errorf("Expected code 200 but got %d", getChat.Code)
	}
	if getChat.Err != nil {
		t.Errorf("Expected no error but got: %s", getChat.Err.Error())
	}

	response := UpdateMessageRead(patient.ID, getChat.Chat.ID)

	if response.Err != nil {
		t.Errorf("Unexpected error: %v", response.Err)
	}

	if response.Code != 200 {
		t.Errorf("Expected code 200, got %d", response.Code)
	}
}

func TestUpdateMessageReadInvalidId(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}

	patient, err := graphql.CreatePatient(model.CreatePatientInput{
		Email:    "test_patient_chat_update_chat_invalid_id@edgar-sante.fr",
		Password: "password",
		Status:   true,
	})
	if err != nil {
		t.Errorf("Error while creating patient: %v", err)
	}

	response := UpdateMessageRead(patient.ID, "test_invalid_id_chat_update")

	if response.Code != 400 || response.Err == nil {
		t.Errorf("Expected error and code 400 but got code %d and err: %s", response.Code, response.Err.Error())
	}
}
