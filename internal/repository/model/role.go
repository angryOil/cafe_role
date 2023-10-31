package model

import (
	"cafe_role/internal/domain"
	"github.com/uptrace/bun"
	"time"
)

type Role struct {
	bun.BaseModel `bun:"table:cafe_role,alias:cr"`

	Id          int       `bun:"id,pk,autoincrement"`
	CafeId      int       `bun:"cafe_id,notnull"`
	Name        string    `bun:"name,notnull"`
	Description string    `bun:"description"`
	CreatedAt   time.Time `bun:"created_at"`
}

func (d Role) ToDomain() domain.Role {
	return domain.Role{
		Id:          d.Id,
		CafeId:      d.CafeId,
		Name:        d.Name,
		Description: d.Description,
		CreatedAt:   d.CreatedAt,
	}
}

func ToDomainList(models []Role) []domain.Role {
	results := make([]domain.Role, len(models))
	for i, m := range models {
		results[i] = m.ToDomain()
	}
	return results
}
