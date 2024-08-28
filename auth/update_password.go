package auth

import (
	"errors"
	"github.com/edgar-care/edgarlib/v2/auth/utils"
	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
	"golang.org/x/crypto/bcrypt"
	//"golang.org/x/crypto/bcrypt"
)

type UpdatePasswordResponse struct {
	Code int
	Err  error
}

type UpdatePasswordInput struct {
	OldPassword string `json:"old_password"`
	NewPassword string `json:"new_password"`
}

func UpdatePassword(input UpdatePasswordInput, ownerId string) UpdatePasswordResponse {

	patient, patientErr := graphql.GetPatientById(ownerId)
	if patientErr == nil {

		err := bcrypt.CompareHashAndPassword([]byte(patient.Password), []byte(input.OldPassword))
		if err != nil {
			return UpdatePasswordResponse{Code: 403, Err: errors.New("old password is not correct")}
		}

		HasheNewpassword := utils.HashPassword(input.NewPassword)

		_, err = graphql.UpdatePatient(ownerId, model.UpdatePatientInput{
			Password: &HasheNewpassword,
		})
		if err != nil {
			return UpdatePasswordResponse{Code: 400, Err: err}
		}
		return UpdatePasswordResponse{Code: 200, Err: nil}
	}

	doctor, doctorErr := graphql.GetDoctorById(ownerId)
	if doctorErr == nil {

		err := bcrypt.CompareHashAndPassword([]byte(doctor.Password), []byte(input.OldPassword))
		if err != nil {
			return UpdatePasswordResponse{Code: 403, Err: errors.New("old password is not correct")}
		}

		HasheNewpassword := utils.HashPassword(input.NewPassword)

		_, err = graphql.UpdateDoctor(ownerId, model.UpdateDoctorInput{
			Password: &HasheNewpassword,
		})
		if err != nil {
			return UpdatePasswordResponse{Code: 400, Err: err}
		}
		return UpdatePasswordResponse{Code: 200, Err: nil}
	}

	return UpdatePasswordResponse{Code: 400, Err: errors.New("owner ID does not correspond to a valid patient or doctor")}
}
