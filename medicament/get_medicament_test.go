package medicament

import (
	"github.com/joho/godotenv"
	"log"
	"testing"
)

func TestGetMedicamentById(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}

	content := CreateMedicamentInput{
		Name:            "test_medicament_get",
		Unit:            "test_unit_get",
		TargetDiseases:  []string{"test_target_disease_get"},
		TreatedSymptoms: []string{"test_treated_symptoms_get"},
		SideEffects:     []string{"test_side_effects_get"},
		Type:            "test_type_get",
		Content:         "test_content_get",
		Quantity:        2,
	}

	getMedicamentByID := CreateMedicament(content)

	if getMedicamentByID.Code != 201 {
		t.Errorf("Expected code 200 but got %d", getMedicamentByID.Code)
	}
	if getMedicamentByID.Err != nil {
		t.Errorf("Expected no error but got: %s", getMedicamentByID.Err.Error())
	}

	response := GetMedicamentById(getMedicamentByID.Medicament.ID)
	if response.Code != 200 {
		t.Errorf("Expected code 200 but got %d", response.Code)
	}
	if response.Err != nil {
		t.Errorf("Expected no error but got: %s", response.Err.Error())
	}
}

func TestGetMedicamentByIdInvalid(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}

	response := GetMedicamentById("test_invalid_id_get_medicament")
	if response.Code != 400 || response.Err == nil {
		t.Errorf("Expected error and code 400 but got code %d and err: %s", response.Code, response.Err.Error())
	}
}

func TestGetMedicaments(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}

	content := CreateMedicamentInput{
		Name:            "test_medicament_get_all",
		Unit:            "test_unit_get_all",
		TargetDiseases:  []string{"test_target_disease_get_all"},
		TreatedSymptoms: []string{"test_treated_symptoms_get_all"},
		SideEffects:     []string{"test_side_effects_get_all"},
		Type:            "test_type_get_all",
		Content:         "test_content_get_all",
		Quantity:        4,
	}

	getMedicamentByID := CreateMedicament(content)

	if getMedicamentByID.Code != 201 {
		t.Errorf("Expected code 200 but got %d", getMedicamentByID.Code)
	}
	if getMedicamentByID.Err != nil {
		t.Errorf("Expected no error but got: %s", getMedicamentByID.Err.Error())
	}

	response := GetMedicaments()
	if response.Code != 200 {
		t.Errorf("Expected code 200 but got %d", response.Code)
	}
	if response.Err != nil {
		t.Errorf("Expected no error but got: %s", response.Err.Error())
	}
}
