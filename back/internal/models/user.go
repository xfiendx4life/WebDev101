package models

import (
	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID
	Name     string
	Password string
	BIO      string
}

type UserStatus string

const (
	StatusCreator UserStatus = "creator"
	StatusParticipant UserStatus = "participant"
)

type UserWithStatuses struct {
	User User
	Statuses []UserStatus
}
