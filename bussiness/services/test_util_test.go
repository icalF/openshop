package services

import (
	"github.com/koneko096/openshop/services"
	"github.com/koneko096/openshop/models/datamodels"
)

func NewTestOrderDetailService() services.OrderDetailService {
	return &testOrderDetailService{}
}
func NewTestProductService() services.ProductService {
	return &testProductService{}
}
func NewTestCouponService() services.CouponService {
	return &testCouponService{}
}

type testOrderDetailService struct{}
type testProductService struct{}
type testCouponService struct{}

// MOCKED METHODS

func (s *testProductService) GetAll() []datamodels.Product {
	return []datamodels.Product{}
}

func (s *testProductService) GetAllPurchasable() []datamodels.Product {
	return []datamodels.Product{}
}

func (s *testProductService) GetByID(id int64) (datamodels.Product, bool) {
	// Mocked
	productMap := s.CreateProductMap(s.GetByIDs([]int64{}))
	return productMap[id], true
}

func (s *testProductService) GetByIDs(ids []int64) []datamodels.Product {
	// Mocked
	return []datamodels.Product{
		{
			ID:    1,
			Price: 35,
			Qty:   3,
		},
		{
			ID:    2,
			Price: 15,
			Qty:   1,
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

func (s *testCouponService) GetAll() []datamodels.Coupon {
	return []datamodels.Coupon{}
}

func (s *testCouponService) GetByID(id int64) (datamodels.Coupon, bool) {
	return datamodels.Coupon{}, true
}

func (s *testCouponService) GetByPromoCode(code string) (datamodels.Coupon, bool) {
	// Mocked
	testSet := map[string]datamodels.Coupon{
		"ABC": {
			Percent: 50,
		},
		"DEF": {
			Nominal: 20000,
		},
	}
	return testSet[code], true
}

func (s *testCouponService) InsertOrUpdate(coupon datamodels.Coupon) (datamodels.Coupon, error) {
	return datamodels.Coupon{}, nil
}

func (s *testCouponService) ValidateCoupon(coupon datamodels.Coupon) bool {
	return true
}

func (s *testCouponService) ValidateAndTakeCoupon(code string) bool {
	return true
}

func (s *testCouponService) DeleteByID(id int64) bool {
	return true
}
