package appointment

import (
	"github.com/edgar-care/edgarlib/graphql"
	"github.com/edgar-care/edgarlib/graphql/model"
	"github.com/joho/godotenv"
	"log"
	"testing"
)

func TestCancelRdv(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}

	appointment, err := graphql.CreateRdv(model.CreateRdvInput{
		IDPatient:         "patient",
		DoctorID:          "doctor",
		StartDate:         0,
		EndDate:           10,
		AppointmentStatus: "WAITING_FOR_REVIEW",
		SessionID:         "",
	})
	if err != nil {
		t.Error("Error while creating appointment")
	}
	if appointment.ID == "" {
		t.Error("Exepected appointment id to not be null")
	}
	reason := "test"
	cancelResponse := CancelRdv(appointment.ID, reason)

	if cancelResponse.Code != 200 {
		t.Errorf("Expected code 200 for successful cancellation, got %d", cancelResponse.Code)
	}
	if cancelResponse.Err != nil {
		t.Errorf("Expected no error, got %v", cancelResponse.Err)
	}

	newAppointment, err := graphql.GetRdvById(appointment.ID)
	if err != nil {
		t.Error("Error while retrieving appointment")
	}
	if newAppointment.AppointmentStatus != model.AppointmentStatusCanceled {
		t.Errorf("Expected empty statusCanceled but got %s", newAppointment.AppointmentStatus)
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
