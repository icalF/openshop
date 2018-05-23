package dao

import (
	"github.com/koneko096/openshop/models/datamodels"
	"github.com/jinzhu/gorm"
)

type PaymentDAO interface {
	Select(query Query) (model datamodels.Payment, found bool)
	SelectMany(query Query, limit int) (results []datamodels.Payment)

	InsertOrUpdate(model datamodels.Payment) (datamodels.Payment, error)
	Delete(query Query) (deleted bool)
}

type paymentRepository struct {
	source *gorm.DB
}

func NewPaymentDAO(connection *gorm.DB) PaymentDAO {
	return &paymentRepository{source: connection}
}

func (r *paymentRepository) Select(query Query) (datamodels.Payment, bool) {
	payment := datamodels.Payment{}
	if err := r.source.Where(query).First(&payment).Error; err != nil {
		return datamodels.Payment{}, false
	}
	return payment, true
}

func (r *paymentRepository) SelectMany(query Query, limit int) (results []datamodels.Payment) {
	payments := new([]datamodels.Payment)
	r.source.Where(query).Find(&payments).Limit(limit)
	return *payments
}

func (r *paymentRepository) InsertOrUpdate(payment datamodels.Payment) (_ datamodels.Payment, err error) {
	var oldPayment datamodels.Payment
	if err := r.source.First(&oldPayment).Error; err != nil {
		r.source.Create(&payment)
	} else {
		r.source.Model(&oldPayment).Update(&payment)
	}

	return payment, err
}

func (r *paymentRepository) Delete(query Query) bool {
	if err := r.source.Delete(datamodels.Payment{}, query).Error; err != nil {
		return false
	}
	return true
}
