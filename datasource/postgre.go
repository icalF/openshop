package datasource

import (
	"github.com/jinzhu/gorm"
	"github.com/icalF/openshop/models"
)

func NewMysqlConnection() (*gorm.DB, error) {
	dbConn, err := gorm.Open("postgres", "host=localhost user=postgres dbname=openshop sslmode=disable password=pgsql")
	if err != nil {
		return &gorm.DB{}, err
	}

	dbConn.AutoMigrate(&models.User{}, &models.Coupon{})
	return dbConn, err
}
