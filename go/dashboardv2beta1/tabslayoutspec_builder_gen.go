// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[TabsLayoutSpec] = (*TabsLayoutSpecBuilder)(nil)

type TabsLayoutSpecBuilder struct {
	internal *TabsLayoutSpec
	errors   cog.BuildErrors
}

func NewTabsLayoutSpecBuilder() *TabsLayoutSpecBuilder {
	resource := NewTabsLayoutSpec()
	builder := &TabsLayoutSpecBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}

	return builder
}

func (builder *TabsLayoutSpecBuilder) Build() (TabsLayoutSpec, error) {
	if err := builder.internal.Validate(); err != nil {
		return TabsLayoutSpec{}, err
	}

	if len(builder.errors) > 0 {
		return TabsLayoutSpec{}, cog.MakeBuildErrors("dashboardv2beta1.tabsLayoutSpec", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *TabsLayoutSpecBuilder) Tabs(tabs []cog.Builder[TabsLayoutTabKind]) *TabsLayoutSpecBuilder {
	tabsResources := make([]TabsLayoutTabKind, 0, len(tabs))
	for _, r1 := range tabs {
		tabsDepth1, err := r1.Build()
		if err != nil {
			builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
			return builder
		}
		tabsResources = append(tabsResources, tabsDepth1)
	}
	builder.internal.Tabs = tabsResources

	return builder
}
