package diagnostic

import (
	"context"
	"github.com/edgar-care/edgarlib"
	"github.com/edgar-care/edgarlib/diagnostic/utils"
	"github.com/edgar-care/edgarlib/graphql"
)

type DiagnoseResponse struct {
	Done     bool
	Question string
	Code     int
	Err      error
}

func Diagnose(id string, sentence string) DiagnoseResponse {
	gqlClient := graphql.CreateClient()
	session, err := graphql.GetSessionById(context.Background(), gqlClient, id)

	symptoms := utils.StringToSymptoms(session.GetSessionById.Symptoms)
	questionSymptom := []string{session.GetSessionById.Last_question}
	if session.GetSessionById.Last_question == "" {
		questionSymptom = []string{}

		tmp := graphql.GetSessionByIdGetSessionByIdSessionLogs{Question: "", Answer: sentence}
		session.GetSessionById.Logs = append(session.GetSessionById.Logs, tmp)
		edgarlib.CheckError(err)
	} else {
		tmp := graphql.GetSessionByIdGetSessionByIdSessionLogs{Question: session.GetSessionById.Last_question, Answer: sentence}
		session.GetSessionById.Logs = append(session.GetSessionById.Logs, tmp)
		edgarlib.CheckError(err)
	}
	newSymptoms := utils.CallNlp(sentence, questionSymptom)
	for _, s := range newSymptoms.Context {
		symptoms = append(symptoms, s)
	}

	exam := utils.CallExam(symptoms)
	if len(exam.Alert) > 0 {
		for _, alert := range exam.Alert {
			session.GetSessionById.Alerts = append(session.GetSessionById.Alerts, alert)
		}
	}
	session.GetSessionById.Symptoms = utils.SymptomsToString(exam.Context)
	if len(exam.Symptoms) > 0 {
		session.GetSessionById.Last_question = exam.Symptoms[0]
	}
	var logs []graphql.LogsInput
	for _, log := range session.GetSessionById.Logs {
		logs = append(logs, graphql.LogsInput{
			Question: log.Question,
			Answer:   log.Answer,
		})
	}
	_, err = graphql.UpdateSession(context.Background(), gqlClient, session.GetSessionById.Id, session.GetSessionById.Symptoms, session.GetSessionById.Age, session.GetSessionById.Height, session.GetSessionById.Weight, session.GetSessionById.Sex, session.GetSessionById.Last_question, logs, session.GetSessionById.Alerts)
	edgarlib.CheckError(err)
	return DiagnoseResponse{
		Done:     exam.Done,
		Question: exam.Question,
		Code:     200,
		Err:      nil,
	}
}
