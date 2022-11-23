package handlers

import (
	"fmt"
	"net/http"
)

func writeError(w http.ResponseWriter, httpStatus int, err error) {
	// Let's assume for the moment that it was not our fault
	w.WriteHeader(httpStatus)
	_, _ = w.Write([]byte(fmt.Sprintf(`{"error":"%s"}`, err.Error())))
}
