package exam

import "github.com/edgar-care/edgarlib/graphql/server/model"

type ExamResponse struct {
	Context  []model.SessionSymptom
	Question string
	Symptoms []string
	Done     bool
	Alert    []string
	Code     int
	Err      error
}

func Exam(context []model.SessionSymptom) ExamResponse {
	question, possibleSymptoms, isDone := GuessQuestion(context)
	alert, err := CheckAlerts(context)
	if err != nil {
		return ExamResponse{Code: 500, Err: err}
	}
	return ExamResponse{
		Context:  context,
		Question: question,
		Symptoms: possibleSymptoms,
		Done:     isDone,
		Alert:    alert,
		Code:     200,
		Err:      nil,
	}
}
