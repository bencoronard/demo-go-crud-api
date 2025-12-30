package config

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/bencoronard/demo-go-crud-api/internal/resource"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/fx"
)

func NewRouter(lc fx.Lifecycle, p *Properties) *echo.Echo {
	e := echo.New()

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			slog.Info(fmt.Sprintf("Process ID: %d, Env: %s", os.Getpid(), p.Env.App.Environment))
			return nil
		},
		OnStop: func(ctx context.Context) error {
			slog.Info("Router shutting down...")
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
