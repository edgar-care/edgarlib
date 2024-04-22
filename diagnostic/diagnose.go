package diagnostic

import (
	"context"
	"github.com/edgar-care/edgarlib"
	"github.com/edgar-care/edgarlib/diagnostic/utils"
	"github.com/edgar-care/edgarlib/graphql"
	"github.com/edgar-care/edgarlib/graphql/server/model"
	"strings"
)

type DiagnoseResponse struct {
	Done     bool
	Question string
	Code     int
	Err      error
}

func nameInList(s utils.Symptom, symptoms []model.SessionSymptom) bool {
	for _, symptom := range symptoms {
		if symptom.Name == s.Name {
			return true
		}
	}
	return false
}

func Diagnose(id string, sentence string) DiagnoseResponse {
	gqlClient := graphql.CreateClient()
	session, err := graphql.GetSessionById(context.Background(), gqlClient, id)

	var symptoms []model.SessionSymptom
	for _, s := range session.GetSessionById.Symptoms {
		var ns model.SessionSymptom
		ns.Name = s.Name
		var b bool
		if s.Presence == true {
			b = true
		} else {
			b = false
		}
		ns.Presence = &b
		dura := s.Duration
		ns.Duration = &dura
		ns.Treated = s.Treated
		symptoms = append(symptoms, ns)
	}

	if session.GetSessionById.Last_question != "" {
		tmp := graphql.GetSessionByIdGetSessionByIdSessionLogs{Question: session.GetSessionById.Last_question, Answer: sentence}
		session.GetSessionById.Logs = append(session.GetSessionById.Logs, tmp)
		edgarlib.CheckError(err)
	}

	questionSymptom := []string{session.GetSessionById.Last_question}

	var newSymptoms utils.NlpResponseBody
	if session.GetSessionById.Last_question == "" || session.GetSessionById.Last_question == "describe symptoms" || session.GetSessionById.Last_question == "describe medicines" {
		questionSymptom = []string{}
	}

	if (len(questionSymptom) == 0) || (len(questionSymptom) != 0 && len(strings.Split(questionSymptom[0], " ")) == 1) {
		newSymptoms = utils.CallNlp(sentence, questionSymptom)
	}

	for _, s := range newSymptoms.Context {
		if nameInList(s, symptoms) {
			continue
		}
		var newSessionSymptom model.SessionSymptom
		newSessionSymptom.Name = s.Name
		newSessionSymptom.Presence = s.Present
		symptoms = append(symptoms, newSessionSymptom)
	}

	var symptomsinput []graphql.SessionSymptomInput
	for _, s := range symptoms {
		var ns graphql.SessionSymptomInput
		ns.Name = s.Name
		if s.Presence != nil && *s.Presence == true {
			ns.Presence = true
		} else {
			ns.Presence = false
		}
		if s.Duration == nil {
			ns.Duration = 0
		} else {
			ns.Duration = *s.Duration
		}
		symptomsinput = append(symptomsinput, ns)
	}

	var anteSymptom string

	var exam utils.ExamResponseBody
	if len(symptomsinput) > 0 {
		symptomsinput, exam.Question, session.GetSessionById.Last_question = utils.CheckSymptomDuration(symptomsinput, session.GetSessionById.Last_question, sentence)
	}
	if len(symptoms) == 0 {
		exam.Question = "Pourriez-vous décrire vos symptomes ?"
		exam.Done = false
		exam.Context = symptoms
		exam.Symptoms = []string{}
		exam.Alert = []string{}
		session.GetSessionById.Last_question = "describe symptoms"
	} else if session.GetSessionById.Medicine[0] == "CanonFlesh" {
		exam.Question = "Avez-vous pris des médicaments récemment ?"
		session.GetSessionById.Last_question = "describe medicines"
		if len(session.GetSessionById.Medicine) > 1 {
			session.GetSessionById.Medicine = session.GetSessionById.Medicine[1:]
		} else {
			session.GetSessionById.Medicine = []string{}
		}

	} else if exam.Question == "" && session.GetSessionById.Last_question == "" {
		exam = utils.CallExam(symptoms)

		if len(exam.Alert) > 0 {
			for _, alert := range exam.Alert {
				session.GetSessionById.Alerts = append(session.GetSessionById.Alerts, alert)
			}
		}

		if len(session.GetSessionById.Medicine) > 0 {
			symptomsinput = utils.CheckTreatments(symptomsinput, session.GetSessionById.Medicine)
		}

		if len(session.GetSessionById.Ante_diseases) > 0 {
			var anteSymptomQuestion string
			anteSymptomQuestion, anteSymptom = utils.CheckAnteDiseaseInSymptoms(session.GetSessionById)
			if anteSymptom != "" {
				exam.Question = anteSymptomQuestion
				session.GetSessionById.Last_question = anteSymptom
			}
		}

		if len(exam.Symptoms) > 0 && anteSymptom == "" {
			session.GetSessionById.Last_question = exam.Symptoms[0]
		}

		if len(exam.Symptoms) == 0 {
			session.GetSessionById.Last_question = ""
		}
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

	_, err = graphql.UpdateSession(context.Background(), gqlClient, session.GetSessionById.Id, diseasesinput, symptomsinput, session.GetSessionById.Age, session.GetSessionById.Height, session.GetSessionById.Weight, session.GetSessionById.Sex, session.GetSessionById.Ante_chirs, session.GetSessionById.Ante_diseases, session.GetSessionById.Medicine, session.GetSessionById.Last_question, logs, session.GetSessionById.Alerts)
	edgarlib.CheckError(err)
	return DiagnoseResponse{
		Done:     exam.Done,
		Question: exam.Question,
		Code:     200,
		Err:      nil,
	}
}
