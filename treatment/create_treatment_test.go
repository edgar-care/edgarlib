package treatment

import (
	"github.com/edgar-care/edgarlib/medical_folder"
	"github.com/joho/godotenv"
	"log"
	"testing"

	"github.com/edgar-care/edgarlib/graphql"
	"github.com/edgar-care/edgarlib/graphql/model"
	"github.com/google/uuid"
)

func TestCreateTreatment(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}
	email := "testuser_create_treatment@example.com"

	patient, err := graphql.CreatePatient(model.CreatePatientInput{
		Email:    email,
		Password: "testpassword",
		Status:   true,
	})
	if err != nil {
		t.Errorf("Failed to create patient: %v", err)
	}

	_ = medical_folder.CreateMedicalInfo(medical_folder.CreateMedicalInfoInput{
		Name:            "bertand",
		Firstname:       "pierre",
		Birthdate:       0,
		Sex:             "",
		Weight:          0,
		Height:          0,
		PrimaryDoctorID: "",
		MedicalAntecedents: []medical_folder.CreateMedicalAntecedentInput{{
			Name:          "",
			Medicines:     nil,
			StillRelevant: false,
		}},
		FamilyMembersMedInfoId: []string{},
	}, patient.ID)

	t.Run("Success without existing DiseaseId", func(t *testing.T) {
		treatmentInput := CreateNewTreatmentInput{
			Name:          "Hypertension",
			DiseaseId:     "",
			StillRelevant: true,
			Treatments: []CreateTreatInput{
				{
					MedicineId: "med123",
					Period:     []string{"MORNING"},
					Day:        []string{"MONDAY"},
					Quantity:   1,
				},
			},
		}

		response := CreateTreatment(treatmentInput, patient.ID)

		if response.Code != 201 {
			t.Errorf("Expected response code 201, got %v", response.Code)
		}
		if response.Err != nil {
			t.Errorf("Unexpected error: %v", response.Err)
		}

		if response.AnteDisease.Name != treatmentInput.Name {
			t.Errorf("Expected AnteDisease name %s, got %s", treatmentInput.Name, response.AnteDisease.Name)
		}
	})

	t.Run("Failure due to missing patient", func(t *testing.T) {
		if err := godotenv.Load(".env.test"); err != nil {
			log.Fatalf("Error loading .env.test file: %v", err)
		}
		nonExistentPatientID := uuid.New().String()

		treatmentInput := CreateNewTreatmentInput{
			Name:          "Diabetes",
			DiseaseId:     "",
			StillRelevant: true,
			Treatments: []CreateTreatInput{
				{
					MedicineId: "med456",
					Period:     []string{"EVENING"},
					Day:        []string{"SUNDAY"},
					Quantity:   1,
				},
			},
		}

		response := CreateTreatment(treatmentInput, nonExistentPatientID)

		if response.Code != 400 {
			t.Errorf("Expected response code 400, got %v", response.Code)
		}
		if response.Err == nil {
			t.Errorf("Expected an error in response")
		}
	})

	t.Run("Failure due to missing medical folder", func(t *testing.T) {
		noMedicalInfoID := uuid.New().String()

		_, err := graphql.CreatePatient(model.CreatePatientInput{
			Email:    "testuser_no_medical_info@example.com",
			Password: "testpassword",
			Status:   true,
		})
		if err != nil {
			t.Errorf("Failed to create patient: %v", err)
		}

		treatmentInput := CreateNewTreatmentInput{
			Name:          "Asthma",
			DiseaseId:     "",
			StillRelevant: true,
			Treatments: []CreateTreatInput{
				{
					MedicineId: uuid.New().String(),
					Period:     []string{"MORNING"},
					Day:        []string{"TUESDAY"},
					Quantity:   3,
				},
			},
		}

		response := CreateTreatment(treatmentInput, noMedicalInfoID)

		if response.Code != 400 {
			t.Errorf("Expected response code 400, got %v", response.Code)
		}
		if response.Err == nil {
			t.Errorf("Expected an error in response")
		}
	})
}
