package slot

import (
	"context"
	"log"
	"testing"

	"github.com/joho/godotenv"

	"github.com/edgar-care/edgarlib/graphql"
)

func TestGetSlots(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}
	gqlClient := graphql.CreateClient()

	patient, err := graphql.CreateDoctor(context.Background(), gqlClient, "test_doctor_getslors_slot_up@edgar-sante.fr", "password", "test_doctor", "first", graphql.AddressInput{"", "", "", ""})
	if err != nil {
		t.Errorf("Error while creating patient: %v", err)
	}

	input := CreateSlotInput{
		StartDate: 14567,
		EndDate:   14578,
	}

	response := CreateSlot(input, patient.CreateDoctor.Id)

	if response.Code != 201 {
		t.Errorf("Expected code 200 but got %d", response.Code)
	}

	getslots := GetSlots(patient.CreateDoctor.Id)

	if getslots.Err != nil {
		t.Errorf("Unexpected error: %v", getslots.Err)
	}

	if getslots.Code != 200 {
		t.Errorf("Expected code 200, got %d", getslots.Code)
	}

}

func TestGetSlotsInvalid(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}

	response := GetSlots("111111111111111111111111")

	if response.Err == nil {
		t.Error("Unexpected null error")
	}

	if response.Code != 400 {
		t.Errorf("Expected code 400, got %d", response.Code)
	}
}

func TestGetSlotById(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}
	gqlClient := graphql.CreateClient()

	patient, err := graphql.CreateDoctor(context.Background(), gqlClient, "test_doctor_getslotsid_slot_up@edgar-sante.fr", "password", "test_doctor", "first", graphql.AddressInput{"", "", "", ""})
	if err != nil {
		t.Errorf("Error while creating patient: %v", err)
	}

	input := CreateSlotInput{
		StartDate: 14567,
		EndDate:   14578,
	}

	response := CreateSlot(input, patient.CreateDoctor.Id)

	responseget := GetSlotById(response.Rdv.ID, patient.CreateDoctor.Id)

	if responseget.Err != nil {
		t.Errorf("Unexpected error: %v", responseget.Err)
	}

	if responseget.Code != 200 {
		t.Errorf("Expected code 200, got %d", responseget.Code)
	}
}

func TestGetSlotByIdInvalid(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}

	response := GetSlotById("111111111111111111111111", "")

	if response.Err == nil {
		t.Error("Unexpected null error")
	}

	if response.Code != 400 {
		t.Errorf("Expected code 400, got %d", response.Code)
	}
}
