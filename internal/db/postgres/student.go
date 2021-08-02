package postgres

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/doniacld/first-step-sql/internal/db"
)

const (
	studentsTable = "students"
)

func (s *studentDB) AddStudent(st db.Student) error {
	sqlStatement := `INSERT INTO ` + studentsTable + `(first_name, last_name, city, phone, email, address, postcode, date_of_birth)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING id;`
	log.Printf("query row %s", sqlStatement)
	row := s.conn.QueryRow(sqlStatement, st.FirstName, st.LastName, st.City, st.Phone, st.Email, st.Address, st.Postcode, st.DateOfBirth)

	id := 0
	err := row.Scan(&id)
	switch err {
	case sql.ErrNoRows:
		return fmt.Errorf("no rows were returned, %s", err)
	case nil:
		log.Printf("New record ID is: %d", id)
	default:
		return fmt.Errorf("something wrong happened, %s", err)
	}

	return nil
}

func (s *studentDB) SearchByName(firstName string, lastName string) (db.Student, error) {
	var student db.Student
	sqlStatement := ` SELECT * FROM ` + studentsTable + `WHERE firstName=$1 AND lastName=$2;`
	row := s.conn.QueryRow(sqlStatement, firstName, lastName)

	err := row.Scan(&student.FirstName, &student.LastName, &student.DateOfBirth, &student.Postcode, &student.Address,
		&student.Email, &student.Phone, &student.City)
	switch err {
	case sql.ErrNoRows:
		return student, fmt.Errorf("no rows were returned, %s", err)
	case nil:
		log.Printf("Added student %#v", student)
	default:
		return student, fmt.Errorf("something wrong happened, %s", err)
	}
	return student, nil
}
