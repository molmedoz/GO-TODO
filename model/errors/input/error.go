package inputerror

// InputError Erro when there a field validation error
type InputError struct {
	Fields  []string `json:"fields"`
	Message string   `json:"message"`
}

func (e *InputError) Error() string {
	return e.Message
}

// New generate a new InputError
func New(fields []string) *InputError {

	return &InputError{Fields: fields, Message: "Invalid Input"}
}
