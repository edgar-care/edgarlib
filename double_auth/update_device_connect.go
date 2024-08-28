package double_auth

import (
	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
)

type UpdateDeviceConnectInput struct {
	DeviceType string `json:"device_type"`
	Browser    string `json:"browser"`
	Ip         string `json:"ip_address"`
	City       string `json:"city"`
	Country    string `json:"country"`
	Date       int    `json:"date"`
}

type UpdateDeviceConnectResponse struct {
	DeviceConnect model.DeviceConnect
	Code          int
	Err           error
}

func UpdateDeviceConnect(input UpdateDeviceConnectInput, deviceId string) UpdateDeviceConnectResponse {

	updatedDevice, err := graphql.UpdateDeviceConnect(deviceId, model.UpdateDeviceConnectInput{
		DeviceType: &input.DeviceType,
		Browser:    &input.Browser,
		IPAddress:  &input.Ip,
		City:       &input.City,
		Country:    &input.Country,
		Date:       &input.Date,
	})
	if err != nil {
		return UpdateDeviceConnectResponse{
			DeviceConnect: model.DeviceConnect{},
			Code:          500,
			Err:           err,
		}
	}

	return UpdateDeviceConnectResponse{
		DeviceConnect: updatedDevice,
		Code:          200,
		Err:           nil,
	}
}
