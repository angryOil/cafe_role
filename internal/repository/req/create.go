package req

import "time"

type Create struct {
	CafeId      int
	Name        string
	Description string
	CreatedAt   time.Time
}
