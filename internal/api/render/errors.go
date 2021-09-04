package render

var (
	// ErrInvalidToken is returned when the api request token is invalid.
	ErrInvalidToken = New("Invalid or missing token")

	// ErrUnauthorized is returned when the user is not authorized.
	ErrUnauthorized = New("Unauthorized")

	// ErrForbidden is returned when user access is forbidden.
	ErrForbidden = New("Forbidden")

	// ErrNotFound is returned when a resource is not found.
	ErrNotFound = New("Not Found")

	// ErrEmptyRequestBody is returned when a request body is empty
	ErrEmptyRequestBody = New("Empty request body")

	// ErrInternal is returned when an unknown error has occured
	ErrInternal = New("Internal Error")

	// ErrTypeQueryParameterNotSpecified is returned when the type query param is empty
	ErrTypeQueryParameterNotSpecified = New("type query parameter must be specified")

	// ErrIdsQueryParameterNotSpecified is returned when the ids query param is empty
	ErrIdsQueryParameterNotSpecified = New("ids query parameter must be specified")

	// ErrIdsQueryParameterInvalid is returned when the ids query param contains an invalid value
	ErrIdsQueryParameterInvalid = New("ids query parameter must contain only integer IDs")
)

// Error represents a json-encoded API error.
type Error struct {
	Message string `json:"message"`
}

func (e *Error) Error() string {
	return e.Message
}

// New returns a new error message.
func New(text string) error {
	return &Error{Message: text}
}
