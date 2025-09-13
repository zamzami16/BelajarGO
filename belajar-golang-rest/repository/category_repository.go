package repository

import (
	"belajar-go-rest/model/domain"
	"context"
	"database/sql"
)

type CategoryRepository interface {
	Save(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category
	Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category
	Delete(ctx context.Context, tx *sql.Tx, category domain.Category)
	FindById(ctx context.Context, tx *sql.Tx, category_id int) (domain.Category, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Category
}