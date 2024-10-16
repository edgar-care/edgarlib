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
			ante, err := graphql.GetAnteDiseaseByIDWithSymptoms(anteId)
			if err != nil {
				return "", "", err
			}
			if ante.StillRelevant == true {
				for _, anteSymptom := range ante.Symptomsclear {
					if anteSymptom.Name != session.LastQuestion {
						if anteSymptom.QuestionAnte != "" {
							question = exam.AddDiscursiveConnector(anteSymptom.QuestionAnte)
						} else {
							question = exam.AddDiscursiveConnector("{{connecteur}}. Ressentez-vous " + anteSymptom.Name + " plus intensément récemment ?")
						}
						questionSymptomName = anteSymptom.Name
					}
					for _, sessionSymptom := range session.Symptoms {
						if anteSymptom.Name == sessionSymptom.Name || anteSymptom.Name == session.LastQuestion {
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
