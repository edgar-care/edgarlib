package dashboard

import (
	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
	"github.com/edgar-care/edgarlib/v2/medical_folder"
	"github.com/joho/godotenv"
	"log"
	"testing"
)

func TestConvertCreateToUpdateMedicalInfoInput(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}

	doctor, err := graphql.CreateDoctor(model.CreateDoctorInput{
		Email:     "doctor_test_convert_create_to_update@edgar-sante.fr",
		Password:  "password",
		Name:      "name",
		Firstname: "first",
		Address:   &model.AddressInput{"", "", "", ""},
		Status:    true,
	})
	if err != nil {
		t.Errorf("Error while creating doctor: %v", err)
	}

	createInput := medical_folder.CreateMedicalInfoInput{
		Name:            "Fabien",
		Firstname:       "leFab",
		Birthdate:       12,
		Sex:             "M",
		Weight:          70,
		Height:          175,
		PrimaryDoctorID: doctor.ID,
		MedicalAntecedents: []medical_folder.CreateMedicalAntecedentInput{
			{
				Name: "Hypertension",
				Medicines: []medical_folder.CreateMedicineInput{
					{
						MedicineID: "med123",
						Period:     []string{"MORNING"},
						Day:        []string{"MONDAY"},
						Quantity:   1,
					},
				},
				StillRelevant: true,
			},
		},
	}

	updateInput := ConvertCreateToUpdateMedicalInfoInput(createInput)

	if updateInput.Name != createInput.Name {
		t.Errorf("Expected Name: %s, got: %s", createInput.Name, updateInput.Name)
	}
	if len(updateInput.MedicalAntecedents) != len(createInput.MedicalAntecedents) {
		t.Errorf("Expected %d antecedents, got: %d", len(createInput.MedicalAntecedents), len(updateInput.MedicalAntecedents))
	}
	if len(updateInput.MedicalAntecedents[0].Medicines) != len(createInput.MedicalAntecedents[0].Medicines) {
		t.Errorf("Expected %d medicines, got: %d", len(createInput.MedicalAntecedents[0].Medicines), len(updateInput.MedicalAntecedents[0].Medicines))
	}
}

func TestCreatePatientFormDoctor_ExistingPatient(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}
	doctor, err := graphql.CreateDoctor(model.CreateDoctorInput{
		Email:     "doctor_test_create_patient_form_doctor_existing@edgar-sante.fr",
		Password:  "password",
		Name:      "name",
		Firstname: "first",
		Address:   &model.AddressInput{"", "", "", ""},
		Status:    true,
	})
	if err != nil {
		t.Errorf("Error while creating doctor: %v", err)
	}

	patient, err := graphql.CreatePatient(model.CreatePatientInput{
		Email:    "existing_patient_create_from_doctor@edgar-sante.fr",
		Password: "password",
	})
	if err != nil {
		t.Errorf("Error while creating patient: %v", err)
	}

	medicalInfo, err := graphql.CreateMedicalFolder(model.CreateMedicalFolderInput{
		Name:                   "Noé",
		Firstname:              "le poisson",
		Birthdate:              12,
		Sex:                    "F",
		Height:                 10,
		Weight:                 1,
		PrimaryDoctorID:        doctor.ID,
		AntecedentDiseaseIds:   []string{},
		FamilyMembersMedInfoID: []string{},
	})
	if err != nil {
		t.Errorf("Error while creating medical folder: %v", err)
	}

	_, err = graphql.UpdatePatient(patient.ID, model.UpdatePatientInput{MedicalInfoID: &medicalInfo.ID})
	if err != nil {
		t.Errorf("Error while updating patient: %v", err)
	}

	newPatient := CreatePatientInput{
		Email: "existing_patient_create_from_doctor@edgar-sante.fr",
		MedicalInfo: medical_folder.CreateMedicalInfoInput{
			Name:            "Noé",
			Firstname:       "le marin",
			Birthdate:       69,
			Sex:             "M",
			Weight:          70,
			Height:          175,
			PrimaryDoctorID: doctor.ID,
			MedicalAntecedents: []medical_folder.CreateMedicalAntecedentInput{
				{
					Name: "Hypertension",
					Medicines: []medical_folder.CreateMedicineInput{
						{
							MedicineID: "med123",
							Period:     []string{"EVENING"},
							Day:        []string{"TUESDAY", "WEDNESDAY"},
							Quantity:   1,
						},
					},
					StillRelevant: true,
				},
			},
			FamilyMembersMedInfoId: []string{},
		},
	}
	doctorID := doctor.ID

	response := CreatePatientFormDoctor(newPatient, doctorID)

	if response.Code != 200 {
		t.Errorf("Expected Code: 200, got: %d", response.Code)
	}

	if response.Err != nil {
		t.Errorf("Expected no error, but got: %v", response.Err)
	}

	if response.Patient.Email != newPatient.Email {
		t.Errorf("Expected Patient Email: %s, got: %s", newPatient.Email, response.Patient.Email)
	}
}

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
		Email: "new_patient_test_create_patient_form_doctor@example.com",
		MedicalInfo: medical_folder.CreateMedicalInfoInput{
			Name:            "Marvin",
			Firstname:       "le bg",
			Birthdate:       8,
			Sex:             "F",
			Weight:          60,
			Height:          165,
			PrimaryDoctorID: doctor.ID,
			MedicalAntecedents: []medical_folder.CreateMedicalAntecedentInput{
				{
					Name: "Asthma",
					Medicines: []medical_folder.CreateMedicineInput{
						{
							MedicineID: "med456",
							Period:     []string{"NOON", "NIGHT"},
							Day:        []string{"SUNDAY"},
							Quantity:   1,
						},
					},
					StillRelevant: true,
				},
			},
			FamilyMembersMedInfoId: []string{},
		},
	}
	doctorID := doctor.ID

	response := CreatePatientFormDoctor(newPatient, doctorID)

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
