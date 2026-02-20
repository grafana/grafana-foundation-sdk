// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[PanelQueryKind] = (*TargetBuilder)(nil)

type TargetBuilder struct {
	internal *PanelQueryKind
	errors   cog.BuildErrors
}

func NewTargetBuilder() *TargetBuilder {
	resource := NewPanelQueryKind()
	builder := &TargetBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}
	builder.internal.Kind = "PanelQuery"

	return builder
}

func (builder *TargetBuilder) Build() (PanelQueryKind, error) {
	if err := builder.internal.Validate(); err != nil {
		return PanelQueryKind{}, err
	}

	if len(builder.errors) > 0 {
		return PanelQueryKind{}, cog.MakeBuildErrors("dashboardv2beta1.target", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *TargetBuilder) RecordError(path string, err error) *TargetBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *TargetBuilder) Query(query cog.Builder[DataQueryKind]) *TargetBuilder {
	queryResource, err := query.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Spec.Query = queryResource

	return builder
}

func (builder *TargetBuilder) RefId(refId string) *TargetBuilder {
	builder.internal.Spec.RefId = refId

	return builder
}

func (builder *TargetBuilder) Hidden(hidden bool) *TargetBuilder {
	builder.internal.Spec.Hidden = hidden

	return builder
}
