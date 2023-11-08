package domain

import (
	"cafe_role/internal/domain/vo"
	"errors"
	"time"
)

var _ Role = (*role)(nil)

type Role interface {
	ValidCreate() error
	ValidUpdate() error

	Update(name, description string) Role

	ToUpdate() vo.Update
	ToInfo() vo.Info
	ToDetail() vo.Detail
}

type role struct {
	id          int
	cafeId      int
	name        string
	description string
	createdAt   time.Time
}

const (
	InvalidId     = "invalid id"
	InvalidCafeId = "invalid cafe id"
	InvalidName   = "invalid name"
)

func (r *role) ValidCreate() error {
	if r.cafeId < 1 {
		return errors.New(InvalidCafeId)
	}
	if r.name == "" {
		return errors.New(InvalidName)
	}
	return nil
}

func (r *role) ValidUpdate() error {
	if r.id < 1 {
		return errors.New(InvalidId)
	}
	if r.cafeId < 1 {
		return errors.New(InvalidCafeId)
	}
	if r.name == "" {
		return errors.New(InvalidName)
	}
	return nil
}

func (r *role) Update(name, description string) Role {
	r.name = name
	r.description = description
	return r
}

func (r *role) ToUpdate() vo.Update {
	return vo.Update{
		Id:          r.id,
		CafeId:      r.cafeId,
		Name:        r.name,
		Description: r.description,
		CreateAt:    r.createdAt,
	}
}

func (r *role) ToInfo() vo.Info {
	return vo.Info{
		Id:   r.id,
		Name: r.name,
	}
}

func (r *role) ToDetail() vo.Detail {
	return vo.Detail{
		Id:          r.id,
		Name:        r.name,
		Description: r.description,
	}
}
