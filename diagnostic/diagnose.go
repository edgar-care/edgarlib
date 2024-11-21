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
	Done     bool
	Question string
	//AutoAnswer *model.AutoAnswer
	Code int
	Err  error
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

	var newSymptoms utils.NlpResponseBody

	symptoms := extractSymptomsFromSession(session)

	updateLastLogAnswer(&session, sentence)

	questionSymptom := getLastQuestionSymptom(session)

	//if len(questionSymptom) > 0 {
	var errCode int
	newSymptoms, errCode = processSymptoms(&session, sentence, questionSymptom, autoAnswer)
	if errCode != 200 {
		return DiagnoseResponse{
			Code: errCode,
			Err:  errors.New("NLP error, please try again"),
		}
	}
	//}

	symptoms = updateSymptomsWithNewData(symptoms, newSymptoms)

	symptomsInput := prepareSymptomsInput(symptoms)

	exam := processExam(&session, symptomsInput, symptoms)
	if exam.Err != nil {
		return DiagnoseResponse{
			Code: 500,
			Err:  exam.Err,
		}
	}

	logsInput := updateLogs(session, sentence, &exam)

	var diseasesInput []*model.SessionDiseasesInput
	if exam.Done {
		diseasesInput, err = updateDiseases(symptoms, session)
		if err != nil {
			return DiagnoseResponse{
				Code: 500,
				Err:  errors.New("error during getSessionDiseases"),
			}
		}
	}

	//var autoAnswerOutput *model.AutoAnswer
	//if exam.AutoAnswer != nil {
	//	autoAnswerOutput = getAutoAnswerOutput(exam)
	//}

	if len(symptomsInput) > 11 {
		exam.Done = true
	}

	_, err = graphql.UpdateSession(session.ID, model.UpdateSessionInput{
		Diseases:     diseasesInput,
		Symptoms:     symptomsInput,
		Medicine:     session.Medicine,
		LastQuestion: &session.LastQuestion,
		Logs:         logsInput,
		Alerts:       session.Alerts,
	})
	edgarlib.CheckError(err)
	return DiagnoseResponse{
		Done:     exam.Done,
		Question: exam.Question,
		//AutoAnswer: autoAnswerOutput,
		Code: 200,
		Err:  nil,
	}
}

//func getAutoAnswerOutput(exam utils.ExamResponseBody) *model.AutoAnswer {
//	var autoAnswerOutput *model.AutoAnswer
//	auA, err := graphql.GetAutoAnswerByName(*exam.AutoAnswer)
//	if err != nil {
//		autoAnswerOutput = nil
//	} else {
//		autoAnswerOutput = &auA
//	}
//	return autoAnswerOutput
//}

func extractSymptomsFromSession(session model.Session) []model.SessionSymptom {
	var symptoms []model.SessionSymptom
	for _, s := range session.Symptoms {
		symptoms = append(symptoms, model.SessionSymptom{
			Name:     s.Name,
			Presence: s.Presence,
			Duration: s.Duration,
			Treated:  s.Treated,
		})
	}
	return symptoms
}

func updateLastLogAnswer(session *model.Session, sentence string) {
	if len(session.Logs) > 0 {
		session.Logs[len(session.Logs)-1].Answer = sentence
	}
}

func getLastQuestionSymptom(session model.Session) []string {
	//if session.LastQuestion == "" || session.LastQuestion == "describe symptoms" || session.LastQuestion == "describe medicines" {
	if session.LastQuestion == "" {
		return []string{}
	}
	return []string{session.LastQuestion}
}

func processSymptoms(session *model.Session, sentence string, questionSymptom []string, autoAnswer *AutoAnswerinfo) (utils.NlpResponseBody, int) {
	//var newSymptoms utils.NlpResponseBody
	//if autoAnswer != nil {
	//	processAutoAnswerSymptoms(autoAnswer, questionSymptom, &newSymptoms)
	//} else {
	isMedicine := session.LastQuestion == "describe medicines"
	durSymptom := getDurationSymptom(session)
	return utils.CallNlp(sentence, questionSymptom, durSymptom, isMedicine)
	//}
	//return newSymptoms, 200
}

func getDurationSymptom(session *model.Session) *string {
	if session.LastQuestion != "" && strings.Split(session.LastQuestion, " | ")[0] == "duration" {
		return &strings.Split(session.LastQuestion, " | ")[1]
	}
	return nil
}

//func processAutoAnswerSymptoms(autoAnswer *AutoAnswerinfo, questionSymptom []string, newSymptoms *utils.NlpResponseBody) {
//	autoA, _ := graphql.GetAutoAnswerByName(autoAnswer.Name)
//	if autoA.Name == "Oui / Non / Ne sais pas" {
//		if len(questionSymptom) > 0 {
//			if autoAnswer.Values[0] == "Oui." {
//				p := true
//				newSymptoms.Context = append(newSymptoms.Context, utils.Symptom{Name: questionSymptom[0], Present: &p})
//			} else if autoAnswer.Values[0] == "Non." {
//				p := false
//				newSymptoms.Context = append(newSymptoms.Context, utils.Symptom{Name: questionSymptom[0], Present: &p})
//			} else if autoAnswer.Values[0] == "Je ne sais pas." {
//				var p *bool
//				p = nil
//				newSymptoms.Context = append(newSymptoms.Context, utils.Symptom{Name: questionSymptom[0], Present: p})
//			}
//		}
//	}
//}

func updateSymptomsWithNewData(symptoms []model.SessionSymptom, newSymptoms utils.NlpResponseBody) []model.SessionSymptom {
	for _, s := range newSymptoms.Context {
		if pres, ite := nameInList(s, symptoms); pres {
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
	return symptoms
}

func prepareSymptomsInput(symptoms []model.SessionSymptom) []*model.SessionSymptomInput {
	var symptomsInput []*model.SessionSymptomInput
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
		symptomsInput = append(symptomsInput, &ns)
	}
	return symptomsInput
}

func processExam(session *model.Session, symptomsInput []*model.SessionSymptomInput, symptoms []model.SessionSymptom) utils.ExamResponseBody {
	var exam utils.ExamResponseBody
	var anteSymptom string
	var err error

	if len(symptomsInput) > 0 {
		symptomsInput, exam.Question, session.LastQuestion = utils.CheckSymptomDuration(symptomsInput, session.LastQuestion)
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
		exam = utils.CallExam(symptoms, float64(session.Weight)/(float64(session.Height)/100.0*(float64(session.Height)/100.0)), session.HereditaryDisease)
		if exam.Err != nil {
			return exam
		}

		if len(exam.Alert) > 0 {
			for _, alert := range exam.Alert {
				session.Alerts = append(session.Alerts, alert)
			}
		}

		//if len(session.Medicine) > 0 {
		//	symptomsInput, err = utils.CheckTreatments(symptomsInput, session.Medicine)
		//	if err != nil {
		//		exam.Err = errors.New("error during checkTreatment")
		//		return exam
		//	}
		//}

		if len(session.MedicalAntecedents) > 0 {
			var anteSymptomQuestion string
			anteSymptomQuestion, anteSymptom, err = utils.CheckAnteDiseaseInSymptoms(*session)
			if err != nil {
				exam.Err = errors.New("error during checkTreatment")
				return exam
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

	return exam
}

func updateLogs(session model.Session, sentence string, exam *utils.ExamResponseBody) []*model.LogsInput {
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
	return logs
}

func updateDiseases(symptoms []model.SessionSymptom, session model.Session) ([]*model.SessionDiseasesInput, error) {
	diseasesInput, err := utils.GetSessionDiseases(symptoms, float64(session.Weight)/(float64(session.Height)/100.0*(float64(session.Height)/100.0)), session.HereditaryDisease)
	if err != nil {
		return nil, err
	}
	return diseasesInput, nil
}
