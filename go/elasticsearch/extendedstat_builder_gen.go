// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ExtendedStat] = (*ExtendedStatBuilder)(nil)

type ExtendedStatBuilder struct {
	internal *ExtendedStat
	errors   map[string]cog.BuildErrors
}

func NewExtendedStatBuilder() *ExtendedStatBuilder {
	resource := &ExtendedStat{}
	builder := &ExtendedStatBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *ExtendedStatBuilder) Build() (ExtendedStat, error) {
	if err := builder.internal.Validate(); err != nil {
		return ExtendedStat{}, err
	}

	return *builder.internal, nil
}

func (builder *ExtendedStatBuilder) Label(label string) *ExtendedStatBuilder {
	builder.internal.Label = label

	return builder
}

func (builder *ExtendedStatBuilder) Value(value ExtendedStatMetaType) *ExtendedStatBuilder {
	builder.internal.Value = value

	return builder
}

func (builder *ExtendedStatBuilder) applyDefaults() {
}
