// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package alertingv1beta1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[SimplifiedRouting] = (*SimplifiedRoutingBuilder)(nil)

type SimplifiedRoutingBuilder struct {
	internal *SimplifiedRouting
	errors   cog.BuildErrors
}

func NewSimplifiedRoutingBuilder() *SimplifiedRoutingBuilder {
	resource := NewSimplifiedRouting()
	builder := &SimplifiedRoutingBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *SimplifiedRoutingBuilder) Build() (SimplifiedRouting, error) {
	if err := builder.internal.Validate(); err != nil {
		return SimplifiedRouting{}, err
	}

	if len(builder.errors) > 0 {
		return SimplifiedRouting{}, cog.MakeBuildErrors("alertingv1beta1.simplifiedRouting", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *SimplifiedRoutingBuilder) RecordError(path string, err error) *SimplifiedRoutingBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *SimplifiedRoutingBuilder) Receiver(receiver string) *SimplifiedRoutingBuilder {
	builder.internal.Receiver = receiver

	return builder
}

func (builder *SimplifiedRoutingBuilder) GroupBy(groupBy []string) *SimplifiedRoutingBuilder {
	builder.internal.GroupBy = groupBy

	return builder
}

func (builder *SimplifiedRoutingBuilder) GroupWait(groupWait PromDuration) *SimplifiedRoutingBuilder {
	builder.internal.GroupWait = &groupWait

	return builder
}

func (builder *SimplifiedRoutingBuilder) GroupInterval(groupInterval PromDuration) *SimplifiedRoutingBuilder {
	builder.internal.GroupInterval = &groupInterval

	return builder
}

func (builder *SimplifiedRoutingBuilder) RepeatInterval(repeatInterval PromDuration) *SimplifiedRoutingBuilder {
	builder.internal.RepeatInterval = &repeatInterval

	return builder
}

func (builder *SimplifiedRoutingBuilder) MuteTimeIntervals(muteTimeIntervals []TimeIntervalRef) *SimplifiedRoutingBuilder {
	builder.internal.MuteTimeIntervals = muteTimeIntervals

	return builder
}

func (builder *SimplifiedRoutingBuilder) ActiveTimeIntervals(activeTimeIntervals []TimeIntervalRef) *SimplifiedRoutingBuilder {
	builder.internal.ActiveTimeIntervals = activeTimeIntervals

	return builder
}
