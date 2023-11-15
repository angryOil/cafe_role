package service

import (
	"cafe_role/internal/domain"
	"cafe_role/internal/domain/vo"
	"cafe_role/internal/page"
	"cafe_role/internal/repository"
	req2 "cafe_role/internal/repository/req"
	"cafe_role/internal/service/req"
	"cafe_role/internal/service/res"
	"context"
	"errors"
	"time"
)

type Service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) Service {
	return Service{repo: repo}
}

func (s Service) Create(ctx context.Context, c req.Create) error {
	cafeId := c.CafeId
	name, description := c.Name, c.Description
	createdAt := time.Now()

	err := domain.NewRoleBuilder().
		CafeId(cafeId).
		Name(name).
		Description(description).
		CreatedAt(createdAt).
		Build().ValidCreate()
	if err != nil {
		return err
	}

	err = s.repo.Create(ctx, req2.Create{
		CafeId:      cafeId,
		Name:        name,
		Description: description,
		CreatedAt:   createdAt,
	})
	return err
}

func (s Service) GetList(ctx context.Context, cafeId int, reqPage page.ReqPage) ([]res.GetList, int, error) {
	domains, total, err := s.repo.GetList(ctx, cafeId, reqPage)
	if err != nil {
		return []res.GetList{}, 0, err
	}
	dto := make([]res.GetList, len(domains))
	for i, d := range domains {
		v := d.ToDetail()
		dto[i] = res.GetList{
			Id:          v.Id,
			Name:        v.Name,
			Description: v.Description,
		}
	}
	return dto, total, err
}

func (s Service) Patch(ctx context.Context, p req.Patch) error {
	id := p.Id
	name, description := p.Name, p.Description

	err := s.repo.Patch(ctx, id,
		func(domains []domain.Role) (domain.Role, error) {
			if len(domains) != 1 {
				return domain.NewRoleBuilder().Build(), errors.New("no rows")
			}
			return domains[0], nil
		},
		func(findD domain.Role) (vo.Update, error) {
			u, err := findD.Update(name, description)
			if err != nil {
				return vo.Update{}, err
			}
			return u.ToUpdate(), nil
		},
	)
	return err
}

func (s Service) Delete(ctx context.Context, cafeId int, roleId int) error {
	err := s.repo.Delete(ctx, cafeId, roleId)
	return err
}
