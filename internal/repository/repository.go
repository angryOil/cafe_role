package repository

import (
	"cafe_role/internal/domain"
	"cafe_role/internal/domain/vo"
	"cafe_role/internal/page"
	"cafe_role/internal/repository/model"
	"cafe_role/internal/repository/req"
	"context"
	"errors"
	"github.com/uptrace/bun"
	"log"
)

type Repository struct {
	db bun.IDB
}

func NewRepository(db bun.IDB) Repository {
	return Repository{db: db}
}

const (
	InternalServerError = "internal server error"
)

func (r Repository) Create(ctx context.Context, c req.Create) error {
	m := model.ToCreateModel(c)
	_, err := r.db.NewInsert().Model(&m).Exec(ctx)

	return err
}

func (r Repository) GetList(ctx context.Context, cafeId int, reqPage page.ReqPage) ([]domain.Role, int, error) {
	var models []model.Role
	cnt, err := r.db.NewSelect().Model(&models).
		ColumnExpr("name,id,substring(cr.description,1,50) as description").
		//Limit(reqPage.Size).Offset(reqPage.OffSet).Order("id desc").
		Where("cafe_id = ?", cafeId).
		ScanAndCount(ctx)

	if err != nil {
		log.Println("GetList NewSelect err: ", err)
		return []domain.Role{}, 0, errors.New(InternalServerError)
	}

	return model.ToDomainList(models), cnt, nil
}

func (r Repository) Patch(ctx context.Context, id int,
	validFunc func(domains []domain.Role) (domain.Role, error),
	mergeFunc func(findDomain domain.Role) (vo.Update, error)) error {

	db, err := r.db.BeginTx(ctx, nil)
	if err != nil {
		log.Println("patch BeginTx err: ", err)
		return errors.New(InternalServerError)
	}
	var models []model.Role
	err = db.NewSelect().Model(&models).Where("id = ? ", id).Scan(ctx)
	if err != nil {
		log.Println("Patch NewSelect err: ", err)
		return errors.New(InternalServerError)
	}

	findDomains := model.ToDomainList(models)
	validDomain, err := validFunc(findDomains)
	if err != nil {
		return err
	}

	mergedVo, err := mergeFunc(validDomain)
	if err != nil {
		return err
	}

	mergedModel := model.ToUpdateModel(req.Update{
		Id:          mergedVo.Id,
		CafeId:      mergedVo.CafeId,
		Name:        mergedVo.Name,
		Description: mergedVo.Description,
		CreatedAt:   mergedVo.CreateAt,
	})

	_, err = db.NewInsert().Model(&mergedModel).On("CONFLICT (id) DO UPDATE").Exec(ctx)

	if err != nil {
		return err
	}

	err = db.Commit()
	if err != nil {
		log.Println("patch db.Commit err: ", err)
		return errors.New(InternalServerError)
	}
	return nil
}

func (r Repository) Delete(ctx context.Context, cafeId int, roleId int) error {
	var m model.Role
	_, err := r.db.NewDelete().Model(&m).Where("id = ? and cafe_id = ?", roleId, cafeId).Exec(ctx)
	if err != nil {
		log.Println("Delete NewDelete err: ", err)
		return errors.New(InternalServerError)
	}
	return nil
}
