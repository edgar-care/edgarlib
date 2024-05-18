package diagnostic

import (
	"context"
	"errors"
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

func nameInList(s utils.Symptom, symptoms []model.SessionSymptom) (bool, int) {
	for i, symptom := range symptoms {
		if symptom.Name == s.Name {
			return true, i
		}
	}
	return false, 0
}

func Diagnose(id string, sentence string) DiagnoseResponse {
	gqlClient := graphql.CreateClient()
	session, err := graphql.GetSessionById(context.Background(), gqlClient, id)
	if err != nil {
		return DiagnoseResponse{
			Code: 400,
			Err:  errors.New("id does not correspond to a session"),
		}
	}

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

	if len(session.GetSessionById.Logs) > 0 {
		session.GetSessionById.Logs[len(session.GetSessionById.Logs)-1].Answer = sentence
	}

	questionSymptom := []string{session.GetSessionById.Last_question}

	var newSymptoms utils.NlpResponseBody
	if session.GetSessionById.Last_question == "" || session.GetSessionById.Last_question == "describe symptoms" || session.GetSessionById.Last_question == "describe medicines" {
		questionSymptom = []string{}
	}

	var durSymptom *string
	if session.GetSessionById.Last_question != "" && strings.Split(session.GetSessionById.Last_question, " ")[0] == "duration" {
		durSymptom = &strings.Split(session.GetSessionById.Last_question, " ")[1]
	}
	var errCode int
	newSymptoms, errCode = utils.CallNlp(sentence, questionSymptom, durSymptom)
	if errCode != 200 {
		return DiagnoseResponse{
			Code: errCode,
			Err:  errors.New("NLP error, please try again"),
		}
	}

	for _, s := range newSymptoms.Context {
		pres, ite := nameInList(s, symptoms)
		if pres == true {
			symptoms[ite].Duration = s.Days
			continue
		}
		var newSessionSymptom model.SessionSymptom
		newSessionSymptom.Name = s.Name
		if s.Present == nil { // todo: temporaire le temps de changer les fonction graphql
			f := false
			newSessionSymptom.Presence = &f
		} else {
			newSessionSymptom.Presence = s.Present
		}
		newSessionSymptom.Duration = s.Days
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
		//} else if len(session.GetSessionById.Medicine) != 0 && session.GetSessionById.Medicine[0] == "CanonFlesh" { todo: uncomment CannonFlesh when this is enabled
		//	exam.Question = "Avez-vous pris des médicaments récemment ?"
		//	session.GetSessionById.Last_question = "describe medicines"
		//	if len(session.GetSessionById.Medicine) > 1 {
		//		session.GetSessionById.Medicine = session.GetSessionById.Medicine[1:]
		//	} else {
		//		session.GetSessionById.Medicine = []string{}
		//	}

	} else if exam.Question == "" && session.GetSessionById.Last_question == "" {
		exam = utils.CallExam(symptoms, float64(session.GetSessionById.Weight)/(float64(session.GetSessionById.Height)/100.0*(float64(session.GetSessionById.Height)/100.0)))
		if exam.Err != nil {
			if exam.Err != nil {
				return DiagnoseResponse{
					Code: 500,
					Err:  exam.Err,
				}
			}
		}

		if len(exam.Alert) > 0 {
			for _, alert := range exam.Alert {
				session.GetSessionById.Alerts = append(session.GetSessionById.Alerts, alert)
			}
		}

		if len(session.GetSessionById.Medicine) > 0 {
			symptomsinput, err = utils.CheckTreatments(symptomsinput, session.GetSessionById.Medicine)
			if err != nil {
				return DiagnoseResponse{
					Code: 500,
					Err:  errors.New("error during checkTreatment"),
				}
			}
		}

		if len(session.GetSessionById.Ante_diseases) > 0 {
			var anteSymptomQuestion string
			anteSymptomQuestion, anteSymptom, err = utils.CheckAnteDiseaseInSymptoms(session.GetSessionById)
			if err != nil {
				return DiagnoseResponse{
					Code: 400, //metter un code correct
					Err:  errors.New("error during checkAnteDiseaseInSymptoms"),
				}
			}
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

	var logs []graphql.LogsInput //
	for _, log := range session.GetSessionById.Logs {
		logs = append(logs, graphql.LogsInput{
			Question: log.Question,
			Answer:   log.Answer,
		})
	}
	if !exam.Done {
		logs = append(logs, graphql.LogsInput{
			Question: exam.Question,
			Answer:   "",
		})
	}

	var diseasesinput []graphql.SessionDiseasesInput
	if exam.Done == true {
		diseasesinput, err = utils.GetSessionDiseases(symptoms)
		if err != nil {
			return DiagnoseResponse{
				Code: 500,
				Err:  errors.New("error during getSessionDiseases"),
			}
		}
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
