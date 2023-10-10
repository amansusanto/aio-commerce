package services

import "aio-commerce/internal/modules/product/response"

type ProductServiceInterface interface {
	GetFeaturedProducts() []response.ProductResponse
	GetStoriesProducts() []response.ProductResponse
	Find(code string) (response.ProductResponse, error)
}
