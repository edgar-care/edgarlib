package utils

import (
	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
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
		medicine, err := graphql.GetMedicineByID(medicineId)
		if err != nil {
			return nil, err
		}
		for _, symptomId := range medicine.TreatedSymptoms {
			symptomT, err := graphql.GetSymptomById(symptomId)
			if err != nil {
				return nil, err
			}
			for i, symptomSy := range symptoms {
				if symptomSy.Name == symptomT.Code && !isSymptomTreated(*symptomSy, symptomT.Code) {
					symptoms[i].Treated = append(symptoms[i].Treated, medicine.Name)
				}
			}
		}
	}
	return symptoms, nil
}
