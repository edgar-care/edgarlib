package appointment

import (
	"errors"
	"github.com/edgar-care/edgarlib/graphql"
	"github.com/edgar-care/edgarlib/graphql/model"
)

type GetDoctorAppointmentResponse struct {
	Appointment model.Rdv
	Code        int
	Err         error
}

type GetAllDoctorAppointmentResponse struct {
	Slots []model.Rdv
	Code  int
	Err   error
}

func GetDoctorAppointment(appointmentId string, doctorId string) GetDoctorAppointmentResponse {

	rdv, err := graphql.GetRdvById(appointmentId)
	if err != nil {
		return GetDoctorAppointmentResponse{model.Rdv{}, 400, errors.New("id does not correspond to an appointment")}
	}

	if rdv.DoctorID != doctorId {
		return GetDoctorAppointmentResponse{model.Rdv{}, 403, errors.New("unauthorized to access this appointment")}
	}

	return GetDoctorAppointmentResponse{
		Appointment: rdv,
		Code:        200,
		Err:         nil,
	}
}

func GetAllDoctorAppointment(doctorId string) GetAllDoctorAppointmentResponse {
	var res []model.Rdv

	_, err := graphql.GetDoctorById(doctorId)
	if err != nil {
		return GetAllDoctorAppointmentResponse{[]model.Rdv{}, 400, errors.New("id does not correspond to a doctor")}
	}

	appointments, err := graphql.GetDoctorRdv(doctorId, nil)
	if err != nil {
		return GetAllDoctorAppointmentResponse{[]model.Rdv{}, 400, errors.New("invalid input: " + err.Error())}
	}

	for _, appointment := range appointments {
		res = append(res, appointment)
	}
	return GetAllDoctorAppointmentResponse{res, 201, nil}
}
