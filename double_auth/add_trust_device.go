package double_auth

import (
	"context"
	"errors"
	"github.com/edgar-care/edgarlib/graphql"
	"github.com/edgar-care/edgarlib/graphql/server/model"
)

type AddTrustDeviceResponse struct {
	Patient model.Patient
	Doctor  model.Doctor
	Code    int
	Err     error
}

func AddTrustDevice(id_device string, id_user string) AddTrustDeviceResponse {
	gqlClient := graphql.CreateClient()

	patient, err := graphql.GetPatientById(context.Background(), gqlClient, id_user)
	if err == nil {
		_, err = graphql.UpdatePatient(context.Background(), gqlClient, id_user, patient.GetPatientById.Email, patient.GetPatientById.Password, patient.GetPatientById.Medical_info_id, patient.GetPatientById.Rendez_vous_ids, patient.GetPatientById.Document_ids, patient.GetPatientById.Treatment_follow_up_ids, patient.GetPatientById.Chat_ids, patient.GetPatientById.Device_connect, patient.GetPatientById.Double_auth_methods_id, append(patient.GetPatientById.Trust_devices, id_device), patient.GetPatientById.Status)
		if err != nil {
			return AddTrustDeviceResponse{Patient: model.Patient{}, Doctor: model.Doctor{}, Code: 400, Err: errors.New("update patient failed: " + err.Error())}
		}

		double_auth, err := graphql.GetDoubleAuthById(context.Background(), gqlClient, patient.GetPatientById.Double_auth_methods_id)
		if err != nil {
			return AddTrustDeviceResponse{Patient: model.Patient{}, Doctor: model.Doctor{}, Code: 400, Err: errors.New("get device_connect failed: " + err.Error())}
		}

		_, err = graphql.UpdateDoubleAuth(context.Background(), gqlClient, patient.GetPatientById.Double_auth_methods_id, double_auth.GetDoubleAuthById.Methods, double_auth.GetDoubleAuthById.Secret, double_auth.GetDoubleAuthById.Url, id_device)
		if err != nil {
			return AddTrustDeviceResponse{Patient: model.Patient{}, Doctor: model.Doctor{}, Code: 400, Err: errors.New("update double_auth failed: " + err.Error())}
		}

		device, err := graphql.GetDeviceConnectById(context.Background(), gqlClient, id_device)
		if err != nil {
			return AddTrustDeviceResponse{Patient: model.Patient{}, Doctor: model.Doctor{}, Code: 400, Err: errors.New("get device_connect failed: " + err.Error())}
		}

		_, err = graphql.UpdateDeviceConnect(context.Background(), gqlClient, id_device, device.GetDeviceConnectById.Device_name, device.GetDeviceConnectById.Ip_address, device.GetDeviceConnectById.Latitude, device.GetDeviceConnectById.Longitude, device.GetDeviceConnectById.Date, true)
		if err != nil {
			return AddTrustDeviceResponse{Patient: model.Patient{}, Doctor: model.Doctor{}, Code: 400, Err: errors.New("update device_connect failed: " + err.Error())}
		}

		return AddTrustDeviceResponse{
			Patient: model.Patient{},
			Doctor:  model.Doctor{},
			Code:    200,
			Err:     nil,
		}
	}

	doctor, err := graphql.GetDoctorById(context.Background(), gqlClient, id_user)
	if err != nil {
		return AddTrustDeviceResponse{Patient: model.Patient{}, Doctor: model.Doctor{}, Code: 400, Err: errors.New("id does not correspond to a patient or doctor")}
	}

	_, err = graphql.UpdateDoctor(context.Background(), gqlClient, doctor.GetDoctorById.Id, doctor.GetDoctorById.Email, doctor.GetDoctorById.Password, doctor.GetDoctorById.Name, doctor.GetDoctorById.Firstname, doctor.GetDoctorById.Rendez_vous_ids, doctor.GetDoctorById.Patient_ids, graphql.AddressInput{Street: doctor.GetDoctorById.Address.Street, Zip_code: doctor.GetDoctorById.Address.Zip_code, Country: doctor.GetDoctorById.Address.Country}, doctor.GetDoctorById.Chat_ids, doctor.GetDoctorById.Device_connect, doctor.GetDoctorById.Double_auth_methods_id, append(doctor.GetDoctorById.Trust_devices, id_device), doctor.GetDoctorById.Status)
	if err != nil {
		return AddTrustDeviceResponse{Patient: model.Patient{}, Doctor: model.Doctor{}, Code: 400, Err: errors.New("update doctor failed: " + err.Error())}
	}

	double_auth, err := graphql.GetDoubleAuthById(context.Background(), gqlClient, doctor.GetDoctorById.Double_auth_methods_id)
	if err != nil {
		return AddTrustDeviceResponse{Patient: model.Patient{}, Doctor: model.Doctor{}, Code: 400, Err: errors.New("get device_connect failed: " + err.Error())}
	}

	_, err = graphql.UpdateDoubleAuth(context.Background(), gqlClient, doctor.GetDoctorById.Double_auth_methods_id, double_auth.GetDoubleAuthById.Methods, double_auth.GetDoubleAuthById.Secret, double_auth.GetDoubleAuthById.Url, id_device)
	if err != nil {
		return AddTrustDeviceResponse{Patient: model.Patient{}, Doctor: model.Doctor{}, Code: 400, Err: errors.New("update double_auth failed: " + err.Error())}
	}

	device, err := graphql.GetDeviceConnectById(context.Background(), gqlClient, id_device)
	if err != nil {
		return AddTrustDeviceResponse{Patient: model.Patient{}, Doctor: model.Doctor{}, Code: 400, Err: errors.New("get device_connect failed: " + err.Error())}
	}

	_, err = graphql.UpdateDeviceConnect(context.Background(), gqlClient, id_device, device.GetDeviceConnectById.Device_name, device.GetDeviceConnectById.Ip_address, device.GetDeviceConnectById.Latitude, device.GetDeviceConnectById.Longitude, device.GetDeviceConnectById.Date, true)
	if err != nil {
		return AddTrustDeviceResponse{Patient: model.Patient{}, Doctor: model.Doctor{}, Code: 400, Err: errors.New("update device_connect failed: " + err.Error())}
	}

	return AddTrustDeviceResponse{
		Patient: model.Patient{},
		Doctor:  model.Doctor{},
		Code:    200,
		Err:     nil,
	}
}
