package config

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/bencoronard/demo-go-crud-api/internal/resource"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/fx"
)

type Container interface {
	Start(lc fx.Lifecycle)
}

type containerImpl struct {
	p *Properties
	r *echo.Echo
	h *resource.ResourceHandler
	l *slog.Logger
}

func NewRouter(h *resource.ResourceHandler, l *slog.Logger, p *Properties) Container {
	r := echo.New()
	registerMiddlewares(r)
	registerRoutes(r, h)
	return &containerImpl{
		p: p,
		r: r,
		h: h,
		l: l,
	}
}

func (c *containerImpl) Start(lc fx.Lifecycle) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			c.l.Info(fmt.Sprintf("Process ID: %d on %s", os.Getpid(), c.p.Env.App.Environment))
			go func() {
				if err := c.r.Start(fmt.Sprintf(":%d", c.p.Env.App.ListenPort)); err != nil && err != http.ErrServerClosed {
					c.l.Error("HTTP server failed to start")
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			c.l.Info("Shutting down HTTP server...")
			return c.r.Shutdown(ctx)
		},
	})
}

func registerMiddlewares(r *echo.Echo) {
	r.Use(middleware.RequestLogger())
	r.Use(middleware.Recover())
}

func registerRoutes(r *echo.Echo, h *resource.ResourceHandler) {
	r.GET("/", h.ListResources)
	r.GET("/", h.RetrieveResource)
	r.POST("/", h.CreateResource)
	r.PUT("/", h.UpdateResource)
	r.DELETE("/", h.DeleteResource)
}
