package helper

import (
	"github.com/abdisetiakawan/go-restfulapi/internal/model"
	"github.com/abdisetiakawan/go-restfulapi/internal/model/web"
)

func ToCategoryResponse(category model.Category) web.CategoryResponse {
	return web.CategoryResponse{ID: category.ID, Name: category.Name}
}

func ToCategoryResponses(categories []model.Category) []web.CategoryResponse {
	var responses []web.CategoryResponse
	for _, category := range categories {
		responses = append(responses, ToCategoryResponse(category))
	}
	return responses
}