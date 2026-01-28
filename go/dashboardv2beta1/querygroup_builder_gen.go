// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[QueryGroupKind] = (*QueryGroupBuilder)(nil)

type QueryGroupBuilder struct {
	internal *QueryGroupKind
	errors   cog.BuildErrors
}

func NewQueryGroupBuilder() *QueryGroupBuilder {
	resource := NewQueryGroupKind()
	builder := &QueryGroupBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}
	builder.internal.Kind = "QueryGroup"

	return builder
}

func (builder *QueryGroupBuilder) Build() (QueryGroupKind, error) {
	if err := builder.internal.Validate(); err != nil {
		return QueryGroupKind{}, err
	}

	if len(builder.errors) > 0 {
		return QueryGroupKind{}, cog.MakeBuildErrors("dashboardv2beta1.queryGroup", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *QueryGroupBuilder) Targets(queries []cog.Builder[PanelQueryKind]) *QueryGroupBuilder {
	queriesResources := make([]PanelQueryKind, 0, len(queries))
	for _, r1 := range queries {
		queriesDepth1, err := r1.Build()
		if err != nil {
			builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
			return builder
		}
		queriesResources = append(queriesResources, queriesDepth1)
	}
	builder.internal.Spec.Queries = queriesResources

	return builder
}

func (builder *QueryGroupBuilder) Target(querie cog.Builder[PanelQueryKind]) *QueryGroupBuilder {
	querieResource, err := querie.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Spec.Queries = append(builder.internal.Spec.Queries, querieResource)

	return builder
}

func (builder *QueryGroupBuilder) Transformations(transformations []cog.Builder[TransformationKind]) *QueryGroupBuilder {
	transformationsResources := make([]TransformationKind, 0, len(transformations))
	for _, r1 := range transformations {
		transformationsDepth1, err := r1.Build()
		if err != nil {
			builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
			return builder
		}
		transformationsResources = append(transformationsResources, transformationsDepth1)
	}
	builder.internal.Spec.Transformations = transformationsResources

	return builder
}

func (builder *QueryGroupBuilder) Transformation(transformation cog.Builder[TransformationKind]) *QueryGroupBuilder {
	transformationResource, err := transformation.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Spec.Transformations = append(builder.internal.Spec.Transformations, transformationResource)

	return builder
}

func (builder *QueryGroupBuilder) QueryOptions(queryOptions cog.Builder[QueryOptionsSpec]) *QueryGroupBuilder {
	queryOptionsResource, err := queryOptions.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Spec.QueryOptions = queryOptionsResource

	return builder
}
