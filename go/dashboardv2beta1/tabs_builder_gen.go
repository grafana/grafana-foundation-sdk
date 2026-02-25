// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package dashboardv2beta1

import (
	cog "github.com/grafana/grafana-foundation-sdk/go/cog"
)

var _ cog.Builder[TabsLayoutKind] = (*TabsBuilder)(nil)

type TabsBuilder struct {
	internal *TabsLayoutKind
	errors   cog.BuildErrors
}

func NewTabsBuilder() *TabsBuilder {
	resource := NewTabsLayoutKind()
	builder := &TabsBuilder{
		internal: resource,
		errors:   make(cog.BuildErrors, 0),
	}
	builder.internal.Kind = "TabsLayout"

	return builder
}

func Tabs() *TabsBuilder {
	builder := NewTabsBuilder()

	return builder
}

func (builder *TabsBuilder) Build() (TabsLayoutKind, error) {
	if err := builder.internal.Validate(); err != nil {
		return TabsLayoutKind{}, err
	}

	if len(builder.errors) > 0 {
		return TabsLayoutKind{}, cog.MakeBuildErrors("dashboardv2beta1.tabs", builder.errors)
	}

	return *builder.internal, nil
}

func (builder *TabsBuilder) RecordError(path string, err error) *TabsBuilder {
	builder.errors = append(builder.errors, cog.MakeBuildErrors(path, err)...)
	return builder
}

func (builder *TabsBuilder) Tabs(tabs []cog.Builder[TabsLayoutTabKind]) *TabsBuilder {
	tabsResources := make([]TabsLayoutTabKind, 0, len(tabs))
	for _, r1 := range tabs {
		tabsDepth1, err := r1.Build()
		if err != nil {
			builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
			return builder
		}
		tabsResources = append(tabsResources, tabsDepth1)
	}
	builder.internal.Spec.Tabs = tabsResources

	return builder
}

func (builder *TabsBuilder) Tab(tab cog.Builder[TabsLayoutTabKind]) *TabsBuilder {
	tabResource, err := tab.Build()
	if err != nil {
		builder.errors = append(builder.errors, err.(cog.BuildErrors)...)
		return builder
	}
	builder.internal.Spec.Tabs = append(builder.internal.Spec.Tabs, tabResource)

	return builder
}
