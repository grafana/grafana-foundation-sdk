// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboard

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[DataSourceRef] = (*DataSourceRefBuilder)(nil)

// Ref to a DataSource instance
type DataSourceRefBuilder struct {
	internal *DataSourceRef
	errors   cog.BuildErrors
}

func NewDataSourceRefBuilder() *DataSourceRefBuilder {
	resource := NewDataSourceRef()
	builder := &DataSourceRefBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *DataSourceRefBuilder) Build() (DataSourceRef, error) {
	if err := builder.internal.Validate(); err != nil {
		return DataSourceRef{}, err
	}

	if len(builder.errors) > 0 {
		return DataSourceRef{}, cog.MakeBuildErrors("dashboard.dataSourceRef", builder.errors)
	}

	return *builder.internal, nil
}

// The plugin type-id
func (builder *DataSourceRefBuilder) Type(typeArg string) *DataSourceRefBuilder {
	builder.internal.Type = &typeArg

	return builder
}

// Specific datasource instance
func (builder *DataSourceRefBuilder) Uid(uid string) *DataSourceRefBuilder {
	builder.internal.Uid = &uid

	return builder
}
