// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[AnnotationQueryKind] = (*AnnotationQueryBuilder)(nil)

type AnnotationQueryBuilder struct {
	internal *AnnotationQueryKind
	errors   cog.BuildErrors
}

func NewAnnotationQueryBuilder() *AnnotationQueryBuilder {
	resource := NewAnnotationQueryKind()
	builder := &AnnotationQueryBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}
	builder.internal.Kind = "AnnotationQuery"

	return builder
}

func (builder *AnnotationQueryBuilder) Build() (AnnotationQueryKind, error) {
	if err := builder.internal.Validate(); err != nil {
		return AnnotationQueryKind{}, err
	}

	if len(builder.errors) > 0 {
		return AnnotationQueryKind{}, cog.MakeBuildErrors("dashboardv2beta1.annotationQuery", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *AnnotationQueryBuilder) Spec(spec cog.Builder[AnnotationQuerySpec]) *AnnotationQueryBuilder {
	specResource, err := spec.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Spec = specResource

	return builder
}

func (builder *AnnotationQueryBuilder) Query(query cog.Builder[DataQueryKind]) *AnnotationQueryBuilder {
	queryResource, err := query.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Spec.Query = queryResource

	return builder
}

func (builder *AnnotationQueryBuilder) Enable(enable bool) *AnnotationQueryBuilder {
	builder.internal.Spec.Enable = enable

	return builder
}

func (builder *AnnotationQueryBuilder) Hide(hide bool) *AnnotationQueryBuilder {
	builder.internal.Spec.Hide = hide

	return builder
}

func (builder *AnnotationQueryBuilder) IconColor(iconColor string) *AnnotationQueryBuilder {
	builder.internal.Spec.IconColor = iconColor

	return builder
}

func (builder *AnnotationQueryBuilder) Name(name string) *AnnotationQueryBuilder {
	builder.internal.Spec.Name = name

	return builder
}

func (builder *AnnotationQueryBuilder) BuiltIn(builtIn bool) *AnnotationQueryBuilder {
	builder.internal.Spec.BuiltIn = &builtIn

	return builder
}

func (builder *AnnotationQueryBuilder) Filter(filter cog.Builder[AnnotationPanelFilter]) *AnnotationQueryBuilder {
	filterResource, err := filter.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Spec.Filter = &filterResource

	return builder
}

// Placement can be used to display the annotation query somewhere else on the dashboard other than the default location.
func (builder *AnnotationQueryBuilder) Placement(placement string) *AnnotationQueryBuilder {
	builder.internal.Spec.Placement = &placement

	return builder
}

// Mappings define how to convert data frame fields to annotation event fields.
func (builder *AnnotationQueryBuilder) Mappings(mappings map[string]cog.Builder[AnnotationEventFieldMapping]) *AnnotationQueryBuilder {
	mappingsResource := make(map[string]AnnotationEventFieldMapping)
	for key1, val1 := range mappings {
		mappingsDepth1, err := val1.Build()
		if err != nil {
			builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
			return builder
		}
		mappingsResource[key1] = mappingsDepth1
	}
	builder.internal.Spec.Mappings = mappingsResource

	return builder
}

// Catch-all field for datasource-specific properties. Should not be available in as code tooling.
func (builder *AnnotationQueryBuilder) LegacyOptions(legacyOptions map[string]any) *AnnotationQueryBuilder {
	builder.internal.Spec.LegacyOptions = legacyOptions

	return builder
}
