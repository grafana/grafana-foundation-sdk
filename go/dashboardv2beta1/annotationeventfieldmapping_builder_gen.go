// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[AnnotationEventFieldMapping] = (*AnnotationEventFieldMappingBuilder)(nil)

// Annotation event field mapping. Defines how to map a data frame field to an annotation event field.
type AnnotationEventFieldMappingBuilder struct {
	internal *AnnotationEventFieldMapping
	errors   cog.BuildErrors
}

func NewAnnotationEventFieldMappingBuilder() *AnnotationEventFieldMappingBuilder {
	resource := NewAnnotationEventFieldMapping()
	builder := &AnnotationEventFieldMappingBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *AnnotationEventFieldMappingBuilder) Build() (AnnotationEventFieldMapping, error) {
	if err := builder.internal.Validate(); err != nil {
		return AnnotationEventFieldMapping{}, err
	}

	if len(builder.errors) > 0 {
		return AnnotationEventFieldMapping{}, cog.MakeBuildErrors("dashboardv2beta1.annotationEventFieldMapping", builder.errors)
	}

	return *builder.internal, nil
}

// Source type for the field value
func (builder *AnnotationEventFieldMappingBuilder) Source(source string) *AnnotationEventFieldMappingBuilder {
	builder.internal.Source = &source

	return builder
}

// Constant value to use when source is "text"
func (builder *AnnotationEventFieldMappingBuilder) Value(value string) *AnnotationEventFieldMappingBuilder {
	builder.internal.Value = &value

	return builder
}

// Regular expression to apply to the field value
func (builder *AnnotationEventFieldMappingBuilder) Regex(regex string) *AnnotationEventFieldMappingBuilder {
	builder.internal.Regex = &regex

	return builder
}
