package double_auth

import (
	"context"
	"errors"
	"github.com/edgar-care/edgarlib/graphql"
	"github.com/edgar-care/edgarlib/graphql/server/model"
)

type RemoveTrustDeviceResponse struct {
	Patient model.Patient
	Doctor  model.Doctor
	Code    int
	Err     error
}

func RemoveTrustDevice(id_device string, id_user string) RemoveTrustDeviceResponse {
	gqlClient := graphql.CreateClient()

	var updateError error
	var doubleAuthID string

	// Attempt to fetch patient by id_user
	patient, err := graphql.GetPatientById(context.Background(), gqlClient, id_user)
	if err == nil {
		// Update patient's trust devices
		_, updateError = graphql.UpdatePatient(context.Background(), gqlClient, patient.GetPatientById.Id, patient.GetPatientById.Email, patient.GetPatientById.Password, patient.GetPatientById.Medical_info_id, patient.GetPatientById.Rendez_vous_ids, patient.GetPatientById.Document_ids, patient.GetPatientById.Treatment_follow_up_ids, patient.GetPatientById.Chat_ids, patient.GetPatientById.Device_connect, patient.GetPatientById.Double_auth_methods_id, removeTrustDevice(patient.GetPatientById.Trust_devices, id_device), patient.GetPatientById.Status)
		if updateError != nil {
			return RemoveTrustDeviceResponse{Code: 400, Err: errors.New("update patient failed: " + updateError.Error())}
		}
		doubleAuthID = patient.GetPatientById.Double_auth_methods_id
	} else {
		// If not a patient, try to fetch doctor by id_user
		doctor, err := graphql.GetDoctorById(context.Background(), gqlClient, id_user)
		if err != nil {
			return RemoveTrustDeviceResponse{Code: 400, Err: errors.New("id does not correspond to a patient or doctor")}
		}
		// Update doctor's trust devices
		_, updateError = graphql.UpdateDoctor(context.Background(), gqlClient, doctor.GetDoctorById.Id, doctor.GetDoctorById.Email, doctor.GetDoctorById.Password, doctor.GetDoctorById.Name, doctor.GetDoctorById.Firstname, doctor.GetDoctorById.Rendez_vous_ids, doctor.GetDoctorById.Patient_ids, graphql.AddressInput{Street: doctor.GetDoctorById.Address.Street, Zip_code: doctor.GetDoctorById.Address.Zip_code, Country: doctor.GetDoctorById.Address.Country}, doctor.GetDoctorById.Chat_ids, doctor.GetDoctorById.Device_connect, doctor.GetDoctorById.Double_auth_methods_id, removeTrustDevice(doctor.GetDoctorById.Trust_devices, id_device), doctor.GetDoctorById.Status)
		if updateError != nil {
			return RemoveTrustDeviceResponse{Code: 400, Err: errors.New("update doctor failed: " + updateError.Error())}
		}
		doubleAuthID = doctor.GetDoctorById.Double_auth_methods_id
	}

	// Update double_auth with removed trust device ID
	doubleAuth, err := graphql.GetDoubleAuthById(context.Background(), gqlClient, doubleAuthID)
	if err != nil {
		return RemoveTrustDeviceResponse{Code: 400, Err: errors.New("get double_auth failed: " + err.Error())}
	}
	_, err = graphql.UpdateDoubleAuth(context.Background(), gqlClient, doubleAuth.GetDoubleAuthById.Id, doubleAuth.GetDoubleAuthById.Methods, doubleAuth.GetDoubleAuthById.Secret, doubleAuth.GetDoubleAuthById.Url, removeTrustDeviceID(doubleAuth.GetDoubleAuthById.Trust_device_id, id_device))
	if err != nil {
		return RemoveTrustDeviceResponse{Code: 400, Err: errors.New("update double_auth failed: " + err.Error())}
	}

	// Update device_connect to mark device as untrusted
	device, err := graphql.GetDeviceConnectById(context.Background(), gqlClient, id_device)
	if err != nil {
		return RemoveTrustDeviceResponse{Code: 400, Err: errors.New("get device_connect failed: " + err.Error())}
	}
	_, err = graphql.UpdateDeviceConnect(context.Background(), gqlClient, id_device, device.GetDeviceConnectById.Device_name, device.GetDeviceConnectById.Ip_address, device.GetDeviceConnectById.Latitude, device.GetDeviceConnectById.Longitude, device.GetDeviceConnectById.Date, false)
	if err != nil {
		return RemoveTrustDeviceResponse{Code: 400, Err: errors.New("update device_connect failed: " + err.Error())}
	}

	return RemoveTrustDeviceResponse{Code: 200, Err: nil}
}

func removeTrustDevice(trustDevices []string, id_device string) []string {
	var updatedDevices []string
	for _, device := range trustDevices {
		if device != id_device {
			updatedDevices = append(updatedDevices, device)
		}
	}
	return updatedDevices
}

func removeTrustDeviceID(trustDeviceID string, id_device string) string {
	if trustDeviceID == id_device {
		return ""
	}
	return trustDeviceID
}
