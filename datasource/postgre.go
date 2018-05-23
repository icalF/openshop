package datasource

import (
	"fmt"
	"os"

	"github.com/koneko096/openshop/models/datamodels"
	"github.com/jinzhu/gorm"
)

func NewPostgreConnection() (*gorm.DB, error) {
	sslmode := "disable"
	if envSslmode := os.Getenv("SSL_MODE"); envSslmode != "" {
	    sslmode = envSslmode
	}
	
	host := "localhost"
	if envHost := os.Getenv("DB_HOST"); envHost != "" {
	    host = envHost
	}
	
	port := "5432"
	if envPort := os.Getenv("DB_PORT"); envPort != "" {
	    port = envPort
	}
	
	user := "postgres"
	if envUser := os.Getenv("DB_USER"); envUser != "" {
	    user = envUser
	}

	db := "openshop"
	if envDb := os.Getenv("DB_NAME"); envDb != "" {
	    db = envDb
	}

	configText := fmt.Sprintf("sslmode=%s host=%s port=%s user=%s dbname=%s password=%s", sslmode, host, port, user, db, os.Getenv("DB_PASS"))
	dbConn, err := gorm.Open("postgres", configText)
	if err != nil {
		return &gorm.DB{}, err
	}

	dbConn.AutoMigrate(&datamodels.User{}, &datamodels.Coupon{},
		&datamodels.Product{}, &datamodels.Shipment{},
		&datamodels.Payment{}, &datamodels.OrderDetail{}, &datamodels.Order{})
	return dbConn, err
}
