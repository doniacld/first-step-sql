package domain

import (
	"net/http"

	"github.com/doniacld/first-step-sql/rest"
)

// AddEndpointMeta defines the endpoint for adding
var AddEndpointMeta = rest.New(
	"/students",
	http.MethodPost,
	http.StatusOK,
)

// AddRequest defines the endpoint's request
type AddRequest struct {
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	City        string `json:"city"`
	Phone       string `json:"phone"`
	Email       string `json:"email"`
	Address     string `json:"address"`
	Postcode    string `json:"postcode"`
	DateOfBirth string `json:"dateOfBirth"`
}

// AddResponse defines the endoint's response
type AddResponse struct {
}
