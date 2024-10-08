package medicament

import (
	"errors"
	"github.com/edgar-care/edgarlib/v2/paging"

	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
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
	medicament, err := graphql.GetMedicineByID(id)
	if err != nil {
		return GetMedicamentByIdResponse{model.Medicine{}, 400, errors.New("id does not correspond to a medicament")}
	}
	return GetMedicamentByIdResponse{medicament, 200, nil}
}

func GetMedicaments(page int, size int) GetMedicamentsResponse {
	medicaments, err := graphql.GetMedicines(paging.CreatePagingOption(page, size))
	if err != nil {
		return GetMedicamentsResponse{[]model.Medicine{}, 400, errors.New("invalid input: " + err.Error())}
	}
	return GetMedicamentsResponse{medicaments, 200, nil}
}
