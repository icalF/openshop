package services

import (
	"github.com/icalF/openshop/models"
	"github.com/icalF/openshop/dao"
)

type PaymentService interface {
	GetAll() []models.Payment
	GetByID(id int64) (models.Payment, bool)
	InsertOrUpdate(payment models.Payment) (models.Payment, error)
	DeleteByID(id int64) bool
}

func NewPaymentService(dao dao.PaymentDAO) PaymentService {
	return &paymentService{
		dao: dao,
	}
}

type paymentService struct {
	dao dao.PaymentDAO
}

func (s *paymentService) GetAll() []models.Payment {
	return s.dao.SelectMany(map[string]string{}, 0)
}

func (s *paymentService) GetByID(id int64) (models.Payment, bool) {
	return s.dao.Select(map[string]string{
		"id": string(id),
	})
}

func (s *paymentService) InsertOrUpdate(payment models.Payment) (models.Payment, error) {
	return s.dao.InsertOrUpdate(payment)
}

func (s *paymentService) DeleteByID(id int64) bool {
	return s.dao.Delete(map[string]string{
		"id": string(id),
	})
}