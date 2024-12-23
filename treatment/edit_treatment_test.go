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
			StartDate: 1234,
			EndDate:   5678,
			Medicines: []medical_folder.CreateAntecedentsMedicines{
				{
					MedicineID: "test",
					Comment:    func(s string) *string { return &s }("comment"),
					Period: []*medical_folder.CreateAntecedentPeriod{
						{
							Quantity:       2,
							Frequency:      2,
							FrequencyRatio: 2,
							FrequencyUnit:  "ANNEE",
							PeriodLength:   func(s int) *int { return &s }(2),
							PeriodUnit:     func(s string) *string { return &s }("JOUR"),
						},
						{
							Quantity:       32,
							Frequency:      31,
							FrequencyRatio: 36,
							FrequencyUnit:  "JOUR",
							PeriodLength:   func(s int) *int { return &s }(23),
							PeriodUnit:     func(s string) *string { return &s }("ANNEE"),
						},
					},
				},
				{
					MedicineID: "test2",
					Comment:    func(s string) *string { return &s }("comment2"),
					Period: []*medical_folder.CreateAntecedentPeriod{{
						Quantity:       13,
						Frequency:      13,
						FrequencyRatio: 13,
						FrequencyUnit:  "MOIS",
						PeriodLength:   func(s int) *int { return &s }(3),
						PeriodUnit:     func(s string) *string { return &s }("ANNEE"),
					}},
				},
			},
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
				FrequencyUnit:  "MOIS",
			}},
		}},
	}

	treat := CreateTreatment(treatmentInput, patient.ID)
	if treat.Err != nil {
		t.Errorf("Error while creating treatment: %v", treat.Err)
	}

	input := UpdateTreatmentInput{
		StartDate: 9876,
		EndDate:   7653,
		Medicines: []UpdateAntecedentsMedicines{
			{
				MedicineID: "AZE",
				Comment:    "comment",
				Period: []UpdateAntecedentPeriod{
					{
						Quantity:       3,
						Frequency:      3,
						FrequencyRatio: 3,
						FrequencyUnit:  "ANNEE",
						PeriodLength:   32,
						PeriodUnit:     "JOUR",
					},
					{
						Quantity:       4,
						Frequency:      4,
						FrequencyRatio: 4,
						FrequencyUnit:  "MOIS",
					},
				},
			},
			//{
			//	MedicineID: "test2",
			//	Comment:    "commenttest",
			//	Period: []UpdateAntecedentPeriod{
			//		{
			//			Quantity:       4,
			//			Frequency:      4,
			//			FrequencyRatio: 4,
			//			FrequencyUnit:  "MOIS",
			//			PeriodLength:   41,
			//			PeriodUnit:     "JOUR",
			//		},
			//	},
			//},
		},
	}

	response := UpdateTreatment(input, patient.ID, ante.MedicalAntecedents[0].Treatments[0].ID)
	if response.Code != 200 {
		t.Errorf("Expected code 200 but got %d", response.Code)
	}
	if response.Err != nil {
		t.Errorf("Expected no error but got: %s", response.Err.Error())
	}

	_ = GetTreatmentById(ante.MedicalAntecedents[0].Treatments[0].ID, patient.ID)
	_ = medical_folder.GetMedicalAntecedents(patient.ID)

}
