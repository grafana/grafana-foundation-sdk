// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;


public class TabBuilder implements com.grafana.foundation.cog.Builder<TabsLayoutTabKind> {
    protected final TabsLayoutTabKind internal;
    
    public TabBuilder() {
        this.internal = new TabsLayoutTabKind();
        this.internal.kind = "TabsLayoutTab";
    }
    public TabBuilder title(String title) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.TabsLayoutTabSpec();
		}
        this.internal.spec.title = title;
        return this;
    }
    
    public TabBuilder gridLayout(com.grafana.foundation.cog.Builder<GridLayoutKind> gridLayoutKind) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.TabsLayoutTabSpec();
		}
    GridLayoutKind gridLayoutKindResource = gridLayoutKind.build();
    GridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind gridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind = new GridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind();
        gridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind.gridLayoutKind = gridLayoutKindResource;
        this.internal.spec.layout = gridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind;
        return this;
    }
    
    public TabBuilder rowsLayout(com.grafana.foundation.cog.Builder<RowsLayoutKind> rowsLayoutKind) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.TabsLayoutTabSpec();
		}
    RowsLayoutKind rowsLayoutKindResource = rowsLayoutKind.build();
    GridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind gridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind = new GridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind();
        gridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind.rowsLayoutKind = rowsLayoutKindResource;
        this.internal.spec.layout = gridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind;
        return this;
    }
    
    public TabBuilder autoGridLayout(com.grafana.foundation.cog.Builder<AutoGridLayoutKind> autoGridLayoutKind) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.TabsLayoutTabSpec();
		}
    AutoGridLayoutKind autoGridLayoutKindResource = autoGridLayoutKind.build();
    GridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind gridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind = new GridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind();
        gridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind.autoGridLayoutKind = autoGridLayoutKindResource;
        this.internal.spec.layout = gridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind;
        return this;
    }
    
    public TabBuilder tabsLayout(com.grafana.foundation.cog.Builder<TabsLayoutKind> tabsLayoutKind) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.TabsLayoutTabSpec();
		}
    TabsLayoutKind tabsLayoutKindResource = tabsLayoutKind.build();
    GridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind gridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind = new GridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind();
        gridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind.tabsLayoutKind = tabsLayoutKindResource;
        this.internal.spec.layout = gridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind;
        return this;
    }
    
    public TabBuilder conditionalRendering(com.grafana.foundation.cog.Builder<ConditionalRenderingGroupKind> conditionalRendering) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.TabsLayoutTabSpec();
		}
    ConditionalRenderingGroupKind conditionalRenderingResource = conditionalRendering.build();
        this.internal.spec.conditionalRendering = conditionalRenderingResource;
        return this;
    }
    
    public TabBuilder repeat(com.grafana.foundation.cog.Builder<TabRepeatOptions> repeat) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.TabsLayoutTabSpec();
		}
    TabRepeatOptions repeatResource = repeat.build();
        this.internal.spec.repeat = repeatResource;
        return this;
    }
    public TabsLayoutTabKind build() {
        return this.internal;
    }
}
