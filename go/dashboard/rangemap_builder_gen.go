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
	resource := &RangeMap{}
	builder := &RangeMapBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()
	builder.internal.Type = "range"

	return builder
}

func (builder *RangeMapBuilder) Build() (RangeMap, error) {
	var errs cog.BuildErrors

	for _, err := range builder.errors {
		errs = append(errs, cog.MakeBuildErrors("RangeMap", err)...)
	}

	if len(errs) != 0 {
		return RangeMap{}, errs
	}

	return *builder.internal, nil
}

// Range to match against and the result to apply when the value is within the range
func (builder *RangeMapBuilder) Options(options struct {
	// Min value of the range. It can be null which means -Infinity
	From *float64 `json:"from"`
	// Max value of the range. It can be null which means +Infinity
	To *float64 `json:"to"`
	// Config to apply when the value is within the range
	Result ValueMappingResult `json:"result"`
}) *RangeMapBuilder {
	builder.internal.Options = options

	return builder
}

func (builder *RangeMapBuilder) applyDefaults() {
}
