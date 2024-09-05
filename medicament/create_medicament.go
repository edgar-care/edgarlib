package medicament

import (
	"errors"

	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
)

type CreateMedicamentInput struct {
	DCI             string   `json:"dci"`
	Name            string   `json:"name"`
	TargetDiseases  []string `json:"target_diseases"`
	TreatedSymptoms []string `json:"treated_symptoms"`
	SideEffects     []string `json:"side_effects"`
	Dosage          int      `json:"dosage"`
	DosageUnit      string   `json:"dosage_unit"`
	Container       string   `json:"container"`
	DosageForm      string   `json:"dosage_form"`
}

type CreateMedicamentResponse struct {
	Medicament model.Medicine
	Code       int
	Err        error
}

func CreateMedicament(input CreateMedicamentInput) CreateMedicamentResponse {
	medicament, err := graphql.CreateMedicine(model.CreateMedicineInput{
		Name:            input.Name,
		Dci:             input.DCI,
		TargetDiseases:  input.TargetDiseases,
		TreatedSymptoms: input.TreatedSymptoms,
		SideEffects:     input.SideEffects,
		Dosage:          input.Dosage,
		DosageUnit:      model.UnitEnum(input.DosageUnit),
		Container:       model.ContainerEnum(input.Container),
		DosageForm:      model.FormEnum(input.DosageForm),
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
