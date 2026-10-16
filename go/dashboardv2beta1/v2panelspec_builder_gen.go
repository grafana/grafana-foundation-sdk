// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[V2PanelSpec] = (*V2PanelSpecBuilder)(nil)

type V2PanelSpecBuilder struct {
	internal *V2PanelSpec
	errors   cog.BuildErrors
}

func NewV2PanelSpecBuilder() *V2PanelSpecBuilder {
	resource := NewV2PanelSpec()
	builder := &V2PanelSpecBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *V2PanelSpecBuilder) Build() (V2PanelSpec, error) {
	if err := builder.internal.Validate(); err != nil {
		return V2PanelSpec{}, err
	}

	if len(builder.errors) > 0 {
		return V2PanelSpec{}, cog.MakeBuildErrors("dashboardv2beta1.v2PanelSpec", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *V2PanelSpecBuilder) RecordError(path string, err error) *V2PanelSpecBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *V2PanelSpecBuilder) Id(id float64) *V2PanelSpecBuilder {
	builder.internal.Id = id

	return builder
}

func (builder *V2PanelSpecBuilder) Title(title string) *V2PanelSpecBuilder {
	builder.internal.Title = title

	return builder
}

// Shown in a info icon tooltip next to panel title
func (builder *V2PanelSpecBuilder) Description(description string) *V2PanelSpecBuilder {
	builder.internal.Description = &description

	return builder
}

// Shown in a sub header below the title.
func (builder *V2PanelSpecBuilder) Subtitle(subtitle string) *V2PanelSpecBuilder {
	builder.internal.Subtitle = &subtitle

	return builder
}

func (builder *V2PanelSpecBuilder) Links(links []cog.Builder[DataLink]) *V2PanelSpecBuilder {
	linksResources := make([]DataLink, 0, len(links))
	for _, r1 := range links {
		linksDepth1, err := r1.Build()
		if err != nil {
			builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
			return builder
		}
		linksResources = append(linksResources, linksDepth1)
	}
	builder.internal.Links = linksResources

	return builder
}

func (builder *V2PanelSpecBuilder) Data(data cog.Builder[V2QueryGroupKind]) *V2PanelSpecBuilder {
	dataResource, err := data.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Data = dataResource

	return builder
}

func (builder *V2PanelSpecBuilder) VizConfig(vizConfig cog.Builder[VizConfigKind]) *V2PanelSpecBuilder {
	vizConfigResource, err := vizConfig.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.VizConfig = vizConfigResource

	return builder
}

func (builder *V2PanelSpecBuilder) Transparent(transparent bool) *V2PanelSpecBuilder {
	builder.internal.Transparent = &transparent

	return builder
}
