package service

func NewServiceError(text string) error {
	return &serviceError{text}
}

type serviceError struct {
	data string
}

func (re *serviceError) Error() string {
	return re.data
}
