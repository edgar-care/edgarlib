package double_auth

import (
	"errors"
	"github.com/edgar-care/edgarlib/graphql"
	"github.com/edgar-care/edgarlib/graphql/model"
)

type RemoveTrustDeviceResponse struct {
	Patient model.Patient
	Doctor  model.Doctor
	Code    int
	Err     error
}

func RemoveTrustDevice(id_device string, id_user string) RemoveTrustDeviceResponse {

	var updateError error
	var doubleAuthID string

	patient, err := graphql.GetPatientById(id_user)
	if err == nil {
		_, updateError = graphql.UpdatePatient(id_user, model.UpdatePatientInput{
			TrustDevices: removeTrustDevice(patient.TrustDevices, &id_device),
		})
		if updateError != nil {
			return RemoveTrustDeviceResponse{Code: 400, Err: errors.New("update patient failed: " + updateError.Error())}
		}
		if patient.DoubleAuthMethodsID == nil {
			return RemoveTrustDeviceResponse{Code: 404, Err: errors.New("double auth not found on patient")}
		}
		doubleAuthID = *patient.DoubleAuthMethodsID
	} else {

		doctor, err := graphql.GetDoctorById(id_user)
		if err != nil {
			return RemoveTrustDeviceResponse{Code: 400, Err: errors.New("id does not correspond to a patient or doctor")}
		}

		_, updateError = graphql.UpdateDoctor(doctor.ID, model.UpdateDoctorInput{TrustDevices: removeTrustDevice(doctor.TrustDevices, &id_device)})
		if updateError != nil {
			return RemoveTrustDeviceResponse{Code: 400, Err: errors.New("update doctor failed: " + updateError.Error())}
		}
		if doctor.DoubleAuthMethodsID == nil {
			return RemoveTrustDeviceResponse{Code: 404, Err: errors.New("double auth not found on doctor")}
		}
		doubleAuthID = *doctor.DoubleAuthMethodsID
	}

	doubleAuth, err := graphql.GetDoubleAuthById(doubleAuthID)
	if err != nil {
		return RemoveTrustDeviceResponse{Code: 400, Err: errors.New("get double_auth failed: " + err.Error())}
	}

	_, err = graphql.UpdateDoubleAuth(doubleAuth.ID, model.UpdateDoubleAuthInput{
		TrustDeviceID: removeTrustDeviceID(&doubleAuth.TrustDeviceID, &id_device),
	})
	if err != nil {
		return RemoveTrustDeviceResponse{Code: 400, Err: errors.New("update double_auth failed: " + err.Error())}
	}

	status := false
	_, err = graphql.UpdateDeviceConnect(id_device, model.UpdateDeviceConnectInput{TrustDevice: &status})
	if err != nil {
		return RemoveTrustDeviceResponse{Code: 400, Err: errors.New("update device_connect failed: " + err.Error())}
	}

	return RemoveTrustDeviceResponse{Code: 200, Err: nil}
}

func removeTrustDevice(trustDevices []*string, id_device *string) []*string {
	var updatedDevices []*string
	for _, device := range trustDevices {
		if device != id_device {
			updatedDevices = append(updatedDevices, device)
		}
	}
	return updatedDevices
}

func removeTrustDeviceID(trustDeviceID *string, id_device *string) *string {
	if trustDeviceID == id_device {
		return nil
	}
	return trustDeviceID
}
