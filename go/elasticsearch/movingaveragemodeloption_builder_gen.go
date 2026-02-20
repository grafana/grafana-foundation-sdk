// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[MovingAverageModelOption] = (*MovingAverageModelOptionBuilder)(nil)

type MovingAverageModelOptionBuilder struct {
	internal *MovingAverageModelOption
	errors   cog.BuildErrors
}

func NewMovingAverageModelOptionBuilder() *MovingAverageModelOptionBuilder {
	resource := NewMovingAverageModelOption()
	builder := &MovingAverageModelOptionBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *MovingAverageModelOptionBuilder) Build() (MovingAverageModelOption, error) {
	if err := builder.internal.Validate(); err != nil {
		return MovingAverageModelOption{}, err
	}

	if len(builder.errors) > 0 {
		return MovingAverageModelOption{}, cog.MakeBuildErrors("elasticsearch.movingAverageModelOption", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *MovingAverageModelOptionBuilder) RecordError(path string, err error) *MovingAverageModelOptionBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *MovingAverageModelOptionBuilder) Label(label string) *MovingAverageModelOptionBuilder {
	builder.internal.Label = label

	return builder
}

func (builder *MovingAverageModelOptionBuilder) Value(value MovingAverageModel) *MovingAverageModelOptionBuilder {
	builder.internal.Value = value

	return builder
}
