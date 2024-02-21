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

func CheckTreatments(symptoms []graphql.SessionSymptomInput, medicines []string) []graphql.SessionSymptomInput {
	gqlClient := graphql.CreateClient()
	for _, medicineId := range medicines {
		medicine, _ := graphql.GetMedicineByID(context.Background(), gqlClient, medicineId)
		for _, symptomId := range medicine.GetMedicineByID.Treated_symptoms {
			symptomT, _ := graphql.GetSymptomById(context.Background(), gqlClient, symptomId)
			for i, symptomSy := range symptoms {
				if symptomSy.Name == symptomT.GetSymptomById.Code && !isSymptomTreated(symptomSy, symptomT.GetSymptomById.Code) {
					symptoms[i].Treated = append(symptoms[i].Treated, medicine.GetMedicineByID.Name)
				}
			}
		}
	}
	return symptoms
}
