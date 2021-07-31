package domain

import (
	"log"

	"github.com/doniacld/first-step-sql/internal/db"
)

// AddStudent inserts a new student in database
func AddStudent(req AddRequest, sdb db.StudentDB) AddResponse {
	log.Printf("Add student %#v", req)

	student := db.Student{
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		City:        req.City,
		Phone:       req.Phone,
		Email:       req.Email,
		Address:     req.Address,
		Postcode:    req.Postcode,
		DateOfBirth: req.DateOfBirth,
	}
	sdb.AddStudent(student)

	return AddResponse{}
}
