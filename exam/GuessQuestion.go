package exam

import (
	"context"
	"github.com/edgar-care/edgarlib/graphql"
	"math/rand"
	"sort"
)

type diseaseCoverage struct {
	coverage int
	present  int
	absent   int
}

type ByCoverage []diseaseCoverage

func (a ByCoverage) Len() int           { return len(a) }
func (a ByCoverage) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByCoverage) Less(i, j int) bool { return a[i].coverage > a[j].coverage }

func findInContext(context []ExamContextItem, symptom string) *ExamContextItem {
	for _, item := range context {
		if item.Symptom == symptom {
			return &item
		}
	}
	return nil
}

func isPresent(context []ExamContextItem, symptom string) *bool {
	item := findInContext(context, symptom)
	if item != nil {
		return item.Present
	}
	return nil
}

func calculCoverage(context []ExamContextItem, disease graphql.GetDiseasesGetDiseasesDisease) diseaseCoverage {
	var coverage int
	var present int
	var absent int
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
		}
	}
	return diseaseCoverage{coverage: coverage * 100 / total, present: present * 100 / total, absent: absent * 100 / total}
}

func GuessQuestion(patientContext []ExamContextItem) (string, []string, bool) {
	gqlClient := graphql.CreateClient()
	diseases, _ := graphql.GetDiseases(context.Background(), gqlClient)
	symptoms := []string{"respiration_difficile", "toux", "respiration_sifflante", "somnolence", "anxiete", "brulure_poitrine", "respiration_difficile", "boule_gorge", "maux_de_tetes", "vision_trouble", "tache_visuel", "abdominalgies", "asthenie", "anorexie", "amaigrissement"}
	next := symptoms[rand.Intn(len(symptoms))]
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
		return next, []string{next}, false
	}
	return "", []string{}, true
}
