// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[RepeatOptions] = (*RepeatOptionsBuilder)(nil)

type RepeatOptionsBuilder struct {
	internal *RepeatOptions
	errors   cog.BuildErrors
}

func NewRepeatOptionsBuilder() *RepeatOptionsBuilder {
	resource := NewRepeatOptions()
	builder := &RepeatOptionsBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *RepeatOptionsBuilder) Build() (RepeatOptions, error) {
	if err := builder.internal.Validate(); err != nil {
		return RepeatOptions{}, err
	}

	if len(builder.errors) > 0 {
		return RepeatOptions{}, cog.MakeBuildErrors("dashboardv2beta1.repeatOptions", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *RepeatOptionsBuilder) Value(value string) *RepeatOptionsBuilder {
	builder.internal.Value = value

	return builder
}

func (builder *RepeatOptionsBuilder) Direction(direction RepeatOptionsDirection) *RepeatOptionsBuilder {
	builder.internal.Direction = &direction

	return builder
}

func (builder *RepeatOptionsBuilder) MaxPerRow(maxPerRow int64) *RepeatOptionsBuilder {
	builder.internal.MaxPerRow = &maxPerRow

	return builder
}
