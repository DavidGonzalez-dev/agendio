package domain

import "time"

type Role string

const (
	RoleClient Role = "client"
	RoleAdmin  Role = "admin"
)

type User struct {
	ID           string
	Name         string
	Email        Email
	PasswordHash string
	Role         Role
	BusinessID   *string
	CreatedAt    time.Time
}
