package double_auth

import (
	"errors"
	"github.com/edgar-care/edgarlib/graphql"
	"github.com/edgar-care/edgarlib/graphql/model"
)

func RemoveDoubleAuthMethod(methodToRemove string, patientId string) error {
	patient, err := graphql.GetPatientById(patientId)
	if err != nil {
		return errors.New("unable to fetch patient")
	}

	if patient.DoubleAuthMethodsID == nil || *patient.DoubleAuthMethodsID == "" {
		return errors.New("patient does not have double_auth_methods_id")
	}

	doubleAuth, err := graphql.GetDoubleAuthById(*patient.DoubleAuthMethodsID)
	if err != nil {
		return errors.New("unable to fetch double_auth")
	}

	methods := doubleAuth.Methods
	found := false
	for i, m := range methods {
		if m == methodToRemove {

			methods = append(methods[:i], methods[i+1:]...)
			found = true
			break
		}
	}

	if !found {
		return errors.New("method to remove not found in current methods")
	}

	_, err = graphql.UpdateDoubleAuth(doubleAuth.ID, model.UpdateDoubleAuthInput{
		Methods: methods,
	})
	if err != nil {
		return errors.New("failed to update double_auth")
	}

	if len(methods) == 0 {
		empty := ""
		_, err = graphql.UpdatePatient(patientId, model.UpdatePatientInput{
			DoubleAuthMethodsID: &empty,
		})
		if err != nil {
			return errors.New("failed to remove double_auth_methods_id from patient")
		}
	}

	return nil
}
