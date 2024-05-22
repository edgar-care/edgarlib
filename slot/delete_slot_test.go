package slot

import (
	"context"
	"log"
	"testing"

	"github.com/edgar-care/edgarlib/graphql"
	"github.com/joho/godotenv"
)

func TestDeleteSlot(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}
	gqlClient := graphql.CreateClient()

	doctor, err := graphql.CreateDoctor(context.Background(), gqlClient, "test_doctor_delete_slot_up@edgar-sante.fr", "password", "test_doctor", "first", graphql.AddressInput{"", "", "", ""})
	if err != nil {
		t.Errorf("Error while creating patient: %v", err)
	}

	input := CreateSlotInput{
		StartDate: 14567,
		EndDate:   14578,
	}

	slot, err := graphql.CreateRdv(context.Background(), gqlClient, "", doctor.CreateDoctor.Id, input.StartDate, input.EndDate, "OPENED", "")
	if err != nil {
		t.Errorf("Error while creating follow up treatment: %v", err)
	}
	slot_upID := slot.CreateRdv.Id
	patientID := doctor.CreateDoctor.Id

	_, err = graphql.UpdateDoctor(context.Background(), gqlClient, patientID, doctor.CreateDoctor.Email, doctor.CreateDoctor.Password, doctor.CreateDoctor.Name, doctor.CreateDoctor.Firstname, append(doctor.CreateDoctor.Rendez_vous_ids, slot.CreateRdv.Id), doctor.CreateDoctor.Patient_ids, graphql.AddressInput{Street: doctor.CreateDoctor.Address.Street, Zip_code: doctor.CreateDoctor.Address.Zip_code, Country: doctor.CreateDoctor.Address.Country}, doctor.CreateDoctor.Chat_ids)
	if err != nil {
		t.Errorf("Error while updating patient: %v", err)
	}
	response := DeleteSlot(slot_upID, patientID)

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
	for _, v := range newPatient.GetPatientById.Treatment_follow_up_ids {
		if v == slot_upID {
			t.Error("slot's id has not been deleted on patient")
		}
	}
}

func TestDelete_follow_upEmptyId(t *testing.T) {
	response := DeleteSlot("", "test")
	if response.Code != 400 || response.Err == nil {
		t.Errorf("Expected error and code 400 but got code %d and err: %s", response.Code, response.Err.Error())
	}
}

func TestDelete_follow_upInvalidId(t *testing.T) {
	response := DeleteSlot("invalid", "test")
	if response.Code != 400 || response.Err == nil {
		t.Errorf("Expected error and code 400 but got code %d and err: %s", response.Code, response.Err.Error())
	}
}
