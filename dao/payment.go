package dao

import (
	"github.com/icalF/openshop/models"
	"github.com/jinzhu/gorm"
)

type PaymentDAO interface {
	Select(query Query) (model models.Payment, found bool)
	SelectMany(query Query, limit int) (results []models.Payment)

	InsertOrUpdate(model models.Payment) (models.Payment, error)
	Delete(query Query) (deleted bool)
}

type paymentRepository struct {
	source *gorm.DB
}

func NewPaymentDAO(connection *gorm.DB) PaymentDAO {
	return &paymentRepository{source: connection}
}

func (r *paymentRepository) Select(query Query) (models.Payment, bool) {
	payment := models.Payment{}
	if err := r.source.Where(query).First(&payment).Error; err != nil {
		return models.Payment{}, false
	}
	return payment, true
}

func (r *paymentRepository) SelectMany(query Query, limit int) (results []models.Payment) {
	payments := new([]models.Payment)
	r.source.Where(query).Find(&payments).Limit(limit)
	return *payments
}

func (r *paymentRepository) InsertOrUpdate(payment models.Payment) (_ models.Payment, err error) {
	var oldPayment models.Payment
	if err := r.source.First(&oldPayment).Error; err != nil {
		r.source.Create(&payment)
	} else {
		r.source.Model(&oldPayment).Update(&payment)
	}

	return payment, err
}

func (r *paymentRepository) Delete(query Query) bool {
	if err := r.source.Delete(models.Payment{}, query).Error; err != nil {
		return false
	}
	return true
}
