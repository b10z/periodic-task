package helper

func NewHelperError(text string) error {
	return &helperError{text}
}

type helperError struct {
	data string
}

func (re *helperError) Error() string {
	return re.data
}
