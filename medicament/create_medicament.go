package medicament

import (
	"errors"

	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
)

type CreateMedicamentInput struct {
	Name            string   `json:"name"`
	Unit            string   `json:"unit"`
	TargetDiseases  []string `json:"target_diseases"`
	TreatedSymptoms []string `json:"treated_symptoms"`
	SideEffects     []string `json:"side_effects"`
	Type            string   `json:"type"`
	Content         string   `json:"content"`
	Quantity        int      `json:"quantity"`
}

type CreateMedicamentResponse struct {
	Medicament model.Medicine
	Code       int
	Err        error
}

func CreateMedicament(input CreateMedicamentInput) CreateMedicamentResponse {
	medicament, err := graphql.CreateMedicine(model.CreateMedicineInput{
		Name:            input.Name,
		Unit:            &input.Unit,
		TargetDiseases:  input.TargetDiseases,
		TreatedSymptoms: input.TreatedSymptoms,
		SideEffects:     input.SideEffects,
		Type:            input.Type,
		Content:         input.Content,
		Quantity:        input.Quantity,
	})
	if err != nil {
		return CreateMedicamentResponse{Medicament: model.Medicine{}, Code: 400, Err: errors.New("unable  (check if you share all information)")}
	}

	return CreateMedicamentResponse{
		Medicament: medicament,
		Code:       201,
		Err:        nil,
	}
}
