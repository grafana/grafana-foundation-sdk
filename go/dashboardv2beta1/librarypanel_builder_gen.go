// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[LibraryPanelKind] = (*LibraryPanelBuilder)(nil)

type LibraryPanelBuilder struct {
	internal *LibraryPanelKind
	errors   cog.BuildErrors
}

func NewLibraryPanelBuilder() *LibraryPanelBuilder {
	resource := NewLibraryPanelKind()
	builder := &LibraryPanelBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}
	builder.internal.Kind = "LibraryPanel"

	return builder
}

func (builder *LibraryPanelBuilder) Build() (LibraryPanelKind, error) {
	if err := builder.internal.Validate(); err != nil {
		return LibraryPanelKind{}, err
	}

	if len(builder.errors) > 0 {
		return LibraryPanelKind{}, cog.MakeBuildErrors("dashboardv2beta1.libraryPanel", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *LibraryPanelBuilder) Spec(spec cog.Builder[LibraryPanelKindSpec]) *LibraryPanelBuilder {
	specResource, err := spec.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Spec = specResource

	return builder
}

// Panel ID for the library panel in the dashboard
func (builder *LibraryPanelBuilder) Id(id float64) *LibraryPanelBuilder {
	builder.internal.Spec.Id = id

	return builder
}

// Title for the library panel in the dashboard
func (builder *LibraryPanelBuilder) Title(title string) *LibraryPanelBuilder {
	builder.internal.Spec.Title = title

	return builder
}

func (builder *LibraryPanelBuilder) LibraryPanel(libraryPanel cog.Builder[LibraryPanelRef]) *LibraryPanelBuilder {
	libraryPanelResource, err := libraryPanel.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Spec.LibraryPanel = libraryPanelResource

	return builder
}
