package controller

import (
	"cafe_role/internal/controller/req"
	"cafe_role/internal/service"
	"context"
)

type Controller struct {
	s service.Service
}

func NewController(s service.Service) Controller {
	return Controller{s: s}
}

func (c Controller) Create(ctx context.Context, cafeId int, d req.CreateDto) error {
	rDomain := d.ToDomain(cafeId)
	err := c.s.Create(ctx, rDomain)
	return err
}
