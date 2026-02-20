// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ValueMappingResult] = (*ValueMappingResultBuilder)(nil)

// Result used as replacement with text and color when the value matches
type ValueMappingResultBuilder struct {
	internal *ValueMappingResult
	errors   cog.BuildErrors
}

func NewValueMappingResultBuilder() *ValueMappingResultBuilder {
	resource := NewValueMappingResult()
	builder := &ValueMappingResultBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *ValueMappingResultBuilder) Build() (ValueMappingResult, error) {
	if err := builder.internal.Validate(); err != nil {
		return ValueMappingResult{}, err
	}

	if len(builder.errors) > 0 {
		return ValueMappingResult{}, cog.MakeBuildErrors("dashboardv2beta1.valueMappingResult", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *ValueMappingResultBuilder) RecordError(path string, err error) *ValueMappingResultBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

// Text to display when the value matches
func (builder *ValueMappingResultBuilder) Text(text string) *ValueMappingResultBuilder {
	builder.internal.Text = &text

	return builder
}

// Text to use when the value matches
func (builder *ValueMappingResultBuilder) Color(color string) *ValueMappingResultBuilder {
	builder.internal.Color = &color

	return builder
}

// Icon to display when the value matches. Only specific visualizations.
func (builder *ValueMappingResultBuilder) Icon(icon string) *ValueMappingResultBuilder {
	builder.internal.Icon = &icon

	return builder
}

// Position in the mapping array. Only used internally.
func (builder *ValueMappingResultBuilder) Index(index int32) *ValueMappingResultBuilder {
	builder.internal.Index = &index

	return builder
}
