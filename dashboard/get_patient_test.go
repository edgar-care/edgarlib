package dashboard

import (
	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
	"github.com/edgar-care/edgarlib/v2/medical_folder"
	"github.com/joho/godotenv"
	"log"
	"testing"
)

func TestGetPatientById(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}

	patient, err := graphql.CreatePatient(model.CreatePatientInput{
		Email:    "test_get_patient_by_id@edgar-sante.fr",
		Password: "password",
	})
	if err != nil {
		t.Errorf("Error while creating patient: %v", err)
	}

	doctor, err := graphql.CreateDoctor(model.CreateDoctorInput{
		Email:     "test_get_patient_by_id@edgar-sante.fr",
		Password:  "password",
		Name:      "name",
		Firstname: "first",
		Address:   &model.AddressInput{"", "", "", ""},
		Status:    false,
	})
	if err != nil {
		t.Errorf("Error while creating doctor: %v", err)
	}

	_, err = graphql.UpdateDoctor(doctor.ID, model.UpdateDoctorInput{
		RendezVousIds: []*string{},
		PatientIds:    []*string{&patient.ID},
	})
	if err != nil {
		t.Errorf("Error while updating doctor: %v", err)
	}

	medical_folder_resp := medical_folder.NewMedicalFolder(medical_folder.CreateNewMedicalInfoInput{
		Name:            "first",
		Firstname:       "first",
		Birthdate:       0,
		Sex:             "",
		Weight:          0,
		Height:          0,
		PrimaryDoctorID: doctor.ID,
		MedicalAntecedents: []medical_folder.CreateNewMedicalAntecedentInput{{
			Name:     "",
			Symptoms: []string{"symptom"},
			Treatments: []medical_folder.CreateTreatInput{
				{
					CreatedBy: "ttt",
					StartDate: 78,
					EndDate:   90,
					Medicines: []medical_folder.CreateAntecedentsMedicines{
						{
							Period: []*medical_folder.CreateAntecedentPeriod{
								{
									Quantity:       6,
									Frequency:      1,
									FrequencyRatio: 3,
									FrequencyUnit:  "MOIS",
									PeriodLength:   4,
									PeriodUnit:     "ANNEE",
									Comment:        "tttt",
								},
							},
						},
					},
				},
			},
		}},
		FamilyMembersMedInfoId: []string{},
	}, patient.ID)

	if medical_folder_resp.Err != nil {
		t.Errorf("Unexpected error while creating medical info: %v", medical_folder_resp.Err)
	}

	response := GetPatientById(patient.ID, doctor.ID)

	if response.Err != nil {
		t.Errorf("Unexpected error: %v", response.Err)
	}

	if response.Code != 200 {
		t.Errorf("Expected code 200, got %d", response.Code)
	}
}

func TestGetPatients(t *testing.T) {
	doctor, err := graphql.CreateDoctor(model.CreateDoctorInput{
		Email:     "test_get_patiets_by_doctor@edgar-sante.fr",
		Password:  "password",
		Name:      "name",
		Firstname: "first",
		Address:   &model.AddressInput{},
	})
	if err != nil {
		t.Errorf("Failed to create doctor: %v", err)
	}

	patient1, err := graphql.CreatePatient(model.CreatePatientInput{
		Email:    "patient1@example.com",
		Password: "password1",
		Status:   true,
	})
	if err != nil {
		t.Errorf("Failed to create patient1: %v", err)
	}

	_ = medical_folder.NewMedicalFolder(medical_folder.CreateNewMedicalInfoInput{
		Name:            "first",
		Firstname:       "first",
		Birthdate:       0,
		Sex:             "",
		Weight:          0,
		Height:          0,
		PrimaryDoctorID: doctor.ID,
		MedicalAntecedents: []medical_folder.CreateNewMedicalAntecedentInput{{
			Name:     "",
			Symptoms: []string{"symptom"},
			Treatments: []medical_folder.CreateTreatInput{
				{
					CreatedBy: "ttt",
					StartDate: 78,
					EndDate:   90,
					Medicines: []medical_folder.CreateAntecedentsMedicines{
						{
							Period: []*medical_folder.CreateAntecedentPeriod{
								{
									Quantity:       6,
									Frequency:      1,
									FrequencyRatio: 3,
									FrequencyUnit:  "MOIS",
									PeriodLength:   4,
									PeriodUnit:     "ANNEE",
									Comment:        "tttt",
								},
							},
						},
					},
				},
			},
		}},
		FamilyMembersMedInfoId: []string{},
	}, patient1.ID)

	patient2, err := graphql.CreatePatient(model.CreatePatientInput{
		Email:    "patient2@example.com",
		Password: "password2",
		Status:   true,
	})
	if err != nil {
		t.Errorf("Failed to create patient2: %v", err)
	}

	_ = medical_folder.NewMedicalFolder(medical_folder.CreateNewMedicalInfoInput{
		Name:            "Second",
		Firstname:       "Second",
		Birthdate:       0,
		Sex:             "",
		Weight:          0,
		Height:          0,
		PrimaryDoctorID: doctor.ID,
		MedicalAntecedents: []medical_folder.CreateNewMedicalAntecedentInput{{
			Name:     "",
			Symptoms: []string{"symptom"},
			Treatments: []medical_folder.CreateTreatInput{
				{
					CreatedBy: "ttt",
					StartDate: 78,
					EndDate:   90,
					Medicines: []medical_folder.CreateAntecedentsMedicines{
						{
							Period: []*medical_folder.CreateAntecedentPeriod{
								{
									Quantity:       6,
									Frequency:      1,
									FrequencyRatio: 3,
									FrequencyUnit:  "MOIS",
									PeriodLength:   4,
									PeriodUnit:     "ANNEE",
									Comment:        "tttt",
								},
							},
						},
					},
				},
			},
		}},
		FamilyMembersMedInfoId: []string{},
	}, patient2.ID)

	_, err = graphql.UpdateDoctorsPatientIDs(doctor.ID, model.UpdateDoctorsPatientIDsInput{
		PatientIds: []*string{&patient1.ID, &patient2.ID},
	})
	if err != nil {
		t.Errorf("Failed to update doctor: %v", err)
	}

	response := GetPatients(doctor.ID)

	if response.Code != 200 {
		t.Errorf("Expected response code 200, got %v", response.Code)
	}
	if response.Err != nil {
		t.Errorf("Unexpected error: %v", response.Err)
	}

	if len(response.PatientsInfo) != 2 {
		t.Errorf("Expected 2 patients, got %v", len(response.PatientsInfo))
	}

	if response.PatientsInfo[0].ID != patient1.ID {
		t.Errorf("Expected patient1 ID to be %v, got %v", patient1.ID, response.PatientsInfo[0].ID)
	}

	if response.PatientsInfo[1].ID != patient2.ID {
		t.Errorf("Expected patient2 ID to be %v, got %v", patient2.ID, response.PatientsInfo[1].ID)
	}
}

func TestGetPatientsWithInvalidDoctorID(t *testing.T) {
	invalidDoctorID := "invalid-doctor-id"

	response := GetPatients(invalidDoctorID)

	if response.Code != 400 {
		t.Errorf("Expected response code 400, got %v", response.Code)
	}
	if response.Err == nil {
		t.Errorf("Expected an error in response, got nil")
	} else if response.Err.Error() != "id does not correspond to a doctor" {
		t.Errorf("Expected 'id does not correspond to a doctor' error, got %v", response.Err.Error())
	}
}

func TestGetPatientsWithMedicalInfoError(t *testing.T) {
	doctor, err := graphql.CreateDoctor(model.CreateDoctorInput{
		Email:     "test_get_patients_with_medical_info_error@edgar-sante.fr",
		Password:  "password",
		Name:      "name",
		Firstname: "first",
		Address:   &model.AddressInput{},
	})
	if err != nil {
		t.Errorf("Failed to create doctor: %v", err)
	}

	patient, err := graphql.CreatePatient(model.CreatePatientInput{
		Email:    "patient_with_error@example.com",
		Password: "password1",
		Status:   true,
	})
	if err != nil {
		t.Errorf("Failed to create patient: %v", err)
	}

	_, err = graphql.UpdateDoctorsPatientIDs(doctor.ID, model.UpdateDoctorsPatientIDsInput{
		PatientIds: []*string{&patient.ID},
	})
	if err != nil {
		t.Errorf("Failed to update doctor: %v", err)
	}

	response := GetPatients(doctor.ID)

	if response.Code != 401 {
		t.Errorf("Expected response code 401 due to medical info error, got %v", response.Code)
	}
	if response.Err == nil {
		t.Errorf("Expected an error in response, got nil")
	} else if response.Err.Error() != "error while retrieving medical info by id" {
		t.Errorf("Expected 'error while retrieving medical info by id', got %v", response.Err.Error())
	}
}
