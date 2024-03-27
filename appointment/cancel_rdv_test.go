package appointment

import (
	"context"
	"github.com/edgar-care/edgarlib/graphql"
	"github.com/joho/godotenv"
	"log"
	"testing"
)

func TestCancelRdv(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}
	gqlClient := graphql.CreateClient()

	appointment, err := graphql.CreateRdv(context.Background(), gqlClient, "patient", "doctor", 0, 10, "WAITING_FOR_REVIEW", "")
	if err != nil {
		t.Error("Error while creating appointment")
	}
	if appointment.CreateRdv.Id == "" {
		t.Error("Exepected appointment id to not be null")
	}
	reason := "test"
	cancelResponse := CancelRdv(appointment.CreateRdv.Id, reason)

	if cancelResponse.Code != 200 {
		t.Errorf("Expected code 200 for successful cancellation, got %d", cancelResponse.Code)
	}
	if cancelResponse.Err != nil {
		t.Errorf("Expected no error, got %v", cancelResponse.Err)
	}

	newAppointment, err := graphql.GetRdvById(context.Background(), gqlClient, appointment.CreateRdv.Id)
	if err != nil {
		t.Error("Error while retrieving appointment")
	}
	if newAppointment.GetRdvById.Appointment_status != graphql.AppointmentStatusCanceled {
		t.Errorf("Expected empty statusCanceled but got %s", newAppointment.GetRdvById.Appointment_status)
	}
	if cancelResponse.Reason != reason {
		t.Errorf("Expected reason: %s. But got: %s", reason, cancelResponse.Reason)
	}
}

func TestCancelRdvEmptyId(t *testing.T) {
	response := CancelRdv("", "test")
	if response.Code != 400 || response.Err == nil {
		t.Errorf("Expected error and code 400 but got code %d and err: %s", response.Code, response.Err.Error())
	}
}

func TestCancelRdvInvalidId(t *testing.T) {
	response := CancelRdv("invalid", "test")
	if response.Code != 400 || response.Err == nil {
		t.Errorf("Expected error and code 400 but got code %d and err: %s", response.Code, response.Err.Error())
	}
}
