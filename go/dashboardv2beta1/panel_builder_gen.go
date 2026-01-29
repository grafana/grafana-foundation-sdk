// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[PanelKind] = (*PanelBuilder)(nil)

type PanelBuilder struct {
	internal *PanelKind
	errors   cog.BuildErrors
}

func NewPanelBuilder() *PanelBuilder {
	resource := NewPanelKind()
	builder := &PanelBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}
	builder.internal.Kind = "Panel"

	return builder
}

func (builder *PanelBuilder) Build() (PanelKind, error) {
	if err := builder.internal.Validate(); err != nil {
		return PanelKind{}, err
	}

	if len(builder.errors) > 0 {
		return PanelKind{}, cog.MakeBuildErrors("dashboardv2beta1.panel", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *PanelBuilder) Id(id float64) *PanelBuilder {
	builder.internal.Spec.Id = id

	return builder
}

func (builder *PanelBuilder) Title(title string) *PanelBuilder {
	builder.internal.Spec.Title = title

	return builder
}

func (builder *PanelBuilder) Description(description string) *PanelBuilder {
	builder.internal.Spec.Description = description

	return builder
}

func (builder *PanelBuilder) Links(links []cog.Builder[DataLink]) *PanelBuilder {
	linksResources := make([]DataLink, 0, len(links))
	for _, r1 := range links {
		linksDepth1, err := r1.Build()
		if err != nil {
			builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
			return builder
		}
		linksResources = append(linksResources, linksDepth1)
	}
	builder.internal.Spec.Links = linksResources

	return builder
}

func (builder *PanelBuilder) Data(data cog.Builder[QueryGroupKind]) *PanelBuilder {
	dataResource, err := data.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Spec.Data = dataResource

	return builder
}

func (builder *PanelBuilder) Visualization(vizConfig cog.Builder[VizConfigKind]) *PanelBuilder {
	vizConfigResource, err := vizConfig.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Spec.VizConfig = vizConfigResource

	return builder
}

func (builder *PanelBuilder) Transparent(transparent bool) *PanelBuilder {
	builder.internal.Spec.Transparent = &transparent

	return builder
}
