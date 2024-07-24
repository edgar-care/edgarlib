package double_auth

import (
	"context"
	"errors"

	"github.com/edgar-care/edgarlib/black_list"
	"github.com/edgar-care/edgarlib/graphql"
)

type DeleteDeviceConnectResponse struct {
	Deleted bool
	Code    int
	Err     error
}

func remElement(slice []string, element string) []string {
	var result []string
	for _, v := range slice {
		if v != element {
			result = append(result, v)
		}
	}
	return result
}

func DeleteDeviceConnect(DeviceId string, patientId string) DeleteDeviceConnectResponse {
	if DeviceId == "" {
		return DeleteDeviceConnectResponse{Deleted: false, Code: 400, Err: errors.New("slot id is required")}
	}
	gqlClient := graphql.CreateClient()

	_, err := graphql.GetDeviceConnectById(context.Background(), gqlClient, DeviceId)
	if err != nil {
		return DeleteDeviceConnectResponse{Deleted: false, Code: 400, Err: errors.New("id does not correspond to a slot")}
	}

	deleted, err := graphql.DeleteDeviceConnect(context.Background(), gqlClient, DeviceId)
	if err != nil {
		return DeleteDeviceConnectResponse{Deleted: false, Code: 500, Err: errors.New("error while deleting slot: " + err.Error())}
	}

	patient, err := graphql.GetPatientById(context.Background(), gqlClient, patientId)
	if err != nil {
		return DeleteDeviceConnectResponse{Deleted: false, Code: 400, Err: errors.New("id does not correspond to a doctor")}
	}

	_, err = graphql.UpdatePatient(context.Background(), gqlClient, patientId, patient.GetPatientById.Email, patient.GetPatientById.Password, patient.GetPatientById.Medical_info_id, patient.GetPatientById.Rendez_vous_ids, patient.GetPatientById.Document_ids, patient.GetPatientById.Treatment_follow_up_ids, patient.GetPatientById.Chat_ids, remElement(patient.GetPatientById.Device_connect, DeviceId), patient.GetPatientById.Double_auth_methods_id, patient.GetPatientById.Trust_devices, patient.GetPatientById.Status)
	if err != nil {
		return DeleteDeviceConnectResponse{Deleted: false, Code: 500, Err: errors.New("error updating patient: " + err.Error())}
	}

	blacklist := black_list.UpdateBlackList(patientId)
	if blacklist.Err != nil {
		return DeleteDeviceConnectResponse{Deleted: false, Code: 500, Err: errors.New("error updating patient: " + blacklist.Err.Error())}
	}

	return DeleteDeviceConnectResponse{
		Deleted: deleted.DeleteDeviceConnect,
		Code:    200,
		Err:     nil}
}
