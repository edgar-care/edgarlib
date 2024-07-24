package appointment

import (
	"context"
	"log"
	"testing"

	"github.com/joho/godotenv"

	"github.com/edgar-care/edgarlib/graphql"
)

func TestDeleteRdv(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}
	gqlClient := graphql.CreateClient()

	patient, err := graphql.CreatePatient(context.Background(), gqlClient, "delete_appointment@edgar-sante.fr", "password")
	if err != nil {
		t.Errorf("Error while creating patient: %v", err)
	}
	appointment, err := graphql.CreateRdv(context.Background(), gqlClient, patient.CreatePatient.Id, "doctorId", 0, 10, "WAITING_FOR_REVIEW", "")
	if err != nil {
		t.Errorf("Error while creating appointment: %v", err)
	}
	appointmentID := appointment.CreateRdv.Id
	patientID := patient.CreatePatient.Id

	_, err = graphql.UpdatePatient(context.Background(), gqlClient, patientID, patient.CreatePatient.Email, patient.CreatePatient.Password, patient.CreatePatient.Medical_info_id, append(patient.CreatePatient.Rendez_vous_ids, "test", appointmentID), patient.CreatePatient.Document_ids, patient.CreatePatient.Treatment_follow_up_ids, patient.CreatePatient.Chat_ids, patient.CreatePatient.Device_connect, patient.CreatePatient.Double_auth_methods_id, patient.CreatePatient.Trust_devices)
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

	newPatient, err := graphql.GetPatientById(context.Background(), gqlClient, patientID)
	if err != nil {
		t.Errorf("Error getting updated patient: %v", err)
	}
	for _, v := range newPatient.GetPatientById.Rendez_vous_ids {
		if v == appointmentID {
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
	gqlClient := graphql.CreateClient()
	appointment, err := graphql.CreateRdv(context.Background(), gqlClient, "test_invalid_Id", "doctorId", 0, 10, "WAITING_FOR_REVIEW", "")
	if err != nil {
		t.Error("Error while creating appointment")
	}

	response := DeleteRdv(appointment.CreateRdv.Id, "invalid")
	if response.Code != 400 || response.Err == nil {
		t.Errorf("Expected error and code 400 but got code %d and err: %s", response.Code, response.Err.Error())
	}
}
