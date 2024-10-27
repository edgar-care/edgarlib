package double_auth

import (
	"errors"
	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
)

func RemoveDoubleAuthMethod(methodToRemove string, ownerId string) error {
	patient, err := graphql.GetPatientById(ownerId)
	if err == nil {
		if patient.DoubleAuthMethodsID == nil || *patient.DoubleAuthMethodsID == "" {
			return errors.New("patient does not have double_auth_methods_id")
		}

		doubleAuth, err := graphql.GetDoubleAuthById(*patient.DoubleAuthMethodsID)
		if err != nil {
			return errors.New("unable to fetch double_auth")
		}

		methods := doubleAuth.Methods
		found := false
		for i, m := range methods {
			if m == methodToRemove {
				methods = append(methods[:i], methods[i+1:]...)
				found = true
				break
			}
		}

		if !found {
			return errors.New("method to remove not found in current methods")
		}

		if methodToRemove == "MOBILE" {
			_, err = graphql.UpdateDoubleAuth(doubleAuth.ID, model.UpdateDoubleAuthInput{
				Methods:       methods,
				TrustDeviceID: nil,
			})
			if err != nil {
				return errors.New("failed to update double_auth")
			}
			status := false
			for _, deviceID := range doubleAuth.TrustDeviceID {
				_, err = graphql.UpdateDeviceConnect(deviceID, model.UpdateDeviceConnectInput{
					TrustDevice: &status,
				})
				if err != nil {
					return errors.New("failed to update device_connect")
				}
				trustDevicePtrs := make([]*string, len(doubleAuth.TrustDeviceID))
				for i, v := range doubleAuth.TrustDeviceID {
					trustDevicePtrs[i] = &v
				}
				remTrust := remove(trustDevicePtrs, &deviceID)
				_, err = graphql.UpdatePatientTrustDevice(patient.ID, model.UpdatePatientTrustDeviceInput{
					TrustDevices: remTrust,
				})
				if err != nil {
					return errors.New("failed to delete trust_device")
				}
			}

		} else if methodToRemove == "AUTHENTIFICATOR" {
			_, err = graphql.UpdateDoubleAuth(doubleAuth.ID, model.UpdateDoubleAuthInput{
				Methods: methods,
				Code:    nil,
				URL:     nil,
			})
		} else {
			_, err = graphql.UpdateDoubleAuth(doubleAuth.ID, model.UpdateDoubleAuthInput{
				Methods: methods,
			})
		}
		if err != nil {
			return errors.New("failed to update double_auth")
		}

		if len(methods) == 0 {
			empty := ""
			_, err = graphql.UpdatePatient(ownerId, model.UpdatePatientInput{
				DoubleAuthMethodsID: &empty,
			})
			if err != nil {
				return errors.New("failed to remove double_auth_methods_id from patient")
			}
		}

		return nil
	}

	doctor, err := graphql.GetDoctorById(ownerId)
	if err == nil {
		if doctor.DoubleAuthMethodsID == nil || *doctor.DoubleAuthMethodsID == "" {
			return errors.New("doctor does not have double_auth_methods_id")
		}

		doubleAuth, err := graphql.GetDoubleAuthById(*doctor.DoubleAuthMethodsID)
		if err != nil {
			return errors.New("unable to fetch double_auth")
		}

		methods := doubleAuth.Methods
		found := false
		for i, m := range methods {
			if m == methodToRemove {
				methods = append(methods[:i], methods[i+1:]...)
				found = true
				break
			}
		}

		if !found {
			return errors.New("method to remove not found in current methods")
		}

		if methodToRemove == "MOBILE" {
			_, err = graphql.UpdateDoubleAuth(doubleAuth.ID, model.UpdateDoubleAuthInput{
				Methods:       methods,
				TrustDeviceID: nil,
			})
			if err != nil {
				return errors.New("failed to update double_auth")
			}
			status := false
			for _, deviceID := range doubleAuth.TrustDeviceID {
				_, err = graphql.UpdateDeviceConnect(deviceID, model.UpdateDeviceConnectInput{
					TrustDevice: &status,
				})
				if err != nil {
					return errors.New("failed to update device_connect")
				}

				trustDevicePtrs := make([]*string, len(doubleAuth.TrustDeviceID))
				for i, v := range doubleAuth.TrustDeviceID {
					trustDevicePtrs[i] = &v
				}
				remTrust := remove(trustDevicePtrs, &deviceID)
				_, err = graphql.UpdateDoctorsTrustDevice(doctor.ID, model.UpdateDoctorsTrustDeviceInput{
					TrustDevices: remTrust,
				})
				if err != nil {
					return errors.New("failed to delete trust_device")
				}
			}

		} else if methodToRemove == "AUTHENTIFICATOR" {
			_, err = graphql.UpdateDoubleAuth(doubleAuth.ID, model.UpdateDoubleAuthInput{
				Methods: methods,
				Code:    nil,
				URL:     nil,
			})
		} else {
			_, err = graphql.UpdateDoubleAuth(doubleAuth.ID, model.UpdateDoubleAuthInput{
				Methods: methods,
			})
		}
		if err != nil {
			return errors.New("failed to update double_auth")
		}

		if len(methods) == 0 {
			empty := ""
			_, err = graphql.UpdateDoctor(ownerId, model.UpdateDoctorInput{
				DoubleAuthMethodsID: &empty,
			})
			if err != nil {
				return errors.New("failed to remove double_auth_methods_id from doctor")
			}
		}

		return nil
	}

	return errors.New("id does not correspond to a patient or doctor")
}

func remove(trustDevices []*string, id_device *string) []*string {
	var updatedDevices []*string
	for _, device := range trustDevices {
		if *device != *id_device {
			updatedDevices = append(updatedDevices, device)
		}
	}
	return updatedDevices
}
