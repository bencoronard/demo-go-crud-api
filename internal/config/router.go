package config

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/bencoronard/demo-go-crud-api/internal/resource"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Router interface {
	Start(lc fx.Lifecycle, props *Properties)
}

type routerImpl struct {
	r   *echo.Echo
	h   *resource.ResourceHandler
	log *zap.Logger
}

func NewRouter(h *resource.ResourceHandler, l *zap.Logger) Router {
	r := echo.New()
	registerMiddlewares(r)
	registerRoutes(r, h)
	return &routerImpl{
		r:   r,
		h:   h,
		log: l,
	}
}

func (r *routerImpl) Start(lc fx.Lifecycle, props *Properties) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			r.log.Info(fmt.Sprintf("Process ID: %d on %s", os.Getpid(), props.Env.App.Environment))
			go func() {
				if err := r.r.Start(fmt.Sprintf(":%d", props.Env.App.ListenPort)); err != nil && err != http.ErrServerClosed {
					r.log.Error("HTTP server failed to start", zap.Error(err))
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			r.log.Info("Shutting down HTTP server...")
			return r.r.Shutdown(ctx)
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
