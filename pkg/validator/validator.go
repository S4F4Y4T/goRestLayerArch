package validator

import (
	"fmt"
	"net/mail"
	"strings"
)

type Error map[string]string

func (e Error) Error() string {
	var errs []string
	for k, v := range e {
		errs = append(errs, fmt.Sprintf("%s: %s", k, v))
	}
	return strings.Join(errs, ", ")
}

func ValidateEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

func IsEmpty(value string) bool {
	return strings.TrimSpace(value) == ""
}
