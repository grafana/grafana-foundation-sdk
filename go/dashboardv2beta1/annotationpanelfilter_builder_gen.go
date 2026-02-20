// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[AnnotationPanelFilter] = (*AnnotationPanelFilterBuilder)(nil)

type AnnotationPanelFilterBuilder struct {
	internal *AnnotationPanelFilter
	errors   cog.BuildErrors
}

func NewAnnotationPanelFilterBuilder() *AnnotationPanelFilterBuilder {
	resource := NewAnnotationPanelFilter()
	builder := &AnnotationPanelFilterBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *AnnotationPanelFilterBuilder) Build() (AnnotationPanelFilter, error) {
	if err := builder.internal.Validate(); err != nil {
		return AnnotationPanelFilter{}, err
	}

	if len(builder.errors) > 0 {
		return AnnotationPanelFilter{}, cog.MakeBuildErrors("dashboardv2beta1.annotationPanelFilter", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *AnnotationPanelFilterBuilder) RecordError(path string, err error) *AnnotationPanelFilterBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

// Should the specified panels be included or excluded
func (builder *AnnotationPanelFilterBuilder) Exclude(exclude bool) *AnnotationPanelFilterBuilder {
	builder.internal.Exclude = &exclude

	return builder
}

// Panel IDs that should be included or excluded
func (builder *AnnotationPanelFilterBuilder) Ids(ids []uint32) *AnnotationPanelFilterBuilder {
	builder.internal.Ids = ids

	return builder
}
