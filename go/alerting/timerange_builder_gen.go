// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package alerting

import (
	"time"

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
	if err := builder.internal.Validate(); err != nil {
		return TimeRange{}, err
	}

	return *builder.internal, nil
}

func (builder *TimeRangeBuilder) From(from time.Time) *TimeRangeBuilder {
	builder.internal.From = &from

	return builder
}

func (builder *TimeRangeBuilder) To(to time.Time) *TimeRangeBuilder {
	builder.internal.To = &to

	return builder
}

func (builder *TimeRangeBuilder) applyDefaults() {
}
