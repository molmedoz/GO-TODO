package autherror

// AuthError Erro when there a field validation error
type AuthError struct {
}

func (e *AuthError) Error() string {
	return "Unhathorithed"
}

// New generate a new InputError
func New() *AuthError {
	return &AuthError{}
}
