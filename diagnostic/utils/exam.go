package utils

import (
	"github.com/edgar-care/edgarlib/v2/exam"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
)

type examRequestBody struct {
	Context []model.SessionSymptom `json:"context"`
}

type ExamResponseBody struct {
	Context    []model.SessionSymptom `json:"context"`
	Done       bool                   `json:"done"`
	Question   string                 `json:"question"`
	Symptoms   []string               `json:"symptoms"`
	AutoAnswer *string                `json:"auto_answer"`
	Alert      []string               `json:"alert"`
	Err        error                  `json:"err"`
}

func CallExam(context []model.SessionSymptom, imc float64, anteChirIds []string, hereditary_disease []string) ExamResponseBody {
	//var rBody = examRequestBody{
	//	Context: context,
	//}

	//var buf = new(bytes.Buffer)
	//err := json.NewEncoder(buf).Encode(rBody)
	//edgarlib.CheckError(err)
	//
	//resp, err := http.Post(os.Getenv("EXAM_URL"), "application/json", buf)
	//edgarlib.CheckError(err)
	//
	//var respBody examResponseBody
	//err = json.NewDecoder(resp.Body).Decode(&respBody)
	//edgarlib.CheckError(err)

	examr := exam.Exam(context, imc, anteChirIds, hereditary_disease)
	if examr.Err != nil {
		return ExamResponseBody{Err: examr.Err}
	}

	var respBody ExamResponseBody
	respBody.Context = examr.Context
	respBody.Done = examr.Done
	respBody.Question = examr.Question
	respBody.Symptoms = examr.Symptoms
	respBody.AutoAnswer = examr.AutoAnswer
	respBody.Alert = examr.Alert
	respBody.Err = nil

	return respBody
}
