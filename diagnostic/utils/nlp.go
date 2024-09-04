package utils

import (
	"bytes"
	"encoding/json"
	"github.com/edgar-care/edgarlib/v2"
	"net/http"
	"os"
)

type nlpRequestBody struct {
	Input    string   `json:"input"`
	Symptoms []string `json:"symptoms"`
	IsTime   bool     `json:"isTime"`
}

type NlpResponseBody struct {
	Context []Symptom `json:"context"`
}

func CallNlp(sentence string, symptoms []string, durSymptom *string) (NlpResponseBody, int) {
	var rBody = nlpRequestBody{
		Input:    sentence,
		Symptoms: symptoms,
		IsTime:   false,
	}

	var respBody NlpResponseBody

	if sentence == "" {
		return respBody, 200
	}

	if durSymptom != nil {
		if len(symptoms) != 0 {
			symptoms[0] = *durSymptom
		}
		rBody.IsTime = true
	}

	var buf = new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(rBody)
	edgarlib.CheckError(err)

	req, err := http.NewRequest("POST", os.Getenv("NLP_URL"), buf)
	req.Header.Add(os.Getenv("NLP_KEY"), os.Getenv("NLP_KEY_VALUE"))
	client := http.Client{}
	resp, err := client.Do(req)
	edgarlib.CheckError(err)

	err = json.NewDecoder(resp.Body).Decode(&respBody)
	edgarlib.CheckError(err)

	return respBody, resp.StatusCode
}

func WakeNlpUp() {
	var rBody = nlpRequestBody{
		Input:    "wake up",
		Symptoms: []string{},
		IsTime:   false,
	}

	var buf = new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(rBody)
	edgarlib.CheckError(err)

	go http.Post(os.Getenv("NLP_URL"), "application/json", buf)
}
