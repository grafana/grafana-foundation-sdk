// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboard

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[AnnotationActions] = (*AnnotationActionsBuilder)(nil)

// +k8s:deepcopy-gen=true
type AnnotationActionsBuilder struct {
	internal *AnnotationActions
	errors   map[string]cog.BuildErrors
}

func NewAnnotationActionsBuilder() *AnnotationActionsBuilder {
	resource := NewAnnotationActions()
	builder := &AnnotationActionsBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	return builder
}

func (builder *AnnotationActionsBuilder) Build() (AnnotationActions, error) {
	if err := builder.internal.Validate(); err != nil {
		return AnnotationActions{}, err
	}

	return *builder.internal, nil
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
