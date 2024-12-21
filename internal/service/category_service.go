package service

import (
	"context"

	"github.com/abdisetiakawan/go-restfulapi/internal/model/web"
)

type CategoryService interface {
	Create(ctx context.Context, request web.CategoryRequest) web.CategoryResponse
	Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse
	Delete(ctx context.Context, id int)
	FindById(ctx context.Context, id int) web.CategoryResponse
	FindAll(ctx context.Context) []web.CategoryResponse
}