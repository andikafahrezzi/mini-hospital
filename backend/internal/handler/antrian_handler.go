package handler

import (
	"net/http"
	"strconv"

	"backend/internal/service"
	"github.com/gin-gonic/gin"
)

type AntrianHandler struct {
	service *service.AntrianService
}

type PoliHandler struct {
	service *service.PoliService
}

func NewAntrianHandler(s *service.AntrianService) *AntrianHandler {
	return &AntrianHandler{service: s}
}

func (h *AntrianHandler) GetAll(c *gin.Context) {
	data, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, data)
}

func (h *AntrianHandler) Create(c *gin.Context) {
	var req struct {
		NamaPasien string `json:"nama_pasien"`
		PoliID     int    `json:"poli_id"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	data, err := h.service.Create(req.NamaPasien, req.PoliID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, data)
}

func (h *AntrianHandler) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))

	err := h.service.Delete(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}


func NewPoliHandler(s *service.PoliService) *PoliHandler {
	return &PoliHandler{service: s}
}

func (h *PoliHandler) GetAll(c *gin.Context) {
	data, err := h.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, data)
}