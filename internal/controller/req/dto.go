package req

import (
	"cafe_role/internal/domain"
	"time"
)

type CreateDto struct {
	Name        string
	Description string
}

func (d CreateDto) ToDomain(cafeId int) domain.Role {
	return domain.Role{
		CafeId:      cafeId,
		Name:        d.Name,
		Description: d.Description,
		CreatedAt:   time.Now(),
	}
}
