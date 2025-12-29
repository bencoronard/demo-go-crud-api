package server

import (
	"github.com/bencoronard/demo-go-crud-api/internal/config"
	"github.com/bencoronard/demo-go-crud-api/internal/resource"
	"go.uber.org/fx"
)

func Start() {
	fx.New(
		fx.Provide(
			config.NewProperties,
			config.NewLogger,
			config.NewDB,
			config.NewJwtVerifier,
			config.NewAuthHeaderResolver,
			resource.NewResourceRepoImpl,
			resource.NewResourceServiceImpl,
			resource.NewResourceHandler,
			config.NewRouter,
		),
		fx.Invoke(
			config.Router.Start,
		),
	).Run()
}
