package domain

import "time"

type Business struct {
	ID        string
	Name      string
	Slug      Slug
	Timezone  Timezone
	CreatedAt time.Time
}
