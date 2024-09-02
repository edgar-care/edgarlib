package ordonnance

import (
	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
	"github.com/edgar-care/edgarlib/v2/medical_folder"
	"github.com/joho/godotenv"
	"log"
	"testing"
)

func TestCreateOrdonnance(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}

	doctor, err := graphql.CreateDoctor(model.CreateDoctorInput{
		Email:     "test_doctor_create_doctor_ordonnace@edgar-sante.fr",
		Password:  "password",
		Name:      "Test",
		Firstname: "Doctor",
		Status:    true,
		Address: &model.AddressInput{
			Street:  "12 rue de Paul",
			ZipCode: "78304",
			Country: "France",
			City:    "Lyon",
		},
	})
	if err != nil {
		t.Errorf("Error while creating doctor: %v", err)
	}

	patient, err := graphql.CreatePatient(model.CreatePatientInput{
		Email:    "test_patient_create_ordonnance@edgar-sante.fr",
		Password: "password",
		Status:   false,
	})
	if err != nil {
		t.Errorf("Error while creating patient: %v", err)
	}

	input := medical_folder.CreateMedicalInfoInput{
		Name:            "test",
		Firstname:       "test",
		Birthdate:       18,
		Sex:             "M",
		Weight:          0,
		Height:          0,
		PrimaryDoctorID: "",
		MedicalAntecedents: []medical_folder.CreateMedicalAntecedentInput{{
			Name: "test",
			Medicines: []medical_folder.CreateMedicineInput{medical_folder.CreateMedicineInput{
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

	_ = medical_folder.CreateMedicalInfo(input, patient.ID)

	medoc, err := graphql.CreateMedicine(model.CreateMedicineInput{
		Dci:             "Para",
		TargetDiseases:  []string{"Test"},
		TreatedSymptoms: []string{"headache"},
		SideEffects:     []string{"prout"},
		Dosage:          2,
		DosageUnit:      "mg",
		Container:       "BOITE",
		Name:            "Doliprane",
		DosageForm:      "GELULE",
	})
	if err != nil {
		t.Errorf("Error while creating medicine: %v", err)
	}

	medoc2, err := graphql.CreateMedicine(model.CreateMedicineInput{
		Dci:             "sdfsdfsd",
		TargetDiseases:  []string{"Tsdfs"},
		TreatedSymptoms: []string{"headache"},
		SideEffects:     []string{"psdffds"},
		Dosage:          2,
		DosageUnit:      "ml",
		Container:       "FLACON",
		Name:            "paracetamol",
		DosageForm:      "SOLUTION_BUVABLE",
	})
	if err != nil {
		t.Errorf("Error while creating medicine: %v", err)
	}

	test := CreateOrdonnaceInput{
		PatientID: patient.ID,
		Medicines: []MedicineInput{
			{
				MedicineID: medoc.ID,
				Qsp:        2,
				QspUnit:    "JOUR",
				Comment:    "test comment",
				Periods: []PeriodInput{
					{
						Quantity:       1,
						Frequency:      2,
						FrequencyRatio: 1,
						FrequencyUnit:  "JOUR",
						//PeriodLength:   2,
						//PeriodUnit:     "ANNEE",
					},
					{
						Quantity:       10,
						Frequency:      1,
						FrequencyRatio: 4,
						FrequencyUnit:  "MOIS",
						PeriodLength:   3,
						PeriodUnit:     "MOIS",
					},
				},
			},
			{
				MedicineID: medoc2.ID,
				Qsp:        5,
				QspUnit:    "MOIS",
				Comment:    "test commentaire generale 2",
				Periods: []PeriodInput{
					{
						Quantity:       10,
						Frequency:      5,
						FrequencyRatio: 4,
						FrequencyUnit:  "JOUR",
						PeriodLength:   6,
						PeriodUnit:     "MOIS",
					},
				},
			},
		},
	}

	response := CreateOrdonnance(test, doctor.ID)
	if response.Code != 200 {
		t.Errorf("Expected code 200, got %d", response.Code)
	}
	if response.Err != nil {
		t.Errorf("Expected no error, got %v", response.Err)
	}
}