package domain

import (
	"fmt"
	"regexp"
	"strings"
)

var emailPattern = regexp.MustCompile(`^[^\s@]+@[^\s@]+\.[^\s@]+$`)

type Email struct {
	value string
}

func NewEmail(raw string) (Email, error) {
	normalized := strings.ToLower(strings.TrimSpace(raw))
	if !emailPattern.MatchString(normalized) {
		return Email{}, fmt.Errorf("invalid email: %q", raw)
	}
	return Email{value: normalized}, nil
}

func (e Email) String() string {
	return e.value
}
