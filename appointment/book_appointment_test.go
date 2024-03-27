package appointment

import (
	"context"
	"github.com/joho/godotenv"
	"log"
	"testing"

	"github.com/edgar-care/edgarlib/graphql"
)

func TestBookAppointment(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}
	gqlClient := graphql.CreateClient()

	patient, err := graphql.CreatePatient(context.Background(), gqlClient, "test_appointment@edgar-sante.fr", "password")
	if err != nil {
		t.Errorf("Error creating patient: %v", err)
	}
	doctor, err := graphql.CreateDoctor(context.Background(), gqlClient, "test_doctor_appointment@edgar-sante.fr", "password", "name", "first", graphql.AddressInput{"", "", "", ""})
	if err != nil {
		t.Errorf("Error while creating doctor: %v", err)
	}
	patientID := patient.CreatePatient.Id
	appointment, err := graphql.CreateRdv(context.Background(), gqlClient, "", doctor.CreateDoctor.Id, 0, 10, "OPENED", "")
	if err != nil {
		t.Errorf("Error creating appointment: %v", err)
	}
	appointmentID := appointment.CreateRdv.Id

	response := BookAppointment(appointmentID, patientID, "session")

	if response.Err != nil {
		t.Errorf("Error booking appointment: %v", response.Err)
	}

	appointmentResult, err := graphql.GetRdvById(context.Background(), gqlClient, appointmentID)
	if err != nil {
		t.Errorf("Error getting appointment by ID: %v", err)
	}

	if appointmentResult.GetRdvById.Id_patient != patientID {
		t.Errorf("Appointment not booked correctly")
	}

	patientResult, err := graphql.GetPatientById(context.Background(), gqlClient, patientID)
	if err != nil {
		t.Errorf("Error getting patient by ID: %v", err)
	}

	found := false
	for _, id := range patientResult.GetPatientById.Rendez_vous_ids {
		if id == appointmentID {
			found = true
			break
		}
	}
	if !found {
		t.Errorf("Appointment ID not added to patient's rendezvous IDs")
	}
}

func TestBookAppointmentEmptyId(t *testing.T) {
	response := BookAppointment("", "", "")
	if response.Code != 400 || response.Err == nil {
		t.Errorf("Expected error and code 400 but got code %d and err: %s", response.Code, response.Err.Error())
	}
}

func TestBookAppointmentEmptySession(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}

	response := BookAppointment("appointment", "patientId", "")
	if response.Code != 400 || response.Err == nil {
		t.Errorf("Expected error and code 400 but got code %d and err: %s", response.Code, response.Err.Error())
	}
}

func TestBookAppointmentInvalidAppointmentId(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}

	response := BookAppointment("invalid", "patientId", "session")
	if response.Code != 400 || response.Err == nil {
		t.Errorf("Expected error and code 400 but got code %d and err: %s", response.Code, response.Err.Error())
	}
}

func TestBookAppointmentAlreadyBooked(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}
	gqlClient := graphql.CreateClient()

	appointment, err := graphql.CreateRdv(context.Background(), gqlClient, "patient", "doctor_id", 0, 10, "OPENED", "session")
	if err != nil {
		t.Errorf("Error creating appointment: %v", err)
	}

	response := BookAppointment(appointment.CreateRdv.Id, "patientId", "session")
	if response.Code != 400 || response.Err == nil {
		t.Errorf("Expected error and code 400 but got code %d and err: %s", response.Code, response.Err.Error())
	}
}

func TestBookAppointmentInvalidPatientId(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}
	gqlClient := graphql.CreateClient()

	appointment, err := graphql.CreateRdv(context.Background(), gqlClient, "", "doctor_id", 0, 10, "OPENED", "session")
	if err != nil {
		t.Errorf("Error creating appointment: %v", err)
	}

	response := BookAppointment(appointment.CreateRdv.Id, "patientId", "session")
	if response.Code != 400 || response.Err == nil {
		t.Errorf("Expected error and code 400 but got code %d and err: %s", response.Code, response.Err.Error())
	}
}

func TestBookAppointmentInvalidDoctorId(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}
	gqlClient := graphql.CreateClient()

	patient, err := graphql.CreatePatient(context.Background(), gqlClient, "test_invalid_doctor_appointment@edgar-sante.fr", "password")
	if err != nil {
		t.Errorf("Error creating patient: %v", err)
	}
	appointment, err := graphql.CreateRdv(context.Background(), gqlClient, "", "invalid", 0, 10, "OPENED", "session")
	if err != nil {
		t.Errorf("Error creating appointment: %v", err)
	}

	response := BookAppointment(appointment.CreateRdv.Id, patient.CreatePatient.Id, "session")
	if response.Code != 400 || response.Err == nil {
		t.Errorf("Expected error and code 400 but got code %d and err: %s", response.Code, response.Err.Error())
	}
}
