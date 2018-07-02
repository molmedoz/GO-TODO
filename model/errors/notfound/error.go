package notfounderror

// NotFoundError Error that represent that a element wasn't found in the data base
type NotFoundError struct {
	model string
}

func (e *NotFoundError) Error() string {
	return "Element Not Found"
}

// New create a new Nor Found Error
func New(model string) *NotFoundError {
	return &NotFoundError{model: model}
}
