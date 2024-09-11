package auth

import (
	"errors"
	"github.com/edgar-care/edgarlib/v2/graphql"
	"github.com/edgar-care/edgarlib/v2/graphql/model"
	"github.com/edgar-care/edgarlib/v2/paging"
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
	doctor, err := graphql.GetDoctorById(id)

	if err != nil {
		return GetDoctorByIdResponse{model.Doctor{}, 400, errors.New("id does not correspond to a doctor")}
	}
	return GetDoctorByIdResponse{doctor, 200, nil}
}

func GetDoctors(page int, size int) GetDoctorsResponse {
	doctors, err := graphql.GetDoctors(paging.CreatePagingOption(page, size))
	if err != nil {
		return GetDoctorsResponse{[]model.Doctor{}, 400, errors.New("invalid input: " + err.Error())}
	}
	return GetDoctorsResponse{doctors, 200, nil}
}
