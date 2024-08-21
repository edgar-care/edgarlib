package medical_folder

import (
	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
	"github.com/joho/godotenv"
	"log"
	"testing"
)

func TestUpdateMedicalFolder(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}

	patient, err := graphql.CreatePatient(model.CreatePatientInput{
		Email:    "test_patient_update_medical_folder_up@edgar-sante.fr",
		Password: "password",
		Status:   false,
	})
	if err != nil {
		t.Errorf("Error while creating patient: %v", err)
	}

	input := CreateMedicalInfoInput{
		Name:            "test",
		Firstname:       "test",
		Birthdate:       0,
		Sex:             "M",
		Weight:          0,
		Height:          0,
		PrimaryDoctorID: "",
		MedicalAntecedents: []CreateMedicalAntecedentInput{{
			Name: "test",
			Medicines: []CreateMedicineInput{CreateMedicineInput{
				MedicineID: "test",
				Period:     []string{"NOON"},
				Day:        []string{"MONDAY"},
				Quantity:   2,
				StartDate:  1234,
				EndDate:    1234,
			}},
			StillRelevant: false,
		},
		},
		FamilyMembersMedInfoId: []string{"test"},
	}

	medicalfolder := CreateMedicalInfo(input, patient.ID)

	inputUpdate := UpdateMedicalInfoInput{
		Name:                   "zetgfdg",
		Firstname:              "zerzer",
		Birthdate:              2,
		Sex:                    "F",
		Weight:                 34,
		Height:                 56,
		PrimaryDoctorID:        "9T",
		MedicalAntecedents:     []UpdateMedicalAntecedentInput{},
		FamilyMembersMedInfoId: []string{"KVFDFD"},
	}
	response := UpdateMedicalFolder(inputUpdate, medicalfolder.MedicalInfo.ID)
	if response.Code != 200 {
		t.Errorf("Expected code 200 but got %d", response.Code)
	}
	if response.Err != nil {
		t.Errorf("Expected no error but got: %s", response.Err.Error())
	}

}
