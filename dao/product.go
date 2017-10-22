package dao

import (
	"github.com/icalF/openshop/models"
	"github.com/jinzhu/gorm"
)

type ProductDAO interface {
	Select(query Query) (model models.Product, found bool)
	SelectMany(query Query, limit int) (results []models.Product)

	InsertOrUpdate(model models.Product) (models.Product, error)
	Delete(query Query) (deleted bool)
}

type productRepository struct {
	source *gorm.DB
}

func NewProductDAO(connection *gorm.DB) ProductDAO {
	return &productRepository{source: connection}
}

func (r *productRepository) Select(query Query) (models.Product, bool) {
	product := models.Product{}
	if err := r.source.Where(query).First(&product).Error; err != nil {
		return models.Product{}, false
	}
	return product, true
}

func (r *productRepository) SelectMany(query Query, limit int) (results []models.Product) {
	product := new([]models.Product)
	r.source.Where(query).Find(&product).Limit(limit)
	return *product
}

func (r *productRepository) InsertOrUpdate(product models.Product) (_ models.Product, err error) {
	var oldProduct models.Product
	if err := r.source.First(&oldProduct).Error; err != nil {
		r.source.Create(&product)
	} else {
		r.source.Model(&oldProduct).Update(&product)
	}

	return product, err
}

func (r *productRepository) Delete(query Query) bool {
	if err := r.source.Delete(models.Product{}, query).Error; err != nil {
		return false
	}
	return true
}
