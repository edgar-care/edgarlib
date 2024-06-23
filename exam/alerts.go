package exam

import (
	"context"
	"github.com/edgar-care/edgarlib/graphql"
	"github.com/edgar-care/edgarlib/graphql/server/model"
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

func coverAlert(context []model.SessionSymptom, alert graphql.GetAlertsGetAlertsAlert) string {
	present := true
	for _, symptom := range alert.Symptoms {
		presence := isAlertPresent(context, symptom)
		if presence == false {
			present = false
		}
	}
	if present == true {
		return alert.Id
	} else {
		return ""
	}
}

func CheckAlerts(patientContext []model.SessionSymptom) ([]string, error) {
	gqlClient := graphql.CreateClient()
	alerts, err := graphql.GetAlerts(context.Background(), gqlClient)
	if err != nil {
		return []string{}, nil
	}
	var present []string
	for _, alert := range alerts.GetAlerts {
		tmp := coverAlert(patientContext, alert)
		if tmp != "" {
			present = append(present, tmp)
		}
	}
	return present, nil
}
