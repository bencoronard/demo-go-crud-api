package resource

import "time"

type Resource struct {
	ID           int64
	Version      int64
	CreatedBy    int64
	CreatedAt    time.Time
	LastUpdated  time.Time
	TextField    *string
	NumberField  *int
	BooleanField *bool
}
