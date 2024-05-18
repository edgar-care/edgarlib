package exam

import (
	"context"
	"github.com/edgar-care/edgarlib/graphql"
	"github.com/edgar-care/edgarlib/graphql/server/model"
	"sort"
)

type DiseaseCoverage struct {
	Disease           string
	Percentage        float64
	PotentialQuestion string
}

type ByCoverage []DiseaseCoverage

func (a ByCoverage) Len() int           { return len(a) }
func (a ByCoverage) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByCoverage) Less(i, j int) bool { return a[i].Percentage > a[j].Percentage }

func isChronic(sessionSymptom model.SessionSymptom) bool {
	gqlClient := graphql.CreateClient()
	symptoms, _ := graphql.GetSymptoms(context.Background(), gqlClient)

	for _, symptom := range symptoms.GetSymptoms {
		if symptom.Code == sessionSymptom.Name && *sessionSymptom.Duration >= symptom.Chronic {
			return true
		}
	}

	return false
}

func CalculPercentage(context []model.SessionSymptom, disease graphql.GetDiseasesGetDiseasesDisease, imc float64) DiseaseCoverage {
	var potentialQuestionSymptom string
	var buf string
	percentage := 0.0

	for _, symptomWeight := range disease.Symptoms_weight {
		lock := 1
		for _, contextSymptom := range context {
			if contextSymptom.Presence != nil {
				if symptomWeight.Symptom == contextSymptom.Name && *contextSymptom.Presence == true {
					if symptomWeight.Chronic && !isChronic(contextSymptom) {
						percentage += symptomWeight.Value * 0.75
					} else if !symptomWeight.Chronic && isChronic(contextSymptom) {
						percentage += symptomWeight.Value * 0.75
					} else {
						percentage += symptomWeight.Value
					}
					lock = 0
					break
				} else if symptomWeight.Symptom == contextSymptom.Name && *contextSymptom.Presence == false {
					if contextSymptom.Treated != nil && len(contextSymptom.Treated) != 0 {
						percentage += symptomWeight.Value * 0.5
					}
					lock = 0
					break
				} else {
					buf = symptomWeight.Symptom
				}
			}
		}
		if lock == 1 {
			potentialQuestionSymptom = buf
		}
	}
	if disease.Overweight_factor != 0 && imc > 25.0 {
		percentage *= disease.Overweight_factor
	}

	return DiseaseCoverage{Disease: disease.Code, Percentage: percentage, PotentialQuestion: potentialQuestionSymptom}
}

func getTheQuestion(symptomName string, symptoms []graphql.GetSymptomsGetSymptomsSymptom) string {
	for _, symptom := range symptoms {
		if symptomName == symptom.Code {
			if symptom.Question != "" {
				return symptom.Question
			} else {
				return "Est-ce que vous avez ce symptôme: " + symptom.Code + " ?"
			}
		}
	}
	return "Est-ce que vous avez ce symptôme: " + symptomName + " ?"
}

func GuessQuestion(mapped []DiseaseCoverage) (string, []string, error) {
	gqlClient := graphql.CreateClient()
	symptoms, err := graphql.GetSymptoms(context.Background(), gqlClient)
	if err != nil {
		return "", nil, err
	}
	i := 0

	for mapped[i].PotentialQuestion == "" {
		i++
	}

	return getTheQuestion(mapped[i].PotentialQuestion, symptoms.GetSymptoms), []string{mapped[i].PotentialQuestion}, nil
}

func Calculi(sessionContext []model.SessionSymptom, imc float64) ([]DiseaseCoverage, bool) {
	gqlClient := graphql.CreateClient()
	diseases, _ := graphql.GetDiseases(context.Background(), gqlClient)
	mapped := make([]DiseaseCoverage, len(diseases.GetDiseases))
	for i, e := range diseases.GetDiseases {
		mapped[i] = CalculPercentage(sessionContext, e, imc)
	}
	sort.Sort(ByCoverage(mapped))

	for _, disease := range mapped {
		if disease.Percentage > 0.7 {
			return mapped, true
		}
	}
	return mapped, false
}
