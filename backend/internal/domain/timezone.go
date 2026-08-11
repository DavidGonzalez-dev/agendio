package domain

import (
	"fmt"
	"time"
)

type Timezone struct {
	value string
}

func NewTimezone(raw string) (Timezone, error) {
	if _, err := time.LoadLocation(raw); err != nil {
		return Timezone{}, fmt.Errorf("invalid timezone: %q", raw)
	}
	return Timezone{value: raw}, nil
}

func (t Timezone) String() string {
	return t.value
}
