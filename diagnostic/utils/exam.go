package utils

import (
	"bytes"
	"encoding/json"
	"github.com/edgar-care/edgarlib/graphql/server/model"
	"net/http"
	"os"

	"github.com/edgar-care/edgarlib"
)

type examRequestBody struct {
	Context []model.SessionSymptom `json:"context"`
}

type examResponseBody struct {
	Context  []model.SessionSymptom `json:"context"`
	Done     bool                   `json:"done"`
	Question string                 `json:"question"`
	Symptoms []string               `json:"symptoms"`
	Alert    []string               `json:"alert"`
}

func CallExam(context []model.SessionSymptom) examResponseBody {
	var rBody = examRequestBody{
		Context: context,
	}

	var buf = new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(rBody)
	edgarlib.CheckError(err)

	resp, err := http.Post(os.Getenv("EXAM_URL"), "application/json", buf)
	edgarlib.CheckError(err)

	var respBody examResponseBody
	err = json.NewDecoder(resp.Body).Decode(&respBody)
	edgarlib.CheckError(err)

	return respBody
}
