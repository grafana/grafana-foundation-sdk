// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[ExtendedStat] = (*ExtendedStatBuilder)(nil)

type ExtendedStatBuilder struct {
	internal *ExtendedStat
	errors   cog.BuildErrors
}

func NewExtendedStatBuilder() *ExtendedStatBuilder {
	resource := NewExtendedStat()
	builder := &ExtendedStatBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *ExtendedStatBuilder) Build() (ExtendedStat, error) {
	if err := builder.internal.Validate(); err != nil {
		return ExtendedStat{}, err
	}

	if len(builder.errors) > 0 {
		return ExtendedStat{}, cog.MakeBuildErrors("elasticsearch.extendedStat", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *ExtendedStatBuilder) RecordError(path string, err error) *ExtendedStatBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *ExtendedStatBuilder) Label(label string) *ExtendedStatBuilder {
	builder.internal.Label = label

	return builder
}

func (builder *ExtendedStatBuilder) Value(value ExtendedStatMetaType) *ExtendedStatBuilder {
	builder.internal.Value = value

	return builder
}
