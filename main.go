// package main

// import (
// 	"log"
// 	"os"

// 	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/controller/routes"
// 	"github.com/gin-gonic/gin"
// 	"github.com/joho/godotenv"
// )

// func main() {
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}

// 	port := os.Getenv("PORT")
// 	if port == "" {
// 		port = "8080"
// 	}

// 	router := gin.Default()
// 	routes.InitRoutes(&router.RouterGroup)

// 	log.Println("Servidor rodando na porta " + port)
// 	router.Run(":" + port)
// }

package main

import (
	"context"
	"log"

	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/database/mongobd"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/configuration/logger"
	"github.com/Lipe-Azevedo/meu-primeio-crud-go/src/controller/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	logger.Info("About to start user application")
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database, err := mongobd.NewMongoDBConnection(context.Background())
	if err != nil {
		log.Fatalf(
			"Error trying to connect to database, error = %s \n",
			err.Error())
		return
	}

	userController, workInfoController, swapController := initDependencies(database)

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup, userController, workInfoController, swapController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
