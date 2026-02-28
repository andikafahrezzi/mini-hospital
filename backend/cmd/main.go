package main

import (
	"backend/internal/handler"
	"backend/internal/repository"
	"backend/internal/service"
	"backend/pkg/database"
	"github.com/joho/godotenv"
    "backend/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env")
	}

	dsn := os.Getenv("DB_DSN")
	database.Connect(dsn)

	r := gin.Default()

	// =============================
	// Dependency Injection
	// =============================
	antrianRepo := repository.NewAntrianRepository()
	antrianService := service.NewAntrianService(antrianRepo)
	antrianHandler := handler.NewAntrianHandler(antrianService)
    poliRepo := repository.NewPoliRepository()
    poliService := service.NewPoliService(poliRepo)
    poliHandler := handler.NewPoliHandler(poliService)
	dokterRepo := repository.NewDokterRepository()
	dokterService := service.NewDokterService(dokterRepo)
	dokterHandler := handler.NewDokterHandler(dokterService)
	// =============================
	// API Routes
	// =============================
	// API routes
	routes.Setup(r, antrianHandler, poliHandler, dokterHandler)

	// =============================
	// Serve Frontend
	// =============================
	r.Static("/static", "../frontend")
	r.GET("/", func(c *gin.Context) {
		c.File("../frontend/index.html")
	})

	r.Run(":8080")
}