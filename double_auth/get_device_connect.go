package double_auth

import (
	"errors"
	"github.com/edgar-care/edgarlib/graphql"
	"github.com/edgar-care/edgarlib/graphql/model"
)

type GetDeviceConnectByIdResponse struct {
	DeviceConnect model.DeviceConnect
	Code          int
	Err           error
}

type GetDevicesConnectResponse struct {
	DevicesConnect []model.DeviceConnect
	Code           int
	Err            error
}

func GetDeviceConnectById(id string) GetDeviceConnectByIdResponse {
	device, err := graphql.GetDeviceConnectById(id)
	if err != nil {
		return GetDeviceConnectByIdResponse{model.DeviceConnect{}, 400, errors.New("id does not correspond to a slot")}
	}

	return GetDeviceConnectByIdResponse{device, 200, nil}
}

func GetDeviceConnect(patientId string) GetDevicesConnectResponse {
	_, err := graphql.GetPatientById(patientId)
	if err != nil {
		return GetDevicesConnectResponse{[]model.DeviceConnect{}, 400, errors.New("id does not correspond to a doctor")}
	}

	devices, err := graphql.GetDevicesConnect(nil)
	if err != nil {
		return GetDevicesConnectResponse{[]model.DeviceConnect{}, 400, errors.New("invalid input: " + err.Error())}
	}

	return GetDevicesConnectResponse{devices, 200, nil}
}
