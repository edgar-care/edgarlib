package double_auth

import (
	"errors"
	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
	"github.com/edgar-care/edgarlib/v2/paging"
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
		return GetDeviceConnectByIdResponse{model.DeviceConnect{}, 400, errors.New("id does not correspond to a device")}
	}

	return GetDeviceConnectByIdResponse{device, 200, nil}
}

func GetDeviceConnect(ownerId string, page int, size int) GetDevicesConnectResponse {

	patient, errPatient := graphql.GetPatientById(ownerId)
	if errPatient == nil {
		devices, err := graphql.GetDevicesConnect(patient.ID, paging.CreatePagingOption(page, size))
		if err != nil {
			return GetDevicesConnectResponse{[]model.DeviceConnect{}, http.StatusBadRequest, errors.New("invalid input: " + err.Error())}
		}
		return GetDevicesConnectResponse{devices, http.StatusOK, nil}
	}

	doctor, errDoctor := graphql.GetDoctorById(ownerId)
	if errDoctor == nil {
		devices, err := graphql.GetDevicesConnect(doctor.ID, paging.CreatePagingOption(page, size))
		if err != nil {
			return GetDevicesConnectResponse{[]model.DeviceConnect{}, http.StatusBadRequest, errors.New("invalid input: " + err.Error())}
		}
		return GetDevicesConnectResponse{devices, http.StatusOK, nil}
	}

	// If neither patient nor doctor retrieval succeeds, return an error
	return GetDevicesConnectResponse{[]model.DeviceConnect{}, http.StatusBadRequest, errors.New("id does not correspond to a patient or doctor")}
}
