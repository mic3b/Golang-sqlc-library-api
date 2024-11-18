package model

import "time"

type Author struct {
	ID        int32
	FirstName string
	LastName  string
	CreatedAt time.Time
}
