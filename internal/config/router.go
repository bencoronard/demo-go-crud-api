package config

import (
	"net/http"

	xhttp "github.com/bencoronard/demo-go-common-libs/http"
	"github.com/bencoronard/demo-go-common-libs/utility"
	"github.com/bencoronard/demo-go-crud-api/internal/resource"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
)

type router struct {
	p *Properties
	e *echo.Echo
	h *resource.ResourceHandler
}

func NewRouter(p *Properties, h *resource.ResourceHandler, v utility.Validator) xhttp.Router {
	e := echo.New()
	e.HTTPErrorHandler = xhttp.GlobalErrorHandler(nil)
	e.Validator = v
	return &router{
		p: p,
		e: e,
		h: h,
	}
}

func (r *router) Port() int {
	return r.p.Env.App.ListenPort
}

func (r *router) Handler() http.Handler {
	return r.e
}

func (r *router) RegisterMiddlewares() {
	r.e.Use(middleware.Recover())
}

func (r *router) RegisterRoutes() {
	api := r.e.Group("/api/resources", middleware.RequestLogger())

	api.GET("/:id", r.h.RetrieveResource)
	api.GET("", r.h.ListResources)
	api.POST("", r.h.CreateResource)
	api.PUT("/:id", r.h.UpdateResource)
	api.DELETE("/:id", r.h.DeleteResource)

	act := r.e.Group("/actuator")
	act.GET("/health", healthCheck)
	act.GET("/readiness", readinessCheck)
}

func healthCheck(c *echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"status": "up"})
}

func readinessCheck(c *echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"status": "ready"})
}
