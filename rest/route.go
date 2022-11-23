package rest

import (
	"net/http"
)

// Route holds all the information of a route
type Route struct {
	Meta
	HandlerFunc http.HandlerFunc
}