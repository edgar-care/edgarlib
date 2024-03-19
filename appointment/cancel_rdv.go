package appointment

import (
	"context"
	"errors"
	"github.com/edgar-care/edgarlib/graphql"
)

type CancelRdvResponse struct {
	Reason string
	Code   int
	Err    error
}

func CancelRdv(id string, reason string) CancelRdvResponse {
	if id == "" {
		return CancelRdvResponse{Reason: "", Code: 400, Err: errors.New("id is required")}
	}
	gqlClient := graphql.CreateClient()

	rdv, err := graphql.GetRdvById(context.Background(), gqlClient, id)
	if err != nil {
		return CancelRdvResponse{Reason: "", Code: 400, Err: errors.New("id does not correspond to an appointment")}
	}

	var appointment_status = graphql.AppointmentStatusCanceled
	_, err = graphql.UpdateRdv(context.Background(), gqlClient, id, "", rdv.GetRdvById.Doctor_id, rdv.GetRdvById.Start_date, rdv.GetRdvById.End_date, reason, appointment_status, rdv.GetRdvById.Session_id)
	if err != nil {
		return CancelRdvResponse{Reason: "", Code: 500, Err: errors.New("unable to update appointment")}
	}
	return CancelRdvResponse{reason, 200, nil}
}
