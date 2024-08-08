package appointment

import (
	"github.com/edgar-care/edgarlib/graphql/model"
	"github.com/joho/godotenv"
	"log"
	"testing"

	"github.com/edgar-care/edgarlib/graphql"
)

func TestGetDoctorRdv(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}

	doctor, err := graphql.CreateDoctor(model.CreateDoctorInput{
		Email:     "test_doctor_get_rdv@edgar-sante.fr",
		Password:  "password",
		Name:      "name",
		Firstname: "first",
		Address: &model.AddressInput{
			Street:  "",
			ZipCode: "",
			Country: "",
			City:    "",
		},
		Status: true,
	})
	if err != nil {
		t.Errorf("Error while creating doctor: %v", err)
	}
	appointment, err := graphql.CreateRdv(model.CreateRdvInput{
		IDPatient:         "patientId",
		DoctorID:          doctor.ID,
		StartDate:         0,
		EndDate:           10,
		AppointmentStatus: "WAITING_FOR_REVIEW",
		SessionID:         "",
	})
	if err != nil {
		t.Errorf("Error creating appointment: %v", err)
	}

	response := GetRdvDoctor(doctor.ID)

	if response.Err != nil {
		t.Errorf("Unexpected error: %v", response.Err)
	}

	if response.Code != 200 {
		t.Errorf("Expected code 200, got %d", response.Code)
	}

	if len(response.Rdv) == 0 {
		t.Errorf("Expected non-empty Rdv slice, got empty slice")
	}
	if response.Rdv[0].ID != appointment.ID {
		t.Errorf("Expected first Rdv slice to have it's ID=%s but go ID=%s", response.Rdv[0].ID, appointment.ID)
	}
}

func TestGetRdvDoctorInvalid(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}

	response := GetRdvDoctor("111111111111111111111111")

	if response.Err == nil {
		t.Error("Unexpected null error")
	}

	if response.Code != 400 {
		t.Errorf("Expected code 400, got %d", response.Code)
	}
}
