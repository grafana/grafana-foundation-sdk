// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package canvas

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ConnectionCoordinates] = (*ConnectionCoordinatesBuilder)(nil)

type ConnectionCoordinatesBuilder struct {
	internal *ConnectionCoordinates
	errors   cog.BuildErrors
}

func NewConnectionCoordinatesBuilder() *ConnectionCoordinatesBuilder {
	resource := NewConnectionCoordinates()
	builder := &ConnectionCoordinatesBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *ConnectionCoordinatesBuilder) Build() (ConnectionCoordinates, error) {
	if err := builder.internal.Validate(); err != nil {
		return ConnectionCoordinates{}, err
	}

	if len(builder.errors) > 0 {
		return ConnectionCoordinates{}, cog.MakeBuildErrors("canvas.connectionCoordinates", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *ConnectionCoordinatesBuilder) RecordError(path string, err error) *ConnectionCoordinatesBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *ConnectionCoordinatesBuilder) X(x float64) *ConnectionCoordinatesBuilder {
	builder.internal.X = x

	return builder
}

func (builder *ConnectionCoordinatesBuilder) Y(y float64) *ConnectionCoordinatesBuilder {
	builder.internal.Y = y

	return builder
}
