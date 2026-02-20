// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package elasticsearch

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[RawData] = (*RawDataBuilder)(nil)

type RawDataBuilder struct {
	internal *RawData
	errors   cog.BuildErrors
}

func NewRawDataBuilder() *RawDataBuilder {
	resource := NewRawData()
	builder := &RawDataBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *RawDataBuilder) Build() (RawData, error) {
	if err := builder.internal.Validate(); err != nil {
		return RawData{}, err
	}

	if len(builder.errors) > 0 {
		return RawData{}, cog.MakeBuildErrors("elasticsearch.rawData", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *RawDataBuilder) RecordError(path string, err error) *RawDataBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *RawDataBuilder) Id(id string) *RawDataBuilder {
	builder.internal.Id = id

	return builder
}

func (builder *RawDataBuilder) Settings(settings cog.Builder[ElasticsearchRawDataSettings]) *RawDataBuilder {
	settingsResource, err := settings.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Settings = &settingsResource

	return builder
}

func (builder *RawDataBuilder) Hide(hide bool) *RawDataBuilder {
	builder.internal.Hide = &hide

	return builder
}
