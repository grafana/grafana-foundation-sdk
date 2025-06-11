// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package common

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[LineStyle] = (*LineStyleBuilder)(nil)

// TODO docs
type LineStyleBuilder struct {
	internal *LineStyle
	errors   cog.BuildErrors
}

func NewLineStyleBuilder() *LineStyleBuilder {
	resource := NewLineStyle()
	builder := &LineStyleBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *LineStyleBuilder) Build() (LineStyle, error) {
	if err := builder.internal.Validate(); err != nil {
		return LineStyle{}, err
	}

	if len(builder.errors) > 0 {
		return LineStyle{}, cog.MakeBuildErrors("common.lineStyle", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *LineStyleBuilder) Fill(fill LineStyleFill) *LineStyleBuilder {
	builder.internal.Fill = &fill

	return builder
}

func (builder *LineStyleBuilder) Dash(dash []float64) *LineStyleBuilder {
	builder.internal.Dash = dash

	return builder
}
