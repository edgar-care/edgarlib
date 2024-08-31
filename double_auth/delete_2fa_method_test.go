package double_auth

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
	"github.com/joho/godotenv"
	"log"
	"testing"
)

func TestRemoveDoubleAuthMethod_Success(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}

	patient, err := graphql.CreatePatient(model.CreatePatientInput{
		Email:    "test_patient_create_remove_double_auth_test@edgar-sante.fr",
		Password: "password",
		Status:   true,
	})
	if err != nil {
		t.Errorf("Error while creating patient: %v", err)
	}

	input := CreateDoubleAuthInput{Methods: "EMAIL"}
	auth := CreateDoubleAuthEmail(input, patient.ID)

	err = RemoveDoubleAuthMethod("EMAIL", patient.ID)

	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	patientx, _ := graphql.GetPatientById(patient.ID)
	spew.Dump(patientx)
	if patientx.DoubleAuthMethodsID != nil {
		doubleAuth, _ := graphql.GetDoubleAuthById(*patientx.DoubleAuthMethodsID)
		for _, method := range doubleAuth.Methods {
			if method == auth.DoubleAuth.ID {
				t.Errorf("Expected method %s to be removed, but it still exists", auth.DoubleAuth.ID)
			}
		}
	}
}

func TestRemoveDoubleAuthMethod_DoctorSuccess(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}

	doctor, err := graphql.CreateDoctor(model.CreateDoctorInput{
		Email:    "test_doctor_create_remove_double_auth_test@edgar-sante.fr",
		Password: "password",
		Status:   true,
		Address: &model.AddressInput{
			Street:  "",
			ZipCode: "",
			Country: "",
			City:    "",
		},
	})
	if err != nil {
		t.Errorf("Error while creating doctor: %v", err)
	}

	input := CreateDoubleAuthInput{Methods: "EMAIL"}
	auth := CreateDoubleAuthEmail(input, doctor.ID)

	err = RemoveDoubleAuthMethod("EMAIL", doctor.ID)

	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	doctorx, _ := graphql.GetDoctorById(doctor.ID)
	if doctorx.DoubleAuthMethodsID != nil {
		doubleAuth, _ := graphql.GetDoubleAuthById(*doctorx.DoubleAuthMethodsID)
		for _, method := range doubleAuth.Methods {
			if method == auth.DoubleAuth.ID {
				t.Errorf("Expected method %s to be removed, but it still exists", auth.DoubleAuth.ID)
			}
		}
	}
}

func TestRemoveDoubleAuthMethod_PatientNotFound(t *testing.T) {
	methodToRemove := "sms"
	invalidPatientId := "invalid_patient_id"

	err := RemoveDoubleAuthMethod(methodToRemove, invalidPatientId)

	if err == nil {
		t.Errorf("Expected an error, got none")
	}
}

func TestRemoveDoubleAuthMethod_DoctorNotFound(t *testing.T) {
	methodToRemove := "sms"
	invalidDoctorId := "invalid_doctor_id"

	err := RemoveDoubleAuthMethod(methodToRemove, invalidDoctorId)

	if err == nil {
		t.Errorf("Expected an error, got none")
	}
}

func TestRemoveDoubleAuthMethod_NoDoubleAuthMethodsID_Patient(t *testing.T) {
	methodToRemove := "sms"
	patientId := "patient_without_double_auth_methods_id"

	err := RemoveDoubleAuthMethod(methodToRemove, patientId)

	if err == nil {
		t.Errorf("Expected an error, got none")
	}
}

func TestRemoveDoubleAuthMethod_NoDoubleAuthMethodsID_Doctor(t *testing.T) {
	methodToRemove := "sms"
	doctorId := "doctor_without_double_auth_methods_id"

	err := RemoveDoubleAuthMethod(methodToRemove, doctorId)

	if err == nil {
		t.Errorf("Expected an error, got none")
	}
}

func TestRemoveDoubleAuthMethod_MethodNotFound_Patient(t *testing.T) {
	methodToRemove := "non_existing_method"
	patientId := "valid_patient_id"

	err := RemoveDoubleAuthMethod(methodToRemove, patientId)

	if err == nil {
		t.Errorf("Expected an error, got none")
	}
}

func TestRemoveDoubleAuthMethod_MethodNotFound_Doctor(t *testing.T) {
	methodToRemove := "non_existing_method"
	doctorId := "valid_doctor_id"

	err := RemoveDoubleAuthMethod(methodToRemove, doctorId)

	if err == nil {
		t.Errorf("Expected an error, got none")
	}
}

func TestRemoveDoubleAuthMethod_RemoveLastMethod_Patient(t *testing.T) {
	patient, err := graphql.CreatePatient(model.CreatePatientInput{
		Email:    "test_patient_create_remove_double_auth_last@edgar-sante.fr",
		Password: "password",
		Status:   true,
	})
	if err != nil {
		t.Errorf("Error while creating patient: %v", err)
	}

	input := CreateDoubleAuthInput{Methods: "EMAIL"}
	_ = CreateDoubleAuthEmail(input, patient.ID)

	err = RemoveDoubleAuthMethod("EMAIL", patient.ID)

	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	updatedPatient, _ := graphql.GetPatientById(patient.ID)
	if updatedPatient.DoubleAuthMethodsID != nil && *updatedPatient.DoubleAuthMethodsID != "" {
		t.Errorf("Expected DoubleAuthMethodsID to be removed, but it still exists: %v", *updatedPatient.DoubleAuthMethodsID)
	}
}

func TestRemoveDoubleAuthMethod_RemoveLastMethod_Doctor(t *testing.T) {
	doctor, err := graphql.CreateDoctor(model.CreateDoctorInput{
		Email:    "test_doctor_create_remove_double_auth_last@edgar-sante.fr",
		Password: "password",
		Status:   true,
		Address: &model.AddressInput{
			Street:  "",
			ZipCode: "",
			Country: "",
			City:    "",
		},
	})
	if err != nil {
		t.Errorf("Error while creating doctor: %v", err)
	}

	input := CreateDoubleAuthInput{Methods: "EMAIL"}
	_ = CreateDoubleAuthEmail(input, doctor.ID)

	err = RemoveDoubleAuthMethod("EMAIL", doctor.ID)

	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	updatedDoctor, _ := graphql.GetDoctorById(doctor.ID)
	if updatedDoctor.DoubleAuthMethodsID != nil && *updatedDoctor.DoubleAuthMethodsID != "" {
		t.Errorf("Expected DoubleAuthMethodsID to be removed, but it still exists: %v", *updatedDoctor.DoubleAuthMethodsID)
	}
}
