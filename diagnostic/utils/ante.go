package utils

import (
	"github.com/edgar-care/edgarlib/v2/exam"
	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
)

func CheckAnteDiseaseInSymptoms(session model.Session) (string, string, error) {
	var question string
	var questionSymptomName string

	for _, anteId := range session.MedicalAntecedents {
		if anteId != "" {
			ante, err := graphql.GetMedicalAntecedentsById(anteId)
			if err != nil {
				return "", "", err
			}
			for _, anteSymptomName := range ante.Symptoms {
				if anteSymptomName != session.LastQuestion {
					anteSymptom, err := graphql.GetSymptomByName(anteSymptomName)
					if err != nil {
						continue
					}
					if anteSymptom.QuestionAnte != "" {
						question = exam.AddDiscursiveConnector(anteSymptom.QuestionAnte)
					} else {
						question = exam.AddDiscursiveConnector("{{connecteur}}. Ressentez-vous " + anteSymptom.Name + " plus intensément récemment ?")
					}
					questionSymptomName = anteSymptom.Name
				}
				for _, sessionSymptom := range session.Symptoms {
					if anteSymptomName == sessionSymptom.Name || anteSymptomName == session.LastQuestion {
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
	return question, questionSymptomName, nil
}
