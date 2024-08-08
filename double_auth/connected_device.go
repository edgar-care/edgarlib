package double_auth

import (
	"errors"

	"github.com/edgar-care/edgarlib/graphql"
	"github.com/edgar-care/edgarlib/graphql/model"
)

type CreateDeviceConnectInput struct {
	DeviceName string  `json:"device_name"`
	Ip         string  `json:"ip"`
	Latitude   float64 `json:"latitude"`
	Longitude  float64 `json:"longitude"`
	Date       int     `json:"date"`
}

type CreateDeviceConnectResponse struct {
	DeviceConnect model.DeviceConnect
	Patient       model.Patient
	Code          int
	Err           error
}

func CreateDeviceConnect(input CreateDeviceConnectInput, patientId string) CreateDeviceConnectResponse {
	device, err := graphql.CreateDeviceConnect(model.CreateDeviceConnectInput{
		DeviceName:  input.DeviceName,
		IPAddress:   input.Ip,
		Latitude:    input.Latitude,
		Longitude:   input.Longitude,
		Date:        input.Date,
		TrustDevice: false,
	})
	if err != nil {
		return CreateDeviceConnectResponse{DeviceConnect: model.DeviceConnect{}, Patient: model.Patient{}, Code: 400, Err: errors.New("unable  (check if you share all information)")}
	}

	patient, err := graphql.GetPatientById(patientId)
	if err != nil {
		return CreateDeviceConnectResponse{DeviceConnect: model.DeviceConnect{}, Patient: model.Patient{}, Code: 400, Err: errors.New("id does not correspond to a doctor")}
	}

	_, err = graphql.UpdatePatient(patientId, model.UpdatePatientInput{DeviceConnect: append(patient.DeviceConnect, &device.ID)})
	if err != nil {
		return CreateDeviceConnectResponse{DeviceConnect: model.DeviceConnect{}, Patient: model.Patient{}, Code: 400, Err: errors.New("update failed" + err.Error())}
	}

	return CreateDeviceConnectResponse{
		DeviceConnect: device,
		Code:          201,
		Err:           nil,
	}
}
