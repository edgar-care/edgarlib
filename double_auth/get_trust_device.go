package double_auth

import (
	"errors"
	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
)

type GetTrustDeviceConnectByIdResponse struct {
	DeviceConnect model.DeviceConnect
	Code          int
	Err           error
}

type GetTrustDevicesConnectResponse struct {
	DevicesConnect []model.DeviceConnect
	Code           int
	Err            error
}

func GetTrustDeviceConnectById(id string) GetDeviceConnectByIdResponse {
	device, err := graphql.GetDeviceConnectById(id)
	if err != nil {
		return GetDeviceConnectByIdResponse{model.DeviceConnect{}, 400, errors.New("id does not correspond to a trust_device")}
	}

	return GetDeviceConnectByIdResponse{device, 200, nil}
}

func GetTrustDeviceConnect(id_user string) GetTrustDevicesConnectResponse {
	var res []model.DeviceConnect

	patient, err := graphql.GetPatientById(id_user)
	if err == nil {
		for _, trustDeviceID := range patient.TrustDevices {
			device, err := graphql.GetDeviceConnectById(*trustDeviceID)
			if err == nil {
				res = append(res, device)
			}
		}
		return GetTrustDevicesConnectResponse{res, 200, nil}
	}

	doctor, err := graphql.GetDoctorById(id_user)
	if err != nil {
		return GetTrustDevicesConnectResponse{[]model.DeviceConnect{}, 400, errors.New("id does not correspond to a patient or doctor")}
	}

	for _, trustDeviceID := range doctor.TrustDevices {
		device, err := graphql.GetDeviceConnectById(*trustDeviceID)
		if err == nil {
			res = append(res, device)
		}
	}
	return GetTrustDevicesConnectResponse{res, 200, nil}
}
