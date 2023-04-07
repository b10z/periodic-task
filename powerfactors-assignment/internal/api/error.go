package api

func NewAPIError(text string) error {
	return &apiError{text}
}

type apiError struct {
	data string
}

func (re *apiError) Error() string {
	return re.data
}
