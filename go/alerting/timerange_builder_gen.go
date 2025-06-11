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
	errors   cog.BuildErrors
}

func NewTimeRangeBuilder() *TimeRangeBuilder {
	resource := NewTimeRange()
	builder := &TimeRangeBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *TimeRangeBuilder) Build() (TimeRange, error) {
	if err := builder.internal.Validate(); err != nil {
		return TimeRange{}, err
	}

	if len(builder.errors) > 0 {
		return TimeRange{}, cog.MakeBuildErrors("alerting.timeRange", builder.errors)
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
