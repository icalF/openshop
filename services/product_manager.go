package services

import "github.com/koneko096/openshop/models/datamodels"

type ProductManager interface {
	GetAll() []datamodels.Product
	GetAllPurchasable() []datamodels.Product
	GetByID(id int64) (datamodels.Product, bool)
	GetByIDs(ids []int64) []datamodels.Product
	InsertOrUpdate(product datamodels.Product) (datamodels.Product, error)
	DeleteByID(id int64) bool
	CreateProductMap(products []datamodels.Product) map[int64]datamodels.Product
}
