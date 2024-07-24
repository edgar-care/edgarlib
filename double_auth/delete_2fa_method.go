package double_auth

import (
	"context"
	"errors"
	"github.com/edgar-care/edgarlib/graphql"
)

func RemoveDoubleAuthMethod(methodToRemove string, patientId string) error {
	gqlClient := graphql.CreateClient()

	patient, err := graphql.GetPatientById(context.Background(), gqlClient, patientId)
	if err != nil {
		return errors.New("unable to fetch patient")
	}

	if patient.GetPatientById.Double_auth_methods_id == "" {
		return errors.New("patient does not have double_auth_methods_id")
	}

	doubleAuth, err := graphql.GetDoubleAuthById(context.Background(), gqlClient, patient.GetPatientById.Double_auth_methods_id)
	if err != nil {
		return errors.New("unable to fetch double_auth")
	}

	methods := doubleAuth.GetDoubleAuthById.Methods
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

	_, err = graphql.UpdateDoubleAuth(context.Background(), gqlClient, doubleAuth.GetDoubleAuthById.Id, methods, doubleAuth.GetDoubleAuthById.Secret, doubleAuth.GetDoubleAuthById.Url, doubleAuth.GetDoubleAuthById.Trust_device_id)
	if err != nil {
		return errors.New("failed to update double_auth")
	}

	if len(methods) == 0 {
		_, err = graphql.UpdatePatient(context.Background(), gqlClient, patientId, patient.GetPatientById.Email, patient.GetPatientById.Password, patient.GetPatientById.Medical_info_id, patient.GetPatientById.Rendez_vous_ids, patient.GetPatientById.Document_ids, patient.GetPatientById.Treatment_follow_up_ids, patient.GetPatientById.Chat_ids, patient.GetPatientById.Device_connect, "", patient.GetPatientById.Trust_devices, patient.GetPatientById.Status)
		if err != nil {
			return errors.New("failed to remove double_auth_methods_id from patient")
		}
	}

	return nil
}
