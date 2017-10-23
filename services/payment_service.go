package services

import (
	"errors"

	"github.com/icalF/openshop/models/datamodels"
	"github.com/icalF/openshop/dao"
)

type PaymentService interface {
	GetAll() []datamodels.Payment
	GetByID(id int64) (datamodels.Payment, bool)
	GetByOrderID(orderId int64) (datamodels.Payment, bool)
	InsertOrUpdate(payment datamodels.Payment) (datamodels.Payment, error)
	UpdatePaymentProof(orderId int64, filename string) (bool, error)
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

func (s *paymentService) GetAll() []datamodels.Payment {
	return s.dao.SelectMany(map[string]string{}, 0)
}

func (s *paymentService) GetByID(id int64) (datamodels.Payment, bool) {
	return s.dao.Select(map[string]string{
		"id": string(id),
	})
}

func (s *paymentService) GetByOrderID(orderId int64) (datamodels.Payment, bool) {
	return s.dao.Select(map[string]string{
		"order_id": string(orderId),
	})
}

func (s *paymentService) InsertOrUpdate(payment datamodels.Payment) (datamodels.Payment, error) {
	return s.dao.InsertOrUpdate(payment)
}

func (s *paymentService) UpdatePaymentProof(orderId int64, filename string) (bool, error) {
	payment, found := s.dao.Select(map[string]string{
		"order_id": string(orderId),
	})
	if !found {
		return false, errors.New("order ID not found")
	}

	payment.Proof = filename
	_, err := s.InsertOrUpdate(payment)
	return true, err
}

func (s *paymentService) DeleteByID(id int64) bool {
	return s.dao.Delete(map[string]string{
		"id": string(id),
	})
}
