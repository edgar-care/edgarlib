package medicament

import (
	"github.com/joho/godotenv"
	"log"
	"testing"
)

func TestCreateMedicament(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}

	content := CreateMedicamentInput{
		Name:            "test_medicament_create",
		Unit:            "test_unit",
		TargetDiseases:  []string{"test_target_disease"},
		TreatedSymptoms: []string{"test_treated_symptoms"},
		SideEffects:     []string{"test_side_effects"},
	}

	response := CreateMedicament(content)

	if response.Code != 201 {
		t.Errorf("Expected code 200 but got %d", response.Code)
	}
	if response.Err != nil {
		t.Errorf("Expected no error but got: %s", response.Err.Error())
	}
}
