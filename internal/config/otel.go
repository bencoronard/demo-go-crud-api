package config

import (
	"go.opentelemetry.io/otel/sdk/resource"
)

func NewResource(p *Properties) (*resource.Resource, error) {
	return resource.Environment(), nil
}
