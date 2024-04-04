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
	Medicament model.Medicine
	Code       int
	Err        error
}

func CreateMedicament(input CreateMedicamentInput) CreateMedicamentResponse {
	gqlClient := graphql.CreateClient()

	medicament, err := graphql.CreateMedicine(context.Background(), gqlClient, input.Name, input.Unit, input.TargetDiseases, input.TreatedSymptoms, input.SideEffects)
	if err != nil {
		return CreateMedicamentResponse{Medicament: model.Medicine{}, Code: 400, Err: errors.New("unable  (check if you share all information)")}
	}

	return CreateMedicamentResponse{
		Medicament: model.Medicine{
			ID:              medicament.CreateMedicine.Id,
			Name:            medicament.CreateMedicine.Name,
			Unit:            model.MedicineUnit(medicament.CreateMedicine.Unit),
			TargetDiseases:  medicament.CreateMedicine.Target_diseases,
			TreatedSymptoms: medicament.CreateMedicine.Treated_symptoms,
			SideEffects:     medicament.CreateMedicine.Side_effects,
		},
		Code: 201,
		Err:  nil,
	}
}
