package appointment

import (
	"context"
	"github.com/joho/godotenv"
	"log"
	"testing"

	"github.com/edgar-care/edgarlib/graphql"
)

func TestGetDoctorRdv(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}
	gqlClient := graphql.CreateClient()

	doctor, err := graphql.CreateDoctor(context.Background(), gqlClient, "test_doctor_get_rdv@edgar-sante.fr", "password", "name", "first", graphql.AddressInput{"", "", "", ""})
	if err != nil {
		t.Errorf("Error while creating doctor: %v", err)
	}
	appointment, err := graphql.CreateRdv(context.Background(), gqlClient, "patientId", doctor.CreateDoctor.Id, 0, 10, "WAITING_FOR_REVIEW", "")
	if err != nil {
		t.Errorf("Error creating appointment: %v", err)
	}

	response := GetRdvDoctor(doctor.CreateDoctor.Id)

	if response.Err != nil {
		t.Errorf("Unexpected error: %v", response.Err)
	}

	if response.Code != 200 {
		t.Errorf("Expected code 200, got %d", response.Code)
	}

	if len(response.Rdv) == 0 {
		t.Errorf("Expected non-empty Rdv slice, got empty slice")
	}
	if response.Rdv[0].ID != appointment.CreateRdv.Id {
		t.Errorf("Expected first Rdv slice to have it's ID=%s but go ID=%s", response.Rdv[0].ID, appointment.CreateRdv.Id)
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
