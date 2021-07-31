package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/doniacld/first-step-sql/internal/domain"
)

// SearchByName calls the endpoint and writes the output
func (s Service) SearchByName(w http.ResponseWriter, r *http.Request) {
	request, err := decodeSearchByNameRequest(r)
	if err != nil {
		return
	}

	resp, err := domain.SearchByName(request, s.db)
	if err != nil {
		w.WriteHeader(http.StatusFailedDependency)
	}
	encodeSearchByNameResponse(resp, w)
}

// decodeSearchByNameRequest decodes the json request
func decodeSearchByNameRequest(r *http.Request) (domain.SearchByNameRequest, error) {
	var req domain.SearchByNameRequest
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return domain.SearchByNameRequest{}, fmt.Errorf("invalid input body, %v", err)
	}
	err = json.Unmarshal(body, &req)
	if err != nil {
		return domain.SearchByNameRequest{}, err
	}
	return req, nil
}

// encodeSearchByNameResponse encodes the response
func encodeSearchByNameResponse(response domain.SearchByNameResponse, w http.ResponseWriter) {
	b, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusFailedDependency)
		return
	}

	w.WriteHeader(domain.SearchByNameEndpointMeta.SuccessCode())
	_, _ = w.Write(b)
}
