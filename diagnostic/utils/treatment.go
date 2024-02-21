package utils

import (
	"context"
	"github.com/edgar-care/edgarlib/graphql"
)

func isSymptomTreated(symptom graphql.SessionSymptomInput, treatmentName string) bool {
	for _, treated := range symptom.Treated {
		if treated == treatmentName {
			return true
		}
	}
	return false
}

func CheckTreatments(symptoms []graphql.SessionSymptomInput, treatments []string) []graphql.SessionSymptomInput {
	gqlClient := graphql.CreateClient()
	for _, treatmentId := range treatments {
		treatment, _ := graphql.GetTreatmentByID(context.Background(), gqlClient, treatmentId)
		for _, symptomId := range treatment.GetTreatmentByID.Symptoms {
			symptomT, _ := graphql.GetSymptomById(context.Background(), gqlClient, symptomId)
			for i, symptomSy := range symptoms {
				if symptomSy.Name == symptomT.GetSymptomById.Code && !isSymptomTreated(symptomSy, symptomT.GetSymptomById.Code) {
					symptoms[i].Treated = append(symptoms[i].Treated, treatment.GetTreatmentByID.Name)
				}
			}
		}
	}
	//tester si symptoms est bien modifié
	return symptoms
}
