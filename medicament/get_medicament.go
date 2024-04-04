package medicament

import (
	"context"
	"errors"

	"github.com/edgar-care/edgarlib/graphql"
	"github.com/edgar-care/edgarlib/graphql/server/model"
)

type GetMedicamentByIdResponse struct {
	Medicine model.Medicine
	Code     int
	Err      error
}

type GetMedicamentsResponse struct {
	Medicines []model.Medicine
	Code      int
	Err       error
}

func GetMedicamentById(id string) GetMedicamentByIdResponse {
	gqlClient := graphql.CreateClient()
	var res model.Medicine

	medicament, err := graphql.GetMedicineByID(context.Background(), gqlClient, id)
	if err != nil {
		return GetMedicamentByIdResponse{model.Medicine{}, 400, errors.New("id does not correspond to a medicament")}
	}
	res = model.Medicine{
		ID:              medicament.GetMedicineByID.Id,
		Name:            medicament.GetMedicineByID.Name,
		Unit:            model.MedicineUnit(medicament.GetMedicineByID.Unit),
		TargetDiseases:  medicament.GetMedicineByID.Target_diseases,
		TreatedSymptoms: medicament.GetMedicineByID.Treated_symptoms,
		SideEffects:     medicament.GetMedicineByID.Side_effects,
	}
	return GetMedicamentByIdResponse{res, 200, nil}
}

func GetMedicaments() GetMedicamentsResponse {
	gqlClient := graphql.CreateClient()
	var res []model.Medicine

	medicaments, err := graphql.GetMedicines(context.Background(), gqlClient)
	if err != nil {
		return GetMedicamentsResponse{[]model.Medicine{}, 400, errors.New("invalid input: " + err.Error())}
	}

	for _, medicament := range medicaments.GetMedicines {
		res = append(res, model.Medicine{
			ID:              medicament.Id,
			Name:            medicament.Name,
			Unit:            model.MedicineUnit(medicament.Unit),
			TargetDiseases:  medicament.Target_diseases,
			TreatedSymptoms: medicament.Treated_symptoms,
			SideEffects:     medicament.Side_effects,
		})
	}
	return GetMedicamentsResponse{res, 200, nil}
}
