package dashboard

import (
	"github.com/edgar-care/edgarlib/graphql"
	"github.com/edgar-care/edgarlib/graphql/model"
	"github.com/edgar-care/edgarlib/medical_folder"
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

	medical_folder_resp := medical_folder.CreateMedicalInfo(medical_folder.CreateMedicalInfoInput{
		Name:            "",
		Firstname:       "",
		Birthdate:       0,
		Sex:             "",
		Weight:          0,
		Height:          0,
		PrimaryDoctorID: "",
		MedicalAntecedents: []medical_folder.CreateMedicalAntecedentInput{{
			Name:          "",
			Medicines:     nil,
			StillRelevant: false,
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

//
//func TestGetPatients(t *testing.T) {
//	if err := godotenv.Load(".env.test"); err != nil {
//		log.Fatalf("Error loading .env.test file: %v", err)
//	}
//	gqlClient := graphql.CreateClient()
//
//	patient1, err := graphql.CreatePatient(context.Background(), gqlClient, "test_get_patient_1@edgar-sante.fr", "password")
//	if err != nil {
//		t.Errorf("Error while creating patient: %v", err)
//	}
//	medical_folder_resp := medical_folder.CreateMedicalInfo(medical_folder.CreateMedicalInfoInput{
//		Name:            "",
//		Firstname:       "",
//		Birthdate:       0,
//		Sex:             "",
//		Weight:          0,
//		Height:          0,
//		PrimaryDoctorID: "",
//		MedicalAntecedents: []medical_folder.CreateMedicalAntecedentInput{{
//			Name:          "",
//			Medicines:     nil,
//			StillRelevant: false,
//		}},
//	}, patient1.CreatePatient.Id)
//
//	if medical_folder_resp.Err != nil {
//		t.Errorf("Unexpected error while creating medical info: %v", medical_folder_resp.Err)
//	}
//	patient2, err := graphql.CreatePatient(context.Background(), gqlClient, "test_get_patient_2@edgar-sante.fr", "password")
//	if err != nil {
//		t.Errorf("Error while creating patient: %v", err)
//	}
//	medical_folder_resp = medical_folder.CreateMedicalInfo(medical_folder.CreateMedicalInfoInput{
//		Name:            "",
//		Firstname:       "",
//		Birthdate:       0,
//		Sex:             "",
//		Weight:          0,
//		Height:          0,
//		PrimaryDoctorID: "",
//		MedicalAntecedents: []medical_folder.CreateMedicalAntecedentInput{{
//			Name:          "",
//			Medicines:     nil,
//			StillRelevant: false,
//		}},
//	}, patient2.CreatePatient.Id)
//
//	if medical_folder_resp.Err != nil {
//		t.Errorf("Unexpected error while creating medical info: %v", medical_folder_resp.Err)
//	}
//	doctor, err := graphql.CreateDoctor(context.Background(), gqlClient, "test_get_patients@edgar-sante.fr", "password", "name", "first", graphql.AddressInput{"", "", "", ""})
//	if err != nil {
//		t.Errorf("Error while creating doctor: %v", err)
//	}
//
//	_, err = graphql.UpdateDoctor(context.Background(), gqlClient, doctor.CreateDoctor.Id, doctor.CreateDoctor.Email, doctor.CreateDoctor.Password, doctor.CreateDoctor.Name, doctor.CreateDoctor.Firstname, []string{}, []string{patient1.CreatePatient.Id, patient2.CreatePatient.Id}, graphql.AddressInput{"", "", "", ""})
//	if err != nil {
//		t.Errorf("Error while updating doctor: %v", err)
//	}
//
//	response := GetPatients(doctor.CreateDoctor.Id)
//
//	if response.Err != nil {
//		t.Errorf("Unexpected error: %v", response.Err)
//	}
//
//	if response.Code != 200 {
//		t.Errorf("Expected code 200, got %d", response.Code)
//	}
//
//	if len(response.PatientsInfo) != 2 {
//		t.Errorf("Expected 2 patients, got %v", response.PatientsInfo)
//	}
//}
