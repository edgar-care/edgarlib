package graphql

import (
	"github.com/jinzhu/copier"
)

/********** Types ***********/

type Rdv struct {
	Id        string `json:"id"`
	DoctorID  string `json:"doctor_id"`
	IdPatient string `json:"id_patient"`
	StartDate int    `json:"start_date"`
	EndDate   int    `json:"end_date"`
}

type RdvOutput struct {
	Id        string  `json:"id"`
	DoctorID  *string `json:"doctor_id"`
	IdPatient *string `json:"id_patient"`
	StartDate *int    `json:"start_date"`
	EndDate   *int    `json:"end_date"`
}

type RdvInput struct {
	Id        string `json:"id"`
	DoctorID  string `json:"doctor_id"`
	IdPatient string `json:"id_patient"`
	StartDate int    `json:"start_date"`
	EndDate   int    `json:"end_date"`
}

type UpdatePatientAppointment struct {
	Id            string   `json:"id"`
	RendezVousIDs []string `json:"rendez_vous_ids"`
}

type UpdatePatientAppointmentInput struct {
	Id            string   `json:"id"`
	RendezVousIDs []string `json:"rendez_vous_ids"`
}

type UpdatePatientAppointmentOutput struct {
	Id            *string   `json:"id"`
	RendezVousIDs *[]string `json:"rendez_vous_ids"`
}

type UpdateDoctorAppointment struct {
	Id            string   `json:"id"`
	RendezVousIDs []string `json:"rendez_vous_ids"`
}

type UpdateDoctorAppointmentInput struct {
	Id            string   `json:"id"`
	RendezVousIDs []string `json:"rendez_vous_ids"`
}

type UpdateDoctorAppointmentOutput struct {
	Id            *string   `json:"id"`
	RendezVousIDs *[]string `json:"rendez_vous_ids"`
}

/**************** GraphQL types *****************/

type updateRdvResponse struct {
	Content RdvOutput `json:"updateRdv"`
}

type getOneRdvByIdResponse struct {
	Content RdvOutput `json:"GetRdvById"`
}

type getAllRdvResponse struct {
	Content []RdvOutput `json:"getPatientRdv"`
}

type updatePatientAppointmentResponse struct {
	Content UpdatePatientAppointmentOutput `json:"updatePatient"`
}

// type deleteRdvByIdResponse struct {
// 	Content RdvOutput `json:"DeleteRdvById"`
// }

type getRdvDoctorResponse struct {
	Content []RdvOutput `json:"getDoctorRdv"`
}

/*************** Implementations *****************/

func UpdateRdv(id_patient string, rdv_id string) (Rdv, error) {
	var rdv updateRdvResponse
	var resp Rdv
	query := `mutation updateRdv($id: String!, $id_patient: String!) {
		updateRdv(id:$id, id_patient:$id_patient) {
                    id,
					doctor_id,
					start_date,
					end_date,
					id_patient
                }
            }`
	err := Query(query, map[string]interface{}{
		"id":         rdv_id,
		"id_patient": id_patient,
	}, &rdv)
	_ = copier.Copy(&resp, &rdv.Content)
	return resp, err
}

func GetRdvById(id string) (Rdv, error) {
	var onerdv getOneRdvByIdResponse
	var resp Rdv
	query := `query getRdvById($id: String!) {
                getRdvById(id: $id) {
                    id,
					doctor_id,
					start_date,
					end_date,
					id_patient
                }
            }`

	err := Query(query, map[string]interface{}{
		"id": id,
	}, &onerdv)
	_ = copier.Copy(&resp, &onerdv.Content)
	return resp, err
}

func GetAllRdv(id string) ([]Rdv, error) {
	var allrdv getAllRdvResponse
	var resp []Rdv
	query := `query getPatientRdv($id_patient: String!){
                getPatientRdv(id_patient: $id_patient) {
                    id,
					doctor_id,
					start_date,
					end_date,
					id_patient
                }
            }`
	err := Query(query, map[string]interface{}{
		"id_patient": id,
	}, &allrdv)
	_ = copier.Copy(&resp, &allrdv.Content)
	return resp, err
}

// ======================================================================================== //

func GetRdvDoctorById(id string) ([]Rdv, error) {
	var allrdv getRdvDoctorResponse
	var resp []Rdv
	query := `query getDoctorRdv($doctor_id: String!){
                getDoctorRdv(doctor_id: $doctor_id) {
                    id,
					doctor_id,
					start_date,
					end_date,
					id_patient
                }
            }`
	err := Query(query, map[string]interface{}{
		"doctor_id": id,
	}, &allrdv)
	_ = copier.Copy(&resp, &allrdv.Content)
	return resp, err
}

// ============================================================================================== //
// Patient
func UpdatePatient(updatePatient UpdatePatientAppointmentInput) (UpdatePatientAppointment, error) {
	var patient updatePatientAppointmentResponse
	var resp UpdatePatientAppointment
	query := `mutation updatePatient($id: String!, $rendez_vous_ids: [String]) {
		updatePatient(id:$id, rendez_vous_ids:$rendez_vous_ids) {
                    id,
					rendez_vous_ids
                }
            }`
	err := Query(query, map[string]interface{}{
		"id":              updatePatient.Id,
		"rendez_vous_ids": updatePatient.RendezVousIDs,
	}, &patient)
	_ = copier.Copy(&resp, &patient.Content)
	return resp, err
}

func GetPatientAppointmentById(id string) (UpdatePatientAppointment, error) {
	var patient updatePatientAppointmentResponse
	var resp UpdatePatientAppointment
	query := `query getPatientById($id: String!) {
                getPatientById(id: $id) {
                    id,
					rendez_vous_ids
                }
            }`

	err := Query(query, map[string]interface{}{
		"id": id,
	}, &patient)
	_ = copier.Copy(&resp, &patient.Content)
	return resp, err
}
