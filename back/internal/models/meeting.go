package models

import (
	"time"

	"github.com/google/uuid"
)


type Meeting struct {
	ID             uuid.UUID
	Name           string
	Place          Place
	Time           time.Time
	AdditionalInfo string
	Users []UserWithStatuses
}

