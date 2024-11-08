package dashboard

import (
	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
	"github.com/edgar-care/edgarlib/v2/medical_folder"
	"github.com/joho/godotenv"
	"log"
	"testing"
)

//func TestCreatePatientFormDoctor_ExistingPatient(t *testing.T) {
//	if err := godotenv.Load(".env.test"); err != nil {
//		log.Fatalf("Error loading .env.test file: %v", err)
//	}
//	doctor, err := graphql.CreateDoctor(model.CreateDoctorInput{
//		Email:     "doctor_test_create_patient_form_doctor_existing@edgar-sante.fr",
//		Password:  "password",
//		Name:      "name",
//		Firstname: "first",
//		Address:   &model.AddressInput{"", "", "", ""},
//		Status:    true,
//	})
//	if err != nil {
//		t.Errorf("Error while creating doctor: %v", err)
//	}
//
//	patient, err := graphql.CreatePatient(model.CreatePatientInput{
//		Email:    "existing_patient_create_from_doctor@edgar-sante.fr",
//		Password: "password",
//	})
//	if err != nil {
//		t.Errorf("Error while creating patient: %v", err)
//	}
//
//	medicalInfo, err := graphql.CreateMedicalFolder(model.CreateMedicalFolderInput{
//		Name:                   "Noé",
//		Firstname:              "le poisson",
//		Birthdate:              12,
//		Sex:                    "F",
//		Height:                 10,
//		Weight:                 1,
//		PrimaryDoctorID:        doctor.ID,
//		AntecedentDiseaseIds:   []string{},
//		FamilyMembersMedInfoID: []string{},
//	})
//	if err != nil {
//		t.Errorf("Error while creating medical folder: %v", err)
//	}
//
//	_, err = graphql.UpdatePatient(patient.ID, model.UpdatePatientInput{MedicalInfoID: &medicalInfo.ID})
//	if err != nil {
//		t.Errorf("Error while updating patient: %v", err)
//	}
//
//	newPatient := CreatePatientInput{
//		Email: "existing_patient_create_from_doctor@edgar-sante.fr",
//		MedicalInfo: medical_folder.CreateNewMedicalInfoInput{
//			Name:            "Noé",
//			Firstname:       "le marin",
//			Birthdate:       69,
//			Sex:             "M",
//			Weight:          70,
//			Height:          175,
//			PrimaryDoctorID: doctor.ID,
//			MedicalAntecedents: []medical_folder.CreateNewMedicalAntecedentInput{
//				{
//					Name:     "Hypertension",
//					Symptoms: []string{"headache"},
//					Treatments: []medical_folder.CreateTreatInput{
//						{
//							CreatedBy: "test",
//							StartDate: 123,
//							EndDate:   234,
//							Medicines: []medical_folder.CreateAntecedentsMedicines{
//								{
//									Period: []*medical_folder.CreateAntecedentPeriod{
//										{
//											Quantity:       1,
//											Frequency:      2,
//											FrequencyRatio: 2,
//											FrequencyUnit:  "JOUR",
//											PeriodLength:   2,
//											PeriodUnit:     "ANNEE",
//											Comment:        "test",
//										},
//										{
//											Quantity:       5,
//											Frequency:      1,
//											FrequencyRatio: 6,
//											FrequencyUnit:  "MOIS",
//											PeriodLength:   1,
//											PeriodUnit:     "ANNEE",
//											Comment:        "test",
//										},
//									},
//								},
//							},
//						},
//					},
//				},
//				{
//					Name:     "test",
//					Symptoms: []string{"test"},
//					Treatments: []medical_folder.CreateTreatInput{
//						{
//							CreatedBy: "test",
//							StartDate: 13,
//							EndDate:   24,
//							Medicines: []medical_folder.CreateAntecedentsMedicines{
//								{
//									Period: []*medical_folder.CreateAntecedentPeriod{
//										{
//											Quantity:       1,
//											Frequency:      2,
//											FrequencyRatio: 2,
//											FrequencyUnit:  "ANNEE",
//											PeriodLength:   2,
//											PeriodUnit:     "JOUR",
//											Comment:        "test",
//										},
//									},
//								},
//							},
//						},
//					},
//				},
//			},
//			FamilyMembersMedInfoId: []string{},
//		},
//	}
//	doctorID := doctor.ID
//
//	response := CreatePatientFromDoctor(doctorID, newPatient)
//
//	if response.Code != 200 {
//		t.Errorf("Expected Code: 200, got: %d", response.Code)
//	}
//
//	if response.Err != nil {
//		t.Errorf("Expected no error, but got: %v", response.Err)
//	}
//
//	if response.Patient.Email != newPatient.Email {
//		t.Errorf("Expected Patient Email: %s, got: %s", newPatient.Email, response.Patient.Email)
//	}
//}

func TestCreatePatientFormDoctor_NewPatient(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}
	doctor, err := graphql.CreateDoctor(model.CreateDoctorInput{
		Email:     "doctor_test_create_patient_form_doctor_new_patient@edgar-sante.fr",
		Password:  "password",
		Name:      "name",
		Firstname: "first",
		Address:   &model.AddressInput{"", "", "", ""},
		Status:    true,
	})
	if err != nil {
		t.Errorf("Error while creating doctor: %v", err)
	}

	newPatient := CreatePatientInput{
		Email: "new_patient_test_create_patient_form_doctor@edgar-sante.fr",
		MedicalInfo: medical_folder.CreateNewMedicalInfoInput{
			Name:            "Noé",
			Firstname:       "le marin",
			Birthdate:       69,
			Sex:             "M",
			Weight:          70,
			Height:          175,
			PrimaryDoctorID: doctor.ID,
			MedicalAntecedents: []medical_folder.CreateNewMedicalAntecedentInput{
				{
					Name:     "Hypertension",
					Symptoms: []string{"headache"},
					Treatments: []medical_folder.CreateTreatInput{
						{
							CreatedBy: "test",
							StartDate: 123,
							EndDate:   234,
							Medicines: []medical_folder.CreateAntecedentsMedicines{
								{
									Period: []*medical_folder.CreateAntecedentPeriod{
										{
											Quantity:       1,
											Frequency:      2,
											FrequencyRatio: 2,
											FrequencyUnit:  "JOUR",
											PeriodLength:   2,
											PeriodUnit:     "ANNEE",
											Comment:        "test",
										},
										{
											Quantity:       5,
											Frequency:      1,
											FrequencyRatio: 6,
											FrequencyUnit:  "MOIS",
											PeriodLength:   1,
											PeriodUnit:     "ANNEE",
											Comment:        "test",
										},
									},
								},
							},
						},
					},
				},
				{
					Name:     "test",
					Symptoms: []string{"test"},
					Treatments: []medical_folder.CreateTreatInput{
						{
							CreatedBy: "test",
							StartDate: 13,
							EndDate:   24,
							Medicines: []medical_folder.CreateAntecedentsMedicines{
								{
									Period: []*medical_folder.CreateAntecedentPeriod{
										{
											Quantity:       1,
											Frequency:      2,
											FrequencyRatio: 2,
											FrequencyUnit:  "ANNEE",
											PeriodLength:   2,
											PeriodUnit:     "JOUR",
											Comment:        "test",
										},
									},
								},
							},
						},
					},
				},
			},
			FamilyMembersMedInfoId: []string{},
		},
	}
	doctorID := doctor.ID

	response := CreatePatientFromDoctor(doctorID, newPatient)

	if response.Code != 201 {
		t.Errorf("Expected Code: 201, got: %d", response.Code)
	}

	if response.Err != nil {
		t.Errorf("Expected no error, but got: %v", response.Err)
	}

	if response.Patient.Email != newPatient.Email {
		t.Errorf("Expected Patient Email: %s, got: %s", newPatient.Email, response.Patient.Email)
	}
}
