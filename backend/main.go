package main

import "github.com/gin-gonic/gin"

func main() {
    r := gin.Default() // bikin router Gin dengan logger & recovery

    // Endpoint GET "/" → Hello World
    r.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "Hello Mini Hospital API!",
        })
    })

    r.Run() // jalankan server di localhost:8080
}