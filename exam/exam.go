package exam

type ExamResponse struct {
	Context  []ExamContextItem
	Question string
	Symptoms []string
	Done     bool
	Alert    []string
	Code     int
	Err      error
}

type ExamContextItem struct {
	Name     string
	Presence *bool
	Duration *int32
}

func Exam(context []ExamContextItem) ExamResponse {
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
