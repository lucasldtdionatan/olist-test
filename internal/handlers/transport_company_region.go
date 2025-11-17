package handlers

import (
	"net/http"
	"olist-project/internal/dto"
	"olist-project/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type TransportCompanyRegionHandler struct {
	service *services.TransportCompanyRegionService
}

func NewTransportCompanyRegionHandler(s *services.TransportCompanyRegionService) *TransportCompanyRegionHandler {
	return &TransportCompanyRegionHandler{s}
}

func (h *TransportCompanyRegionHandler) CreateForCompany(c *gin.Context) {
	idParam := c.Param("id")
	companyID, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var body dto.CreateTransportCompanyRegionRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	region, err := h.service.CreateForCompany(companyID, body)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, region)
}

func (h *TransportCompanyRegionHandler) ListForCompany(c *gin.Context) {
	idParam := c.Param("id")
	companyID, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	regions, err := h.service.ListForCompany(companyID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, regions)
}
