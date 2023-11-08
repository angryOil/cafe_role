package controller

import (
	"cafe_role/internal/controller/req"
	"cafe_role/internal/controller/res"
	"cafe_role/internal/page"
	"cafe_role/internal/service"
	req2 "cafe_role/internal/service/req"
	res2 "cafe_role/internal/service/res"
	"context"
)

type Controller struct {
	s service.Service
}

func NewController(s service.Service) Controller {
	return Controller{s: s}
}

func (c Controller) Create(ctx context.Context, cafeId int, d req.CreateDto) error {
	err := c.s.Create(ctx, req2.Create{
		CafeId:      cafeId,
		Name:        d.Name,
		Description: d.Description,
	})
	return err
}

func (c Controller) GetList(ctx context.Context, cafeId int, reqPage page.ReqPage) ([]res.RoleDto, int, error) {
	listArr, total, err := c.s.GetList(ctx, cafeId, reqPage)
	if err != nil {
		return []res.RoleDto{}, 0, err
	}
	return convertGetListToDtoArr(listArr), total, err
}

func convertGetListToDtoArr(list []res2.GetList) []res.RoleDto {
	result := make([]res.RoleDto, len(list))
	for i, l := range list {
		result[i] = res.RoleDto{
			Id:          l.Id,
			Name:        l.Name,
			Description: l.Description,
		}
	}
	return result
}

func (c Controller) Patch(ctx context.Context, cafeId int, roleId int, d req.PatchDto) error {
	err := c.s.Patch(ctx, req2.Patch{
		Id:          roleId,
		CafeId:      cafeId,
		Name:        d.Name,
		Description: d.Description,
	})
	return err
}

func (c Controller) Delete(ctx context.Context, cafeId int, roleId int) error {
	err := c.s.Delete(ctx, cafeId, roleId)
	return err
}
