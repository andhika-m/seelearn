package main

import (
	"crypto/rand"
	"crypto/rsa"
	"log"
	"seelearn/internal/api/v1/rest"
	"seelearn/internal/database"
	"seelearn/internal/middleware"
	uRepo "seelearn/internal/repository/user"
	vRepo "seelearn/internal/repository/video"
	vUsecase "seelearn/internal/usecase/video"
	"time"

	"github.com/joho/godotenv"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	// database connect
	db, err := database.ConnectDB()
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	secret := "AES256Key-32Characters1234567890"
	signKey, err := rsa.GenerateKey(rand.Reader, 4096)
	if err != nil {
		panic(err)
	}

	storagePath := "./public/upload/videos"

	videoRepo := vRepo.GetRepository(db, storagePath)
	userRepo, err := uRepo.GetRepository(db, secret, 1, 64*1024, 4, 32, signKey, 60*time.Hour)
	if err != nil {
		panic(err)
	}

	videoUsecase := vUsecase.NewVideoUsecase(videoRepo, userRepo)

	videoHandler := rest.NewVideoHandler(videoUsecase)

	middleware.LoadMiddlewares(e)
	rest.InitVideoRoutes(e, videoHandler)

	e.Start(":8080")
}
