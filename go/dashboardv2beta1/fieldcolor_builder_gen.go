// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[FieldColor] = (*FieldColorBuilder)(nil)

// Map a field to a color.
type FieldColorBuilder struct {
	internal *FieldColor
	errors   cog.BuildErrors
}

func NewFieldColorBuilder() *FieldColorBuilder {
	resource := NewFieldColor()
	builder := &FieldColorBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *FieldColorBuilder) Build() (FieldColor, error) {
	if err := builder.internal.Validate(); err != nil {
		return FieldColor{}, err
	}

	if len(builder.errors) > 0 {
		return FieldColor{}, cog.MakeBuildErrors("dashboardv2beta1.fieldColor", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *FieldColorBuilder) RecordError(path string, err error) *FieldColorBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

// The main color scheme mode.
func (builder *FieldColorBuilder) Mode(mode FieldColorModeId) *FieldColorBuilder {
	builder.internal.Mode = mode

	return builder
}

// The fixed color value for fixed or shades color modes.
func (builder *FieldColorBuilder) FixedColor(fixedColor string) *FieldColorBuilder {
	builder.internal.FixedColor = &fixedColor

	return builder
}

// Some visualizations need to know how to assign a series color from by value color schemes.
func (builder *FieldColorBuilder) SeriesBy(seriesBy FieldColorSeriesByMode) *FieldColorBuilder {
	builder.internal.SeriesBy = &seriesBy

	return builder
}
