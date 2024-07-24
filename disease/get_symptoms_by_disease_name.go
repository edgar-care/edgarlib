package disease

import (
	"context"
	"errors"
	"github.com/edgar-care/edgarlib/graphql"
)

type GetSymptomsByDiseaseNameResponse struct {
	Symptoms []string
	Code     int
	Err      error
}

func GetSymptomsByDiseaseName(name string) GetSymptomsByDiseaseNameResponse {
	gqlClient := graphql.CreateClient()
	var res []string

	symptoms, err := graphql.GetSymptomsByDiseaseName(context.Background(), gqlClient, name)
	if err != nil {
		return GetSymptomsByDiseaseNameResponse{[]string{}, 400, errors.New("name does not correspond to a disease")}
	}
	res = symptoms.GetSymptomsByDiseaseName.Symptoms
	return GetSymptomsByDiseaseNameResponse{res, 200, nil}
}
