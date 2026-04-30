// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[DatasourceControlSourceRef] = (*DatasourceControlSourceRefBuilder)(nil)

// Source information for controls (e.g. variables or links)
type DatasourceControlSourceRefBuilder struct {
	internal *DatasourceControlSourceRef
	errors   cog.BuildErrors
}

func NewDatasourceControlSourceRefBuilder() *DatasourceControlSourceRefBuilder {
	resource := NewDatasourceControlSourceRef()
	builder := &DatasourceControlSourceRefBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}
	builder.internal.Type = "datasource"

	return builder
}

func (builder *DatasourceControlSourceRefBuilder) Build() (DatasourceControlSourceRef, error) {
	if err := builder.internal.Validate(); err != nil {
		return DatasourceControlSourceRef{}, err
	}

	if len(builder.errors) > 0 {
		return DatasourceControlSourceRef{}, cog.MakeBuildErrors("dashboardv2.datasourceControlSourceRef", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *DatasourceControlSourceRefBuilder) RecordError(path string, err error) *DatasourceControlSourceRefBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

// The plugin type-id
func (builder *DatasourceControlSourceRefBuilder) Group(group string) *DatasourceControlSourceRefBuilder {
	builder.internal.Group = group

	return builder
}
