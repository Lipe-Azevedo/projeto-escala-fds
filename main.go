package main

import (
	"context"
	"log" // log padrão para erros fatais de inicialização
	"os"  // Para os.Getenv

	"github.com/Lipe-Azevedo/escala-fds/src/configuration/database/mongobd"
	"github.com/Lipe-Azevedo/escala-fds/src/configuration/logger" // Seu logger customizado
	"github.com/Lipe-Azevedo/escala-fds/src/controller/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	logger.Info("About to start application") // Usando seu logger customizado
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file") // log.Fatal para erro crítico na inicialização
	}

	// Validar JWT_SECRET_KEY no início
	jwtSecret := os.Getenv("JWT_SECRET_KEY")
	if jwtSecret == "" {
		log.Fatal("FATAL: JWT_SECRET_KEY environment variable not set or empty.")
	}
	logger.Info("JWT_SECRET_KEY loaded.")

	database, err := mongobd.NewMongoDBConnection(context.Background())
	if err != nil {
		// Usar logger.Error do seu pacote e depois log.Fatalf para consistência
		// Ou apenas log.Fatalf que já inclui logar e sair.
		logger.Error("Error trying to connect to database", err)    // log.Fatalf já faz o print e exit
		log.Fatalf("Database connection error: %s \n", err.Error()) // Para garantir a saída
		return                                                      // Redundante devido ao Fatalf, mas bom para clareza
	}
	logger.Info("Database connected successfully")

	// Inicializar dependências, incluindo o novo CommentController
	userController, workInfoController, swapController, commentController := initDependencies(database)
	logger.Info("Dependencies initialized successfully")

	router := gin.Default()
	// Configurar o RouterGroup base, se necessário (ex: r := router.Group("/v1"))
	// Por enquanto, usando o RouterGroup padrão do Gin que InitRoutes espera.
	// InitRoutes espera *gin.RouterGroup. router.RouterGroup é o grupo raiz.
	routes.InitRoutes(&router.RouterGroup, userController, workInfoController, swapController, commentController)
	logger.Info("Routes initialized successfully")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Porta padrão
	}

	logger.Info("Server starting on port " + port)
	if err := router.Run(":" + port); err != nil {
		log.Fatal(err) // log.Fatal para erro crítico ao iniciar o servidor
	}
}
