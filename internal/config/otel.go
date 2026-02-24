package config

import (
	"context"

	"github.com/bencoronard/demo-go-common-libs/otel"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/trace"
)

func NewTracerProvider(p *Properties) (*trace.TracerProvider, error) {
	return otel.NewTracerProvider(
		context.Background(),
		p.Env.Otel.TracesEndpoint, p.Env.Otel.TracesBatchTimeoutInSec,
		p.Env.Otel.TracesSamplingProbability,
	)
}

func NewMeterProvider(p *Properties) (*metric.MeterProvider, error) {
	return otel.NewMeterProvider(
		context.Background(),
		p.Env.Otel.MetricsEndpoint,
		p.Env.Otel.MetricsSamplingFreqInSec,
	)
}
