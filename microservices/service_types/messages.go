package service_types

import "fmt"

const (
	LoginMsg           = "Please login"
	EmailNotRegistered = "The email is not registered"
)

const (
	ErrEmailNotRegistered = "Cannot find an account registered with the provided email"
)

func CannotCreateToken(val string, err error) string {
	return fmt.Sprintf("cannot create a token for %s, error: %+v", val, err)
}
