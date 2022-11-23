package handlers

import (
	"github.com/gorilla/mux"

	"github.com/doniacld/first-step-sql/internal/db"
	"github.com/doniacld/first-step-sql/internal/domain"
	"github.com/doniacld/first-step-sql/rest"
)

// Service holds service dependency
type Service struct {
	db db.StudentDB
}

// Router returns the service's endpoints
func (s Service) Router() *mux.Router {
	router := mux.NewRouter()

	// list all the endpoints
	routes := []rest.Route{
		{domain.AddEndpointMeta, s.Add},
		{domain.SearchByNameEndpointMeta, s.SearchByName},
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
func NewService(db db.StudentDB) Service {
	return Service{db: db}
}
