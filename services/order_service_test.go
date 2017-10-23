package services_test

import (
	"testing"

	"github.com/icalF/openshop/services"
	"github.com/icalF/openshop/models/datamodels"
)

func NewTestOrderDetailService() services.OrderDetailService {
	return &testOrderDetailService{}
}
func NewTestProductService() services.ProductService {
	return &testProductService{}
}

type testOrderDetailService struct{}
type testProductService struct{}

func TestOrderService_GetTotalAmount(t *testing.T) {
	orderDetailService := NewTestOrderDetailService()
	productService := NewTestProductService()
	orderService := services.NewOrderService(nil, nil, orderDetailService, productService, nil)

	result := orderService.GetTotalAmount(1)
	if result != 105 {
		t.Errorf("Total amount: expected %d, actual %d", 105, result)
	}
}

// MOCKED METHODS

func (s *testProductService) GetAll() []datamodels.Product {
	return []datamodels.Product{}
}

func (s *testProductService) GetAllPurchasable() []datamodels.Product {
	return []datamodels.Product{}
}

func (s *testProductService) GetByID(id int64) (datamodels.Product, bool) {
	return datamodels.Product{}, true
}

func (s *testProductService) GetByIDs(ids []int64) []datamodels.Product {
	// Mocked
	return []datamodels.Product{
		{
			ID:    1,
			Price: 35,
		},
		{
			ID:    2,
			Price: 15,
		},
		{
			ID:    3,
			Price: 5,
		},
	}
}

func (s *testProductService) InsertOrUpdate(product datamodels.Product) (datamodels.Product, error) {
	return datamodels.Product{}, nil
}

func (s *testProductService) DeleteByID(id int64) bool {
	return true
}

func (s *testProductService) CreateProductMap(products []datamodels.Product) map[int64]datamodels.Product {
	// Real method
	productMap := map[int64]datamodels.Product{}
	for _, product := range products {
		productMap[product.ID] = product
	}
	return productMap
}

func (s *testOrderDetailService) GetAll() []datamodels.OrderDetail {
	return []datamodels.OrderDetail{}
}

func (s *testOrderDetailService) GetByID(id int64) (datamodels.OrderDetail, bool) {
	return datamodels.OrderDetail{}, true
}

func (s *testOrderDetailService) GetByOrderID(id int64) []datamodels.OrderDetail {
	// Mocked
	return []datamodels.OrderDetail{
		{
			ProductID: 1,
			Qty:       2,
		},
		{
			ProductID: 2,
			Qty:       2,
		},
		{
			ProductID: 3,
			Qty:       1,
		},
	}
}

func (s *testOrderDetailService) InsertOrUpdate(orderDetail datamodels.OrderDetail) (datamodels.OrderDetail, error) {
	return datamodels.OrderDetail{}, nil
}

func (s *testOrderDetailService) DeleteByID(id int64) bool {
	return true
}

func (s *testOrderDetailService) ValidatePurchase(orderDetail datamodels.OrderDetail) bool {
	return true
}
