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
	StatusBool := true
	patient, err := graphql.GetPatientById(id_user)
	if err == nil {
		updatedPatient, err := graphql.UpdatePatient(id_user, model.UpdatePatientInput{
			TrustDevices: append(patient.TrustDevices, &id_device),
		})
		if err != nil {
			return AddTrustDeviceResponse{Patient: &model.Patient{}, Doctor: nil, Code: 400, Err: errors.New("update patient failed: " + err.Error())}
		}

		if patient.DoubleAuthMethodsID == nil {
			return AddTrustDeviceResponse{Patient: nil, Doctor: nil, Code: 404, Err: errors.New("double auth not found on patient")}
		}
		_, err = graphql.UpdateDoubleAuth(*patient.DoubleAuthMethodsID, model.UpdateDoubleAuthInput{
			TrustDeviceID: &id_device,
		})
		if err != nil {
			return AddTrustDeviceResponse{Patient: &model.Patient{}, Doctor: nil, Code: 400, Err: errors.New("update double_auth failed: " + err.Error())}
		}

		_, err = graphql.UpdateDeviceConnect(id_device, model.UpdateDeviceConnectInput{
			TrustDevice: &StatusBool,
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

	updatedDoctor, err := graphql.UpdateDoctor(id_user, model.UpdateDoctorInput{
		TrustDevices: append(patient.TrustDevices, &id_device),
	})
	if err != nil {
		return AddTrustDeviceResponse{Patient: nil, Doctor: &model.Doctor{}, Code: 400, Err: errors.New("update doctor failed: " + err.Error())}
	}

	if doctor.DoubleAuthMethodsID == nil {
		return AddTrustDeviceResponse{Patient: nil, Doctor: nil, Code: 404, Err: errors.New("double auth not found on doctor")}
	}
	_, err = graphql.UpdateDoubleAuth(*doctor.DoubleAuthMethodsID, model.UpdateDoubleAuthInput{
		TrustDeviceID: &id_device,
	})
	if err != nil {
		return AddTrustDeviceResponse{Patient: nil, Doctor: &model.Doctor{}, Code: 400, Err: errors.New("update double_auth failed: " + err.Error())}
	}

	_, err = graphql.UpdateDeviceConnect(id_device, model.UpdateDeviceConnectInput{TrustDevice: &StatusBool})
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
