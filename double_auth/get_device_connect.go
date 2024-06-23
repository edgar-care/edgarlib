package double_auth

import (
	"context"
	"errors"
	"github.com/edgar-care/edgarlib/graphql"
	"github.com/edgar-care/edgarlib/graphql/server/model"
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

func GetDeviceConnectById(id string, doctorId string) GetDeviceConnectByIdResponse {
	gqlClient := graphql.CreateClient()
	var res model.DeviceConnect

	device, err := graphql.GetDeviceConnectById(context.Background(), gqlClient, id)
	if err != nil {
		return GetDeviceConnectByIdResponse{model.DeviceConnect{}, 400, errors.New("id does not correspond to a slot")}
	}
	res = model.DeviceConnect{
		ID:          device.GetDeviceConnectById.Id,
		DeviceName:  device.GetDeviceConnectById.Device_name,
		IPAddress:   device.GetDeviceConnectById.Ip_address,
		Latitude:    device.GetDeviceConnectById.Latitude,
		Longitude:   device.GetDeviceConnectById.Longitude,
		Date:        device.GetDeviceConnectById.Date,
		TrustDevice: device.GetDeviceConnectById.Trust_device,
	}
	return GetDeviceConnectByIdResponse{res, 200, nil}
}

func GetDeviceConnect(patientId string) GetDevicesConnectResponse {
	gqlClient := graphql.CreateClient()
	var res []model.DeviceConnect

	_, err := graphql.GetPatientById(context.Background(), gqlClient, patientId)
	if err != nil {
		return GetDevicesConnectResponse{[]model.DeviceConnect{}, 400, errors.New("id does not correspond to a doctor")}
	}

	devices, err := graphql.GetDevicesConnect(context.Background(), gqlClient)
	if err != nil {
		return GetDevicesConnectResponse{[]model.DeviceConnect{}, 400, errors.New("invalid input: " + err.Error())}
	}

	for _, device := range devices.GetDevicesConnect {
		res = append(res, model.DeviceConnect{
			ID:          device.Id,
			DeviceName:  device.Device_name,
			IPAddress:   device.Ip_address,
			Latitude:    device.Latitude,
			Longitude:   device.Longitude,
			Date:        device.Date,
			TrustDevice: device.Trust_device,
		})
	}
	return GetDevicesConnectResponse{res, 200, nil}
}
