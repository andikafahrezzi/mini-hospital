package routes

import (
    "backend/controllers"
    "github.com/gin-gonic/gin"
)

func Setup(r *gin.Engine) {
    r.POST("/antrian", controllers.AddAntrian)
    r.GET("/antrian", controllers.GetAntrian)
    r.DELETE("/antrian/:id", controllers.DeleteAntrian)
}