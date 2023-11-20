package graphql

import (
	"github.com/jinzhu/copier"
)

/********** Types ***********/

type UpdatePatientOnboarding struct {
	Id                 string `json:"id"`
	OnboardingInfoID   string `json:"onboarding_info_id"`
	OnboardingHealthID string `json:"onboarding_health_id"`
}

type UpdatePatientOnboardingInput struct {
	Id                 string `json:"id"`
	OnboardingInfoID   string `json:"onboarding_info_id"`
	OnboardingHealthID string `json:"onboarding_health_id"`
}

type UpdatePatientOnboardingOutput struct {
	Id                 *string `json:"id"`
	OnboardingInfoID   *string `json:"onboarding_info_id"`
	OnboardingHealthID *string `json:"onboarding_health_id"`
}

/**************** GraphQL types *****************/

type updateOnboardingResponse struct {
	Content UpdatePatientOnboardingOutput `json:"updatePatient"`
}

type getUpdatePatientOnboardingByIdResponse struct {
	Content UpdatePatientOnboardingOutput `json:"getUpdatePatientOnboardingById"`
}

/*************** Implementations *****************/

func GetUpdatePatientOnboardingById(id string) (UpdatePatientOnboarding, error) {
	var patient getUpdatePatientOnboardingByIdResponse
	var resp UpdatePatientOnboarding
	query := `query getPatientById($id: String!) {
                getPatientById(id: $id) {
                    id,
					onboarding_info_id,
					onboarding_health_id
                }
            }`

	err := Query(query, map[string]interface{}{
		"id": id,
	}, &patient)
	_ = copier.Copy(&resp, &patient.Content)
	return resp, err
}
