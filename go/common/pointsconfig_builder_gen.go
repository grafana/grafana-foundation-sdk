// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package common

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[PointsConfig] = (*PointsConfigBuilder)(nil)

// TODO docs
type PointsConfigBuilder struct {
	internal *PointsConfig
	errors   cog.BuildErrors
}

func NewPointsConfigBuilder() *PointsConfigBuilder {
	resource := NewPointsConfig()
	builder := &PointsConfigBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *PointsConfigBuilder) Build() (PointsConfig, error) {
	if err := builder.internal.Validate(); err != nil {
		return PointsConfig{}, err
	}

	if len(builder.errors) > 0 {
		return PointsConfig{}, cog.MakeBuildErrors("common.pointsConfig", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *PointsConfigBuilder) ShowPoints(showPoints VisibilityMode) *PointsConfigBuilder {
	builder.internal.ShowPoints = &showPoints

	return builder
}

func (builder *PointsConfigBuilder) PointSize(pointSize float64) *PointsConfigBuilder {
	builder.internal.PointSize = &pointSize

	return builder
}

func (builder *PointsConfigBuilder) PointColor(pointColor string) *PointsConfigBuilder {
	builder.internal.PointColor = &pointColor

	return builder
}

func (builder *PointsConfigBuilder) PointSymbol(pointSymbol string) *PointsConfigBuilder {
	builder.internal.PointSymbol = &pointSymbol

	return builder
}
