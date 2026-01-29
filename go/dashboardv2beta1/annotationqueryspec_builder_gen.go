// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[AnnotationQuerySpec] = (*AnnotationQuerySpecBuilder)(nil)

type AnnotationQuerySpecBuilder struct {
	internal *AnnotationQuerySpec
	errors   cog.BuildErrors
}

func NewAnnotationQuerySpecBuilder() *AnnotationQuerySpecBuilder {
	resource := NewAnnotationQuerySpec()
	builder := &AnnotationQuerySpecBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *AnnotationQuerySpecBuilder) Build() (AnnotationQuerySpec, error) {
	if err := builder.internal.Validate(); err != nil {
		return AnnotationQuerySpec{}, err
	}

	if len(builder.errors) > 0 {
		return AnnotationQuerySpec{}, cog.MakeBuildErrors("dashboardv2beta1.annotationQuerySpec", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *AnnotationQuerySpecBuilder) Query(query cog.Builder[DataQueryKind]) *AnnotationQuerySpecBuilder {
	queryResource, err := query.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Query = queryResource

	return builder
}

func (builder *AnnotationQuerySpecBuilder) Enable(enable bool) *AnnotationQuerySpecBuilder {
	builder.internal.Enable = enable

	return builder
}

func (builder *AnnotationQuerySpecBuilder) Hide(hide bool) *AnnotationQuerySpecBuilder {
	builder.internal.Hide = hide

	return builder
}

func (builder *AnnotationQuerySpecBuilder) IconColor(iconColor string) *AnnotationQuerySpecBuilder {
	builder.internal.IconColor = iconColor

	return builder
}

func (builder *AnnotationQuerySpecBuilder) Name(name string) *AnnotationQuerySpecBuilder {
	builder.internal.Name = name

	return builder
}

func (builder *AnnotationQuerySpecBuilder) BuiltIn(builtIn bool) *AnnotationQuerySpecBuilder {
	builder.internal.BuiltIn = &builtIn

	return builder
}

func (builder *AnnotationQuerySpecBuilder) Filter(filter cog.Builder[AnnotationPanelFilter]) *AnnotationQuerySpecBuilder {
	filterResource, err := filter.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Filter = &filterResource

	return builder
}

// Placement can be used to display the annotation query somewhere else on the dashboard other than the default location.
func (builder *AnnotationQuerySpecBuilder) Placement(placement string) *AnnotationQuerySpecBuilder {
	builder.internal.Placement = &placement

	return builder
}

// Mappings define how to convert data frame fields to annotation event fields.
func (builder *AnnotationQuerySpecBuilder) Mappings(mappings map[string]cog.Builder[AnnotationEventFieldMapping]) *AnnotationQuerySpecBuilder {
	mappingsResource := make(map[string]AnnotationEventFieldMapping)
	for key1, val1 := range mappings {
		mappingsDepth1, err := val1.Build()
		if err != nil {
			builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
			return builder
		}
		mappingsResource[key1] = mappingsDepth1
	}
	builder.internal.Mappings = mappingsResource

	return builder
}

// Catch-all field for datasource-specific properties. Should not be available in as code tooling.
func (builder *AnnotationQuerySpecBuilder) LegacyOptions(legacyOptions map[string]any) *AnnotationQuerySpecBuilder {
	builder.internal.LegacyOptions = legacyOptions

	return builder
}
