package services

import (
	"github.com/icalF/openshop/dao"
	"github.com/icalF/openshop/models/datamodels"
)

type OrderDetailService interface {
	GetAll() []datamodels.OrderDetail
	GetByID(id int64) (datamodels.OrderDetail, bool)
	GetByOrderID(id int64) []datamodels.OrderDetail
	InsertOrUpdate(orderDetail datamodels.OrderDetail) (datamodels.OrderDetail, error)
	DeleteByID(id int64) bool
	ValidatePurchase(orderDetail datamodels.OrderDetail) bool
}

func NewOrderDetailService(dao dao.OrderDetailDAO, productService ProductService) OrderDetailService {
	return &orderDetailService{
		dao:            dao,
		productService: productService,
	}
}

type orderDetailService struct {
	dao            dao.OrderDetailDAO
	productService ProductService
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
	product, found := s.productService.GetByID(orderDetail.ProductID)
	if !found {
		return false
	}

	productQty := product.Qty
	return productQty >= orderDetail.Qty
}