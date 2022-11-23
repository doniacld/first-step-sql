package domain

import (
	"github.com/doniacld/first-step-sql/internal/db"
	"net/http"

	"github.com/doniacld/first-step-sql/rest"
)

// SearchByNameEndpointMeta defines the endpoint for searching by name
var SearchByNameEndpointMeta = rest.New(
	"/students/search",
	http.MethodPost,
	http.StatusOK,
)

// SearchByNameRequest defines the endpoint's request
type SearchByNameRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

// SearchByNameResponse defines the endpoint's response
// TODO change the model and makes aliases
type SearchByNameResponse db.Student
