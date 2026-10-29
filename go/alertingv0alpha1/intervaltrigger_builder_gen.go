// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package alertingv0alpha1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[IntervalTrigger] = (*IntervalTriggerBuilder)(nil)

type IntervalTriggerBuilder struct {
	internal *IntervalTrigger
	errors   cog.BuildErrors
}

func NewIntervalTriggerBuilder() *IntervalTriggerBuilder {
	resource := NewIntervalTrigger()
	builder := &IntervalTriggerBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *IntervalTriggerBuilder) Build() (IntervalTrigger, error) {
	if err := builder.internal.Validate(); err != nil {
		return IntervalTrigger{}, err
	}

	if len(builder.errors) > 0 {
		return IntervalTrigger{}, cog.MakeBuildErrors("alertingv0alpha1.intervalTrigger", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *IntervalTriggerBuilder) RecordError(path string, err error) *IntervalTriggerBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *IntervalTriggerBuilder) Interval(interval PromDuration) *IntervalTriggerBuilder {
	builder.internal.Interval = interval

	return builder
}
