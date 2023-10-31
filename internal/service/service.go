package service

import (
	"cafe_role/internal/domain"
	"cafe_role/internal/page"
	"cafe_role/internal/repository"
	"context"
	"errors"
)

type Service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) Service {
	return Service{repo: repo}
}

func (s Service) Create(ctx context.Context, d domain.Role) error {
	if d.Name == "" {
		return errors.New("invalid name")
	}
	err := s.repo.Create(ctx, d)
	return err
}

func (s Service) GetList(ctx context.Context, cafeId int, reqPage page.ReqPage) ([]domain.Role, int, error) {
	domains, total, err := s.repo.GetList(ctx, cafeId, reqPage)
	return domains, total, err
}

func (s Service) Patch(ctx context.Context, rDomain domain.Role) error {
	if rDomain.Name == "" {
		return errors.New("invalid name")
	}
	err := s.repo.Patch(ctx, rDomain.CafeId, rDomain.Id,
		func(domains []domain.Role) (domain.Role, error) {
			if len(domains) == 0 {
				return domain.Role{}, errors.New("no rows")
			}
			return domains[0], nil
		},
		func(findDomain domain.Role) domain.Role {
			return domain.Role{
				Id:          findDomain.Id,
				CafeId:      findDomain.CafeId,
				Name:        rDomain.Name,
				Description: rDomain.Description,
				CreatedAt:   findDomain.CreatedAt,
			}
		},
	)
	return err
}

func (s Service) Delete(ctx context.Context, cafeId int, roleId int) error {
	err := s.repo.Delete(ctx, cafeId, roleId)
	return err
}
