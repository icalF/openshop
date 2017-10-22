package services

import (
	"github.com/icalF/openshop/models"
	"github.com/icalF/openshop/dao"
)

type ProductService interface {
	GetAll() []models.Product
	GetByID(id int64) (models.Product, bool)
	InsertOrUpdate(product models.Product) (models.Product, error)
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

func (s *productService) GetAll() []models.Product {
	return s.dao.SelectMany(map[string]string{}, 0)
}

func (s *productService) GetByID(id int64) (models.Product, bool) {
	return s.dao.Select(map[string]string{
		"id": string(id),
	})
}

func (s *productService) InsertOrUpdate(product models.Product) (models.Product, error) {
	return s.dao.InsertOrUpdate(product)
}

func (s *productService) DeleteByID(id int64) bool {
	return s.dao.Delete(map[string]string{
		"id": string(id),
	})
}