package datasource

import (
	"fmt"
	"os"
	
	"github.com/jinzhu/gorm"
	"github.com/icalF/openshop/models/datamodels"
)

func NewPostgreConnection() (*gorm.DB, error) {
	configText := fmt.Sprintf("host=%s user=%s dbname=%s port=5432 password=%s", os.Getenv("HOST"), os.Getenv("USER"),  os.Getenv("DB"), os.Getenv("PASS"))
	dbConn, err := gorm.Open("postgres", configText)
	if err != nil {
		return &gorm.DB{}, err
	}

	dbConn.AutoMigrate(&datamodels.User{}, &datamodels.Coupon{},
		&datamodels.Product{}, &datamodels.Shipment{},
		&datamodels.Payment{}, &datamodels.OrderDetail{}, &datamodels.Order{})
	return dbConn, err
}
