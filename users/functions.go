package users

import "fmt"

var usrs Users

// New doc ...
func New(u Users) Users {
	return u
}

// Register doc ...
func Register(u Users) {
	usrs = u
}

// ValidateRegisterImplement doc ...
func ValidateRegisterImplement() error {
	if usrs != nil {
		return nil
	}
	return fmt.Errorf("The implementation of the 'Users' interface has not been registered")
}
