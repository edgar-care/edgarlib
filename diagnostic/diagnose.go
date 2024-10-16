package diagnostic

import (
	"errors"
	"github.com/edgar-care/edgarlib/v2"
	"github.com/edgar-care/edgarlib/v2/diagnostic/utils"
	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
	"strings"
)

type DiagnoseResponse struct {
	Done       bool
	Question   string
	AutoAnswer *model.AutoAnswer
	Code       int
	Err        error
}

type AutoAnswerinfo struct {
	Name   string
	Values []string
}

func nameInList(s utils.Symptom, symptoms []model.SessionSymptom) (bool, int) {
	for i, symptom := range symptoms {
		if symptom.Name == s.Name {
			return true, i
		}
	}
	return false, 0
}

func Diagnose(id string, sentence string, autoAnswer *AutoAnswerinfo) DiagnoseResponse {
	session, err := graphql.GetSessionById(id)
	if err != nil {
		return DiagnoseResponse{
			Code: 400,
			Err:  errors.New("id does not correspond to a session"),
		}
	}

	var symptoms []model.SessionSymptom
	for _, s := range session.Symptoms {
		var ns model.SessionSymptom
		ns.Name = s.Name
		ns.Presence = s.Presence
		ns.Duration = s.Duration
		ns.Treated = s.Treated
		symptoms = append(symptoms, ns)
	}

	if len(session.Logs) > 0 {
		session.Logs[len(session.Logs)-1].Answer = sentence
	}

	questionSymptom := []string{session.LastQuestion}

	var newSymptoms utils.NlpResponseBody
	if session.LastQuestion == "" {
		questionSymptom = []string{}
	}
	if session.LastQuestion == "describe symptoms" || session.LastQuestion == "describe medicines" {
		questionSymptom = []string{}
	} else {

		var durSymptom *string
		if session.LastQuestion != "" && strings.Split(session.LastQuestion, " ")[0] == "duration" {
			durSymptom = &strings.Split(session.LastQuestion, " ")[1]
		}

		if autoAnswer != nil {
			autoA, _ := graphql.GetAutoAnswerByName(autoAnswer.Name)
			if autoA.Name == "Oui / Non / Ne sais pas" {
				if len(questionSymptom) > 0 {
					if autoAnswer.Values[0] == "Oui." {
						p := true
						newSymptoms.Context = append(newSymptoms.Context, utils.Symptom{Name: questionSymptom[0], Present: &p})
					} else if autoAnswer.Values[0] == "Non." {
						p := false
						newSymptoms.Context = append(newSymptoms.Context, utils.Symptom{Name: questionSymptom[0], Present: &p})
					} else if autoAnswer.Values[0] == "Je ne sais pas." {
						var p *bool
						p = nil
						newSymptoms.Context = append(newSymptoms.Context, utils.Symptom{Name: questionSymptom[0], Present: p})
					}
				}
			}
		} else {
			var errCode int
			isMedicine := false
			if session.LastQuestion == "describe medicines" {
				isMedicine = true
			}
			newSymptoms, errCode = utils.CallNlp(sentence, questionSymptom, durSymptom, isMedicine)
			if errCode != 200 {
				return DiagnoseResponse{
					Code: errCode,
					Err:  errors.New("NLP error, please try again"),
				}
			}
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
			newSessionSymptom.Presence = 0
		} else if *s.Present {
			newSessionSymptom.Presence = 1
		} else {
			newSessionSymptom.Presence = 2
		}
		newSessionSymptom.Duration = s.Days
		symptoms = append(symptoms, newSessionSymptom)
	}

	var symptomsinput []*model.SessionSymptomInput
	for _, s := range symptoms {
		var ns model.SessionSymptomInput
		ns.Name = s.Name
		ns.Presence = s.Presence
		if s.Duration == nil {
			duration := 0
			ns.Duration = &duration
		} else {
			ns.Duration = s.Duration
		}
		symptomsinput = append(symptomsinput, &ns)
	}

	var anteSymptom string

	var exam utils.ExamResponseBody
	if len(symptomsinput) > 0 {
		symptomsinput, exam.Question, session.LastQuestion = utils.CheckSymptomDuration(symptomsinput, session.LastQuestion)
	}
	if len(symptoms) == 0 {
		exam.Question = "Pourriez-vous décrire vos symptomes ?"
		exam.Done = false
		exam.Context = symptoms
		exam.Symptoms = []string{}
		exam.Alert = []string{}
		session.LastQuestion = "describe symptoms"
	} else if len(session.Medicine) != 0 && session.Medicine[0] == "CanonFlesh" { //todo: uncomment CannonFlesh when this is enabled
		exam.Question = "Avez-vous pris des médicaments récemment ?"
		session.LastQuestion = "describe medicines"
		if len(session.Medicine) > 1 {
			session.Medicine = session.Medicine[1:]
		} else {
			session.Medicine = []string{}
		}

	} else if exam.Question == "" && session.LastQuestion == "" {
		exam = utils.CallExam(symptoms, float64(session.Weight)/(float64(session.Height)/100.0*(float64(session.Height)/100.0)), session.AnteChirs, session.HereditaryDisease)
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
				session.Alerts = append(session.Alerts, alert)
			}
		}

		if len(session.Medicine) > 0 {
			symptomsinput, err = utils.CheckTreatments(symptomsinput, session.Medicine)
			if err != nil {
				return DiagnoseResponse{
					Code: 500,
					Err:  errors.New("error during checkTreatment"),
				}
			}
		}

		if len(session.AnteDiseases) > 0 {
			var anteSymptomQuestion string
			anteSymptomQuestion, anteSymptom, err = utils.CheckAnteDiseaseInSymptoms(session)
			if err != nil {
				return DiagnoseResponse{
					Code: 400, //metter un code correct
					Err:  errors.New("error during checkAnteDiseaseInSymptoms"),
				}
			}
			if anteSymptom != "" {
				exam.Question = anteSymptomQuestion
				session.LastQuestion = anteSymptom
			}
		}

		if len(exam.Symptoms) > 0 && anteSymptom == "" {
			session.LastQuestion = exam.Symptoms[0]
		}

		if len(exam.Symptoms) == 0 {
			session.LastQuestion = ""
		}
	}

	var logs []*model.LogsInput //
	for _, log := range session.Logs {
		logs = append(logs, &model.LogsInput{
			Question: log.Question,
			Answer:   log.Answer,
		})
	}
	if len(logs) == 0 && sentence != "" {
		logs = append(logs, &model.LogsInput{
			Question: "Pourriez-vous décrire vos symptomes ?",
			Answer:   sentence,
		})
	}
	if !exam.Done {
		logs = append(logs, &model.LogsInput{
			Question: exam.Question,
			Answer:   "",
		})
	}

	var diseasesinput []*model.SessionDiseasesInput
	if exam.Done == true {
		diseasesinput, err = utils.GetSessionDiseases(symptoms, float64(session.Weight)/(float64(session.Height)/100.0*(float64(session.Height)/100.0)), session.AnteChirs, session.HereditaryDisease)
		if err != nil {
			return DiagnoseResponse{
				Code: 500,
				Err:  errors.New("error during getSessionDiseases"),
			}
		}
	}

	var autoAnswerOutput *model.AutoAnswer
	if exam.AutoAnswer != nil {
		auA, err := graphql.GetAutoAnswerByName(*exam.AutoAnswer)
		if err != nil {
			autoAnswerOutput = nil
		} else {
			autoAnswerOutput = &auA
		}
	}

	if len(symptomsinput) > 11 {
		exam.Done = true
	}

	_, err = graphql.UpdateSession(session.ID, model.UpdateSessionInput{
		Diseases:     diseasesinput,
		Symptoms:     symptomsinput,
		Medicine:     session.Medicine,
		LastQuestion: &session.LastQuestion,
		Logs:         logs,
		Alerts:       session.Alerts,
	})
	edgarlib.CheckError(err)
	return DiagnoseResponse{
		Done:       exam.Done,
		Question:   exam.Question,
		AutoAnswer: autoAnswerOutput,
		Code:       200,
		Err:        nil,
	}
}
