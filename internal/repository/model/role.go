package model

import (
	"cafe_role/internal/domain"
	"cafe_role/internal/repository/req"
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

func ToCreateModel(c req.Create) Role {
	return Role{
		CafeId:      c.CafeId,
		Name:        c.Name,
		Description: c.Description,
		CreatedAt:   c.CreatedAt,
	}
}
func ToUpdateModel(u req.Update) Role {
	return Role{
		Id:          u.Id,
		CafeId:      u.CafeId,
		Name:        u.Name,
		Description: u.Description,
		CreatedAt:   u.CreatedAt,
	}
}
func (d Role) ToDomain() domain.Role {
	return domain.NewRoleBuilder().
		Id(d.Id).
		CafeId(d.CafeId).
		Name(d.Name).
		Description(d.Description).
		CreatedAt(d.CreatedAt).
		Build()
}

func ToDomainList(models []Role) []domain.Role {
	results := make([]domain.Role, len(models))
	for i, m := range models {
		results[i] = m.ToDomain()
	}
	return results
}
