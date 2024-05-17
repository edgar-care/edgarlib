package utils

import (
	"github.com/edgar-care/edgarlib/exam"
	"github.com/edgar-care/edgarlib/graphql/server/model"
)

type examRequestBody struct {
	Context []model.SessionSymptom `json:"context"`
}

type ExamResponseBody struct {
	Context  []model.SessionSymptom `json:"context"`
	Done     bool                   `json:"done"`
	Question string                 `json:"question"`
	Symptoms []string               `json:"symptoms"`
	Alert    []string               `json:"alert"`
	Err      error                  `json:"err"`
}

func CallExam(context []model.SessionSymptom, imc float64) ExamResponseBody {
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

	examr := exam.Exam(context, imc)
	if examr.Err != nil {
		return ExamResponseBody{Err: examr.Err}
	}

	var respBody ExamResponseBody
	respBody.Context = examr.Context
	respBody.Done = examr.Done
	respBody.Question = examr.Question
	respBody.Symptoms = examr.Symptoms
	respBody.Alert = examr.Alert
	respBody.Err = nil

	return respBody
}
