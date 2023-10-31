package repository

import (
	"cafe_role/internal/domain"
	"cafe_role/internal/page"
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

func (r Repository) GetList(ctx context.Context, cafeId int, reqPage page.ReqPage) ([]domain.Role, int, error) {
	var models []model.Role
	cnt, err := r.db.NewSelect().Model(&models).
		ColumnExpr("name,id,substring(cr.description,1,50) as description").
		Limit(reqPage.Size).Offset(reqPage.OffSet).Order("id desc").
		Where("cafe_id = ?", cafeId).
		ScanAndCount(ctx)

	if err != nil {
		log.Println("GetList NewSelect err: ", err)
		return []domain.Role{}, 0, errors.New("internal server error")
	}

	return model.ToDomainList(models), cnt, nil
}

func (r Repository) Patch(ctx context.Context, cafeId int, roleId int,
	validFunc func(domains []domain.Role) (domain.Role, error),
	mergeFunc func(findDomain domain.Role) domain.Role) error {
	var models []model.Role
	err := r.db.NewSelect().Model(&models).Where("id = ? and cafe_id = ?", roleId, cafeId).Scan(ctx)
	if err != nil {
		log.Println("Patch NewSelect err: ", err)
		return errors.New("internal server error")
	}

	findDomains := model.ToDomainList(models)
	validDomain, err := validFunc(findDomains)
	if err != nil {
		return err
	}
	mergedDomain := mergeFunc(validDomain)
	mergedModel := model.ToModel(mergedDomain)
	_, err = r.db.NewInsert().Model(&mergedModel).On("CONFLICT (id) DO UPDATE").Exec(ctx)

	if err != nil {
		log.Println("Patch NewInsert err: ", err)
		return errors.New("internal server error")
	}
	return nil
}

func (r Repository) Delete(ctx context.Context, cafeId int, roleId int) error {
	var m model.Role
	_, err := r.db.NewDelete().Model(&m).Where("id = ? and cafe_id = ?", roleId, cafeId).Exec(ctx)
	if err != nil {
		log.Println("Delete NewDelete err: ", err)
		return errors.New("invalid server error")
	}
	return nil
}
