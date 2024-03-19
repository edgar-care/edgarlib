package appointment

import (
	"context"
	"errors"
	"github.com/edgar-care/edgarlib/graphql"
	"github.com/edgar-care/edgarlib/graphql/server/model"
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
	gqlClient := graphql.CreateClient()

	rdv, err := graphql.GetRdvById(context.Background(), gqlClient, appointmentId)
	if err != nil {
		return GetDoctorAppointmentResponse{model.Rdv{}, 400, errors.New("id does not correspond to an appointment")}
	}

	if rdv.GetRdvById.Doctor_id != doctorId {
		return GetDoctorAppointmentResponse{model.Rdv{}, 403, errors.New("unauthorized to access this appointment")}
	}

	return GetDoctorAppointmentResponse{
		Appointment: model.Rdv{
			ID:                rdv.GetRdvById.Id,
			DoctorID:          rdv.GetRdvById.Doctor_id,
			IDPatient:         rdv.GetRdvById.Id_patient,
			StartDate:         rdv.GetRdvById.Start_date,
			EndDate:           rdv.GetRdvById.End_date,
			CancelationReason: &rdv.GetRdvById.Cancelation_reason,
			AppointmentStatus: model.AppointmentStatus(rdv.GetRdvById.Appointment_status),
			SessionID:         rdv.GetRdvById.Session_id,
		},
		Code: 200,
		Err:  nil,
	}
}

func GetAllDoctorAppointment(doctorId string) GetAllDoctorAppointmentResponse {
	gqlClient := graphql.CreateClient()
	var res []model.Rdv

	_, err := graphql.GetDoctorById(context.Background(), gqlClient, doctorId)
	if err != nil {
		return GetAllDoctorAppointmentResponse{[]model.Rdv{}, 400, errors.New("id does not correspond to a doctor")}
	}

	appointments, err := graphql.GetDoctorRdv(context.Background(), gqlClient, doctorId)
	if err != nil {
		return GetAllDoctorAppointmentResponse{[]model.Rdv{}, 400, errors.New("invalid input: " + err.Error())}
	}

	for _, appointment := range appointments.GetDoctorRdv {
		temp := appointment.Cancelation_reason
		res = append(res, model.Rdv{
			ID:                appointment.Id,
			DoctorID:          appointment.Doctor_id,
			IDPatient:         appointment.Id_patient,
			StartDate:         appointment.Start_date,
			EndDate:           appointment.End_date,
			CancelationReason: &temp,
			AppointmentStatus: model.AppointmentStatus(appointment.Appointment_status),
			SessionID:         appointment.Session_id,
		})
	}
	return GetAllDoctorAppointmentResponse{res, 201, nil}
}
