package double_auth

import (
	"errors"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
	"net/http"

	"github.com/edgar-care/edgarlib/v2/black_list"
	"github.com/edgar-care/edgarlib/v2/graphql"
)

type DeleteDeviceConnectResponse struct {
	Deleted bool
	Code    int
	Err     error
}

func remElement(slice []*string, element *string) []*string {
	var result []*string
	for _, v := range slice {
		if *v != *element {
			result = append(result, v)
		}
	}
	return result
}

func DeleteDeviceConnect(DeviceId string, ownerId string) DeleteDeviceConnectResponse {
	if DeviceId == "" {
		return DeleteDeviceConnectResponse{Deleted: false, Code: http.StatusBadRequest, Err: errors.New("device id is required")}
	}

	_, err := graphql.GetDeviceConnectById(DeviceId)
	if err != nil {
		return DeleteDeviceConnectResponse{Deleted: false, Code: http.StatusBadRequest, Err: errors.New("id does not correspond to a device")}
	}

	deleted, err := graphql.DeleteDeviceConnect(DeviceId)
	if err != nil {
		return DeleteDeviceConnectResponse{Deleted: false, Code: http.StatusInternalServerError, Err: errors.New("error while deleting device: " + err.Error())}
	}

	patient, errPatient := graphql.GetPatientById(ownerId)
	if errPatient == nil {

		_, err := graphql.UpdatePatientsDeviceConnect(ownerId, model.UpdatePatientsDeviceConnectInput{DeviceConnect: remElement(patient.DeviceConnect, &DeviceId), TrustDevices: remElement(patient.TrustDevices, &DeviceId)})

		if err != nil {
			return DeleteDeviceConnectResponse{Deleted: false, Code: http.StatusInternalServerError, Err: errors.New("error updating patient: " + err.Error())}
		}

		blacklist := black_list.UpdateBlackList(ownerId)
		if blacklist.Err != nil {
			return DeleteDeviceConnectResponse{Deleted: false, Code: http.StatusInternalServerError, Err: errors.New("error updating blacklist: " + blacklist.Err.Error())}
		}

		return DeleteDeviceConnectResponse{
			Deleted: deleted,
			Code:    http.StatusOK,
			Err:     nil,
		}
	}

	doctor, errDoctor := graphql.GetDoctorById(ownerId)
	if errDoctor == nil {

		_, err := graphql.UpdateDoctorsDeviceConnect(ownerId, model.UpdateDoctorsDeviceConnectInput{DeviceConnect: remElement(doctor.DeviceConnect, &DeviceId), TrustDevices: remElement(doctor.TrustDevices, &DeviceId)})
		if err != nil {
			return DeleteDeviceConnectResponse{Deleted: false, Code: http.StatusInternalServerError, Err: errors.New("error updating doctor: " + err.Error())}
		}

		blacklist := black_list.UpdateBlackList(ownerId)
		if blacklist.Err != nil {
			return DeleteDeviceConnectResponse{Deleted: false, Code: http.StatusInternalServerError, Err: errors.New("error updating blacklist: " + blacklist.Err.Error())}
		}

		return DeleteDeviceConnectResponse{
			Deleted: deleted,
			Code:    http.StatusOK,
			Err:     nil,
		}
	}

	return DeleteDeviceConnectResponse{Deleted: false, Code: http.StatusBadRequest, Err: errors.New("id does not correspond to a patient or doctor")}
}
