package value_objects

import (
	"errors"
	"regexp"
)

type Email string

func NewEmail(email string) (Email, error) {
	if !isValidEmail(email) {
		return "", errors.New("Invalid email!")
	}
	return Email(email), nil
}

func isValidEmail(email string) bool {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	regex := regexp.MustCompile(pattern)
	return regex.MatchString(email)
}
