// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;


public class TabsLayoutTabBuilder implements com.grafana.foundation.cog.Builder<TabsLayoutTabKind> {
    protected final TabsLayoutTabKind internal;
    
    public TabsLayoutTabBuilder(String title) {
        this.internal = new TabsLayoutTabKind();
        this.internal.kind = "TabsLayoutTab";
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.TabsLayoutTabSpecBuilder().build();
		}
        this.internal.spec.title = title;
    }
    public TabsLayoutTabBuilder title(String title) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.TabsLayoutTabSpecBuilder().build();
		}
        this.internal.spec.title = title;
        return this;
    }
    
    public TabsLayoutTabBuilder gridLayout(com.grafana.foundation.cog.Builder<GridLayoutKind> gridLayoutKind) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.TabsLayoutTabSpecBuilder().build();
		}
    GridLayoutKind gridLayoutKindResource = gridLayoutKind.build();
    GridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind gridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind = new GridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind();
        gridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind.gridLayoutKind = gridLayoutKindResource;
        this.internal.spec.layout = gridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind;
        return this;
    }
    
    public TabsLayoutTabBuilder rowsLayout(com.grafana.foundation.cog.Builder<RowsLayoutKind> rowsLayoutKind) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.TabsLayoutTabSpecBuilder().build();
		}
    RowsLayoutKind rowsLayoutKindResource = rowsLayoutKind.build();
    GridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind gridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind = new GridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind();
        gridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind.rowsLayoutKind = rowsLayoutKindResource;
        this.internal.spec.layout = gridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind;
        return this;
    }
    
    public TabsLayoutTabBuilder autoGridLayout(com.grafana.foundation.cog.Builder<AutoGridLayoutKind> autoGridLayoutKind) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.TabsLayoutTabSpecBuilder().build();
		}
    AutoGridLayoutKind autoGridLayoutKindResource = autoGridLayoutKind.build();
    GridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind gridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind = new GridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind();
        gridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind.autoGridLayoutKind = autoGridLayoutKindResource;
        this.internal.spec.layout = gridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind;
        return this;
    }
    
    public TabsLayoutTabBuilder tabsLayout(com.grafana.foundation.cog.Builder<TabsLayoutKind> tabsLayoutKind) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.TabsLayoutTabSpecBuilder().build();
		}
    TabsLayoutKind tabsLayoutKindResource = tabsLayoutKind.build();
    GridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind gridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind = new GridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind();
        gridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind.tabsLayoutKind = tabsLayoutKindResource;
        this.internal.spec.layout = gridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind;
        return this;
    }
    
    public TabsLayoutTabBuilder conditionalRendering(com.grafana.foundation.cog.Builder<ConditionalRenderingGroupKind> conditionalRendering) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.TabsLayoutTabSpecBuilder().build();
		}
    ConditionalRenderingGroupKind conditionalRenderingResource = conditionalRendering.build();
        this.internal.spec.conditionalRendering = conditionalRenderingResource;
        return this;
    }
    
    public TabsLayoutTabBuilder repeat(com.grafana.foundation.cog.Builder<TabRepeatOptions> repeat) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.TabsLayoutTabSpecBuilder().build();
		}
    TabRepeatOptions repeatResource = repeat.build();
        this.internal.spec.repeat = repeatResource;
        return this;
    }
    public TabsLayoutTabKind build() {
        return this.internal;
    }
}
