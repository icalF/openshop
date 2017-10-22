package dao

import (
	"github.com/icalF/openshop/models/datamodels"
	"github.com/jinzhu/gorm"
)

type ProductDAO interface {
	Select(query Query) (model datamodels.Product, found bool)
	SelectMany(query Query, limit int) (results []datamodels.Product)

	InsertOrUpdate(model datamodels.Product) (datamodels.Product, error)
	Delete(query Query) (deleted bool)
}

type productRepository struct {
	source *gorm.DB
}

func NewProductDAO(connection *gorm.DB) ProductDAO {
	return &productRepository{source: connection}
}

func (r *productRepository) Select(query Query) (datamodels.Product, bool) {
	product := datamodels.Product{}
	if err := r.source.Where(query).First(&product).Error; err != nil {
		return datamodels.Product{}, false
	}
	return product, true
}

func (r *productRepository) SelectMany(query Query, limit int) (results []datamodels.Product) {
	product := new([]datamodels.Product)
	r.source.Where(query).Find(&product).Limit(limit)
	return *product
}

func (r *productRepository) InsertOrUpdate(product datamodels.Product) (_ datamodels.Product, err error) {
	var oldProduct datamodels.Product
	if err := r.source.First(&oldProduct).Error; err != nil {
		r.source.Create(&product)
	} else {
		r.source.Model(&oldProduct).Update(&product)
	}

	return product, err
}

func (r *productRepository) Delete(query Query) bool {
	if err := r.source.Delete(datamodels.Product{}, query).Error; err != nil {
		return false
	}
	return true
}
