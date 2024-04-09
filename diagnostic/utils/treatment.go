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

func CheckTreatments(symptoms []graphql.SessionSymptomInput, medicines []string) ([]graphql.SessionSymptomInput, error) {
	gqlClient := graphql.CreateClient()
	for _, medicineId := range medicines {
		medicine, err := graphql.GetMedicineByID(context.Background(), gqlClient, medicineId)
		if err != nil {
			return nil, err
		}
		for _, symptomId := range medicine.GetMedicineByID.Treated_symptoms {
			symptomT, err := graphql.GetSymptomById(context.Background(), gqlClient, symptomId)
			if err != nil {
				return nil, err
			}
			for i, symptomSy := range symptoms {
				if symptomSy.Name == symptomT.GetSymptomById.Code && !isSymptomTreated(symptomSy, symptomT.GetSymptomById.Code) {
					symptoms[i].Treated = append(symptoms[i].Treated, medicine.GetMedicineByID.Name)
				}
			}
		}
	}
	return symptoms, nil
}
