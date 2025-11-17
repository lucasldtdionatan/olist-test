package database

import (
	"log"
	"olist-project/internal/config"
	"olist-project/internal/entities"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitPostgres(cfg config.Config) *gorm.DB {
	db, err := gorm.Open(postgres.Open(cfg.DBUrl), &gorm.Config{})
	if err != nil {
		log.Fatalf("Erro ao conectar no banco: %v", err)
	}

	db.AutoMigrate(&entities.Package{}, &entities.TransportCompany{}, &entities.TransportCompanyRegion{}, &entities.Shipment{})

	return db
}
