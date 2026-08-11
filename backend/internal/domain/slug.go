package domain

import (
	"fmt"
	"regexp"
)

var slugPattern = regexp.MustCompile(`^[a-z0-9]+(-[a-z0-9]+)*$`)

type Slug struct {
	value string
}

func NewSlug(raw string) (Slug, error) {
	if !slugPattern.MatchString(raw) {
		return Slug{}, fmt.Errorf("invalid slug: %q (expected lowercase letters, numbers and hyphens)", raw)
	}
	return Slug{value: raw}, nil
}

func (s Slug) String() string {
	return s.value
}
