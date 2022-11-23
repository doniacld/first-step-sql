package rest

// Meta holds endpoint metadata
type Meta interface {
	Path() string
	Verb() string
	SuccessCode() int
}

// New creates a new endpoint structure
func New(path, verb string, code int) Meta {
	e := endpoint{
		path:        path,
		verb:        verb,
		successCode: code,
	}
	return &e
}

// endpoint holds information about it
type endpoint struct {
	path        string
	verb        string
	successCode int
}

// Path returns the path
func (e *endpoint) Path() string {
	return e.path
}

// Verb returns the verb
func (e *endpoint) Verb() string {
	return e.verb
}

// SuccessCode returns the success code
func (e *endpoint) SuccessCode() int {
	return e.successCode
}
