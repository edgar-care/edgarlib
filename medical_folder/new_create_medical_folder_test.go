package medical_folder

import (
	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
	"github.com/joho/godotenv"
	"log"
	"testing"
)

func TestAddMedicalAntecedent(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}

	// Create a patient
	patient, err := graphql.CreatePatient(model.CreatePatientInput{
		Email:    "test_patient_add_medical_antecedent@edgar-sante.fr",
		Password: "password",
		Status:   false,
	})
	if err != nil {
		t.Errorf("Error while creating patient: %v", err)
	}

	// Create a medical folder for the patient
	medicalFolderInput := CreateNewMedicalInfoInput{
		Name:            "test",
		Firstname:       "test",
		Birthdate:       32,
		Sex:             "M",
		Weight:          123,
		Height:          123,
		PrimaryDoctorID: "test",
		MedicalAntecedents: []CreateNewMedicalAntecedentInput{
			{
				Name:     "Hypertension",
				Symptoms: []string{"headache"},
				Treatments: []CreateTreatInput{
					{
						StartDate: 123,
						EndDate:   234,
						Medicines: []CreateAntecedentsMedicines{
							{
								Comment:    func(s string) *string { return &s }("comment"),
								MedicineID: "test",
								Period: []*CreateAntecedentPeriod{
									{
										Quantity:       1,
										Frequency:      2,
										FrequencyRatio: 2,
										FrequencyUnit:  "JOUR",
										PeriodLength:   func(i int) *int { return &i }(2),
										PeriodUnit:     func(s string) *string { return &s }("ANNEE"),
									},
									{
										Quantity:       5,
										Frequency:      1,
										FrequencyRatio: 6,
										FrequencyUnit:  "MOIS",
										PeriodLength:   func(i int) *int { return &i }(1),
										PeriodUnit:     func(s string) *string { return &s }("ANNEE"),
									},
								},
							},
						},
					},
				},
			},
		},
		FamilyMembersMedInfoId: []string{""},
	}

	//jsonInput, err := json.Marshal(medicalFolderInput)
	//if err != nil {
	//	log.Fatalf("Error marshaling input: %v", err)
	//}
	//fmt.Println("Serialized JSON:", string(jsonInput))

	_ = NewMedicalFolder(medicalFolderInput, patient.ID)
	//spew.Dump(tt)
	antecedentInput := CreateNewMedicalAntecedentInput{
		Name:     "new_antecedent",
		Symptoms: []string{"symptoms"},
		Treatments: []CreateTreatInput{{
			StartDate: 1234,
			EndDate:   5678,
			Medicines: []CreateAntecedentsMedicines{
				{
					MedicineID: "test",
					Period: []*CreateAntecedentPeriod{{
						Quantity:       2,
						Frequency:      2,
						FrequencyRatio: 2,
						FrequencyUnit:  "JOUR",
					}},
				},
				{
					MedicineID: "test",
					Period: []*CreateAntecedentPeriod{{
						Quantity:       3,
						Frequency:      23,
						FrequencyRatio: 2,
						FrequencyUnit:  "MOIS",
					},
					},
				}},
		}},
	}

	response := AddMedicalAntecedent(antecedentInput, patient.ID)
	_ = GetMedicalFolder(patient.ID)
	if response.Code != 201 {
		t.Errorf("Expected code 201 but got %d", response.Code)
	}
	if response.Err != nil {
		t.Errorf("Expected no error but got: %s", response.Err.Error())
	}
	if len(response.MedicalAntecedents) == 0 {
		t.Errorf("Expected at least one medical antecedent but got none")
	}
}
