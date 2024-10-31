// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package common

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[LineStyle] = (*LineStyleBuilder)(nil)

// TODO docs
type LineStyleBuilder struct {
	internal *LineStyle
	errors   map[string]cog.BuildErrors
}

func NewLineStyleBuilder() *LineStyleBuilder {
	resource := &LineStyle{}
	builder := &LineStyleBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *LineStyleBuilder) Build() (LineStyle, error) {
	if err := builder.internal.Validate(); err != nil {
		return LineStyle{}, err
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

func (builder *LineStyleBuilder) applyDefaults() {
}
