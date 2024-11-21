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
	var accountType string

	patient, err := graphql.GetPatientById(id_user)
	if err == nil {
		accountType = "patient"
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
		accountType = "doctor"
		doctor, err := graphql.GetDoctorById(id_user)
		if err != nil {
			return RemoveTrustDeviceResponse{Code: 400, Err: errors.New("id does not correspond to a patient or doctor")}
		}
		accountType = "doctor"
		_, updateError := graphql.UpdateDoctorsTrustDevice(id_user, model.UpdateDoctorsTrustDeviceInput{
			TrustDevices: removeTrustDevice(doctor.TrustDevices, &id_device),
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

	newDeviceList := removeTrustDevice2fa(doubleAuth.TrustDeviceID, id_device)

	updatedDeviceList := make([]string, len(newDeviceList))
	copy(updatedDeviceList, newDeviceList)

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

	check, err := graphql.GetDoubleAuthById(doubleAuthID)
	if err != nil {
		return RemoveTrustDeviceResponse{Code: 400, Err: errors.New("get double_auth failed: " + err.Error())}
	}
	if check.Methods == nil || len(check.Methods) == 0 {
		_, err := graphql.DeleteDoubleAuth(doubleAuthID)
		if err != nil {
			return RemoveTrustDeviceResponse{Code: 400, Err: errors.New("delete double_auth failed: " + err.Error())}
		}
		empty := ""
		if accountType == "doctor" {
			_, err := graphql.UpdateDoctor(id_user, model.UpdateDoctorInput{DoubleAuthMethodsID: &empty})
			if err != nil {
				return RemoveTrustDeviceResponse{Code: 400, Err: errors.New("failed to remove double_auth_methods_id from doctor")}
			}
		} else {
			empty := ""
			_, err = graphql.UpdatePatient(id_user, model.UpdatePatientInput{
				DoubleAuthMethodsID: &empty,
			})
			if err != nil {
				return RemoveTrustDeviceResponse{Code: 400, Err: errors.New("failed to remove double_auth_methods_id from patient")}
			}
		}
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

func removeTrustDevice2fa(trustDevices []string, id_device string) []string {
	var updatedDevices []string
	for _, device := range trustDevices {
		if device != id_device {
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
