package medical_folder

import (
	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
	"testing"
)

func TestGetMedicalAntecedentByIdSuccessfully(t *testing.T) {
	patient, err := graphql.CreatePatient(model.CreatePatientInput{
		Email:    "test_patient_get_medical_antecedentbyId@edgar-sante.fr",
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
		Name:     "new_antecedent_get_byid",
		Symptoms: []string{"symptoms"},
		Treatments: []CreateTreatInput{{
			StartDate: 1234,
			EndDate:   5678,
			Medicines: []CreateAntecedentsMedicines{{
				Comment:    "comment",
				MedicineID: "test",
				Period: []*CreateAntecedentPeriod{{
					Quantity:       2,
					Frequency:      2,
					FrequencyRatio: 2,
					FrequencyUnit:  "JOUR",
					PeriodLength:   2,
					PeriodUnit:     "JOUR",
				}},
			}},
		}},
	}

	antecedent := AddMedicalAntecedent(antecedentInput, patient.ID)
	if antecedent.Err != nil {
		t.Errorf("Error while adding medical antecedent: %v", err)
	}

	antecedentID := antecedent.MedicalAntecedents[0].ID

	response := GetMedicalAntecedentById(antecedentID, patient.ID)

	if response.Code != 200 {
		t.Errorf("Expected code 200 but got %d", response.Code)
	}
	if response.Err != nil {
		t.Errorf("Expected no error but got: %s", response.Err.Error())
	}
	if response.MedicalAntecedent.ID != antecedentID {
		t.Errorf("Expected antecedent ID %s but got %s", antecedentID, response.MedicalAntecedent.ID)
	}
}

func TestGetMedicalAntecedentByIdWithInvalidPatientID(t *testing.T) {
	patientID := "invalid_patient_id"
	antecedentID := "valid_antecedent_id"

	response := GetMedicalAntecedentById(antecedentID, patientID)

	if response.Code == 200 {
		t.Errorf("Expected non-200 code but got %d", response.Code)
	}
	if response.Err == nil {
		t.Errorf("Expected error but got none")
	}
}

func TestGetMedicalAntecedentByIdWithInvalidAntecedentID(t *testing.T) {
	patientID := "valid_patient_id"
	antecedentID := "invalid_antecedent_id"

	response := GetMedicalAntecedentById(antecedentID, patientID)

	if response.Code == 200 {
		t.Errorf("Expected non-200 code but got %d", response.Code)
	}
	if response.Err == nil {
		t.Errorf("Expected error but got none")
	}
}

func TestGetMedicalAntecedentsSuccessfully(t *testing.T) {
	patient, err := graphql.CreatePatient(model.CreatePatientInput{
		Email:    "test_patient_get_medical_antecedents@edgar-sante.fr",
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
		Name:     "new_antecedent_get",
		Symptoms: []string{"symptoms"},
		Treatments: []CreateTreatInput{{
			StartDate: 1234,
			EndDate:   5678,
			Medicines: []CreateAntecedentsMedicines{{
				Comment:    "comment",
				MedicineID: "test",
				Period: []*CreateAntecedentPeriod{{
					Quantity:       2,
					Frequency:      2,
					FrequencyRatio: 2,
					FrequencyUnit:  "JOUR",
					PeriodLength:   2,
					PeriodUnit:     "JOUR",
				}},
			}},
		}},
	}

	antecedent := AddMedicalAntecedent(antecedentInput, patient.ID)
	if antecedent.Err != nil {
		t.Errorf("Error while adding medical antecedent: %v", err)
	}

	response := GetMedicalAntecedents(patient.ID)

	if response.Code != 200 {
		t.Errorf("Expected code 200 but got %d", response.Code)
	}
	if response.Err != nil {
		t.Errorf("Expected no error but got: %s", response.Err.Error())
	}
	if len(response.MedicalAntecedents) == 0 {
		t.Errorf("Expected at least one medical antecedent but got none")
	}
}

func TestGetMedicalAntecedentsWithInvalidPatientID(t *testing.T) {
	patientID := "invalid_patient_id"

	response := GetMedicalAntecedents(patientID)

	if response.Code == 200 {
		t.Errorf("Expected non-200 code but got %d", response.Code)
	}
	if response.Err == nil {
		t.Errorf("Expected error but got none")
	}
}

func TestGetMedicalAntecedentsWithNoMedicalFolder(t *testing.T) {
	patientID := "valid_patient_id_no_medical_folder"

	response := GetMedicalAntecedents(patientID)

	if response.Code == 200 {
		t.Errorf("Expected non-200 code but got %d", response.Code)
	}
	if response.Err == nil {
		t.Errorf("Expected error but got none")
	}
}
