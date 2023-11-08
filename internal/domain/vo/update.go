package vo

import "time"

type Update struct {
	Id          int
	CafeId      int
	Name        string
	Description string
	CreateAt    time.Time
}
