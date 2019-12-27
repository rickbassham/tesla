package tesla

import (
	"errors"
	"fmt"
)

// HTTPStatusError is returned if the response code received from the API is non-200.
type HTTPStatusError struct {
	statusCode int
	message    string
}

func (err HTTPStatusError) Error() string {
	if err.message == "" {
		return fmt.Sprintf("http status error: %d", err.statusCode)
	}

	return fmt.Sprintf("http status error \"%s\": %d", err.message, err.statusCode)
}

var (
	// ErrMissingRefreshToken is returned when an API call is made without the required refresh token.
	ErrMissingRefreshToken = errors.New("missing refresh token")
	// ErrMissingAccessToken is returned when an API call is made without the required access token.
	ErrMissingAccessToken = errors.New("missing access token, authenticate first")
	// ErrCommandError is returned with executing a command against the vehicle and the Tesla API returns an error message.
	ErrCommandError = errors.New("error executing command")
)
