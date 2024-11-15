package treatment

import (
	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
	"github.com/edgar-care/edgarlib/v2/medical_folder"
	"github.com/joho/godotenv"
	"log"
	"testing"
)

func TestDeleteTreatment(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}
	email := "testuser_delete_treatment@example.com"

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
				MedicineID: "test",
				Comment:    "comment",
				Period: []*medical_folder.CreateAntecedentPeriod{{
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

	ante := medical_folder.AddMedicalAntecedent(antecedentInput, patient.ID)

	treatmentInput := CreateTreatInput{
		MedicalantecedentID: ante.MedicalAntecedents[0].ID,
		StartDate:           1234,
		EndDate:             5678,
		Medicines: []CreateAntecedentsMedicines{{
			MedicineID: "test",
			Comment:    "comment",
			Period: []CreateAntecedentPeriod{{
				Quantity:       2,
				Frequency:      2,
				FrequencyRatio: 2,
				FrequencyUnit:  "JOUR",
				PeriodLength:   2,
				PeriodUnit:     "JOUR",
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
	_, err = graphql.GetMedicalAntecedentsById(ante.MedicalAntecedents[0].ID)
	if err != nil {
		t.Errorf("Failed to get antecedent treatments: %v", err)
	}

	response2 := DeleteTreatment(response.Treatment[0].ID)
	if response2.Code != 200 {
		t.Errorf("Expected response code 200, got %v", response2.Code)
	}

	_, err = graphql.GetMedicalAntecedentsById(ante.MedicalAntecedents[0].ID)
	if err != nil {
		t.Errorf("Failed to get antecedent treatments: %v", err)
	}

}

//func TestDeleteTreatmentWithValidInput(t *testing.T) {
//	//treatmentID := uuid.New().String()
//
//	anteDisease, err := graphql.CreateAnteDisease(model.CreateAnteDiseaseInput{
//		Name:          "Hypertension",
//		TreatmentIds:  []string{""},
//		SurgeryIds:    []string{},
//		Symptoms:      []string{},
//		StillRelevant: true,
//	})
//	if err != nil {
//		t.Errorf("Failed to create AnteDisease: %v", err)
//	}
//
//	treatment, err := graphql.CreateTreatment(model.CreateTreatmentInput{
//		MedicineID: "test",
//		Quantity:   1,
//		Period:     []model.Period{"MORNING"},
//		Day:        []model.Day{"MONDAY"},
//		StartDate:  1234,
//		EndDate:    2344,
//	})
//	if err != nil {
//		t.Errorf("Failed to create treatment: %v", err)
//	}
//
//	treatment2, err := graphql.CreateTreatment(model.CreateTreatmentInput{
//		MedicineID: "testSecond",
//		Quantity:   8,
//		Period:     []model.Period{"NOON"},
//		Day:        []model.Day{"MONDAY"},
//		StartDate:  56,
//		EndDate:    78,
//	})
//	if err != nil {
//		t.Errorf("Failed to create treatment: %v", err)
//	}
//
//	input := model.UpdateAnteDiseaseInput{
//		TreatmentIds: []string{treatment.ID},
//	}
//
//	ttt, err := graphql.UpdateAnteDisease(anteDisease.ID, input)
//	if err != nil {
//		t.Errorf("Failed to update antedisease: %v", err)
//	}
//
//	input2 := model.UpdateAnteDiseaseInput{
//		TreatmentIds: append(ttt.TreatmentIds, treatment2.ID),
//	}
//	_, err = graphql.UpdateAnteDisease(anteDisease.ID, input2)
//	if err != nil {
//		t.Errorf("Failed to update antedisease: %v", err)
//	}
//
//	_, err = graphql.GetAnteDiseaseByID(anteDisease.ID)
//	if err != nil {
//		t.Errorf("Failed to retrieve AnteDisease: %v", err)
//	}
//	response := DeleteTreatment(treatment.ID)
//
//	if response.Code != 200 {
//		t.Errorf("Expected response code 200, got %v", response.Code)
//	}
//	if response.Err != nil {
//		t.Errorf("Unexpected error: %v", response.Err)
//	}
//	if !response.Deleted {
//		t.Errorf("Expected treatment to be deleted, but it was not")
//	}
//
//	_, err = graphql.GetAnteDiseaseByID(anteDisease.ID)
//	if err != nil {
//		t.Errorf("Failed to retrieve AnteDisease: %v", err)
//	}
//
//	//if contains(updatedAnteDisease.TreatmentIds, treatment.ID) {
//	//	t.Errorf("Treatment ID was not removed from AnteDisease")
//	//}
//}
//
//func TestDeleteTreatmentWithEmptyTreatmentID(t *testing.T) {
//	treatmentID := ""
//
//	response := DeleteTreatment(treatmentID)
//
//	if response.Code != 400 {
//		t.Errorf("Expected response code 400, got %v", response.Code)
//	}
//	if response.Err == nil {
//		t.Errorf("Expected an error in response")
//	} else if response.Err.Error() != "treatment id is required" {
//		t.Errorf("Expected 'treatment id is required' error, got %v", response.Err.Error())
//	}
//	if response.Deleted {
//		t.Errorf("Expected treatment not to be deleted")
//	}
//}
