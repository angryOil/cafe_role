package repository

import (
	"cafe_role/internal/domain"
	"cafe_role/internal/repository/model"
	"context"
	"errors"
	"github.com/uptrace/bun"
	"log"
	"strings"
)

type Repository struct {
	db bun.IDB
}

func NewRepository(db bun.IDB) Repository {
	return Repository{db: db}
}

func (r Repository) Create(ctx context.Context, d domain.Role) error {
	m := model.ToModel(d)
	_, err := r.db.NewInsert().Model(&m).Exec(ctx)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			return err
		}
		log.Println("Create NewInsert err: ", err)
		return errors.New("internal server error")
	}
	return nil
}
