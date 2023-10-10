package repositories

import (
	ProductModel "aio-commerce/internal/modules/product/models"
)

type ProductRepositoryInterface interface {
	Featured(limit int) []ProductModel.Product
	List(limit int) []ProductModel.Product
	Find(code string) ProductModel.Product
}
