// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[CellKind] = (*CellKindBuilder)(nil)

// A cell holds non-panel narrative content (markdown text, code) in a notebook layout.
// Panel cells are not represented here — they reuse V2PanelKind.
type CellKindBuilder struct {
	internal *CellKind
	errors   cog.BuildErrors
}

func NewCellKindBuilder() *CellKindBuilder {
	resource := NewCellKind()
	builder := &CellKindBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}
	builder.internal.Kind = "Cell"

	return builder
}

func (builder *CellKindBuilder) Build() (CellKind, error) {
	if err := builder.internal.Validate(); err != nil {
		return CellKind{}, err
	}

	if len(builder.errors) > 0 {
		return CellKind{}, cog.MakeBuildErrors("dashboardv2beta1.cellKind", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *CellKindBuilder) RecordError(path string, err error) *CellKindBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *CellKindBuilder) Spec(spec cog.Builder[CellSpec]) *CellKindBuilder {
	specResource, err := spec.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Spec = specResource

	return builder
}
