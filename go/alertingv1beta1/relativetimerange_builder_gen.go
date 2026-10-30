// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package alertingv1beta1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[RelativeTimeRange] = (*RelativeTimeRangeBuilder)(nil)

type RelativeTimeRangeBuilder struct {
	internal *RelativeTimeRange
	errors   cog.BuildErrors
}

func NewRelativeTimeRangeBuilder() *RelativeTimeRangeBuilder {
	resource := NewRelativeTimeRange()
	builder := &RelativeTimeRangeBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *RelativeTimeRangeBuilder) Build() (RelativeTimeRange, error) {
	if err := builder.internal.Validate(); err != nil {
		return RelativeTimeRange{}, err
	}

	if len(builder.errors) > 0 {
		return RelativeTimeRange{}, cog.MakeBuildErrors("alertingv1beta1.relativeTimeRange", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *RelativeTimeRangeBuilder) RecordError(path string, err error) *RelativeTimeRangeBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *RelativeTimeRangeBuilder) From(from PromDurationWMillis) *RelativeTimeRangeBuilder {
	builder.internal.From = from

	return builder
}

func (builder *RelativeTimeRangeBuilder) To(to PromDurationWMillis) *RelativeTimeRangeBuilder {
	builder.internal.To = to

	return builder
}
