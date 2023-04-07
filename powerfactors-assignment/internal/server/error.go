package server

func NewServerError(text string, err error) error {
	return &serverError{text, err}
}

type serverError struct {
	data string
	err  error
}

func (re *serverError) Error() string {
	return re.data
}
