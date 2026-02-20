// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[DataLink] = (*DataLinkBuilder)(nil)

type DataLinkBuilder struct {
	internal *DataLink
	errors   cog.BuildErrors
}

func NewDataLinkBuilder() *DataLinkBuilder {
	resource := NewDataLink()
	builder := &DataLinkBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *DataLinkBuilder) Build() (DataLink, error) {
	if err := builder.internal.Validate(); err != nil {
		return DataLink{}, err
	}

	if len(builder.errors) > 0 {
		return DataLink{}, cog.MakeBuildErrors("dashboardv2beta1.dataLink", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *DataLinkBuilder) RecordError(path string, err error) *DataLinkBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *DataLinkBuilder) Title(title string) *DataLinkBuilder {
	builder.internal.Title = title

	return builder
}

func (builder *DataLinkBuilder) Url(url string) *DataLinkBuilder {
	builder.internal.Url = url

	return builder
}

func (builder *DataLinkBuilder) TargetBlank(targetBlank bool) *DataLinkBuilder {
	builder.internal.TargetBlank = &targetBlank

	return builder
}
