package router

import (
	"olist-project/internal/handlers"
	"olist-project/internal/repositories"
	"olist-project/internal/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	packageRepo := repositories.NewPackageRepository(db)
	packageService := services.NewPackageService(packageRepo)
	packageHandler := handlers.NewPackageHandler(packageService)

	transportCompanyRepo := repositories.NewTransportCompanyRepository(db)
	transportCompanyService := services.NewTransportCompanyService(transportCompanyRepo)
	transportCompanyHandler := handlers.NewTransportCompanyHandler(transportCompanyService)

	transportCompanyRegionRepo := repositories.NewTransportCompanyRegionRepository(db)
	transportCompanyRegionService := services.NewTransportCompanyRegionService(transportCompanyRegionRepo)
	transportCompanyRegionHandler := handlers.NewTransportCompanyRegionHandler(transportCompanyRegionService)

	v1 := r.Group("/api/v1")
	{
		v1.POST("/packages", packageHandler.Create)
		v1.GET("/packages", packageHandler.List)
		v1.GET("/packages/:id", packageHandler.Get)
		v1.PUT("/packages/:id", packageHandler.Update)
		v1.DELETE("/packages/:id", packageHandler.Delete)

		tc := v1.Group("/transport-companies")

		tc.POST("", transportCompanyHandler.Create)
		tc.GET("", transportCompanyHandler.List)
		tc.GET("/:id", transportCompanyHandler.Get)
		tc.PUT("/:id", transportCompanyHandler.Update)
		tc.DELETE("/:id", transportCompanyHandler.Delete)

		regions := tc.Group("/:id/regions")
		{
			regions.POST("", transportCompanyRegionHandler.CreateForCompany)
			regions.GET("", transportCompanyRegionHandler.ListForCompany)
		}

		shippingService := services.NewShippingService(transportCompanyRegionRepo)
		shippingHandler := handlers.NewShippingHandler(shippingService)

		v1.POST("/shipping/quote", shippingHandler.CalculateQuote)

		shipmentRepo := repositories.NewShipmentRepository(db)
		shipmentService := services.NewShipmentService(shipmentRepo)
		shipmentHandler := handlers.NewShipmentHandler(shipmentService)

		shipments := v1.Group("/shipments")
		{
			shipments.POST("", shipmentHandler.Create)
			shipments.GET("", shipmentHandler.List)
			shipments.GET("/:id", shipmentHandler.Get)
			shipments.PUT("/:id", shipmentHandler.Update)
			shipments.DELETE("/:id", shipmentHandler.Delete)
		}
	}

	return r
}
