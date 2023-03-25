package models

import "github.com/google/uuid"

type Team struct {
	ID uuid.UUID
	Name string
	Logo string
	AdditionalInfo string
	Users []UserWithStatuses
}