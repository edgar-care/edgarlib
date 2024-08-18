package exam

import (
	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
)

func isAlertPresent(context []model.SessionSymptom, symptom string) bool {
	for _, e := range context {
		if e.Presence != 0 {
			if e.Name == symptom && e.Presence == 1 {
				return true
			}
		}
	}
	return false
}

func coverAlert(context []model.SessionSymptom, alert model.Alert) string {
	present := true
	for _, symptom := range alert.Symptoms {
		presence := isAlertPresent(context, symptom)
		if presence == false {
			present = false
		}
	}
	if present == true {
		return alert.ID
	} else {
		return ""
	}
}

func CheckAlerts(patientContext []model.SessionSymptom) ([]string, error) {
	alerts, err := graphql.GetAlerts(nil)
	if err != nil {
		return []string{}, nil
	}
	var present []string
	for _, alert := range alerts {
		tmp := coverAlert(patientContext, alert)
		if tmp != "" {
			present = append(present, tmp)
		}
	}
	return present, nil
}
