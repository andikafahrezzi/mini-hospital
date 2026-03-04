package routes

import (
	"backend/internal/handler"
	"backend/internal/middleware"
	"github.com/gin-gonic/gin"
)

func Setup(r *gin.Engine,
	antrianHandler *handler.AntrianHandler,
	poliHandler *handler.PoliHandler,
	dokterHandler *handler.DokterHandler,

) {

	r.GET("/antrian", antrianHandler.GetAll)
	r.POST("/antrian", antrianHandler.Create)
	r.DELETE("/antrian/:id", antrianHandler.Delete)

	r.GET("/poli", poliHandler.GetAll)
	r.GET("/dokter", dokterHandler.GetByPoli)
	r.PUT("/antrian/:id/status", antrianHandler.UpdateStatus)
	r.POST("/login", handler.Login)

auth := r.Group("/")
auth.Use(middleware.AuthMiddleware())
{
	auth.GET("/protected",
	middleware.RequireRole("dokter"),
	func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "dokter area"})
	})
}
}