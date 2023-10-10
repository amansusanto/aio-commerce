package response

import ProductModel "aio-commerce/internal/modules/product/models"

type ProductResponse struct {
	Code        string
	Title       string
	Description string
	Qty         float32
	Price       float32
	Price1      float32
	ShortNotes  string
	Notes       string
	UserName    string
}

func ToProduct(product ProductModel.Product) ProductResponse {
	return ProductResponse{
		Code:        product.Code,
		Title:       product.Title,
		Description: product.Description,
		Qty:         product.Qty,
		Price:       product.Price,
		Price1:      product.Price1,
		ShortNotes:  product.ShortNotes,
		Notes:       product.Notes,
	}
}
func ToProducts(products []ProductModel.Product) []ProductResponse {
	var response []ProductResponse

	for _, product := range products {
		response = append(response, ToProduct(product))
	}
	return response
}
