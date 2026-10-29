// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package alertingv0alpha1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[PanelRef] = (*PanelRefBuilder)(nil)

type PanelRefBuilder struct {
	internal *PanelRef
	errors   cog.BuildErrors
}

func NewPanelRefBuilder() *PanelRefBuilder {
	resource := NewPanelRef()
	builder := &PanelRefBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *PanelRefBuilder) Build() (PanelRef, error) {
	if err := builder.internal.Validate(); err != nil {
		return PanelRef{}, err
	}

	if len(builder.errors) > 0 {
		return PanelRef{}, cog.MakeBuildErrors("alertingv0alpha1.panelRef", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *PanelRefBuilder) RecordError(path string, err error) *PanelRefBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *PanelRefBuilder) DashboardUID(dashboardUID string) *PanelRefBuilder {
	builder.internal.DashboardUID = dashboardUID

	return builder
}

func (builder *PanelRefBuilder) PanelID(panelID int64) *PanelRefBuilder {
	builder.internal.PanelID = panelID

	return builder
}
