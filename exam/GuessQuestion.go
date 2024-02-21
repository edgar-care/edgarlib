package exam

import (
	"context"
	"github.com/edgar-care/edgarlib/graphql"
	"github.com/edgar-care/edgarlib/graphql/server/model"
	"sort"
)

type diseaseCoverage struct {
	disease           string
	coverage          int
	present           int
	absent            int
	potentialQuestion string
}

type ByCoverage []diseaseCoverage

func (a ByCoverage) Len() int           { return len(a) }
func (a ByCoverage) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByCoverage) Less(i, j int) bool { return a[i].coverage > a[j].coverage }

func findInContext(context []model.SessionSymptom, symptom string) *model.SessionSymptom {
	for _, item := range context {
		if item.Name == symptom {
			return &item
		}
	}
	return nil
}

func isPresent(context []model.SessionSymptom, symptom string) *bool {
	item := findInContext(context, symptom)
	if item != nil {
		return item.Presence
	}
	return nil
}

func calculCoverage(context []model.SessionSymptom, disease graphql.GetDiseasesGetDiseasesDisease) diseaseCoverage {
	var coverage int
	var present int
	var absent int
	var potentialquestionSymptom string
	total := len(disease.Symptoms)

	for _, symptom := range disease.Symptoms {
		presence := isPresent(context, symptom)
		if presence != nil {
			coverage += 1
			if *presence == true {
				present += 1
			} else {
				absent += 1
			}
		} else {
			potentialquestionSymptom = symptom
		}
	}
	return diseaseCoverage{disease: disease.Code, coverage: coverage * 100 / total, present: present * 100 / total, absent: absent * 100 / total, potentialQuestion: potentialquestionSymptom}
}

func getTheQuestion(symptomName string, symptoms []graphql.GetSymptomsGetSymptomsSymptom) string {
	for _, symptom := range symptoms {
		if symptomName == symptom.Name {
			return symptom.Question
		}
	}
	return "Est-ce que vous avez ce symptôme: " + symptomName + " ?"
}

func GuessQuestion(patientContext []model.SessionSymptom) (string, []string, bool) {
	gqlClient := graphql.CreateClient()
	diseases, _ := graphql.GetDiseases(context.Background(), gqlClient)
	symptoms, _ := graphql.GetSymptoms(context.Background(), gqlClient)
	mapped := make([]diseaseCoverage, len(diseases.GetDiseases))
	for i, e := range diseases.GetDiseases {
		mapped[i] = calculCoverage(patientContext, e)
	}

	if len(patientContext) == 0 {
		return "Pourriez-vous décrire vos symptomes ?", []string{}, false
	}

	sort.Sort(ByCoverage(mapped))

	for _, disease := range mapped {
		if disease.absent >= 40 {
			continue
		}
		if disease.present >= 70 {
			return "", []string{}, true
		}
		return getTheQuestion(disease.potentialQuestion, symptoms.GetSymptoms), []string{disease.potentialQuestion}, false
	}
	return "", []string{}, true
}
