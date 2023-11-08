package req

import "time"

type Update struct {
	Id          int
	CafeId      int
	Name        string
	Description string
	CreatedAt   time.Time
}
