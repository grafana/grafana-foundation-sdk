// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboard

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[RangeMap] = (*RangeMapBuilder)(nil)

// Maps numerical ranges to a display text and color.
// For example, if a value is within a certain range, you can configure a range value mapping to display Low or High rather than the number.
type RangeMapBuilder struct {
	internal *RangeMap
	errors   map[string]cog.BuildErrors
}

func NewRangeMapBuilder() *RangeMapBuilder {
	resource := NewRangeMap()
	builder := &RangeMapBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}
	builder.internal.Type = "range"

	return builder
}

func (builder *RangeMapBuilder) Build() (RangeMap, error) {
	if err := builder.internal.Validate(); err != nil {
		return RangeMap{}, err
	}

	return *builder.internal, nil
}

// Range to match against and the result to apply when the value is within the range
func (builder *RangeMapBuilder) Options(options cog.Builder[DashboardRangeMapOptions]) *RangeMapBuilder {
	optionsResource, err := options.Build()
	if err != nil {
		builder.errors["options"] = err.(cog.BuildErrors)
		return builder
	}
	builder.internal.Options = optionsResource

	return builder
}
