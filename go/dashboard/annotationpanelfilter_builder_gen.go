// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboard

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[AnnotationPanelFilter] = (*AnnotationPanelFilterBuilder)(nil)

type AnnotationPanelFilterBuilder struct {
	internal *AnnotationPanelFilter
	errors   map[string]cog.BuildErrors
}

func NewAnnotationPanelFilterBuilder() *AnnotationPanelFilterBuilder {
	resource := &AnnotationPanelFilter{}
	builder := &AnnotationPanelFilterBuilder{
		internal: resource,
		errors:   make(map[string]cog.BuildErrors),
	}

	builder.applyDefaults()

	return builder
}

func (builder *AnnotationPanelFilterBuilder) Build() (AnnotationPanelFilter, error) {
	if err := builder.internal.Validate(); err != nil {
		return AnnotationPanelFilter{}, err
	}

	return *builder.internal, nil
}

// Should the specified panels be included or excluded
func (builder *AnnotationPanelFilterBuilder) Exclude(exclude bool) *AnnotationPanelFilterBuilder {
	builder.internal.Exclude = &exclude

	return builder
}

// Panel IDs that should be included or excluded
func (builder *AnnotationPanelFilterBuilder) Ids(ids []uint8) *AnnotationPanelFilterBuilder {
	builder.internal.Ids = ids

	return builder
}

func (builder *AnnotationPanelFilterBuilder) applyDefaults() {
	builder.Exclude(false)
}
