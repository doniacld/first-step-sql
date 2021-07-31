package db

// Student defines a student identity
type Student struct {
	FirstName   string
	LastName    string
	City        string
	Phone       string
	Email       string
	Address     string
	Postcode    int
	DateOfBirth string
}

// StudentDB provides different methods to the db
type StudentDB interface {
	AddStudent(s Student)
	SearchByName(firstName string, lastName string) (Student, error)

	Close() error
}
