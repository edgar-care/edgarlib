package medical_folder

import (
	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
	"testing"
)

func TestUpdateMedicalAntecedentSuccessfully(t *testing.T) {
	patient, err := graphql.CreatePatient(model.CreatePatientInput{
		Email:    "test_patient_update_medical_antecedent@edgar-sante.fr",
		Password: "password",
		Status:   false,
	})
	if err != nil {
		t.Errorf("Error while creating patient: %v", err)
	}

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
			StartDate: 1234,
			EndDate:   5678,
			Medicines: []CreateAntecedentsMedicines{{
				Comment:    func(s string) *string { return &s }("comment"),
				MedicineID: "test",
				Period: []*CreateAntecedentPeriod{{
					Quantity:       2,
					Frequency:      2,
					FrequencyRatio: 2,
					FrequencyUnit:  "JOUR",
					PeriodLength:   func(i int) *int { return &i }(2),
					PeriodUnit:     func(s string) *string { return &s }("JOUR"),
				}},
			}},
		}},
	}
	antecedent := AddMedicalAntecedent(antecedentInput, patient.ID)

	updatedName := "updated_antecedent"

	input := UpdateMedicalFolderPatientInput{
		MedicalAntecedentInput: model.UpdateMedicalAntecedentsInput{
			Name:     &updatedName,
			Symptoms: []string{"updated_symptoms"},
		},
	}

	response := UpdateMedicalAntecedent(patient.ID, input, antecedent.MedicalAntecedents[0].ID)
	if response.Code != 200 {
		t.Errorf("Expected code 200 but got %d", response.Code)
	}
	if response.Err != nil {
		t.Errorf("Expected no error but got: %s", response.Err.Error())
	}
}

func TestUpdateMedicalFolderSuccessfully(t *testing.T) {
	patient, err := graphql.CreatePatient(model.CreatePatientInput{
		Email:    "test_patient_update_medical_folder@edgar-sante.fr",
		Password: "password",
		Status:   false,
	})
	if err != nil {
		t.Errorf("Error while creating patient: %v", err)
	}

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
			StartDate: 1234,
			EndDate:   5678,
			Medicines: []CreateAntecedentsMedicines{{
				Comment:    func(s string) *string { return &s }("comment"),
				MedicineID: "test",
				Period: []*CreateAntecedentPeriod{{
					Quantity:       2,
					Frequency:      2,
					FrequencyRatio: 2,
					FrequencyUnit:  "JOUR",
					PeriodLength:   func(s int) *int { return &s }(2),
					PeriodUnit:     func(s string) *string { return &s }("JOUR"),
				}},
			}},
		}},
	}
	_ = AddMedicalAntecedent(antecedentInput, patient.ID)

	updatedName := "updated_name"
	updatedFirstname := "updated_firstname"
	updatedBirthdate := 12
	updatedSex := "F"
	updatedHeight := 179
	updatedWeight := 98
	updatedDoctorID := "updated_doctor_id"
	updatedAntecedentID := []string{"updated_antecedent_id"}
	updatedOnboardingStatus := "DONE"
	updatedFamilyMemberID := []string{"updated_family_member_id"}

	input := model.UpdateMedicalFolderInput{
		Name:                   &updatedName,
		Firstname:              &updatedFirstname,
		Birthdate:              &updatedBirthdate,
		Sex:                    &updatedSex,
		Height:                 &updatedHeight,
		Weight:                 &updatedWeight,
		PrimaryDoctorID:        &updatedDoctorID,
		AntecedentDiseaseIds:   updatedAntecedentID,
		OnboardingStatus:       (*model.OnboardingStatus)(&updatedOnboardingStatus),
		FamilyMembersMedInfoID: updatedFamilyMemberID,
	}

	response := UpdateMedicalFolderPatient(patient.ID, input)
	if response.Code != 200 {
		t.Errorf("Expected code 200 but got %d", response.Code)
	}
	if response.Err != nil {
		t.Errorf("Expected no error but got: %s", response.Err.Error())
	}
}

func TestUpdateMedicalAntecedentWithInvalidPatientID(t *testing.T) {
	patientID := "invalid_patient_id"
	updatedName := "updated_antecedent"
	input := UpdateMedicalFolderPatientInput{
		MedicalAntecedentInput: model.UpdateMedicalAntecedentsInput{
			Name:     &updatedName,
			Symptoms: []string{"updated_symptoms"},
		},
	}

	response := UpdateMedicalAntecedent(patientID, input, "valid_antecedent_id")

	if response.Code == 200 {
		t.Errorf("Expected non-200 code but got %d", response.Code)
	}
	if response.Err == nil {
		t.Errorf("Expected error but got none")
	}
}

func TestUpdateMedicalAntecedentWithInvalidAntecedentID(t *testing.T) {
	patientID := "invalid_patient_id"
	updatedName := "updated_antecedent"
	input := UpdateMedicalFolderPatientInput{
		MedicalAntecedentInput: model.UpdateMedicalAntecedentsInput{
			Name:     &updatedName,
			Symptoms: []string{"updated_symptoms"},
		},
	}

	response := UpdateMedicalAntecedent(patientID, input, "valid_antecedent_id")

	if response.Code == 200 {
		t.Errorf("Expected non-200 code but got %d", response.Code)
	}
	if response.Err == nil {
		t.Errorf("Expected error but got none")
	}
}

func TestUpdateMedicalAntecedentWithNoMedicalFolder(t *testing.T) {
	patientID := "invalid_patient_id"
	updatedName := "updated_antecedent"
	input := UpdateMedicalFolderPatientInput{
		MedicalAntecedentInput: model.UpdateMedicalAntecedentsInput{
			Name:     &updatedName,
			Symptoms: []string{"updated_symptoms"},
		},
	}

	response := UpdateMedicalAntecedent(patientID, input, "valid_antecedent_id")

	if response.Code == 200 {
		t.Errorf("Expected non-200 code but got %d", response.Code)
	}
	if response.Err == nil {
		t.Errorf("Expected error but got none")
	}
}

func TestUpdateMedicalAntecedentWithInvalidInput(t *testing.T) {
	patientID := "invalid_patient_id"
	updatedName := "updated_antecedent"
	input := UpdateMedicalFolderPatientInput{
		MedicalAntecedentInput: model.UpdateMedicalAntecedentsInput{
			Name:     &updatedName,
			Symptoms: []string{"updated_symptoms"},
		},
	}

	response := UpdateMedicalAntecedent(patientID, input, "valid_antecedent_id")

	if response.Code == 200 {
		t.Errorf("Expected non-200 code but got %d", response.Code)
	}
	if response.Err == nil {
		t.Errorf("Expected error but got none")
	}
}
