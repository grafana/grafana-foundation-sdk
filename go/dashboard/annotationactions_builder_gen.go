// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboard

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[AnnotationActions] = (*AnnotationActionsBuilder)(nil)

type AnnotationActionsBuilder struct {
	internal *AnnotationActions
	errors   cog.BuildErrors
}

func NewAnnotationActionsBuilder() *AnnotationActionsBuilder {
	resource := NewAnnotationActions()
	builder := &AnnotationActionsBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *AnnotationActionsBuilder) Build() (AnnotationActions, error) {
	if err := builder.internal.Validate(); err != nil {
		return AnnotationActions{}, err
	}

	if len(builder.errors) > 0 {
		return AnnotationActions{}, cog.MakeBuildErrors("dashboard.annotationActions", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *AnnotationActionsBuilder) RecordError(path string, err error) *AnnotationActionsBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *AnnotationActionsBuilder) CanAdd(canAdd bool) *AnnotationActionsBuilder {
	builder.internal.CanAdd = &canAdd

	return builder
}

func (builder *AnnotationActionsBuilder) CanDelete(canDelete bool) *AnnotationActionsBuilder {
	builder.internal.CanDelete = &canDelete

	return builder
}

func (builder *AnnotationActionsBuilder) CanEdit(canEdit bool) *AnnotationActionsBuilder {
	builder.internal.CanEdit = &canEdit

	return builder
}
