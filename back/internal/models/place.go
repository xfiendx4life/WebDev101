package models

import (
	"time"

	"github.com/google/uuid"
)

type Place struct {
	ID             uuid.UUID
	Name           string
	Location       [2]float64
	Opens          time.Time
	Closes         time.Time
	AdditionalInfo string
}
