package services

import (
	ProductRepository "aio-commerce/internal/modules/product/repositories"
	ProductResponse "aio-commerce/internal/modules/product/response"
	"errors"
)

type ProductService struct {
	productRepository ProductRepository.ProductRepositoryInterface
}

func New() *ProductService {
	return &ProductService{
		productRepository: ProductRepository.New(),
	}
}

func (productService ProductService) GetFeaturedProducts() []ProductResponse.ProductResponse {
	products := productService.productRepository.Featured(20)
	return ProductResponse.ToProducts(products)

}

func (productService ProductService) GetStoriesProducts() []ProductResponse.ProductResponse {
	products := productService.productRepository.List(100)
	return ProductResponse.ToProducts(products)
}

func (productService ProductService) Find(code string) (ProductResponse.ProductResponse, error) {
	var response ProductResponse.ProductResponse
	product := productService.productRepository.Find(code)
	if product.Code == "" {
		return response, errors.New("Product Not Found")
	}
	return ProductResponse.ToProduct(product), nil
}
