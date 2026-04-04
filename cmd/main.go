package main

import (
	"github.com/bencoronard/demo-go-common-libs/http"

	"github.com/bencoronard/demo-go-common-libs/otel"
	"github.com/bencoronard/demo-go-common-libs/rdb"
	"github.com/bencoronard/demo-go-common-libs/validator"

	"github.com/bencoronard/demo-go-crud-api/internal/config"
	"github.com/bencoronard/demo-go-crud-api/internal/resource"
	"go.uber.org/fx"
)

func main() {
	fx.New(
		fx.Provide(
			config.NewProperties,
			config.NewDB,
			rdb.NewTransactionManager,
			config.NewJwtVerifier,
			config.NewAuthHeaderResolver,
			validator.New,
			config.NewRouter,
			config.NewAppErrorHandler,
		),
		fx.Provide(
			resource.NewResourceRepo,
			resource.NewResourceService,
			resource.NewResourceHandler,
		),
		fx.Provide(
			config.NewResource,
			otel.NewTracerProvider,
			otel.NewMeterProvider,
			otel.NewLoggerProvider,
		),
		fx.Provide(
			http.NewGlobalErrorHandler,
		),
		fx.Invoke(
			config.ConfigureLogger,
			http.Router.RegisterMiddlewares,
			http.Router.RegisterRoutes,
		),
		fx.Invoke(http.StartServer),
	).Run()
}
