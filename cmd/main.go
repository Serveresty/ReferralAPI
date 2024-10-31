package main

import (
	"ReferralAPI/configs"
	"ReferralAPI/database"
	"ReferralAPI/internal/controllers"
	"ReferralAPI/internal/repository"
	"ReferralAPI/internal/routes"
	"ReferralAPI/internal/service"
	"context"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	if err := godotenv.Load("./configs/.env"); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	dbConf := configs.LoadDBConfig()
	serverConf := configs.LoadServerConfig()

	dbConn, err := database.DBInit(dbConf)
	if err != nil {
		log.Fatalf("error connecting to db: %s", err.Error())
	}
	defer dbConn.Close(context.Background())

	database.RunMigrations(dbConn)

	refRepo := repository.NewRefRepository(dbConn)
	refServ := service.NewRefService(refRepo)
	refCont := controllers.NewRefController(refServ)
	refRout := routes.NewRefRouteController(refCont)

	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	refRout.RefRoutes(router)

	if err := router.Run(serverConf.Host + ":" + serverConf.Port); err != nil {
		log.Fatalf("error starting server: %s", err.Error())
	}
}
