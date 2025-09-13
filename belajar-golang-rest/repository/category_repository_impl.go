package repository

import (
	"belajar-go-rest/helper"
	"belajar-go-rest/model/domain"
	"context"
	sqlpkg "database/sql"
	"errors"
)

type CategoryRepositoryImpl struct {
}

func NewCategoryRepository() CategoryRepository {
	return &CategoryRepositoryImpl{}
}

func (repository *CategoryRepositoryImpl) Save(ctx context.Context, tx *sqlpkg.Tx, category domain.Category) domain.Category {
	sql := "insert into category (name) values ($1) RETURNING id;"
	err := tx.QueryRowContext(ctx, sql, category.Name).Scan(&category.Id)
	helper.PanicIfError(err)
	return category
}

func (repository *CategoryRepositoryImpl) Update(ctx context.Context, tx *sqlpkg.Tx, category domain.Category) domain.Category {
	sql := "update category set name = $1 where id = $2;"
	_, err := tx.ExecContext(ctx, sql, category.Name, category.Id)
	helper.PanicIfError(err)

	return category
}

func (repository *CategoryRepositoryImpl) Delete(ctx context.Context, tx *sqlpkg.Tx, category domain.Category) {
	sql := "delete from category where id = $1;"
	_, err := tx.ExecContext(ctx, sql, category.Id)
	helper.PanicIfError(err)
}

func (repository *CategoryRepositoryImpl) FindById(ctx context.Context, tx *sqlpkg.Tx, category_id int) (domain.Category, error) {
	sql := "select id, name from category where id = $1;"
	row := tx.QueryRowContext(ctx, sql, category_id)

	category := domain.Category{}
	err := row.Scan(&category.Id, &category.Name)
	if err != nil {
		if err == sqlpkg.ErrNoRows {
			return category, errors.New("category not found")
		}
		helper.PanicIfError(err)
	}
	return category, nil
}

func (repository *CategoryRepositoryImpl) FindAll(ctx context.Context, tx *sqlpkg.Tx) []domain.Category {
	sql := "select id, name from category;"
	row, err := tx.QueryContext(ctx, sql)
	helper.PanicIfError(err)

	categories := []domain.Category{}
	for row.Next() {
		category := domain.Category{}
		err := row.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)

		categories = append(categories, category)
	}

	return categories
}
