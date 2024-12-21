package repository

import (
	"context"
	"database/sql"

	"github.com/abdisetiakawan/go-restfulapi/internal/model"
)

type CategoryRepository interface {
    Save(ctx context.Context, tx sql.Tx, category model.Category) model.Category
    Update(ctx context.Context, tx sql.Tx, category model.Category) model.Category
    Delete(ctx context.Context, tx sql.Tx, category model.Category)
    FindById(ctx context.Context, tx sql.Tx, id uint) model.Category
    FindAll(ctx context.Context, tx sql.Tx) []model.Category
}