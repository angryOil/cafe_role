package service

import (
	"cafe_role/internal/domain"
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
