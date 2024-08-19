package double_auth

import (
	"errors"
	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
)

func RemoveDoubleAuthMethod(methodToRemove string, ownerId string) error {
	patient, err := graphql.GetPatientById(ownerId)
	if err == nil {
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
			_, err = graphql.UpdatePatient(ownerId, model.UpdatePatientInput{
				DoubleAuthMethodsID: &empty,
			})
			if err != nil {
				return errors.New("failed to remove double_auth_methods_id from patient")
			}
		}

		return nil
	}

	doctor, err := graphql.GetDoctorById(ownerId)
	if err == nil {
		if doctor.DoubleAuthMethodsID == nil || *doctor.DoubleAuthMethodsID == "" {
			return errors.New("doctor does not have double_auth_methods_id")
		}

		doubleAuth, err := graphql.GetDoubleAuthById(*doctor.DoubleAuthMethodsID)
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
			_, err = graphql.UpdateDoctor(ownerId, model.UpdateDoctorInput{
				DoubleAuthMethodsID: &empty,
			})
			if err != nil {
				return errors.New("failed to remove double_auth_methods_id from doctor")
			}
		}

		return nil
	}

	return errors.New("id does not correspond to a patient or doctor")
}
