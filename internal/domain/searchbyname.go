package domain

import (
	"fmt"
	"log"

	"github.com/doniacld/first-step-sql/internal/db"
)

// SearchByName search a student by its name
func SearchByName(req SearchByNameRequest, sdb db.StudentDB) (SearchByNameResponse, error) {
	log.Printf("Get the student info with the name %#v", req)

	studentDB, err := sdb.SearchByName(req.FirstName, req.LastName)
	if err != nil {
		return SearchByNameResponse{}, fmt.Errorf("unable to find the student: %s", err)
	}

	return SearchByNameResponse{
		FirstName:   studentDB.FirstName,
		LastName:    studentDB.LastName,
		City:        studentDB.City,
		Phone:       studentDB.Phone,
		Email:       studentDB.Email,
		Address:     studentDB.Address,
		Postcode:    studentDB.Postcode,
		DateOfBirth: studentDB.DateOfBirth,
	}, nil
}
