package main

import (
	"context"
	"log"
	"os"

	"github.com/Lipe-Azevedo/escala-fds/src/configuration/database/mongobd"
	"github.com/Lipe-Azevedo/escala-fds/src/configuration/logger"
	"github.com/Lipe-Azevedo/escala-fds/src/controller/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	logger.Info("About to start application")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	jwtSecret := os.Getenv("JWT_SECRET_KEY")
	if jwtSecret == "" {
		log.Fatal("FATAL: JWT_SECRET_KEY environment variable not set or empty.")
	}
	logger.Info("JWT_SECRET_KEY loaded.")

	database, err := mongobd.NewMongoDBConnection(context.Background())
	if err != nil {
		logger.Error("Error trying to connect to database", err)
		log.Fatalf("Database connection error: %s \n", err.Error())
		return
	}
	logger.Info("Database connected successfully")

	userController, workInfoController, swapController, commentController := initDependencies(database)
	logger.Info("Dependencies initialized successfully")

	router := gin.Default()

	config := cors.Config{
		// Exemplo: AllowOrigins: []string{"https://meufrontend.com"},
		AllowOrigins: []string{"http://localhost:3000", "http://127.0.0.1:3000"}, // Para desenvolvimento local
		AllowMethods: []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders: []string{
			"Origin",
			"Content-Type",
			"Accept",
			"Authorization",
			"X-Requested-With",
		},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		// MaxAge: 12 * time.Hour, // Opcional: quanto tempo o resultado de OPTIONS pode ser cacheado
	}
	// Se AllowCredentials = true, AllowOrigins NÃO PODE ser "*".

	router.Use(cors.New(config)) // Aplica o middleware CORS

	routes.InitRoutes(&router.RouterGroup, userController, workInfoController, swapController, commentController)
	logger.Info("Routes initialized successfully")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	logger.Info("Server starting on port " + port)
	if err := router.Run(":" + port); err != nil {
		log.Fatal(err)
	}
}
