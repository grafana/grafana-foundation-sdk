// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package alerting

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[MuteTiming] = (*MuteTimingBuilder)(nil)

type MuteTimingBuilder struct {
	internal *MuteTiming
	errors   map[string]cog.BuildErrors
}

func NewMuteTimingBuilder() *MuteTimingBuilder {
	resource := &MuteTiming{}
	builder := &MuteTimingBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *MuteTimingBuilder) Build() (MuteTiming, error) {
	if err := builder.internal.Validate(); err != nil {
		return MuteTiming{}, err
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
			builder.errors["time_intervals"] = err.(cog.BuildErrors)
			return builder
		}
		timeIntervalsResources = append(timeIntervalsResources, timeIntervalsDepth1)
	}
	builder.internal.TimeIntervals = timeIntervalsResources

	return builder
}

func (builder *MuteTimingBuilder) applyDefaults() {
}
