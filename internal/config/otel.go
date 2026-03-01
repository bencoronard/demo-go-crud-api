package config

import (
	"context"

	"github.com/bencoronard/demo-go-common-libs/otel"
	"go.opentelemetry.io/otel/sdk/log"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.39.0"
)

func NewResource(p *Properties) (*resource.Resource, error) {
	return resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceNameKey.String(p.Env.App.Name),
		semconv.DeploymentEnvironmentNameKey.String(p.Env.App.Environment),
	), nil
}

func NewTracerProvider(res *resource.Resource) (*trace.TracerProvider, error) {
	return otel.NewTracerProvider(context.Background(), res)
}

func NewMeterProvider(res *resource.Resource) (*metric.MeterProvider, error) {
	return otel.NewMeterProvider(context.Background(), res)
}

func NewLoggerProvider(res *resource.Resource) (*log.LoggerProvider, error) {
	return otel.NewLoggerProvider(context.Background(), res)
}
