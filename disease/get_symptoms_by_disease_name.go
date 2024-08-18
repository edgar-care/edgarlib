package disease

import (
	"errors"
	"github.com/edgar-care/edgarlib/v2/graphql"
)

type GetSymptomsByDiseaseNameResponse struct {
	Symptoms []string
	Code     int
	Err      error
}

func GetSymptomsByDiseaseName(name string) GetSymptomsByDiseaseNameResponse {
	var res []string

	symptoms, err := graphql.GetSymptomsByDiseaseName(name)
	if err != nil {
		return GetSymptomsByDiseaseNameResponse{[]string{}, 400, errors.New("name does not correspond to a disease")}
	}
	res = symptoms.Symptoms
	return GetSymptomsByDiseaseNameResponse{res, 200, nil}
}
