package auth

import (
	"errors"
	"github.com/edgar-care/edgarlib/graphql"
	"github.com/edgar-care/edgarlib/graphql/model"
)

type GetDoctorByIdResponse struct {
	Doctor model.Doctor
	Code   int
	Err    error
}

type GetDoctorsResponse struct {
	Doctors []model.Doctor
	Code    int
	Err     error
}

func GetDoctorById(id string) GetDoctorByIdResponse {
	var res model.Doctor

	doctor, err := graphql.GetDoctorById(id)

	if err != nil {
		return GetDoctorByIdResponse{model.Doctor{}, 400, errors.New("id does not correspond to a doctor")}
	}

	doctorAddress := &model.Address{
		Street:  doctor.Address.Street,
		ZipCode: doctor.Address.ZipCode,
		Country: doctor.Address.Country,
		City:    doctor.Address.City,
	}

	res = model.Doctor{
		ID:        doctor.ID,
		Email:     doctor.Email,
		Name:      doctor.Name,
		Firstname: doctor.Firstname,
		Address:   doctorAddress,
	}
	return GetDoctorByIdResponse{res, 200, nil}
}

func GetDoctors() GetDoctorsResponse {
	var res []model.Doctor

	doctors, err := graphql.GetDoctors(nil)
	if err != nil {
		return GetDoctorsResponse{[]model.Doctor{}, 400, errors.New("invalid input: " + err.Error())}
	}

	for _, doctor := range doctors {
		doctorAddress := &model.Address{
			Street:  doctor.Address.Street,
			ZipCode: doctor.Address.ZipCode,
			Country: doctor.Address.Country,
			City:    doctor.Address.City,
		}
		res = append(res, model.Doctor{
			ID:        doctor.ID,
			Email:     doctor.Email,
			Name:      doctor.Name,
			Firstname: doctor.Firstname,
			Address:   doctorAddress,
		})
	}
	return GetDoctorsResponse{res, 200, nil}
}
