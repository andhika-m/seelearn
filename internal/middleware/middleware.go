package middleware

import (
	"context"
	"seelearn/internal/model/constant"
	"seelearn/internal/usecase/video"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type authMiddleware struct {
	videoUsecase video.Usecase
}

func LoadMiddlewares(e *echo.Echo) {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{

		AllowOrigins: []string{"*"},
	}))
}

func GetAuthMiddleware(videoUsecase video.Usecase) *authMiddleware {
	return &authMiddleware{
		videoUsecase: videoUsecase,
	}
}

func (am *authMiddleware) CheckAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sessionData, err := GetSessionData(c.Request())
		if err != nil {
			return &echo.HTTPError{
				Code:     401,
				Message:  err.Error(),
				Internal: err,
			}
		}

		userID, err := am.videoUsecase.CheckSession(sessionData)
		if err != nil {
			return &echo.HTTPError{
				Code:     401,
				Message:  err.Error(),
				Internal: err,
			}
		}

		authContext := context.WithValue(c.Request().Context(), constant.AuthContextKey, userID)
		c.SetRequest(c.Request().WithContext(authContext))

		if err := next(c); err != nil {
			return err
		}

		return nil
	}
}
