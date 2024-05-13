// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package alerting

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[TimeRange] = (*TimeRangeBuilder)(nil)

// Redefining this to avoid an import cycle
type TimeRangeBuilder struct {
	internal *TimeRange
	errors   map[string]cog.BuildErrors
}

func NewTimeRangeBuilder() *TimeRangeBuilder {
	resource := &TimeRange{}
	builder := &TimeRangeBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *TimeRangeBuilder) Build() (TimeRange, error) {
	var errs cog.BuildErrors

	for _, err := range builder.errors {
		errs = append(errs, cog.MakeBuildErrors("TimeRange", err)...)
	}

	if len(errs) != 0 {
		return TimeRange{}, errs
	}

	return *builder.internal, nil
}

// Redefining this to avoid an import cycle
func (builder *TimeRangeBuilder) From(from string) *TimeRangeBuilder {
	builder.internal.From = &from

	return builder
}

// Redefining this to avoid an import cycle
func (builder *TimeRangeBuilder) To(to string) *TimeRangeBuilder {
	builder.internal.To = &to

	return builder
}

func (builder *TimeRangeBuilder) applyDefaults() {
}
