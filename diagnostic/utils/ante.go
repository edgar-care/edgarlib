package utils

import (
	"context"
	"github.com/edgar-care/edgarlib/graphql"
)

func CheckAnteDiseaseInSymptoms(session graphql.GetSessionByIdGetSessionByIdSession) (string, string, error) {
	gqlClient := graphql.CreateClient()
	var question string
	var questionSymptomName string

	for _, anteId := range session.Ante_diseases {
		if anteId != "" {
			ante, err := graphql.GetAnteDiseaseByID(context.Background(), gqlClient, anteId)
			if err != nil {
				return "", "", err
			}
			if ante.GetAnteDiseaseByID.Still_relevant == true && len(ante.GetGetAnteDiseaseByID().Symptoms) > 0 {
				for _, anteSymptomId := range ante.GetAnteDiseaseByID.Symptoms {
					if anteSymptomId == "" {
						continue
					}
					anteSymptom, err := graphql.GetSymptomById(context.Background(), gqlClient, anteSymptomId)
					if err != nil {
						return "", "", err
					}
					if anteSymptom.GetSymptomById.Code != session.Last_question {
						if anteSymptom.GetSymptomById.Question_ante != "" {
							question = anteSymptom.GetSymptomById.Question_ante
						} else {
							question = "Ressentez-vous " + anteSymptom.GetSymptomById.Name + " plus intensément récemment ?"
						}
						questionSymptomName = anteSymptom.GetSymptomById.Code
					}
					for _, sessionSymptom := range session.Symptoms {
						if anteSymptom.GetSymptomById.Code == sessionSymptom.Name || anteSymptom.GetSymptomById.Code == session.Last_question {
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
