// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[MovingAverageModelOption] = (*MovingAverageModelOptionBuilder)(nil)

type MovingAverageModelOptionBuilder struct {
	internal *MovingAverageModelOption
	errors   map[string]cog.BuildErrors
}

func NewMovingAverageModelOptionBuilder() *MovingAverageModelOptionBuilder {
	resource := &MovingAverageModelOption{}
	builder := &MovingAverageModelOptionBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *MovingAverageModelOptionBuilder) Build() (MovingAverageModelOption, error) {
	if err := builder.internal.Validate(); err != nil {
		return MovingAverageModelOption{}, err
	}

	return *builder.internal, nil
}

func (builder *MovingAverageModelOptionBuilder) Label(label string) *MovingAverageModelOptionBuilder {
	builder.internal.Label = label

	return builder
}

func (builder *MovingAverageModelOptionBuilder) Value(value MovingAverageModel) *MovingAverageModelOptionBuilder {
	builder.internal.Value = value

	return builder
}

func (builder *MovingAverageModelOptionBuilder) applyDefaults() {
}
