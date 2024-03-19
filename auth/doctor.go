package auth

import (
	"context"
	"errors"
	"github.com/edgar-care/edgarlib/graphql"
	"github.com/edgar-care/edgarlib/graphql/server/model"
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
	gqlClient := graphql.CreateClient()
	var res model.Doctor

	doctor, err := graphql.GetDoctorById(context.Background(), gqlClient, id)
	if err != nil {
		return GetDoctorByIdResponse{model.Doctor{}, 400, errors.New("id does not correspond to a doctor")}
	}

	doctorAddress := &model.Address{
		Street:  doctor.GetDoctorById.Address.Street,
		ZipCode: doctor.GetDoctorById.Address.Zip_code,
		Country: doctor.GetDoctorById.Address.Country,
		City:    doctor.GetDoctorById.Address.City,
	}

	res = model.Doctor{
		ID:        doctor.GetDoctorById.Id,
		Email:     doctor.GetDoctorById.Email,
		Name:      doctor.GetDoctorById.Name,
		Firstname: doctor.GetDoctorById.Firstname,
		Address:   doctorAddress,
	}
	return GetDoctorByIdResponse{res, 200, nil}
}

func GetDoctors() GetDoctorsResponse {
	gqlClient := graphql.CreateClient()
	var res []model.Doctor

	doctors, err := graphql.GetDoctors(context.Background(), gqlClient)
	if err != nil {
		return GetDoctorsResponse{[]model.Doctor{}, 400, errors.New("invalid input: " + err.Error())}
	}

	for _, doctor := range doctors.GetDoctors {
		doctorAddress := &model.Address{
			Street:  doctor.Address.Street,
			ZipCode: doctor.Address.Zip_code,
			Country: doctor.Address.Country,
			City:    doctor.Address.City,
		}
		res = append(res, model.Doctor{
			ID:        doctor.Id,
			Email:     doctor.Email,
			Name:      doctor.Name,
			Firstname: doctor.Firstname,
			Address:   doctorAddress,
		})
	}
	return GetDoctorsResponse{res, 200, nil}
}
