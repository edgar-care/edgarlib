package auth

import (
	"net/http"

	lib "github.com/edgar-care/edgarlib/http"
)

func AuthenticateRequest(w http.ResponseWriter, req *http.Request) (string, bool) {
	patientID := AuthMiddleware(w, req)
	if patientID == "" {
		lib.WriteResponse(w, map[string]string{
			"message": "Not authenticated",
		}, http.StatusUnauthorized)
		return "", false
	}

	return patientID, true
}
