package utils

import (
	"github.com/edgar-care/edgarlib/v2/exam"
	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
	"sort"
)

type ByCoverage []exam.DiseaseCoverage

func (a ByCoverage) Len() int           { return len(a) }
func (a ByCoverage) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByCoverage) Less(i, j int) bool { return a[i].Percentage > a[j].Percentage }

func GetSessionDiseases(sessionContext []model.SessionSymptom, imc float64, hereditaryDisease []string) ([]*model.SessionDiseasesInput, error) {
	diseases, err := graphql.GetDiseases(nil)
	if err != nil {
		return nil, err
	}
	mapped := make([]exam.DiseaseCoverage, len(diseases))
	var sortedDiseases []*model.SessionDiseasesInput
	for i, e := range diseases {
		mapped[i] = exam.CalculPercentage(sessionContext, e, imc, hereditaryDisease)
	}
	sort.Sort(ByCoverage(mapped))

	for _, disease := range mapped {
		if disease.Percentage >= 0.30 {
			newsorted := model.SessionDiseasesInput{Name: disease.Disease, Presence: disease.Percentage, UnknownPresence: disease.Unknown}
			sortedDiseases = append(sortedDiseases, &newsorted)
		}
	}
	return sortedDiseases, nil
}
