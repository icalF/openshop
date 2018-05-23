package services

import (
	"github.com/koneko096/openshop/models/datamodels"
	"github.com/koneko096/openshop/dao"
)


func NewProductManager(dao dao.ProductDAO) ProductManager {
	return &productService{
		dao: dao,
	}
}


type productService struct {
	dao dao.ProductDAO
}



func (s *productService) GetAll() []datamodels.Product {
	return s.dao.SelectMany(map[string]string{}, 0, 0)
}

func (s *productService) GetAllPurchasable() []datamodels.Product {
	return s.dao.SelectMany(map[string]string{}, 0, 1)
}

func (s *productService) GetByID(id int64) (datamodels.Product, bool) {
	return s.dao.Select(map[string]string{
		"id": string(id),
	})
}

func (s *productService) GetByIDs(ids []int64) []datamodels.Product {
	return s.dao.SelectMany(ids, 0, 0)
}

func (s *productService) InsertOrUpdate(product datamodels.Product) (datamodels.Product, error) {
	return s.dao.InsertOrUpdate(product)
}

func (s *productService) DeleteByID(id int64) bool {
	return s.dao.Delete(map[string]string{
		"id": string(id),
	})
}

func (s *productService) CreateProductMap(products []datamodels.Product) map[int64]datamodels.Product {
	productMap := map[int64]datamodels.Product{}
	for _, product := range products {
		productMap[product.ID] = product
	}
	return productMap
}