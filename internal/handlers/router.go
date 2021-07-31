package handlers

import (
	"github.com/doniacld/first-step-sql/domain"
	"github.com/doniacld/first-step-sql/rest"
	"github.com/gorilla/mux"
)

// Service holds service dependency
type Service struct {
}

// Router returns the service's endpoints
func (s Service) Router() *mux.Router {
	router := mux.NewRouter()

	// list all the endpoints
	routes := []rest.Route{
		{
			domain.AddEndpointMeta, s.Add,
		},
	}

	for _, r := range routes {
		router.
			Methods(r.Verb()).
			Path(r.Path()).
			Handler(r.HandlerFunc)
	}
	return router
}

// NewService instantiates a new service
func NewService() Service {
	return Service{}
}
