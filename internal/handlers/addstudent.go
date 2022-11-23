package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/doniacld/first-step-sql/internal/domain"
)

// Add calls the endpoint and writes the output
func (s Service) Add(w http.ResponseWriter, r *http.Request) {
	log.Printf("Attempting to decode AddStudentRequest")
	request, err := decodeAddRequest(r)
	if err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	_, err = domain.AddStudent(request, s.db)
	if err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}
	log.Printf("student %s %s has been added", request.FirstName, request.LastName)

	encodeAddResponse(w)
}

// decodeSearchRequest decodes the json request
func decodeAddRequest(r *http.Request) (domain.AddRequest, error) {
	var req domain.AddRequest
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return domain.AddRequest{}, fmt.Errorf("invalid input body, %v", err)
	}
	err = json.Unmarshal(body, &req)
	if err != nil {
		return domain.AddRequest{}, fmt.Errorf("unable to unmarshal the body, %v", err)
	}
	return req, nil
}

// encodeAddResponse encodes the response
func encodeAddResponse(w http.ResponseWriter) {
	w.WriteHeader(domain.AddEndpointMeta.SuccessCode())
}
