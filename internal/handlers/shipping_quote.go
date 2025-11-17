package handlers

import (
	"net/http"
	"olist-project/internal/dto"
	"olist-project/internal/services"

	"github.com/gin-gonic/gin"
)

type ShippingHandler struct {
	service services.ShippingService
}

func NewShippingHandler(service services.ShippingService) *ShippingHandler {
	return &ShippingHandler{service}
}

func (h *ShippingHandler) CalculateQuote(c *gin.Context) {
	var req dto.ShippingQuoteRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	quote, err := h.service.GetQuotesByRegion(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, quote)
}
