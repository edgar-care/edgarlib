package utils

import (
	"fmt"
	"github.com/edgar-care/edgarlib/graphql"
	"strconv"
	"strings"
)

type Symptom struct {
	Name    string `json:"symptom"`
	Present *bool  `json:"present"`
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

func CheckSymptomDuration(symptoms []graphql.SessionSymptomInput, lastQuestion string, sentence string) ([]graphql.SessionSymptomInput, string, string) {

	list := strings.Split(lastQuestion, " ")
	var duration int
	question := ""
	nextLastQuestion := ""
	var symptomName string

	if list[0] == "duration" {
		symptomName = list[1]
		duration, _ = strconv.Atoi(sentence)
	}
	for i, symptom := range symptoms {
		if symptom.Duration == 0 && symptom.Presence == true && symptomName != symptom.Name {
			question = "Depuis combien de jours souffrez-vous de " + symptom.Name
			nextLastQuestion = "duration " + symptom.Name
		}
		if symptom.Name == symptomName {
			symptoms[i].Duration = duration
		}

	}

	return symptoms, question, nextLastQuestion
}
