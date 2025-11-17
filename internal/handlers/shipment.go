package handlers

import (
	"errors"
	"net/http"
	"olist-project/internal/dto"
	"olist-project/internal/services"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
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

func (h *ShipmentHandler) List(c *gin.Context) {
	var filters dto.ShipmentFilter

	if err := c.ShouldBindQuery(&filters); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "query params inválidos"})
		return
	}

	shipments, err := h.service.List(filters)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := dto.ToShipmentResponseList(shipments)

	c.JSON(http.StatusOK, response)
}

func (h *ShipmentHandler) Get(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	shipment, err := h.service.GetByID(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "shipment not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	response := dto.ToShipmentResponse(*shipment)
	c.JSON(http.StatusOK, response)
}

func (h *ShipmentHandler) Update(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var req dto.UpdateShipmentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	shipment, err := h.service.Update(c.Request.Context(), id, req)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "shipment not found"})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, dto.ToShipmentResponse(*shipment))
}

func (h *ShipmentHandler) Delete(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	err = h.service.Delete(c.Request.Context(), id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "shipment not found"})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
