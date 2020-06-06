package http

import "net/http"

const (
	StatusOK        = http.StatusOK        // Response with content
	StatusNoContent = http.StatusNoContent // Response without content

	StatusBadRequest   = http.StatusBadRequest   // Invalid request
	StatusUnauthorized = http.StatusUnauthorized // Need to login
	StatusForbidden    = http.StatusForbidden    // No rights to operate
	StatusNotFound     = http.StatusNotFound     // Not found

	StatusInternalServerError = http.StatusInternalServerError // Server error
)
