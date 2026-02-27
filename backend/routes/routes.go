package routes

import (
	"backend/internal/handler"

	"github.com/gin-gonic/gin"
)

func Setup(r *gin.Engine,
	antrianHandler *handler.AntrianHandler,
	poliHandler *handler.PoliHandler,
) {

	r.GET("/antrian", antrianHandler.GetAll)
	r.POST("/antrian", antrianHandler.Create)
	r.DELETE("/antrian/:id", antrianHandler.Delete)

	r.GET("/poli", poliHandler.GetAll)
}