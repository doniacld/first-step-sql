package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/doniacld/first-step-sql/domain"
)

// Add calls the endpoint and writes the output
func (s Service) Add(w http.ResponseWriter, r *http.Request) {
	// Search calls the endpoint and writes the output
	request, err := decodeRequest(r)
	if err != nil {
		return
	}

	// TODO implement the business
	fmt.Println(request)

	encodeResponse(w)
}

// decodeSearchRequest decodes the json request
func decodeRequest(r *http.Request) (domain.AddRequest, error) {
	var req domain.AddRequest
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return domain.AddRequest{}, fmt.Errorf("invalid input body, %v", err)
	}
	err = json.Unmarshal(body, &req)
	if err != nil {
		return domain.AddRequest{}, err
	}
	return req, nil
}

// encodeResponse encodes the response
func encodeResponse(w http.ResponseWriter) {
	w.WriteHeader(domain.AddEndpointMeta.SuccessCode())
}
