package exam

import (
	"context"
	"github.com/edgar-care/edgarlib/graphql"
	"github.com/edgar-care/edgarlib/graphql/server/model"
	"sort"
)

type diseaseCoverage struct {
	disease           string
	percentage        float64
	potentialQuestion string
}

type ByCoverage []diseaseCoverage

func (a ByCoverage) Len() int           { return len(a) }
func (a ByCoverage) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByCoverage) Less(i, j int) bool { return a[i].percentage > a[j].percentage }

func calculPercentage(context []model.SessionSymptom, disease graphql.GetDiseasesGetDiseasesDisease) diseaseCoverage {
	var potentialQuestionSymptom string
	var buf string
	percentage := 0.0

	for _, symptomWeight := range disease.Symptoms_weight {
		lock := 1
		for _, contextSymptom := range context {
			if symptomWeight.Symptom == contextSymptom.Name && *contextSymptom.Presence == true {
				percentage += symptomWeight.Value
				lock = 0
				break
			} else if symptomWeight.Symptom == contextSymptom.Name && *contextSymptom.Presence == false {
				if contextSymptom.Treated != nil && len(contextSymptom.Treated) != 0 {
					percentage += symptomWeight.Value / 2
				}
				lock = 0
				break
			} else {
				buf = symptomWeight.Symptom
			}
		}
		if lock == 1 {
			potentialQuestionSymptom = buf
		}
	}

	return diseaseCoverage{disease: disease.Code, percentage: percentage, potentialQuestion: potentialQuestionSymptom}
}

func getTheQuestion(symptomName string, symptoms []graphql.GetSymptomsGetSymptomsSymptom) string {
	for _, symptom := range symptoms {
		if symptomName == symptom.Code {
			return symptom.Question
		}
	}
	return "Est-ce que vous avez ce symptôme: " + symptomName + " ?"
}

func GuessQuestion(mapped []diseaseCoverage) (string, []string) {
	gqlClient := graphql.CreateClient()
	symptoms, _ := graphql.GetSymptoms(context.Background(), gqlClient)
	i := 0

	for mapped[i].potentialQuestion == "" {
		i++
	}

	return getTheQuestion(mapped[i].potentialQuestion, symptoms.GetSymptoms), []string{mapped[i].potentialQuestion}
}

func Calculi(sessionContext []model.SessionSymptom) ([]diseaseCoverage, bool) {
	gqlClient := graphql.CreateClient()
	diseases, _ := graphql.GetDiseases(context.Background(), gqlClient)
	mapped := make([]diseaseCoverage, len(diseases.GetDiseases))
	for i, e := range diseases.GetDiseases {
		mapped[i] = calculPercentage(sessionContext, e)
	}
	sort.Sort(ByCoverage(mapped))

	for _, disease := range mapped {
		if disease.percentage > 0.7 {
			return mapped, true
		}
	}
	return mapped, false
}
