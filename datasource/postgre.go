package datasource

import (
	"github.com/jinzhu/gorm"
	"github.com/icalF/openshop/models/datamodels"
)

func NewMysqlConnection() (*gorm.DB, error) {
	dbConn, err := gorm.Open("postgres", "host=localhost user=postgres dbname=openshop sslmode=disable password=pgsql")
	if err != nil {
		return &gorm.DB{}, err
	}

	dbConn.AutoMigrate(&datamodels.User{}, &datamodels.Coupon{},
		&datamodels.Product{}, &datamodels.Shipment{},
		&datamodels.Payment{}, &datamodels.OrderDetail{}, &datamodels.Order{})
	return dbConn, err
}
