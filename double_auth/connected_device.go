package double_auth

import (
	"errors"

	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
)

type CreateDeviceConnectInput struct {
	DeviceType string `json:"device_type"`
	Browser    string `json:"browser"`
	Ip         string `json:"ip"`
	City       string `json:"city"`
	Country    string `json:"country"`
	Date       int    `json:"date"`
}

type CreateDeviceConnectResponse struct {
	DeviceConnect model.DeviceConnect
	Code          int
	Err           error
}

func CreateDeviceConnect(input CreateDeviceConnectInput, ownerId string) CreateDeviceConnectResponse {
	device, err := graphql.CreateDeviceConnect(model.CreateDeviceConnectInput{
		DeviceType:  input.DeviceType,
		Browser:     input.Browser,
		IPAddress:   input.Ip,
		City:        input.City,
		Country:     input.Country,
		Date:        input.Date,
		TrustDevice: false,
	})
	if err != nil {
		return CreateDeviceConnectResponse{
			DeviceConnect: model.DeviceConnect{},
			Code:          400,
			Err:           errors.New("unable to create device connection (check if you shared all information)"),
		}
	}

	patient, patientErr := graphql.GetPatientById(ownerId)
	if patientErr == nil {
		_, err = graphql.UpdatePatient(ownerId, model.UpdatePatientInput{
			DeviceConnect: append(patient.DeviceConnect, &device.ID),
		})
		if err != nil {
			return CreateDeviceConnectResponse{
				DeviceConnect: model.DeviceConnect{},
				Code:          400,
				Err:           errors.New("failed to update patient device connection: " + err.Error()),
			}
		}
		return CreateDeviceConnectResponse{
			DeviceConnect: device,
			Code:          201,
			Err:           nil,
		}
	}

	doctor, doctorErr := graphql.GetDoctorById(ownerId)
	if doctorErr == nil {
		_, err = graphql.UpdateDoctor(ownerId, model.UpdateDoctorInput{
			DeviceConnect: append(doctor.DeviceConnect, &device.ID),
		})
		if err != nil {
			return CreateDeviceConnectResponse{
				DeviceConnect: model.DeviceConnect{},
				Code:          400,
				Err:           errors.New("failed to update doctor device connection: " + err.Error()),
			}
		}
		return CreateDeviceConnectResponse{
			DeviceConnect: device,
			Code:          201,
			Err:           nil,
		}
	}

	return CreateDeviceConnectResponse{
		DeviceConnect: model.DeviceConnect{},
		Code:          400,
		Err:           errors.New("owner ID does not correspond to a valid patient or doctor"),
	}
}
