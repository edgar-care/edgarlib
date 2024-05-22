package medicament

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

	input := CreateMedicamentInput{
		Name:            "test32",
		Unit:            "APPLICATION",
		TargetDiseases:  []string{"test"},
		TreatedSymptoms: []string{"test"},
		SideEffects:     []string{"test"},
	}

	response := CreateMedicament(input)

	if response.Code != 201 {
		t.Errorf("Expected code 200 but got %d", response.Code)
	}
	if response.Err != nil {
		t.Errorf("Expected no error but got: %s", response.Err.Error())
	}

	_, err := graphql.GetMedicineByID(context.Background(), gqlClient, response.Medicament.ID)

	if err != nil {
		t.Errorf("Error while retrieving created appointment: %s", err.Error())
	}

}

func TestCreateSlotInvalidId(t *testing.T) {

	input := CreateMedicamentInput{
		Name:            "test32",
		Unit:            "APPLICATION",
		TargetDiseases:  []string{"test"},
		TreatedSymptoms: []string{"test"},
		SideEffects:     []string{"test"},
	}

	response := CreateMedicament(input)
	if response.Code != 400 || response.Err == nil {
		t.Errorf("Expected error and code 400 but got code %d and err: %s", response.Code, response.Err.Error())
	}
}
