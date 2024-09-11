package ordonnance

import (
	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
	"github.com/edgar-care/edgarlib/v2/medical_folder"
	"github.com/joho/godotenv"
	"log"
	"testing"
)

func TestGetOrdonnancebyID(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}

	doctor, err := graphql.CreateDoctor(model.CreateDoctorInput{
		Email:    "test_doctor_get_doctor_ordonnace@edgar-sante.fr",
		Password: "password",
		Status:   true,
		Address: &model.AddressInput{
			Street:  "sqdqsd",
			ZipCode: "dsfsdf",
			Country: "fdgdfg",
			City:    "azeazeaze",
		},
	})
	if err != nil {
		t.Errorf("Error while creating doctor: %v", err)
	}

	medoc, err := graphql.CreateMedicine(model.CreateMedicineInput{
		Dci:             "MOLO",
		TargetDiseases:  []string{"Test"},
		TreatedSymptoms: []string{"headache"},
		SideEffects:     []string{"prout"},
		Dosage:          2,
		DosageUnit:      "g",
		Container:       "TUBE",
		Name:            "daphalgan",
		DosageForm:      "CREME",
	})
	if err != nil {
		t.Errorf("Error while creating medicine: %v", err)
	}

	patient, err := graphql.CreatePatient(model.CreatePatientInput{
		Email:    "test_patient_get_ordonnance@edgar-sante.fr",
		Password: "password",
		Status:   false,
	})
	if err != nil {
		t.Errorf("Error while creating patient: %v", err)
	}

	test := medical_folder.CreateMedicalInfoInput{
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

	_ = medical_folder.CreateMedicalInfo(test, patient.ID)

	input := CreateOrdonnaceInput{
		PatientID: patient.ID,
		Medicines: []MedicineInput{
			{
				MedicineID: medoc.ID,
				Qsp:        2,
				QspUnit:    "JOUR",
				Comment:    "test commentaire sgénérale",
				Periods: []PeriodInput{
					{
						Quantity:       1,
						Frequency:      2,
						FrequencyRatio: 1,
						FrequencyUnit:  "JOUR",
						PeriodLength:   2,
						PeriodUnit:     "JOUR",
					},
				},
			},
		},
	}

	ordonnance := CreateOrdonnance(input, doctor.ID)
	response := GetOrdonnancebyID(ordonnance.Ordonnance.ID)
	if response.Code != 200 {
		t.Errorf("Expected code 200, got %d", response.Code)
	}
	if response.Err != nil {
		t.Errorf("Expected no error, got %v", response.Err)
	}
}

func TestGetOrdonnancesDoctor(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}

	doctor, err := graphql.CreateDoctor(model.CreateDoctorInput{
		Email:    "test_doctor_get_doctor_all_ordonnace@edgar-sante.fr",
		Password: "password",
		Status:   true,
		Address: &model.AddressInput{
			Street:  "dfsdfsdfsd",
			ZipCode: "fdgdfgdg",
			Country: "azfazf",
			City:    "dsfsdfdsf",
		},
	})
	if err != nil {
		t.Errorf("Error while creating doctor: %v", err)
	}

	medoc, err := graphql.CreateMedicine(model.CreateMedicineInput{
		Dci:             "MOLO",
		TargetDiseases:  []string{"Test"},
		TreatedSymptoms: []string{"headache"},
		SideEffects:     []string{"prout"},
		Dosage:          2,
		DosageUnit:      "g",
		Container:       "TUBE",
		Name:            "daphalgan",
		DosageForm:      "CREME",
	})
	if err != nil {
		t.Errorf("Error while creating medicine: %v", err)
	}

	patient, err := graphql.CreatePatient(model.CreatePatientInput{
		Email:    "test_patient_get_all_ordonnance@edgar-sante.fr",
		Password: "password",
		Status:   false,
	})
	if err != nil {
		t.Errorf("Error while creating patient: %v", err)
	}

	test := medical_folder.CreateMedicalInfoInput{
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

	_ = medical_folder.CreateMedicalInfo(test, patient.ID)

	input := CreateOrdonnaceInput{
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
						PeriodLength:   2,
						PeriodUnit:     "JOUR",
					},
				},
			},
		},
	}

	_ = CreateOrdonnance(input, doctor.ID)
	response := GetOrdonnancesDoctor(doctor.ID)
	if response.Code != 200 {
		t.Errorf("Expected code 200, got %d", response.Code)
	}
	if response.Err != nil {
		t.Errorf("Expected no error, got %v", response.Err)
	}
}
