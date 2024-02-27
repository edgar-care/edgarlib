package medicament

import (
	"context"
	"errors"

	"github.com/edgar-care/edgarlib/graphql"
	"github.com/edgar-care/edgarlib/graphql/server/model"
)

type GetMedicamentByIdResponse struct {
	Medicament model.Medicament
	Code       int
	Err        error
}

type GetMedicamentsResponse struct {
	Medicaments []model.Medicament
	Code        int
	Err         error
}

func GetMedicamentById(id string) GetMedicamentByIdResponse {
	gqlClient := graphql.CreateClient()
	var res model.Medicament

	medicament, err := graphql.GetMedicamentByID(context.Background(), gqlClient, id)
	if err != nil {
		return GetMedicamentByIdResponse{model.Medicament{}, 400, errors.New("id does not correspond to a medicament")}
	}
	res = model.Medicament{
		ID:              medicament.GetMedicamentByID.Id,
		Name:            medicament.GetMedicamentByID.Name,
		Unit:            model.Unit(medicament.GetMedicamentByID.Unit),
		TargetDiseases:  medicament.GetMedicamentByID.Target_diseases,
		TreatedSymptoms: medicament.GetMedicamentByID.Treated_symptoms,
		SideEffects:     medicament.GetMedicamentByID.Side_effects,
	}
	return GetMedicamentByIdResponse{res, 200, nil}
}

func GetMedicaments() GetMedicamentsResponse {
	gqlClient := graphql.CreateClient()
	var res []model.Medicament

	medicaments, err := graphql.GetMedicaments(context.Background(), gqlClient)
	if err != nil {
		return GetMedicamentsResponse{[]model.Medicament{}, 400, errors.New("invalid input: " + err.Error())}
	}

	for _, medicament := range medicaments.GetMedicaments {
		res = append(res, model.Medicament{
			ID:              medicament.Id,
			Name:            medicament.Name,
			Unit:            model.Unit(medicament.Unit),
			TargetDiseases:  medicament.Target_diseases,
			TreatedSymptoms: medicament.Treated_symptoms,
			SideEffects:     medicament.Side_effects,
		})
	}
	return GetMedicamentsResponse{res, 200, nil}
}
