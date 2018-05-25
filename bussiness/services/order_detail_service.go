package services

import (
	"github.com/koneko096/openshop/dao"
	"github.com/koneko096/openshop/models/datamodels"
	"github.com/koneko096/openshop/bussiness/usecases"
)

func NewOrderDetailManager(dao dao.OrderDetailDAO, productManager usecases.ProductManager) usecases.OrderDetailManager {
	return &orderDetailService{
		dao:            dao,
		productManager: productManager,
	}
}

func NewPurchaseValidator(dao dao.OrderDetailDAO, productManager usecases.ProductManager) usecases.PurchaseValidator {
	return &orderDetailService{
		dao:            dao,
		productManager: productManager,
	}
}

type orderDetailService struct {
	dao            dao.OrderDetailDAO
	productManager usecases.ProductManager
}

func (s *orderDetailService) GetAll() []datamodels.OrderDetail {
	return s.dao.SelectMany(map[string]string{}, 0)
}

func (s *orderDetailService) GetByID(id int64) (datamodels.OrderDetail, bool) {
	return s.dao.Select(map[string]string{
		"id": string(id),
	})
}

func (s *orderDetailService) GetByOrderID(id int64) []datamodels.OrderDetail {
	return s.dao.SelectMany(map[string]string{
		"order_id": string(id),
	}, 0)
}

func (s *orderDetailService) InsertOrUpdate(orderDetail datamodels.OrderDetail) (datamodels.OrderDetail, error) {
	return s.dao.InsertOrUpdate(orderDetail)
}

func (s *orderDetailService) DeleteByID(id int64) bool {
	return s.dao.Delete(map[string]string{
		"id": string(id),
	})
}

func (s *orderDetailService) ValidatePurchase(orderDetail datamodels.OrderDetail) bool {
	product, found := s.productManager.GetByID(orderDetail.ProductID)
	if !found {
		return false
	}

	productQty := product.Qty
	return productQty >= orderDetail.Qty
}
