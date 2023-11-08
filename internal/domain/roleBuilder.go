package domain

import "time"

var _ RoleBuilder = (*roleBuilder)(nil)

func NewRoleBuilder() RoleBuilder {
	return &roleBuilder{}
}

type RoleBuilder interface {
	Id(id int) RoleBuilder
	CafeId(cafeId int) RoleBuilder
	Name(name string) RoleBuilder
	Description(description string) RoleBuilder
	CreatedAt(createdAt time.Time) RoleBuilder

	Build() Role
}

type roleBuilder struct {
	id          int
	cafeId      int
	name        string
	description string
	createdAt   time.Time
}

func (r *roleBuilder) Id(id int) RoleBuilder {
	r.id = id
	return r
}

func (r *roleBuilder) CafeId(cafeId int) RoleBuilder {
	r.cafeId = cafeId
	return r
}

func (r *roleBuilder) Name(name string) RoleBuilder {
	r.name = name
	return r
}

func (r *roleBuilder) Description(description string) RoleBuilder {
	r.description = description
	return r
}

func (r *roleBuilder) CreatedAt(createdAt time.Time) RoleBuilder {
	r.createdAt = createdAt
	return r
}

func (r *roleBuilder) Build() Role {
	return &role{
		id:          r.id,
		cafeId:      r.cafeId,
		name:        r.name,
		description: r.description,
		createdAt:   r.createdAt,
	}
}
