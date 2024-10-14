package appointment

import (
	"github.com/edgar-care/edgarlib/v2/dashboard"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
	"github.com/edgar-care/edgarlib/v2/medical_folder"
	"github.com/joho/godotenv"
	"log"
	"testing"

	"github.com/edgar-care/edgarlib/v2/graphql"
)

func TestBookAppointment(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}

	patient, err := graphql.CreatePatient(model.CreatePatientInput{Email: "test_appointment@edgar-sante.fr", Password: "password"})
	if err != nil {
		t.Errorf("Error creating patient: %v", err)
	}
	doctor, err := graphql.CreateDoctor(model.CreateDoctorInput{Email: "test_doctor_appointment@edgar-sante.fr", Password: "password", Name: "name", Firstname: "first", Address: &model.AddressInput{"", "", "", ""}})
	if err != nil {
		t.Errorf("Error while creating doctor: %v", err)
	}
	patientID := patient.ID
	appointment, err := graphql.CreateRdv(model.CreateRdvInput{
		IDPatient:         "",
		DoctorID:          doctor.ID,
		StartDate:         0,
		EndDate:           10,
		AppointmentStatus: "OPENED",
		SessionID:         "",
	})
	if err != nil {
		t.Errorf("Error creating appointment: %v", err)
	}
	appointmentID := appointment.ID

	response := BookAppointment(appointmentID, patientID, "session")

	if response.Err != nil {
		t.Errorf("Error booking appointment: %v", response.Err)
	}

	appointmentResult, err := graphql.GetRdvById(appointmentID)
	if err != nil {
		t.Errorf("Error getting appointment by ID: %v", err)
	}

	if appointmentResult.IDPatient != patientID {
		t.Errorf("Appointment not booked correctly")
	}

	patientResult, err := graphql.GetPatientById(patientID)
	if err != nil {
		t.Errorf("Error getting patient by ID: %v", err)
	}

	found := false
	for _, id := range patientResult.RendezVousIds {
		if *id == appointmentID {
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

	appointment, err := graphql.CreateRdv(model.CreateRdvInput{
		IDPatient:         "patient",
		DoctorID:          "doctor_id",
		StartDate:         0,
		EndDate:           10,
		AppointmentStatus: "OPENED",
		SessionID:         "session",
	})
	if err != nil {
		t.Errorf("Error creating appointment: %v", err)
	}

	response := BookAppointment(appointment.ID, "patientId", "session")
	if response.Code != 400 || response.Err == nil {
		t.Errorf("Expected error and code 400 but got code %d and err: %s", response.Code, response.Err.Error())
	}
}

func TestBookAppointmentInvalidPatientId(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}

	appointment, err := graphql.CreateRdv(model.CreateRdvInput{"", "doctor_id", 0, 10, "OPENED", "session"})
	if err != nil {
		t.Errorf("Error creating appointment: %v", err)
	}

	response := BookAppointment(appointment.ID, "patientId", "session")
	if response.Code != 400 || response.Err == nil {
		t.Errorf("Expected error and code 400 but got code %d and err: %s", response.Code, response.Err.Error())
	}
}

func TestBookAppointmentInvalidDoctorId(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}

	patient, err := graphql.CreatePatient(model.CreatePatientInput{Email: "test_invalid_doctor_appointment@edgar-sante.fr", Password: "password"})
	if err != nil {
		t.Errorf("Error creating patient: %v", err)
	}
	appointment, err := graphql.CreateRdv(model.CreateRdvInput{"", "invalid", 0, 10, "OPENED", "session"})
	if err != nil {
		t.Errorf("Error creating appointment: %v", err)
	}

	response := BookAppointment(appointment.ID, patient.ID, "session")
	if response.Code != 400 || response.Err == nil {
		t.Errorf("Expected error and code 400 but got code %d and err: %s", response.Code, response.Err.Error())
	}
}

func TestBookAppointmentCheckDoctor(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}

	patient, err := graphql.CreatePatient(model.CreatePatientInput{Email: "test_appointment_check_dupli@edgar-sante.fr", Password: "password"})
	if err != nil {
		t.Errorf("Error creating patient: %v", err)
	}

	input := medical_folder.CreateMedicalInfoInput{
		Name:            "test",
		Firstname:       "test",
		Birthdate:       0,
		Sex:             "M",
		Weight:          0,
		Height:          0,
		PrimaryDoctorID: "",
		MedicalAntecedents: []medical_folder.CreateMedicalAntecedentInput{{
			Name: "test",
			Medicines: []medical_folder.CreateMedicineInput{medical_folder.CreateMedicineInput{
				MedicineID: "test",
				Period:     []string{"NOON"},
				Day:        []string{"MONDAY"},
				Quantity:   2,
				StartDate:  1234,
				EndDate:    1234,
			}},
			StillRelevant: false,
		},
		},
		FamilyMembersMedInfoId: []string{"test"},
	}

	_ = medical_folder.CreateMedicalInfo(input, patient.ID)

	doctor, err := graphql.CreateDoctor(model.CreateDoctorInput{Email: "test_doctor_appointment_check_dupli@edgar-sante.fr", Password: "password", Name: "name", Firstname: "first", Address: &model.AddressInput{"", "", "", ""}})
	if err != nil {
		t.Errorf("Error while creating doctor: %v", err)
	}
	patientID := patient.ID
	appointment, err := graphql.CreateRdv(model.CreateRdvInput{
		IDPatient:         "",
		DoctorID:          doctor.ID,
		StartDate:         0,
		EndDate:           10,
		AppointmentStatus: "OPENED",
		SessionID:         "",
	})
	if err != nil {
		t.Errorf("Error creating appointment: %v", err)
	}
	appointmentID := appointment.ID

	response := BookAppointment(appointmentID, patientID, "session")

	appointment2, err := graphql.CreateRdv(model.CreateRdvInput{
		IDPatient:         "",
		DoctorID:          doctor.ID,
		StartDate:         45,
		EndDate:           65,
		AppointmentStatus: "OPENED",
		SessionID:         "",
	})
	if err != nil {
		t.Errorf("Error creating appointment: %v", err)
	}
	appointment2ID := appointment2.ID

	_ = BookAppointment(appointment2ID, patientID, "session")

	if response.Err != nil {
		t.Errorf("Error booking appointment: %v", response.Err)
	}

	appointmentResult, err := graphql.GetRdvById(appointmentID)
	if err != nil {
		t.Errorf("Error getting appointment by ID: %v", err)
	}

	if appointmentResult.IDPatient != patientID {
		t.Errorf("Appointment not booked correctly")
	}

	_, err = graphql.GetDoctorById(doctor.ID)
	if err != nil {
		t.Errorf("Error getting doctor by ID: %v", err)
	}

	_ = dashboard.GetPatients(doctor.ID)

}
