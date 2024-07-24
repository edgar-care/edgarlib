package double_auth

import (
	"context"
	"errors"
	"github.com/edgar-care/edgarlib/graphql"
	"github.com/edgar-care/edgarlib/graphql/server/model"
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
	gqlClient := graphql.CreateClient()
	var res model.DeviceConnect

	device, err := graphql.GetDeviceConnectById(context.Background(), gqlClient, id)
	if err != nil {
		return GetDeviceConnectByIdResponse{model.DeviceConnect{}, 400, errors.New("id does not correspond to a trust_device")}
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

func GetTrustDeviceConnect(id_user string) GetTrustDevicesConnectResponse {
	gqlClient := graphql.CreateClient()
	var res []model.DeviceConnect

	// Essayer de récupérer le patient par son id_user
	patient, err := graphql.GetPatientById(context.Background(), gqlClient, id_user)
	if err == nil {
		// Obtenir les périphériques de confiance associés au patient
		for _, trustDeviceID := range patient.GetPatientById.Trust_devices {
			device, err := graphql.GetDeviceConnectById(context.Background(), gqlClient, trustDeviceID)
			if err == nil {
				res = append(res, model.DeviceConnect{
					ID:          device.GetDeviceConnectById.Id,
					DeviceName:  device.GetDeviceConnectById.Device_name,
					IPAddress:   device.GetDeviceConnectById.Ip_address,
					Latitude:    device.GetDeviceConnectById.Latitude,
					Longitude:   device.GetDeviceConnectById.Longitude,
					Date:        device.GetDeviceConnectById.Date,
					TrustDevice: device.GetDeviceConnectById.Trust_device,
				})
			}
		}
		return GetTrustDevicesConnectResponse{res, 200, nil}
	}

	// Si la récupération du patient échoue, essayer de récupérer le docteur par son id_user
	doctor, err := graphql.GetDoctorById(context.Background(), gqlClient, id_user)
	if err != nil {
		return GetTrustDevicesConnectResponse{[]model.DeviceConnect{}, 400, errors.New("id does not correspond to a patient or doctor")}
	}

	// Obtenir les périphériques de confiance associés au docteur
	for _, trustDeviceID := range doctor.GetDoctorById.Trust_devices {
		device, err := graphql.GetDeviceConnectById(context.Background(), gqlClient, trustDeviceID)
		if err == nil {
			res = append(res, model.DeviceConnect{
				ID:          device.GetDeviceConnectById.Id,
				DeviceName:  device.GetDeviceConnectById.Device_name,
				IPAddress:   device.GetDeviceConnectById.Ip_address,
				Latitude:    device.GetDeviceConnectById.Latitude,
				Longitude:   device.GetDeviceConnectById.Longitude,
				Date:        device.GetDeviceConnectById.Date,
				TrustDevice: device.GetDeviceConnectById.Trust_device,
			})
		}
	}
	return GetTrustDevicesConnectResponse{res, 200, nil}
}
