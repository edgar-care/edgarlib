package utils

import (
	"github.com/edgar-care/edgarlib/v2/exam"
	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
)

func CheckAnteDiseaseInSymptoms(session model.Session) (string, string, error) {
	var question string
	var questionSymptomName string

	for _, anteId := range session.AnteDiseases {
		if anteId != "" {
			ante, err := graphql.GetAnteDiseaseByID(anteId)
			if err != nil {
				return "", "", err
			}
			if ante.StillRelevant == true && len(ante.Symptoms) > 0 {
				for _, anteSymptomId := range ante.Symptoms {
					if anteSymptomId == "" {
						continue
					}
					anteSymptom, err := graphql.GetSymptomById(anteSymptomId)
					if err != nil {
						return "", "", err
					}
					if anteSymptom.Code != session.LastQuestion {
						if anteSymptom.QuestionAnte != "" {
							question = exam.AddDiscursiveConnector(anteSymptom.QuestionAnte)
						} else {
							question = exam.AddDiscursiveConnector("{{connecteur}}. Ressentez-vous " + anteSymptom.Name + " plus intensément récemment ?")
						}
						questionSymptomName = anteSymptom.Code
					}
					for _, sessionSymptom := range session.Symptoms {
						if anteSymptom.Code == sessionSymptom.Name || anteSymptom.Code == session.LastQuestion {
							question = ""
							questionSymptomName = ""
						}
					}
					if question != "" {
						return question, questionSymptomName, nil
					}
				}
			}
		}
	}
	return question, questionSymptomName, nil
}
