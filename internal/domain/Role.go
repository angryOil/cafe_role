package domain

import "time"

type Role struct {
	Id          int
	CafeId      int
	Name        string
	Description string
	CreatedAt   time.Time
}
