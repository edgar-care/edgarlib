package appointment

import (
	"github.com/edgar-care/edgarlib/v2/graphql/model"
	"log"
	"testing"

	"github.com/joho/godotenv"

	"github.com/edgar-care/edgarlib/v2/graphql"
)

func TestDeleteRdv(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}

	patient, err := graphql.CreatePatient(model.CreatePatientInput{
		Email:    "delete_appointment@edgar-sante.fr",
		Password: "password",
	})
	if err != nil {
		t.Errorf("Error while creating patient: %v", err)
	}
	doctor, err := graphql.CreateDoctor(model.CreateDoctorInput{
		Email:     "test_delete_doctor_appointment_success@edgar-sante.fr",
		Password:  "password",
		Name:      "name",
		Firstname: "first",
		Address: &model.AddressInput{
			Street:  "",
			ZipCode: "",
			Country: "",
			City:    "",
		},
		Status: false,
	})
	appointment, err := graphql.CreateRdv(model.CreateRdvInput{patient.ID, doctor.ID, 0, 10, "WAITING_FOR_REVIEW", ""})
	if err != nil {
		t.Errorf("Error while creating appointment: %v", err)
	}
	appointmentID := appointment.ID
	patientID := patient.ID
	test := "test"
	_, err = graphql.UpdatePatient(patientID, model.UpdatePatientInput{
		RendezVousIds: append(patient.RendezVousIds, &test, &appointmentID),
	})
	if err != nil {
		t.Errorf("Error while updating patient: %v", err)
	}
	response := DeleteRdv(appointmentID, patientID)

	if response.Err != nil {
		t.Errorf("Unexpected error: %v", response.Err)
	}

	if response.Code != 200 {
		t.Errorf("Expected code 200, got %d", response.Code)
	}

	newPatient, err := graphql.GetPatientById(patientID)
	if err != nil {
		t.Errorf("Error getting updated patient: %v", err)
	}
	for _, v := range newPatient.RendezVousIds {
		if v == &appointmentID {
			t.Error("Appointment's id has not been deleted on patient")
		}
	}
}

func TestDeleteRdvEmptyId(t *testing.T) {
	response := DeleteRdv("", "test")
	if response.Code != 400 || response.Err == nil {
		t.Errorf("Expected error and code 400 but got code %d and err: %s", response.Code, response.Err.Error())
	}
}

func TestDeleteRdvInvalidId(t *testing.T) {
	response := DeleteRdv("invalid", "test")
	if response.Code != 400 || response.Err == nil {
		t.Errorf("Expected error and code 400 but got code %d and err: %s", response.Code, response.Err.Error())
	}
}

func TestDeleteRdvInvalidPatient(t *testing.T) {
	appointment, err := graphql.CreateRdv(model.CreateRdvInput{"test_invalid_Id", "doctorId", 0, 10, "WAITING_FOR_REVIEW", ""})
	if err != nil {
		t.Error("Error while creating appointment")
	}

	response := DeleteRdv(appointment.ID, "invalid")
	if response.Code != 400 || response.Err == nil {
		t.Errorf("Expected error and code 400 but got code %d and err: %s", response.Code, response.Err.Error())
	}
}
