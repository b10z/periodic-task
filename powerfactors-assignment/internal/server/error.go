package server

func NewListenerError(text string, err error) error {
	return &ListenerError{text, err}
}

type ListenerError struct {
	data string
	err  error
}

func (re *ListenerError) Error() string {
	return re.data
}
