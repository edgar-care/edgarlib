package treatment

import (
	"github.com/edgar-care/edgarlib/v2/medical_folder"
	"testing"

	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
	"github.com/google/uuid"
)

func TestUpdateTreatmentWithValidInput(t *testing.T) {
	patient, err := graphql.CreatePatient(model.CreatePatientInput{
		Email:    "test_update_treatment_valid@edgar-sante.fr",
		Password: "password",
	})
	if err != nil {
		t.Errorf("Failed to create patient: %v", err)
	}

	_ = medical_folder.CreateMedicalInfo(medical_folder.CreateMedicalInfoInput{
		Name:            "valjean",
		Firstname:       "jean",
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

	treatment, err := graphql.CreateTreatment(model.CreateTreatmentInput{
		MedicineID: uuid.New().String(),
		Quantity:   1,
		Period:     []model.Period{"MORNING"},
		Day:        []model.Day{"FRIDAY"},
		StartDate:  1234,
		EndDate:    2344,
	})
	if err != nil {
		t.Errorf("Failed to create treatment: %v", err)
	}

	input := UpdateTreatmentInput{
		Treatments: []TreatmentsInput{
			{
				ID:         treatment.ID,
				MedicineId: uuid.New().String(),
				Period:     []string{"EVENING"},
				Day:        []string{"TUESDAY"},
				Quantity:   2,
				StartDate:  123,
				EndDate:    2344,
			},
		},
	}

	response := UpdateTreatment(input, patient.ID)

	if response.Code != 200 {
		t.Errorf("Expected response code 200, got %v", response.Code)
	}
	if response.Err != nil {
		t.Errorf("Unexpected error: %v", response.Err)
	}

	updatedTreatment, err := graphql.GetTreatmentByID(treatment.ID)
	if err != nil {
		t.Errorf("Failed to retrieve updated treatment: %v", err)
	}

	if updatedTreatment.Quantity != 2 || updatedTreatment.Period[0] != "EVENING" || updatedTreatment.Day[0] != "TUESDAY" {
		t.Errorf("Treatment was not updated correctly")
	}
}

func TestUpdateTreatmentWithInvalidPatientID(t *testing.T) {
	invalidPatientID := uuid.New().String()

	input := UpdateTreatmentInput{
		Treatments: []TreatmentsInput{
			{
				MedicineId: uuid.New().String(),
				Period:     []string{"MORNING"},
				Day:        []string{"MONDAY"},
				Quantity:   1,
				StartDate:  123,
				EndDate:    2344,
			},
		},
	}

	response := UpdateTreatment(input, invalidPatientID)

	if response.Code != 400 {
		t.Errorf("Expected response code 400, got %v", response.Code)
	}
	if response.Err == nil {
		t.Errorf("Expected an error in response")
	}
}

func TestUpdateTreatmentWithNonExistentTreatmentID(t *testing.T) {
	patientID := uuid.New().String()
	nonExistentTreatmentID := uuid.New().String()

	_, err := graphql.CreatePatient(model.CreatePatientInput{
		Email:    "test_update_treatment_with_non_existent@edgar-sante.fr",
		Password: "password",
	})
	if err != nil {
		t.Errorf("Failed to create patient: %v", err)
	}

	input := UpdateTreatmentInput{
		Treatments: []TreatmentsInput{
			{
				ID:         nonExistentTreatmentID,
				MedicineId: uuid.New().String(),
				Period:     []string{"MORNING"},
				Day:        []string{"MONDAY"},
				Quantity:   1,
				StartDate:  123,
				EndDate:    2344,
			},
		},
	}

	response := UpdateTreatment(input, patientID)

	if response.Code != 400 {
		t.Errorf("Expected response code 400, got %v", response.Code)
	}
	if response.Err == nil {
		t.Errorf("Expected an error in response")
	}
}

func TestUpdateTreatmentWithNoMedicalFolder(t *testing.T) {
	patientID := uuid.New().String()
	treatmentID := uuid.New().String()

	_, err := graphql.CreatePatient(model.CreatePatientInput{
		Email:    "test_update_treatment_with_no_medical_folder@edgar-sante.fr",
		Password: "password",
	})
	if err != nil {
		t.Errorf("Failed to create patient: %v", err)
	}

	input := UpdateTreatmentInput{
		Treatments: []TreatmentsInput{
			{
				ID:         treatmentID,
				MedicineId: uuid.New().String(),
				Period:     []string{"MORNING"},
				Day:        []string{"MONDAY"},
				Quantity:   1,
				StartDate:  123,
				EndDate:    234,
			},
		},
	}

	response := UpdateTreatment(input, patientID)

	if response.Code != 400 {
		t.Errorf("Expected response code 400, got %v", response.Code)
	}
	if response.Err == nil {
		t.Errorf("Expected an error in response")
	}
}
