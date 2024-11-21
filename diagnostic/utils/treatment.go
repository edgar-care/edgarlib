package utils

import (
	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
	"log"
)

func isSymptomTreated(symptom model.SessionSymptomInput, treatmentName string) bool {
	for _, treated := range symptom.Treated {
		if treated == treatmentName {
			return true
		}
	}
	return false
}

func CheckTreatments(symptoms []*model.SessionSymptomInput, medicines []string) ([]*model.SessionSymptomInput, error) {
	for _, medicineId := range medicines {
		log.Println(medicineId)
		medicine, err := graphql.GetMedicineByIDWithSymptoms(medicineId)
		log.Println(medicine, err)
		if err != nil {
			return nil, err
		}
		for _, symptomT := range medicine.Symptoms {
			for i, symptomSy := range symptoms {
				if symptomSy.Name == symptomT.Name && !isSymptomTreated(*symptomSy, symptomT.Name) {
					symptoms[i].Treated = append(symptoms[i].Treated, medicine.Name)
				}
			}
		}
	}
	return symptoms, nil
}
