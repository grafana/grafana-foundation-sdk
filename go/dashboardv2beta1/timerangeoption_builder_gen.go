// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[TimeRangeOption] = (*TimeRangeOptionBuilder)(nil)

type TimeRangeOptionBuilder struct {
	internal *TimeRangeOption
	errors   cog.BuildErrors
}

func NewTimeRangeOptionBuilder() *TimeRangeOptionBuilder {
	resource := NewTimeRangeOption()
	builder := &TimeRangeOptionBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *TimeRangeOptionBuilder) Build() (TimeRangeOption, error) {
	if err := builder.internal.Validate(); err != nil {
		return TimeRangeOption{}, err
	}

	if len(builder.errors) > 0 {
		return TimeRangeOption{}, cog.MakeBuildErrors("dashboardv2beta1.timeRangeOption", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *TimeRangeOptionBuilder) Display(display string) *TimeRangeOptionBuilder {
	builder.internal.Display = display

	return builder
}

func (builder *TimeRangeOptionBuilder) From(from string) *TimeRangeOptionBuilder {
	builder.internal.From = from

	return builder
}

func (builder *TimeRangeOptionBuilder) To(to string) *TimeRangeOptionBuilder {
	builder.internal.To = to

	return builder
}
