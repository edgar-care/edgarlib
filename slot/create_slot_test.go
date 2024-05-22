package slot

import (
	"context"
	"log"
	"testing"

	"github.com/edgar-care/edgarlib/graphql"
	"github.com/joho/godotenv"
)

func TestCreateSlot(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}
	gqlClient := graphql.CreateClient()

	patient, err := graphql.CreateDoctor(context.Background(), gqlClient, "test_doctor_create_slot_up@edgar-sante.fr", "password", "test_doctor", "first", graphql.AddressInput{"", "", "", ""})
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
	if response.Err != nil {
		t.Errorf("Expected no error but got: %s", response.Err.Error())
	}

	_, err = graphql.GetSlotById(context.Background(), gqlClient, response.Rdv.ID)

	if err != nil {
		t.Errorf("Error while retrieving created appointment: %s", err.Error())
	}

}

func TestCreateSlotInvalidId(t *testing.T) {

	input := CreateSlotInput{
		StartDate: 14567,
		EndDate:   14578,
	}

	response := CreateSlot(input, "")
	if response.Code != 400 || response.Err == nil {
		t.Errorf("Expected error and code 400 but got code %d and err: %s", response.Code, response.Err.Error())
	}
}
