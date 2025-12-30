package config

import (
	"context"

	"github.com/bencoronard/demo-go-crud-api/internal/resource"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/fx"
)

func NewRouter(lc fx.Lifecycle) *echo.Echo {
	e := echo.New()

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return e.Shutdown(ctx)
		},
	})

	return e
}

func RegisterMiddlewares(e *echo.Echo) {
	e.Use(middleware.RequestLogger())
	e.Use(middleware.Recover())
}

func RegisterRoutes(e *echo.Echo, h *resource.ResourceHandler) {
	e.GET("/", h.ListResources)
	e.GET("/", h.RetrieveResource)
	e.POST("/", h.CreateResource)
	e.PUT("/", h.UpdateResource)
	e.DELETE("/", h.DeleteResource)
}
