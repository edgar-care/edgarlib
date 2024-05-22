package appointment

import (
	"context"
	"github.com/joho/godotenv"
	"log"
	"testing"

	"github.com/edgar-care/edgarlib/graphql"
)

func TestGetWaitingReview(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}
	gqlClient := graphql.CreateClient()
	doctor, err := graphql.CreateDoctor(context.Background(), gqlClient, "test_doctor_get_wating4review_rdv@edgar-sante.fr", "password", "name", "first", graphql.AddressInput{"", "", "", ""})
	if err != nil {
		t.Errorf("Error while creating doctor: %v", err)
	}
	session, err := graphql.CreateSession(context.Background(), gqlClient, []graphql.SessionDiseasesInput{}, []graphql.SessionSymptomInput{}, 3, 10, 10, "M", []string{}, []string{}, []string{}, "", []graphql.LogsInput{}, []string{})
	if err != nil {
		t.Errorf("Error while creating session: %v", err)
	}
	appointment, err := graphql.CreateRdv(context.Background(), gqlClient, "patientId", doctor.CreateDoctor.Id, 0, 10, "WAITING_FOR_REVIEW", session.CreateSession.Id)
	if err != nil {
		t.Errorf("Error creating appointment: %v", err)
	}

	response := GetWaitingReview(doctor.CreateDoctor.Id)

	if response.Err != nil {
		t.Errorf("Unexpected error: %v", response.Err)
	}

	if response.Code != 200 {
		t.Errorf("Expected code 200, got %d", response.Code)
	}

	if len(response.RdvWithSession) == 0 {
		t.Errorf("Expected non-empty Rdv slice, got empty slice")
	}
	if response.RdvWithSession[0].Rdv.ID != appointment.CreateRdv.Id {
		t.Errorf("Expected first Rdv slice to have it's ID=%s but go ID=%s", response.RdvWithSession[0].Rdv.ID, appointment.CreateRdv.Id)
	}
}

func TestGetWaitingReviewInvalid(t *testing.T) {
	if err := godotenv.Load(".env.test"); err != nil {
		log.Fatalf("Error loading .env.test file: %v", err)
	}

	response := GetWaitingReview("111111111111111111111111")

	if response.Err == nil {
		t.Error("Unexpected null error")
	}

	if response.Code != 400 {
		t.Errorf("Expected code 400, got %d", response.Code)
	}
}
