package main

import (
    "backend/database"
    "backend/routes"

    "github.com/gin-gonic/gin"
)

func main() {
    database.Connect()
    r := gin.Default()

    // Serve frontend
    r.Static("/static", "../frontend")
    r.GET("/", func(c *gin.Context) {
        c.File("../frontend/index.html")
    })

    routes.Setup(r)

    r.Run() // :8080
}