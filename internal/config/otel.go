package config

import (
	"context"

	"github.com/bencoronard/demo-go-common-libs/otel"
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

func NewTracerProvider(res *resource.Resource, p *Properties) (*trace.TracerProvider, error) {
	return otel.NewTracerProvider(
		context.Background(),
		res,
		p.Env.Otel.TracesEndpoint, p.Env.Otel.TracesBatchTimeoutInSec,
		p.Env.Otel.TracesSamplingProbability,
	)
}

func NewMeterProvider(res *resource.Resource, p *Properties) (*metric.MeterProvider, error) {
	return otel.NewMeterProvider(
		context.Background(),
		res,
		p.Env.Otel.MetricsEndpoint,
		p.Env.Otel.MetricsSamplingFreqInSec,
	)
}
