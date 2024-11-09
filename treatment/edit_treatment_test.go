package treatment

import (
	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
	"github.com/edgar-care/edgarlib/v2/medical_folder"
	"github.com/joho/godotenv"
	"log"
	"testing"
)

func TestUpdateTreatment(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}
	email := "testuser_update_treatment@example.com"

	patient, err := graphql.CreatePatient(model.CreatePatientInput{
		Email:    email,
		Password: "testpassword",
		Status:   true,
	})
	if err != nil {
		t.Errorf("Failed to create patient: %v", err)
	}

	medicalFolderInput := medical_folder.CreateNewMedicalInfoInput{
		Name:                   "test_get_treatmentby_id",
		Firstname:              "test",
		Birthdate:              0,
		Sex:                    "M",
		Weight:                 0,
		Height:                 0,
		PrimaryDoctorID:        "",
		MedicalAntecedents:     []medical_folder.CreateNewMedicalAntecedentInput{},
		FamilyMembersMedInfoId: []string{""},
	}

	_ = medical_folder.NewMedicalFolder(medicalFolderInput, patient.ID)

	antecedentInput := medical_folder.CreateNewMedicalAntecedentInput{
		Name:     "new_antecedent_get_byid",
		Symptoms: []string{"symptoms"},
		Treatments: []medical_folder.CreateTreatInput{{
			CreatedBy: "test1_get_treatmentby_id",
			StartDate: 1234,
			EndDate:   5678,
			Medicines: []medical_folder.CreateAntecedentsMedicines{{
				Period: []*medical_folder.CreateAntecedentPeriod{{
					Quantity:       2,
					Frequency:      2,
					FrequencyRatio: 2,
					FrequencyUnit:  "ANNEE",
					PeriodLength:   2,
					PeriodUnit:     "JOUR",
					Comment:        "comment",
				}},
			}},
		}},
	}

	ante := medical_folder.AddMedicalAntecedent(antecedentInput, patient.ID)

	treatmentInput := CreateTreatInput{
		MedicalantecedentID: ante.MedicalAntecedents[0].ID,
		CreatedBy:           "testgetbyid",
		StartDate:           1234,
		EndDate:             5678,
		Medicines: []CreateAntecedentsMedicines{{
			Period: []CreateAntecedentPeriod{{
				Quantity:       2,
				Frequency:      2,
				FrequencyRatio: 2,
				FrequencyUnit:  "MOIS",
				PeriodLength:   2,
				PeriodUnit:     "JOUR",
				Comment:        "comment",
			}},
		}},
	}

	treat := CreateTreatment(treatmentInput, patient.ID)
	if treat.Err != nil {
		t.Errorf("Error while creating treatment: %v", treat.Err)
	}

	input := UpdateTreatmentInput{
		ID:        treat.Treatment[0].ID,
		CreatedBy: "testupdate",
		StartDate: 9876,
		EndDate:   7653,
		Medicines: []UpdateAntecedentsMedicines{{
			Period: []UpdateAntecedentPeriod{{
				Quantity:       3,
				Frequency:      3,
				FrequencyRatio: 3,
				FrequencyUnit:  "JOUR",
				PeriodLength:   3,
				PeriodUnit:     "JOUR",
				Comment:        "comment",
			}},
		}},
	}

	response := UpdateTreatment(input, patient.ID, ante.MedicalAntecedents[0].ID)
	if response.Code != 200 {
		t.Errorf("Expected code 200 but got %d", response.Code)
	}
	if response.Err != nil {
		t.Errorf("Expected no error but got: %s", response.Err.Error())
	}

}
