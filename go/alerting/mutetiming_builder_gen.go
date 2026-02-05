// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package alerting

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[MuteTiming] = (*MuteTimingBuilder)(nil)

type MuteTimingBuilder struct {
	internal *MuteTiming
	errors   cog.BuildErrors
}

func NewMuteTimingBuilder() *MuteTimingBuilder {
	resource := NewMuteTiming()
	builder := &MuteTimingBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *MuteTimingBuilder) Build() (MuteTiming, error) {
	if err := builder.internal.Validate(); err != nil {
		return MuteTiming{}, err
	}

	if len(builder.errors) > 0 {
		return MuteTiming{}, cog.MakeBuildErrors("alerting.muteTiming", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *MuteTimingBuilder) Name(name string) *MuteTimingBuilder {
	builder.internal.Name = &name

	return builder
}

func (builder *MuteTimingBuilder) TimeIntervals(timeIntervals []cog.Builder[TimeInterval]) *MuteTimingBuilder {
	timeIntervalsResources := make([]TimeInterval, 0, len(timeIntervals))
	for _, r1 := range timeIntervals {
		timeIntervalsDepth1, err := r1.Build()
		if err != nil {
			builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
			return builder
		}
		timeIntervalsResources = append(timeIntervalsResources, timeIntervalsDepth1)
	}
	builder.internal.TimeIntervals = timeIntervalsResources

	return builder
}
