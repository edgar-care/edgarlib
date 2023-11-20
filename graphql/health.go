package graphql

import (
	"github.com/jinzhu/copier"
)

/********** Types ***********/

type Health struct {
	Id                    string    `json:"id"`
	PatientsAllergies     *[]string `json:"patients_allergies,omitempty"`
	PatientsIllness       *[]string `json:"patients_illness,omitempty"`
	PatientsTreatments    *[]string `json:"patients_treatments,omitempty"`
	PatientsPrimaryDoctor string    `json:"patients_primary_doctor,omitempty"`
}

type HealthInput struct {
	PatientsAllergies     *[]string `json:"patients_allergies,omitempty"`
	PatientsIllness       *[]string `json:"patients_illness,omitempty"`
	PatientsTreatments    *[]string `json:"patients_treatments,omitempty"`
	PatientsPrimaryDoctor string    `json:"patients_primary_doctor,omitempty"`
}

type HealthOutput struct {
	Id                    string    `json:"id"`
	PatientsAllergies     *[]string `json:"patients_allergies,omitempty"`
	PatientsIllness       *[]string `json:"patients_illness,omitempty"`
	PatientsTreatments    *[]string `json:"patients_treatments,omitempty"`
	PatientsPrimaryDoctor *string   `json:"patients_primary_doctor,omitempty"`
}

/**************** GraphQL types *****************/

type createHealthResponse struct {
	Content HealthOutput `json:"createHealth"`
}

type getHealthByIdResponse struct {
	Content HealthOutput `json:"getHealthById"`
}

type updateHealthResponse struct {
	Content HealthOutput `json:"updateHealth"`
}

/*************** Implementations *****************/

func CreateHealth(newHealth HealthInput) (Health, error) {
	var health createHealthResponse
	var resp Health
	query := `mutation createHealth($patients_allergies: [String!], $patients_illness: [String!], $patients_treatments: [String!], $patients_primary_doctor: String!) {
            createHealth(patients_allergies:$patients_allergies, patients_illness:$patients_illness, patients_treatments:$patients_treatments, patients_primary_doctor:$patients_primary_doctor) {
                    id,
					patients_allergies,
					patients_illness,
					patients_primary_doctor,
					patients_treatments
                }
            }`
	err := Query(query, map[string]interface{}{
		"patients_allergies":      newHealth.PatientsAllergies,
		"patients_illness":        newHealth.PatientsIllness,
		"patients_treatments":     newHealth.PatientsTreatments,
		"patients_primary_doctor": newHealth.PatientsPrimaryDoctor,
	}, &health)

	_ = copier.Copy(&resp, &health.Content)
	return resp, err
}

func GetHealthById(id string) (Health, error) {
	var health getHealthByIdResponse
	var resp Health
	query := `query getHealthById($id: String!) {
                getHealthById(id: $id) {
                    id,
					patients_allergies,
					patients_illness,
					patients_primary_doctor,
					patients_treatments
                }
            }`
	err := Query(query, map[string]interface{}{
		"id": id,
	}, &health)
	_ = copier.Copy(&resp, &health.Content)
	return resp, err
}

func AddOnboardingHealthID(updatePatient UpdatePatientOnboardingInput) (UpdatePatientOnboarding, error) {
	var patient updateOnboardingResponse
	var resp UpdatePatientOnboarding
	query := `mutation updatePatient($id: String!, $onboarding_health_id: String) {
		updatePatient(id:$id, onboarding_health_id:$onboarding_health_id) {
                    id,
					onboarding_health_id
                }
            }`
	err := Query(query, map[string]interface{}{
		"id":                   updatePatient.Id,
		"onboarding_health_id": updatePatient.OnboardingHealthID,
	}, &patient)
	_ = copier.Copy(&resp, &patient.Content)
	return resp, err
}
