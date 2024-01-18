package rest

import (
	"seelearn/internal/middleware"

	"github.com/labstack/echo/v4"
)

func InitVideoRoutes(e *echo.Echo, handler *handler) {
	authMiddleware := middleware.GetAuthMiddleware(handler.videoUsecase)

	videoGroup := e.Group("/api/v1/videos")
	videoGroup.POST("", handler.CreateVideo,
		authMiddleware.CheckAuth,
	)
	videoGroup.GET("", handler.GetVideos)
	videoGroup.GET("/:id", handler.GetVideoByID)
	videoGroup.GET("/download/:id", handler.DownloadVideo)
	videoGroup.PUT("/:id", handler.EditVideo,
		authMiddleware.CheckAuth,
	)
	videoGroup.DELETE("/:id", handler.DeleteVideo,
		authMiddleware.CheckAuth,
	)

	userGroup := e.Group("/api/v1/user")
	userGroup.POST("/register", handler.RegisterUser)
	userGroup.POST("/login", handler.Login)
}
