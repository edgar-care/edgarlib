package utils

import (
	"bytes"
	"encoding/json"
	"github.com/edgar-care/edgarlib"
	"net/http"
	"os"
)

type nlpRequestBody struct {
	Input    string   `json:"input"`
	Symptoms []string `json:"symptoms"`
}

type nlpResponseBody struct {
	Context []Symptom `json:"context"`
}

func CallNlp(sentence string, symptoms []string) nlpResponseBody {
	var rBody = nlpRequestBody{
		Input:    sentence,
		Symptoms: symptoms,
	}

	var buf = new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(rBody)
	edgarlib.CheckError(err)

	resp, err := http.Post(os.Getenv("NLP_URL"), "application/json", buf)
	edgarlib.CheckError(err)

	var respBody nlpResponseBody
	err = json.NewDecoder(resp.Body).Decode(&respBody)
	edgarlib.CheckError(err)

	return respBody
}

func WakeNlpUp() {
	var rBody = nlpRequestBody{
		Input:    "wake up",
		Symptoms: []string{},
	}

	var buf = new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(rBody)
	edgarlib.CheckError(err)

	go http.Post(os.Getenv("NLP_URL"), "application/json", buf)
}
