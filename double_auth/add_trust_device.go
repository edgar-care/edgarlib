package double_auth

import (
	"errors"
	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
)

type AddTrustDeviceResponse struct {
	Patient *model.Patient
	Doctor  *model.Doctor
	Code    int
	Err     error
}

func AddTrustDevice(id_device string, id_user string) AddTrustDeviceResponse {

	patient, err := graphql.GetPatientById(id_user)
	if err == nil {
		for _, device := range patient.TrustDevices {
			if *device == id_device {
				return AddTrustDeviceResponse{Patient: nil, Doctor: nil, Code: 400, Err: errors.New("device already trusted")}
			}
		}

		updatedPatient, err := graphql.UpdatePatient(id_user, model.UpdatePatientInput{
			TrustDevices: append(patient.TrustDevices, &id_device),
		})
		if err != nil {
			return AddTrustDeviceResponse{Patient: &model.Patient{}, Doctor: nil, Code: 400, Err: errors.New("update patient failed: " + err.Error())}
		}

		if patient.DoubleAuthMethodsID == nil {
			return AddTrustDeviceResponse{Patient: nil, Doctor: nil, Code: 404, Err: errors.New("double auth not found on patient")}
		}

		getDoubleAuth, err := graphql.GetDoubleAuthById(*patient.DoubleAuthMethodsID)
		if err != nil {
			return AddTrustDeviceResponse{Patient: nil, Doctor: nil, Code: 404, Err: errors.New("double auth not found on patient")}
		}

		_, err = graphql.UpdateDoubleAuth(*patient.DoubleAuthMethodsID, model.UpdateDoubleAuthInput{
			Methods:       getDoubleAuth.Methods,
			TrustDeviceID: append(getDoubleAuth.TrustDeviceID, id_device),
		})
		if err != nil {
			return AddTrustDeviceResponse{Patient: &model.Patient{}, Doctor: nil, Code: 400, Err: errors.New("update double_auth failed: " + err.Error())}
		}

		status := true
		_, err = graphql.UpdateDeviceConnect(id_device, model.UpdateDeviceConnectInput{
			TrustDevice: &status,
		})
		if err != nil {
			return AddTrustDeviceResponse{Patient: &model.Patient{}, Doctor: nil, Code: 400, Err: errors.New("update device_connect failed: " + err.Error())}
		}

		return AddTrustDeviceResponse{
			Patient: &updatedPatient,
			Doctor:  nil,
			Code:    200,
			Err:     nil,
		}
	}

	doctor, err := graphql.GetDoctorById(id_user)
	if err != nil {
		return AddTrustDeviceResponse{Patient: nil, Doctor: &model.Doctor{}, Code: 400, Err: errors.New("id does not correspond to a patient or doctor")}
	}
	for _, device := range doctor.TrustDevices {
		if *device == id_device {
			return AddTrustDeviceResponse{Patient: nil, Doctor: nil, Code: 400, Err: errors.New("device already trusted")}
		}
	}

	updatedDoctor, err := graphql.UpdateDoctorsTrustDevice(id_user, model.UpdateDoctorsTrustDeviceInput{
		TrustDevices: append(doctor.TrustDevices, &id_device),
	})
	if err != nil {
		return AddTrustDeviceResponse{Patient: nil, Doctor: &model.Doctor{}, Code: 400, Err: errors.New("update doctor failed: " + err.Error())}
	}

	if doctor.DoubleAuthMethodsID == nil {
		return AddTrustDeviceResponse{Patient: nil, Doctor: nil, Code: 404, Err: errors.New("double auth not found on doctor")}
	}

	getDoubleAuth, err := graphql.GetDoubleAuthById(*doctor.DoubleAuthMethodsID)
	if err != nil {
		return AddTrustDeviceResponse{Patient: nil, Doctor: nil, Code: 404, Err: errors.New("double auth not found on patient")}
	}
	_, err = graphql.UpdateDoubleAuth(*doctor.DoubleAuthMethodsID, model.UpdateDoubleAuthInput{
		Methods:       getDoubleAuth.Methods,
		TrustDeviceID: append(getDoubleAuth.TrustDeviceID, id_device),
	})
	if err != nil {
		return AddTrustDeviceResponse{Patient: nil, Doctor: &model.Doctor{}, Code: 400, Err: errors.New("update double_auth failed: " + err.Error())}
	}

	status := true
	_, err = graphql.UpdateDeviceConnect(id_device, model.UpdateDeviceConnectInput{
		TrustDevice: &status,
	})
	if err != nil {
		return AddTrustDeviceResponse{Patient: nil, Doctor: &model.Doctor{}, Code: 400, Err: errors.New("update device_connect failed: " + err.Error())}
	}

	return AddTrustDeviceResponse{
		Patient: nil,
		Doctor:  &updatedDoctor,
		Code:    200,
		Err:     nil,
	}
}
