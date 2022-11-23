package domain

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/doniacld/first-step-sql/internal/db"
)

// AddStudent inserts a new student in database
func AddStudent(req AddRequest, sdb db.StudentDB) (AddResponse, error) {
	log.Printf("Add student %#v", req)

	postcode, err := strconv.Atoi(req.Postcode)
	if err != nil {
		return AddResponse{}, fmt.Errorf("unable to convert postcode field to int, %v", err)
	}

	var datetime = time.Now()
	dt := datetime.Format(time.RFC3339)

	student := db.Student{
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		City:        req.City,
		Phone:       req.Phone,
		Email:       req.Email,
		Address:     req.Address,
		Postcode:    postcode,
		DateOfBirth: dt,
	}
	err = sdb.AddStudent(student)
	if err != nil {
		return AddResponse{}, fmt.Errorf("unable to add student, %v", err)
	}

	return AddResponse{}, nil
}
