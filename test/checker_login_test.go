package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/edgar-care/edgarlib/auth"
	"github.com/stretchr/testify/assert"
)

func TestAuthenticateRequest(t *testing.T) {
	req, err := http.NewRequest("GET", "/test", nil)
	assert.NoError(t, err)

	w := httptest.NewRecorder()

	patientID, authenticated := auth.AuthenticateRequest(w, req)
	assert.True(t, authenticated)
	assert.NotEmpty(t, patientID)

	w = httptest.NewRecorder()
	req.Header.Set("Authorization", "invalid_token")
	patientID, authenticated = auth.AuthenticateRequest(w, req)
	assert.False(t, authenticated)
	assert.Empty(t, patientID)
	assert.Equal(t, http.StatusUnauthorized, w.Code)
}
