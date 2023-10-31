package controller

import (
	"cafe_role/internal/controller/req"
	"cafe_role/internal/controller/res"
	"cafe_role/internal/page"
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

func (c Controller) GetList(ctx context.Context, cafeId int, reqPage page.ReqPage) ([]res.RoleDto, int, error) {
	domains, total, err := c.s.GetList(ctx, cafeId, reqPage)
	if err != nil {
		return []res.RoleDto{}, 0, err
	}
	return res.ToRoleDtoList(domains), total, err
}

func (c Controller) Patch(ctx context.Context, cafeId int, roleId int, d req.PatchDto) error {
	rDomain := d.ToDomain(cafeId, roleId)
	err := c.s.Patch(ctx, rDomain)
	return err
}

func (c Controller) Delete(ctx context.Context, cafeId int, roleId int) error {
	err := c.s.Delete(ctx, cafeId, roleId)
	return err
}
