package double_auth

import (
	"errors"
	"github.com/edgar-care/edgarlib/graphql/model"

	"github.com/edgar-care/edgarlib/black_list"
	"github.com/edgar-care/edgarlib/graphql"
)

type DeleteDeviceConnectResponse struct {
	Deleted bool
	Code    int
	Err     error
}

func remElement(slice []*string, element *string) []*string {
	var result []*string
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

	_, err := graphql.GetDeviceConnectById(DeviceId)
	if err != nil {
		return DeleteDeviceConnectResponse{Deleted: false, Code: 400, Err: errors.New("id does not correspond to a slot")}
	}

	deleted, err := graphql.DeleteDeviceConnect(DeviceId)
	if err != nil {
		return DeleteDeviceConnectResponse{Deleted: false, Code: 500, Err: errors.New("error while deleting slot: " + err.Error())}
	}

	patient, err := graphql.GetPatientById(patientId)
	if err != nil {
		return DeleteDeviceConnectResponse{Deleted: false, Code: 400, Err: errors.New("id does not correspond to a doctor")}
	}

	_, err = graphql.UpdatePatient(patientId, model.UpdatePatientInput{DeviceConnect: remElement(patient.DeviceConnect, &DeviceId)})
	if err != nil {
		return DeleteDeviceConnectResponse{Deleted: false, Code: 500, Err: errors.New("error updating patient: " + err.Error())}
	}

	blacklist := black_list.UpdateBlackList(patientId)
	if blacklist.Err != nil {
		return DeleteDeviceConnectResponse{Deleted: false, Code: 500, Err: errors.New("error updating patient: " + blacklist.Err.Error())}
	}

	return DeleteDeviceConnectResponse{
		Deleted: deleted,
		Code:    200,
		Err:     nil}
}
