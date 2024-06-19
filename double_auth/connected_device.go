package double_auth

import (
	"context"
	"errors"

	"github.com/edgar-care/edgarlib/graphql"
	"github.com/edgar-care/edgarlib/graphql/server/model"
)

type CreateDeviceConnectInput struct {
	DeviceName string  `json:"device_name"`
	Ip         string  `json:"ip"`
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
	Date       int     `json:"date"`
	//TrustDevice bool   `json:"trust_device"`
}

type CreateDeviceConnectResponse struct {
	DeviceConnect model.DeviceConnect
	Patient       model.Patient
	Code          int
	Err           error
}

func CreateDeviceConnect(input CreateDeviceConnectInput, patientId string) CreateDeviceConnectResponse {
	gqlClient := graphql.CreateClient()

	device, err := graphql.CreateDeviceConnect(context.Background(), gqlClient, input.DeviceName, input.Ip, input.Latitude, input.Longitude, input.Date, false)
	if err != nil {
		return CreateDeviceConnectResponse{DeviceConnect: model.DeviceConnect{}, Patient: model.Patient{}, Code: 400, Err: errors.New("unable  (check if you share all information)")}
	}

	patient, err := graphql.GetPatientById(context.Background(), gqlClient, patientId)
	if err != nil {
		return CreateDeviceConnectResponse{DeviceConnect: model.DeviceConnect{}, Patient: model.Patient{}, Code: 400, Err: errors.New("id does not correspond to a doctor")}
	}

	_, err = graphql.UpdatePatient(context.Background(), gqlClient, patientId, patient.GetPatientById.Email, patient.GetPatientById.Password, patient.GetPatientById.Medical_info_id, patient.GetPatientById.Rendez_vous_ids, patient.GetPatientById.Document_ids, patient.GetPatientById.Treatment_follow_up_ids, patient.GetPatientById.Chat_ids, append(patient.GetPatientById.Device_connect, device.CreateDeviceConnect.Id))
	if err != nil {
		return CreateDeviceConnectResponse{DeviceConnect: model.DeviceConnect{}, Patient: model.Patient{}, Code: 400, Err: errors.New("update failed" + err.Error())}
	}

	return CreateDeviceConnectResponse{
		DeviceConnect: model.DeviceConnect{
			ID:          device.CreateDeviceConnect.Id,
			DeviceName:  device.CreateDeviceConnect.Device_name,
			IPAddress:   device.CreateDeviceConnect.Ip_address,
			Latitude:    device.CreateDeviceConnect.Latitude,
			Longitude:   device.CreateDeviceConnect.Longitude,
			Date:        device.CreateDeviceConnect.Date,
			TrustDevice: device.CreateDeviceConnect.Trust_device,
		},
		Code: 200,
		Err:  nil,
	}
}
