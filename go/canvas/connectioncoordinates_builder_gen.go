// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package canvas

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ConnectionCoordinates] = (*ConnectionCoordinatesBuilder)(nil)

type ConnectionCoordinatesBuilder struct {
	internal *ConnectionCoordinates
	errors   map[string]cog.BuildErrors
}

func NewConnectionCoordinatesBuilder() *ConnectionCoordinatesBuilder {
	resource := &ConnectionCoordinates{}
	builder := &ConnectionCoordinatesBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *ConnectionCoordinatesBuilder) Build() (ConnectionCoordinates, error) {
	if err := builder.internal.Validate(); err != nil {
		return ConnectionCoordinates{}, err
	}

	return *builder.internal, nil
}

func (builder *ConnectionCoordinatesBuilder) X(x float64) *ConnectionCoordinatesBuilder {
	builder.internal.X = x

	return builder
}

func (builder *ConnectionCoordinatesBuilder) Y(y float64) *ConnectionCoordinatesBuilder {
	builder.internal.Y = y

	return builder
}

func (builder *ConnectionCoordinatesBuilder) applyDefaults() {
}
