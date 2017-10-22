package services

import (
	"github.com/icalF/openshop/models/datamodels"
	"github.com/icalF/openshop/dao"
)

type ProductService interface {
	GetAll() []datamodels.Product
	GetAllPurchaseable() []datamodels.Product
	GetByID(id int64) (datamodels.Product, bool)
	InsertOrUpdate(product datamodels.Product) (datamodels.Product, error)
	DeleteByID(id int64) bool
}

func NewProductService(dao dao.ProductDAO) ProductService {
	return &productService{
		dao: dao,
	}
}

type productService struct {
	dao dao.ProductDAO
}

func (s *productService) GetAll() []datamodels.Product {
	return s.dao.SelectMany(map[string]string{}, 1, 0)
}

func (s *productService) GetAllPurchaseable() []datamodels.Product {
	return s.dao.SelectMany(map[string]string{}, 1, 1)
}

func (s *productService) GetByID(id int64) (datamodels.Product, bool) {
	return s.dao.Select(map[string]string{
		"id": string(id),
	})
}

func (s *productService) InsertOrUpdate(product datamodels.Product) (datamodels.Product, error) {
	return s.dao.InsertOrUpdate(product)
}

func (s *productService) DeleteByID(id int64) bool {
	return s.dao.Delete(map[string]string{
		"id": string(id),
	})
}