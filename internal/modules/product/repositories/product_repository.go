package repositories

import (
	ProductModel "aio-commerce/internal/modules/product/models"
	"aio-commerce/pkg/database"
	"database/sql"
	"gorm.io/gorm"
	"log"
	"strconv"
)

type ProductRepository struct {
	DB *gorm.DB
}

func New() *ProductRepository {
	return &ProductRepository{
		DB: database.Connection(),
	}
}

func (productRepository ProductRepository) Featured(limit int) []ProductModel.Product {
	var products []ProductModel.Product
	Query := "select top " + strconv.Itoa(limit) + " s.code, substring(s.description,1,50) as Title, s.description, sum(sp.qty) as qty, s.Price, " +
		"s.Price1, substring(s.Notes,1,100) as ShortNotes, s.Notes from stock s left outer join StockPos sp on s.code=sp.code  " +
		"where s.code in (select distinct code from jualdtl where jam >dateadd(d,-30,GETDATE())) " +
		"group by s.code, s.description, s.Price, s.Price1, s.Notes"

	productRepository.DB.Raw(Query).Scan(&products)
	return products
}

func (productRepository ProductRepository) List(limit int) []ProductModel.Product {
	var products []ProductModel.Product
	Query := "select s.code, substring(s.description,1,50) as Title, s.description, sum(sp.qty) as qty, s.Price, s.Price1, substring(s.Notes,1,100) as ShortNotes, " +
		"s.Notes from stock s left outer join StockPos sp on " +
		"s.code=sp.code where s.code in (select top " + strconv.Itoa(limit) + " code from (select code, sum(qty) as qty " +
		"from jualdtl where jam >dateadd(d,-60,GETDATE()) group by code) a " +
		"order by qty desc) group by s.code, s.description, s.Price, s.Price1, s.Notes"
	productRepository.DB.Raw(Query).Scan(&products)
	return products
}

func (productRepository ProductRepository) Find(code string) ProductModel.Product {
	var product ProductModel.Product
	qry := "select s.code, substring(s.description,1,50) as Title, s.description, sum(sp.qty) as qty, s.Price, s.Price1, substring(s.Notes,1,100) as ShortNotes, s.Notes from stock s left outer join " +
		"StockPos sp on s.code=sp.code where s.code=@code " +
		"group by s.code, s.description, s.Price, s.Price1, s.Notes"
	log.Println(qry)
	productRepository.DB.Raw(qry, sql.Named("code", code)).Scan(&product)
	return product
}
