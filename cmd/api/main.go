package main

import (
	"log"
	"olist-project/internal/config"
	"olist-project/internal/database"
	"olist-project/internal/router"
)

func main() {
	cfg := config.LoadConfig()

	db := database.InitPostgres(cfg)

	r := router.SetupRouter(db)

	log.Println("Server running... :", cfg.Port)
	r.Run(":" + cfg.Port)
}
