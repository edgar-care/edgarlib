package medicament

import (
	"log"
	"testing"

	"github.com/joho/godotenv"
)

func TestGetMedicaments(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}

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

	getmedicament := GetMedicaments()

	if getmedicament.Err != nil {
		t.Errorf("Unexpected error: %v", getmedicament.Err)
	}

	if getmedicament.Code != 200 {
		t.Errorf("Expected code 200, got %d", getmedicament.Code)
	}

}

func TestGetMedicamentsInvalid(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}

	response := GetMedicaments()

	if response.Err == nil {
		t.Error("Unexpected null error")
	}

	if response.Code != 400 {
		t.Errorf("Expected code 400, got %d", response.Code)
	}
}

func TestGetMedicamentById(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}

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

	getmedicamentbyid := GetMedicamentById(response.Medicament.ID)

	if getmedicamentbyid.Err != nil {
		t.Errorf("Unexpected error: %v", getmedicamentbyid.Err)
	}

	if getmedicamentbyid.Code != 200 {
		t.Errorf("Expected code 200, got %d", getmedicamentbyid.Code)
	}
}

func TestGetMedicamentByIdInvalid(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}

	response := GetMedicamentById("111111111111111111111111")

	if response.Err == nil {
		t.Error("Unexpected null error")
	}

	if response.Code != 400 {
		t.Errorf("Expected code 400, got %d", response.Code)
	}
}
