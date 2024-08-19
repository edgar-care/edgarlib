package double_auth

import (
	"errors"
	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
	"net/http"
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

func GetDeviceConnect(ownerId string) GetDevicesConnectResponse {

	_, errPatient := graphql.GetPatientById(ownerId)
	if errPatient == nil {

		devices, err := graphql.GetDevicesConnect(nil)
		if err != nil {
			return GetDevicesConnectResponse{[]model.DeviceConnect{}, http.StatusBadRequest, errors.New("invalid input: " + err.Error())}
		}
		return GetDevicesConnectResponse{devices, http.StatusOK, nil}
	}

	_, errDoctor := graphql.GetDoctorById(ownerId)
	if errDoctor == nil {
		// Fetch devices connected to the doctor
		devices, err := graphql.GetDevicesConnect(nil)
		if err != nil {
			return GetDevicesConnectResponse{[]model.DeviceConnect{}, http.StatusBadRequest, errors.New("invalid input: " + err.Error())}
		}
		return GetDevicesConnectResponse{devices, http.StatusOK, nil}
	}

	// If neither patient nor doctor retrieval succeeds, return an error
	return GetDevicesConnectResponse{[]model.DeviceConnect{}, http.StatusBadRequest, errors.New("id does not correspond to a patient or doctor")}
}
