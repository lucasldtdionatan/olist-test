package handlers

import (
	"net/http"
	"olist-project/internal/dto"
	"olist-project/internal/services"

	"github.com/gin-gonic/gin"
)

type ShipmentHandler struct {
	service *services.ShipmentService
}

func NewShipmentHandler(s *services.ShipmentService) *ShipmentHandler {
	return &ShipmentHandler{s}
}

func (h *ShipmentHandler) Create(c *gin.Context) {
	var req dto.CreateShipmentRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.service.CreateShipment(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, resp)
}
