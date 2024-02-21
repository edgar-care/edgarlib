package diagnostic

import (
	"context"
	"fmt"
	"github.com/edgar-care/edgarlib"
	"github.com/edgar-care/edgarlib/diagnostic/utils"
	"github.com/edgar-care/edgarlib/graphql"
	"github.com/edgar-care/edgarlib/graphql/server/model"
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

	var symptoms []model.SessionSymptom
	for _, s := range session.GetSessionById.Symptoms {
		var ns model.SessionSymptom
		ns.Name = s.Name
		ns.Presence = &s.Presence
		ns.Duration = &s.Duration
		ns.Treated = s.Treated
		symptoms = append(symptoms, ns)
	}

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

	fmt.Println(newSymptoms)

	for _, s := range newSymptoms.Context {
		var newSessionSymptom model.SessionSymptom
		newSessionSymptom.Name = s.Name
		newSessionSymptom.Presence = s.Present
		symptoms = append(symptoms, newSessionSymptom)
	}

	var symptomsinput []graphql.SessionSymptomInput

	var exam utils.ExamResponseBody
	if len(symptoms) == 0 {
		exam.Question = "Pourriez-vous décrire vos symptomes ?"
		exam.Done = false
		exam.Context = symptoms
		exam.Symptoms = []string{}
		exam.Alert = []string{}
	} else if session.GetSessionById.Treatments[0] == "CanonFlesh" {
		exam.Question = "Avez-vous pris des médicaments récemment ?"
		if len(session.GetSessionById.Treatments) > 1 {
			session.GetSessionById.Treatments = session.GetSessionById.Treatments[1:]
		} else {
			session.GetSessionById.Treatments = []string{}
		}
	} else {
		exam = utils.CallExam(symptoms)

		if len(exam.Alert) > 0 {
			for _, alert := range exam.Alert {
				session.GetSessionById.Alerts = append(session.GetSessionById.Alerts, alert)
			}
		}
		for _, s := range exam.Context {
			var ns graphql.SessionSymptomInput
			ns.Name = s.Name
			if s.Presence != nil && *s.Presence == true {
				ns.Presence = true
			} else {
				ns.Presence = false
			}
			ns.Duration = 0
			symptomsinput = append(symptomsinput, ns)
		}

		if len(session.GetSessionById.Treatments) > 0 {
			symptomsinput = utils.CheckTreatments(symptomsinput, session.GetSessionById.Treatments)
		}

		if len(session.GetSessionById.Ante_diseases) > 0 {
			anteSymptomQuestion, anteSymptom := utils.CheckAnteDiseaseInSymptoms(session.GetSessionById)
			if anteSymptom != "" {
				exam.Question = anteSymptomQuestion
				session.GetSessionById.Last_question = anteSymptom
			} else if len(exam.Symptoms) > 0 {
				session.GetSessionById.Last_question = exam.Symptoms[0]
			}
		}
	}

	if len(exam.Symptoms) == 0 {
		session.GetSessionById.Last_question = ""
	}

	var logs []graphql.LogsInput
	for _, log := range session.GetSessionById.Logs {
		logs = append(logs, graphql.LogsInput{
			Question: log.Question,
			Answer:   log.Answer,
		})
	}

	var diseasesinput []graphql.SessionDiseasesInput
	if exam.Done == true {
		diseasesinput = utils.GetSessionDiseases(symptoms)
	}

	_, err = graphql.UpdateSession(context.Background(), gqlClient, session.GetSessionById.Id, diseasesinput, symptomsinput, session.GetSessionById.Age, session.GetSessionById.Height, session.GetSessionById.Weight, session.GetSessionById.Sex, session.GetSessionById.Ante_chirs, session.GetSessionById.Ante_diseases, session.GetSessionById.Treatments, session.GetSessionById.Last_question, logs, session.GetSessionById.Alerts)
	edgarlib.CheckError(err)
	return DiagnoseResponse{
		Done:     exam.Done,
		Question: exam.Question,
		Code:     200,
		Err:      nil,
	}
}
