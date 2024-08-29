package double_auth

import (
	"errors"
	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
)

type RemoveTrustDeviceResponse struct {
	Patient model.Patient
	Doctor  model.Doctor
	Code    int
	Err     error
}

func RemoveTrustDevice(id_device string, id_user string) RemoveTrustDeviceResponse {
	var doubleAuthID string

	patient, err := graphql.GetPatientById(id_user)
	if err == nil {
		_, updateError := graphql.UpdatePatientTrustDevice(id_user, model.UpdatePatientTrustDeviceInput{
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

		_, updateError := graphql.UpdateDoctorsTrustDevice(id_user, model.UpdateDoctorsTrustDeviceInput{
			TrustDevices: removeTrustDevice(patient.TrustDevices, &id_device),
		})
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

	codeList := make([]*string, len(doubleAuth.TrustDeviceID))
	for i, v := range doubleAuth.TrustDeviceID {
		codeList[i] = &v
	}

	newDeviceList := removeTrustDevice(codeList, &id_device)

	updatedDeviceList := make([]string, len(newDeviceList))
	for i, v := range newDeviceList {
		updatedDeviceList[i] = *v
	}

	availableMethods := doubleAuth.Methods
	if len(updatedDeviceList) == 0 {
		availableMethods = removeMethods(availableMethods, "MOBILE")
	}

	_, err = graphql.UpdateDoubleAuth(doubleAuth.ID, model.UpdateDoubleAuthInput{
		Methods:       availableMethods,
		TrustDeviceID: updatedDeviceList,
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
		if *device != *id_device {
			updatedDevices = append(updatedDevices, device)
		}
	}
	return updatedDevices
}

func removeMethods(slice []string, element string) []string {
	var result []string
	for _, v := range slice {
		if v != element {
			result = append(result, v)
		}
	}
	return result
}
