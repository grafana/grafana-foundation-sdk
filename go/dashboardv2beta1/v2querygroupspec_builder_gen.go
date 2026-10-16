// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[V2QueryGroupSpec] = (*V2QueryGroupSpecBuilder)(nil)

type V2QueryGroupSpecBuilder struct {
	internal *V2QueryGroupSpec
	errors   cog.BuildErrors
}

func NewV2QueryGroupSpecBuilder() *V2QueryGroupSpecBuilder {
	resource := NewV2QueryGroupSpec()
	builder := &V2QueryGroupSpecBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *V2QueryGroupSpecBuilder) Build() (V2QueryGroupSpec, error) {
	if err := builder.internal.Validate(); err != nil {
		return V2QueryGroupSpec{}, err
	}

	if len(builder.errors) > 0 {
		return V2QueryGroupSpec{}, cog.MakeBuildErrors("dashboardv2beta1.v2QueryGroupSpec", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *V2QueryGroupSpecBuilder) RecordError(path string, err error) *V2QueryGroupSpecBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *V2QueryGroupSpecBuilder) Queries(queries []cog.Builder[PanelQueryKind]) *V2QueryGroupSpecBuilder {
	queriesResources := make([]PanelQueryKind, 0, len(queries))
	for _, r1 := range queries {
		queriesDepth1, err := r1.Build()
		if err != nil {
			builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
			return builder
		}
		queriesResources = append(queriesResources, queriesDepth1)
	}
	builder.internal.Queries = queriesResources

	return builder
}

func (builder *V2QueryGroupSpecBuilder) Transformations(transformations []cog.Builder[V2TransformationKind]) *V2QueryGroupSpecBuilder {
	transformationsResources := make([]V2TransformationKind, 0, len(transformations))
	for _, r1 := range transformations {
		transformationsDepth1, err := r1.Build()
		if err != nil {
			builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
			return builder
		}
		transformationsResources = append(transformationsResources, transformationsDepth1)
	}
	builder.internal.Transformations = transformationsResources

	return builder
}

func (builder *V2QueryGroupSpecBuilder) QueryOptions(queryOptions cog.Builder[QueryOptionsSpec]) *V2QueryGroupSpecBuilder {
	queryOptionsResource, err := queryOptions.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.QueryOptions = queryOptionsResource

	return builder
}
