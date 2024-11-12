package medical_folder

import (
	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
	"github.com/joho/godotenv"
	"log"
	"testing"
)

func TestDeleteMedicalAntecedentSuccessfully(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}

	// Create a patient
	patient, err := graphql.CreatePatient(model.CreatePatientInput{
		Email:    "test_patient_delete_medical_antecedent@edgar-sante.fr",
		Password: "password",
		Status:   false,
	})
	if err != nil {
		t.Errorf("Error while creating patient: %v", err)
	}

	// Create a medical folder for the patient
	medicalFolderInput := CreateNewMedicalInfoInput{
		Name:                   "test",
		Firstname:              "test",
		Birthdate:              0,
		Sex:                    "M",
		Weight:                 0,
		Height:                 0,
		PrimaryDoctorID:        "",
		MedicalAntecedents:     []CreateNewMedicalAntecedentInput{},
		FamilyMembersMedInfoId: []string{""},
	}

	_ = NewMedicalFolder(medicalFolderInput, patient.ID)

	antecedentInput := CreateNewMedicalAntecedentInput{
		Name:     "new_antecedent",
		Symptoms: []string{"symptoms"},
		Treatments: []CreateTreatInput{{
			CreatedBy: "test",
			StartDate: 1234,
			EndDate:   5678,
			Medicines: []CreateAntecedentsMedicines{{
				Period: []*CreateAntecedentPeriod{{
					Quantity:       2,
					Frequency:      2,
					FrequencyRatio: 2,
					FrequencyUnit:  "JOUR",
					PeriodLength:   2,
					PeriodUnit:     "JOUR",
					Comment:        "comment",
				}},
			}},
		}},
	}

	ante := AddMedicalAntecedent(antecedentInput, patient.ID)
	//spew.Dump(getMedicalFolder.MedicalInfo.AntecedentDiseaseIds)

	response := DeleteMedicalAntecedent(ante.MedicalAntecedents[0].ID, patient.ID)

	_ = GetMedicalFolder(patient.ID)
	//spew.Dump(getMedicalFolder.MedicalInfo)
	if response.Code != 200 {
		t.Errorf("Expected code 200 but got %d", response.Code)
	}
	if response.Err != nil {
		t.Errorf("Expected no error but got: %s", response.Err.Error())
	}
	if !response.Deleted {
		t.Errorf("Expected antecedent to be deleted but it was not")
	}
}

func TestDeleteMedicalAntecedentWithInvalidPatientID(t *testing.T) {
	patientID := "invalid_patient_id"
	antecedentID := "valid_antecedent_id"

	response := DeleteMedicalAntecedent(antecedentID, patientID)

	if response.Code == 200 {
		t.Errorf("Expected non-200 code but got %d", response.Code)
	}
	if response.Err == nil {
		t.Errorf("Expected error but got none")
	}
}

func TestDeleteMedicalAntecedentWithInvalidAntecedentID(t *testing.T) {
	patientID := "valid_patient_id"
	antecedentID := "invalid_antecedent_id"

	response := DeleteMedicalAntecedent(antecedentID, patientID)

	if response.Code == 200 {
		t.Errorf("Expected non-200 code but got %d", response.Code)
	}
	if response.Err == nil {
		t.Errorf("Expected error but got none")
	}
}

func TestDeleteMedicalAntecedentWithNoMedicalFolder(t *testing.T) {
	patientID := "valid_patient_id_no_medical_folder"
	antecedentID := "valid_antecedent_id"

	response := DeleteMedicalAntecedent(antecedentID, patientID)

	if response.Code == 200 {
		t.Errorf("Expected non-200 code but got %d", response.Code)
	}
	if response.Err == nil {
		t.Errorf("Expected error but got none")
	}
}

func TestDeleteMedicalAntecedentNotFoundInMedicalFolder(t *testing.T) {
	patientID := "valid_patient_id"
	antecedentID := "non_existent_antecedent_id"

	response := DeleteMedicalAntecedent(antecedentID, patientID)

	if response.Code == 200 {
		t.Errorf("Expected non-200 code but got %d", response.Code)
	}
	if response.Err == nil {
		t.Errorf("Expected error but got none")
	}
}
