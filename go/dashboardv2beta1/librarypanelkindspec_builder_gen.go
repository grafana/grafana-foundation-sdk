// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[LibraryPanelKindSpec] = (*LibraryPanelKindSpecBuilder)(nil)

type LibraryPanelKindSpecBuilder struct {
	internal *LibraryPanelKindSpec
	errors   cog.BuildErrors
}

func NewLibraryPanelKindSpecBuilder() *LibraryPanelKindSpecBuilder {
	resource := NewLibraryPanelKindSpec()
	builder := &LibraryPanelKindSpecBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *LibraryPanelKindSpecBuilder) Build() (LibraryPanelKindSpec, error) {
	if err := builder.internal.Validate(); err != nil {
		return LibraryPanelKindSpec{}, err
	}

	if len(builder.errors) > 0 {
		return LibraryPanelKindSpec{}, cog.MakeBuildErrors("dashboardv2beta1.libraryPanelKindSpec", builder.errors)
	}

	return *builder.internal, nil
}

// Panel ID for the library panel in the dashboard
func (builder *LibraryPanelKindSpecBuilder) Id(id float64) *LibraryPanelKindSpecBuilder {
	builder.internal.Id = id

	return builder
}

// Title for the library panel in the dashboard
func (builder *LibraryPanelKindSpecBuilder) Title(title string) *LibraryPanelKindSpecBuilder {
	builder.internal.Title = title

	return builder
}

func (builder *LibraryPanelKindSpecBuilder) LibraryPanel(libraryPanel cog.Builder[LibraryPanelRef]) *LibraryPanelKindSpecBuilder {
	libraryPanelResource, err := libraryPanel.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.LibraryPanel = libraryPanelResource

	return builder
}
