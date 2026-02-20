// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[DataQueryKind] = (*DataQueryKindBuilder)(nil)

type DataQueryKindBuilder struct {
	internal *DataQueryKind
	errors   cog.BuildErrors
}

func NewDataQueryKindBuilder() *DataQueryKindBuilder {
	resource := NewDataQueryKind()
	builder := &DataQueryKindBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}
	builder.internal.Kind = "DataQuery"

	return builder
}

func (builder *DataQueryKindBuilder) Build() (DataQueryKind, error) {
	if err := builder.internal.Validate(); err != nil {
		return DataQueryKind{}, err
	}

	if len(builder.errors) > 0 {
		return DataQueryKind{}, cog.MakeBuildErrors("dashboardv2beta1.dataQueryKind", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *DataQueryKindBuilder) RecordError(path string, err error) *DataQueryKindBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *DataQueryKindBuilder) Group(group string) *DataQueryKindBuilder {
	builder.internal.Group = group

	return builder
}

func (builder *DataQueryKindBuilder) Version(version string) *DataQueryKindBuilder {
	builder.internal.Version = version

	return builder
}

// New type for datasource reference
// Not creating a new type until we figure out how to handle DS refs for group by, adhoc, and every place that uses DataSourceRef in TS.
func (builder *DataQueryKindBuilder) Datasource(datasource cog.Builder[Dashboardv2beta1DataQueryKindDatasource]) *DataQueryKindBuilder {
	datasourceResource, err := datasource.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Datasource = &datasourceResource

	return builder
}

func (builder *DataQueryKindBuilder) Spec(spec any) *DataQueryKindBuilder {
	builder.internal.Spec = &spec

	return builder
}
