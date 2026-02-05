// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[LibraryPanelRef] = (*LibraryPanelRefBuilder)(nil)

// A library panel is a reusable panel that you can use in any dashboard.
// When you make a change to a library panel, that change propagates to all instances of where the panel is used.
// Library panels streamline reuse of panels across multiple dashboards.
type LibraryPanelRefBuilder struct {
	internal *LibraryPanelRef
	errors   cog.BuildErrors
}

func NewLibraryPanelRefBuilder() *LibraryPanelRefBuilder {
	resource := NewLibraryPanelRef()
	builder := &LibraryPanelRefBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *LibraryPanelRefBuilder) Build() (LibraryPanelRef, error) {
	if err := builder.internal.Validate(); err != nil {
		return LibraryPanelRef{}, err
	}

	if len(builder.errors) > 0 {
		return LibraryPanelRef{}, cog.MakeBuildErrors("dashboardv2beta1.libraryPanelRef", builder.errors)
	}

	return *builder.internal, nil
}

// Library panel name
func (builder *LibraryPanelRefBuilder) Name(name string) *LibraryPanelRefBuilder {
	builder.internal.Name = name

	return builder
}

// Library panel uid
func (builder *LibraryPanelRefBuilder) Uid(uid string) *LibraryPanelRefBuilder {
	builder.internal.Uid = uid

	return builder
}
