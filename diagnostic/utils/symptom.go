package utils

import (
	"fmt"
	"github.com/edgar-care/edgarlib/v2/exam"
	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
	"strings"
)

type Symptom struct {
	Name    string `json:"symptom"`
	Present *bool  `json:"present"`
	Days    *int   `json:"days"`
}

func SymptomsToString(symptoms []Symptom) []string {
	var strings = []string{}
	for _, symptom := range symptoms {
		if symptom.Present == nil {
			strings = append(strings, fmt.Sprintf("?%s", symptom.Name))
		} else if *symptom.Present == true {
			strings = append(strings, fmt.Sprintf("*%s", symptom.Name))
		} else if *symptom.Present == false {
			strings = append(strings, fmt.Sprintf("!%s", symptom.Name))
		}
	}
	return strings
}

func pointerToBool(val bool) *bool {
	return &val
}

func StringToSymptoms(strings []string) []Symptom {
	var newSymptoms = []Symptom{}
	for _, symptom := range strings {
		if symptom[0] == '*' {
			newSymptoms = append(newSymptoms, Symptom{Name: symptom[1:], Present: pointerToBool(true)})
		} else if symptom[0] == '!' {
			newSymptoms = append(newSymptoms, Symptom{Name: symptom[1:], Present: pointerToBool(false)})
		} else if symptom[0] == '?' {
			newSymptoms = append(newSymptoms, Symptom{Name: symptom[1:], Present: nil})
		}
	}
	return newSymptoms
}

func CheckSymptomDuration(symptoms []*model.SessionSymptomInput, lastQuestion string) ([]*model.SessionSymptomInput, string, string) {
	//allSymptoms, _ := graphql.GetSymptoms(nil)
	list := strings.Split(lastQuestion, " ")
	question := ""
	nextLastQuestion := ""
	var symptomName string

	if list[0] == "duration" {
		symptomName = list[1]
	}
	for _, symptom := range symptoms {
		if symptom.Duration != nil && *symptom.Duration == 0 && symptom.Presence == 1 && symptomName != symptom.Name {
			sy, _ := graphql.GetSymptomByCode(symptom.Name)
			if sy.QuestionDuration != "" {
				question = exam.AddDiscursiveConnector(sy.QuestionDuration)
			} else {
				question = exam.AddDiscursiveConnector("{{connecteur}}. Depuis combien de jours souffrez-vous de " + symptom.Name)
			}
			nextLastQuestion = "duration " + symptom.Name
		}

	}
	return symptoms, question, nextLastQuestion
}
