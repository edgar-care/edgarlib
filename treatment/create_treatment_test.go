package treatment

import (
	"github.com/edgar-care/edgarlib/v2/medical_folder"
	"github.com/joho/godotenv"
	"log"
	"testing"

	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
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

	medicalFolderInput := medical_folder.CreateNewMedicalInfoInput{
		Name:                   "test",
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
		Name:     "new_antecedent",
		Symptoms: []string{"symptoms"},
		Treatments: []medical_folder.CreateTreatInput{{
			StartDate: 1234,
			EndDate:   5678,
			Medicines: []medical_folder.CreateAntecedentsMedicines{{
				Comment:    func(s string) *string { return &s }("comment"),
				MedicineID: "test",
				Period: []*medical_folder.CreateAntecedentPeriod{{
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

	ante := medical_folder.AddMedicalAntecedent(antecedentInput, patient.ID)

	treatmentInput := CreateTreatInput{
		MedicalantecedentID: ante.MedicalAntecedents[0].ID,
		StartDate:           1234,
		EndDate:             5678,
		Medicines: []CreateAntecedentsMedicines{{
			MedicineID: "test",
			Period: []CreateAntecedentPeriod{{
				Quantity:       2,
				Frequency:      2,
				FrequencyRatio: 2,
				FrequencyUnit:  "JOUR",
			}},
		}},
	}

	response := CreateTreatment(treatmentInput, patient.ID)
	if response.Code != 201 {
		t.Errorf("Expected response code 400, got %v", response.Code)
	}

	_, err = graphql.GetMedicalAntecedentsById(ante.MedicalAntecedents[0].ID)
	if err != nil {
		t.Errorf("Failed to get medical antecedent: %v", err)
	}

	_, test := graphql.GetAntecedentTreatmentByID(response.Treatment[0].ID)
	if test != nil {
		t.Errorf("Failed to get treatment: %v", test)
	}
	//spew.Dump(check)

}
