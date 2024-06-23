package utils

import (
	"context"
	"github.com/edgar-care/edgarlib/exam"
	"github.com/edgar-care/edgarlib/graphql"
	"github.com/edgar-care/edgarlib/graphql/server/model"
	"sort"
)

type ByCoverage []exam.DiseaseCoverage

func (a ByCoverage) Len() int           { return len(a) }
func (a ByCoverage) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByCoverage) Less(i, j int) bool { return a[i].Percentage > a[j].Percentage }

func GetSessionDiseases(sessionContext []model.SessionSymptom, imc float64, anteChirIds []string) ([]graphql.SessionDiseasesInput, error) {
	gqlClient := graphql.CreateClient()
	diseases, err := graphql.GetDiseases(context.Background(), gqlClient)
	if err != nil {
		return nil, err
	}
	mapped := make([]exam.DiseaseCoverage, len(diseases.GetDiseases))
	var sortedDiseases []graphql.SessionDiseasesInput
	for i, e := range diseases.GetDiseases {
		mapped[i] = exam.CalculPercentage(sessionContext, e, imc, anteChirIds)
	}
	sort.Sort(ByCoverage(mapped))

	for _, disease := range mapped {
		if disease.Percentage >= 0.30 {
			newsorted := graphql.SessionDiseasesInput{Name: disease.Disease, Presence: disease.Percentage, Unknown_presence: disease.Unknown}
			sortedDiseases = append(sortedDiseases, newsorted)
		}
	}
	return sortedDiseases, nil
}
