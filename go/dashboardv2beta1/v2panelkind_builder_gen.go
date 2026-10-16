// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[V2PanelKind] = (*V2PanelKindBuilder)(nil)

// The notebook's own panel chain. It is a copy of the dashboard one down to the transformation,
// which follows the dashboard v2 shape rather than the v2beta1 shape in this package. The chain has
// to be forked rather than shared because PanelKind reaches TransformationKind through
// QueryGroupKind, and those three are what Dashboard v2beta1 serves. Everything the chain does not
// change (DataLink, VizConfigKind, PanelQueryKind, QueryOptionsSpec) stays shared.
type V2PanelKindBuilder struct {
	internal *V2PanelKind
	errors   cog.BuildErrors
}

func NewV2PanelKindBuilder() *V2PanelKindBuilder {
	resource := NewV2PanelKind()
	builder := &V2PanelKindBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}
	builder.internal.Kind = "Panel"

	return builder
}

func (builder *V2PanelKindBuilder) Build() (V2PanelKind, error) {
	if err := builder.internal.Validate(); err != nil {
		return V2PanelKind{}, err
	}

	if len(builder.errors) > 0 {
		return V2PanelKind{}, cog.MakeBuildErrors("dashboardv2beta1.v2PanelKind", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *V2PanelKindBuilder) RecordError(path string, err error) *V2PanelKindBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *V2PanelKindBuilder) Spec(spec cog.Builder[V2PanelSpec]) *V2PanelKindBuilder {
	specResource, err := spec.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Spec = specResource

	return builder
}
