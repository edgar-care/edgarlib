package exam

import (
	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
	"math/rand"
	"sort"
	"strings"
	"time"
)

type DiseaseCoverage struct {
	Disease           string
	Percentage        float64
	Unknown           float64
	PotentialQuestion string
}

type ByCoverage []DiseaseCoverage

func (a ByCoverage) Len() int           { return len(a) }
func (a ByCoverage) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByCoverage) Less(i, j int) bool { return a[i].Percentage > a[j].Percentage }

func isChronic(sessionSymptom model.SessionSymptom) bool {
	symptoms, _ := graphql.GetSymptoms(nil)

	for _, symptom := range symptoms {
		if symptom.Name == sessionSymptom.Name && *sessionSymptom.Duration >= *symptom.Chronic {
			return true
		}
	}

	return false
}

func isInStringArray(array []string, str string) bool {
	for _, s := range array {
		if s != "" && str != "" && s == str {
			return true
		}
	}
	return false
}

func isAnteChir(symptomName string, anteChirs []model.AnteChir) float64 {
	for _, anteChir := range anteChirs {
		for _, iS := range anteChir.InducedSymptoms {
			if iS.Symptom == symptomName {
				return iS.Factor
			}
		}
	}
	return 1.0
}

func CalculPercentage(context_ []model.SessionSymptom, disease model.Disease, imc float64, anteChirIds []string, hereditary_disease []string) DiseaseCoverage {
	var anteChirs []model.AnteChir
	var potentialQuestionSymptom string
	var buf string
	percentage := 0.0
	unknown := 0.0

	if anteChirIds != nil && len(anteChirIds) > 0 {
		for _, chirId := range anteChirIds {
			var nAC model.AnteChir
			anteChir, _ := graphql.GetAnteChirByID(chirId)
			nAC.ID = anteChir.ID
			nAC.Name = anteChir.Name
			for _, iS := range anteChir.InducedSymptoms {
				var nIS model.ChirInducedSymptom
				nIS.Symptom = iS.Symptom
				nIS.Factor = iS.Factor
				nAC.InducedSymptoms = append(nAC.InducedSymptoms, &nIS)
			}
			anteChirs = append(anteChirs, nAC)
		}
	}

	for _, symptomWeight := range disease.SymptomsWeight {
		lock := 1
		for _, contextSymptom := range context_ {
			if symptomWeight.Symptom == contextSymptom.Name && contextSymptom.Presence == 0 {
				unknown += symptomWeight.Value
				lock = 0
				break
			} else if symptomWeight.Symptom == contextSymptom.Name && contextSymptom.Presence == 1 {
				if symptomWeight.Chronic && !isChronic(contextSymptom) {
					percentage += symptomWeight.Value * 0.75 * isAnteChir(contextSymptom.Name, anteChirs)
				} else if !symptomWeight.Chronic && isChronic(contextSymptom) {
					percentage += symptomWeight.Value * 0.75 * isAnteChir(contextSymptom.Name, anteChirs)
				} else {
					percentage += symptomWeight.Value * isAnteChir(contextSymptom.Name, anteChirs)
				}
				lock = 0
				break
			} else if symptomWeight.Symptom == contextSymptom.Name && contextSymptom.Presence == 2 {
				if contextSymptom.Treated != nil && len(contextSymptom.Treated) != 0 {
					percentage += symptomWeight.Value * 0.5
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
	if disease.OverweightFactor != 0 && imc > 25.0 {
		percentage *= disease.OverweightFactor
	}
	if disease.HeredityFactor != 0 && isInStringArray(hereditary_disease, disease.Name) {
		percentage *= disease.HeredityFactor
	}

	return DiseaseCoverage{Disease: disease.Name, Percentage: percentage, Unknown: unknown, PotentialQuestion: potentialQuestionSymptom}
}

func AddDiscursiveConnector(question string) string {
	rand.Seed(time.Now().UnixNano())
	connectorList := []string{"Je comprends", "Très bien", "Je note", "Je vois", "C'est noté", "Je vous suis", "D'accord", "Ensuite", "D'accord, je vois", "C'est entendu", "Ça marche", "Compris", "Oui, je comprends", "Entendu"}
	it := rand.Intn(len(connectorList) - 1)
	question = strings.Replace(question, "{{connecteur}}", connectorList[it], 1)
	return question
}

func getTheQuestion(symptomName string, symptoms []model.Symptom) (string, *string) {
	autoA := "Oui / Non / Ne sais pas"
	for _, symptom := range symptoms {
		if symptomName == symptom.Name {
			if symptom.QuestionBasic != "" {
				return AddDiscursiveConnector(symptom.QuestionBasic), &autoA
			} else {
				return AddDiscursiveConnector("{{connecteur}}. Est-ce que vous avez ce symptôme: " + symptom.Name + " ?"), &autoA
			}
		}
	}
	return AddDiscursiveConnector("{{connecteur}}. Est-ce que vous avez ce symptôme: " + symptomName + " ?"), &autoA
}

func GuessQuestion(mapped []DiseaseCoverage) (string, *string, []string, error) {
	symptoms, err := graphql.GetSymptoms(nil)
	if err != nil {
		return "", nil, nil, err
	}
	i := 0

	for mapped[i].PotentialQuestion == "" {
		i++
	}
	question, autoA := getTheQuestion(mapped[i].PotentialQuestion, symptoms)

	return question, autoA, []string{mapped[i].PotentialQuestion}, nil
}

func Calculi(sessionContext []model.SessionSymptom, imc float64, anteChirIds []string, hereditary_disease []string) ([]DiseaseCoverage, bool) {
	diseases, _ := graphql.GetDiseases(nil)
	mapped := make([]DiseaseCoverage, len(diseases))
	for i, e := range diseases {
		mapped[i] = CalculPercentage(sessionContext, e, imc, anteChirIds, hereditary_disease)
	}
	sort.Sort(ByCoverage(mapped))

	for _, disease := range mapped {
		if disease.Percentage > 0.7 {
			return mapped, true
		}
	}
	return mapped, false
}
