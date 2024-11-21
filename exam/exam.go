package exam

import (
	"github.com/edgar-care/edgarlib/v2/graphql/model"
)

type ExamResponse struct {
	Context    []model.SessionSymptom
	Question   string
	Symptoms   []string
	AutoAnswer *string
	Done       bool
	Alert      []string
	Code       int
	Err        error
}

func Exam(context []model.SessionSymptom, imc float64, hereditaryDisease []string) ExamResponse {
	var question string
	var possibleSymptoms []string
	var autoA *string
	mappedDiseaseCoverage, isDone := Calculi(context, imc, hereditaryDisease)
	if isDone == false {
		var err error
		question, autoA, possibleSymptoms, err = GuessQuestion(mappedDiseaseCoverage)
		if err != nil {
			return ExamResponse{Code: 500, Err: err}
		}
	}
	alert, err := CheckAlerts(context)
	if err != nil {
		return ExamResponse{Code: 500, Err: err}
	}
	return ExamResponse{
		Context:    context,
		Question:   question,
		Symptoms:   possibleSymptoms,
		AutoAnswer: autoA,
		Done:       isDone,
		Alert:      alert,
		Code:       200,
		Err:        nil,
	}
}
