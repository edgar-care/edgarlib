package medicament

import (
	"context"
	"errors"

	"github.com/edgar-care/edgarlib/graphql"
	"github.com/edgar-care/edgarlib/graphql/server/model"
)

type CreateMedicamentInput struct {
	Name            string   `json:"name"`
	Unit            string   `json:"unit"`
	TargetDiseases  []string `json:"target_diseases"`
	TreatedSymptoms []string `json:"treated_symptoms"`
	SideEffects     []string `json:"side_effects"`
}

type CreateMedicamentResponse struct {
	Medicament model.Medicament
	Code       int
	Err        error
}

func CreateMedicament(input CreateMedicamentInput) CreateMedicamentResponse {
	gqlClient := graphql.CreateClient()

	medicament, err := graphql.CreateMedicament(context.Background(), gqlClient, input.Name, input.Unit, input.TargetDiseases, input.TreatedSymptoms, input.SideEffects)
	if err != nil {
		return CreateMedicamentResponse{Medicament: model.Medicament{}, Code: 400, Err: errors.New("unable  (check if you share all information)")}
	}

	return CreateMedicamentResponse{
		Medicament: model.Medicament{
			ID:              medicament.CreateMedicament.Id,
			Name:            medicament.CreateMedicament.Name,
			Unit:            model.Unit(medicament.CreateMedicament.Unit),
			TargetDiseases:  medicament.CreateMedicament.Target_diseases,
			TreatedSymptoms: medicament.CreateMedicament.Treated_symptoms,
			SideEffects:     medicament.CreateMedicament.Side_effects,
		},
		Code: 201,
		Err:  nil,
	}
}
