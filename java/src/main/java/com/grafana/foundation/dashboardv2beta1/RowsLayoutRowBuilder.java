// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;


public class RowsLayoutRowBuilder implements com.grafana.foundation.cog.Builder<RowsLayoutRowKind> {
    protected final RowsLayoutRowKind internal;
    
    public RowsLayoutRowBuilder(String title) {
        this.internal = new RowsLayoutRowKind();
        this.internal.kind = "RowsLayoutRow";
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.RowsLayoutRowSpecBuilder().build();
		}
        this.internal.spec.title = title;
    }
    public RowsLayoutRowBuilder title(String title) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.RowsLayoutRowSpecBuilder().build();
		}
        this.internal.spec.title = title;
        return this;
    }
    
    public RowsLayoutRowBuilder collapse(Boolean collapse) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.RowsLayoutRowSpecBuilder().build();
		}
        this.internal.spec.collapse = collapse;
        return this;
    }
    
    public RowsLayoutRowBuilder hideHeader(Boolean hideHeader) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.RowsLayoutRowSpecBuilder().build();
		}
        this.internal.spec.hideHeader = hideHeader;
        return this;
    }
    
    public RowsLayoutRowBuilder fillScreen(Boolean fillScreen) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.RowsLayoutRowSpecBuilder().build();
		}
        this.internal.spec.fillScreen = fillScreen;
        return this;
    }
    
    public RowsLayoutRowBuilder conditionalRendering(com.grafana.foundation.cog.Builder<ConditionalRenderingGroupKind> conditionalRendering) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.RowsLayoutRowSpecBuilder().build();
		}
    ConditionalRenderingGroupKind conditionalRenderingResource = conditionalRendering.build();
        this.internal.spec.conditionalRendering = conditionalRenderingResource;
        return this;
    }
    
    public RowsLayoutRowBuilder repeat(com.grafana.foundation.cog.Builder<RowRepeatOptions> repeat) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.RowsLayoutRowSpecBuilder().build();
		}
    RowRepeatOptions repeatResource = repeat.build();
        this.internal.spec.repeat = repeatResource;
        return this;
    }
    
    public RowsLayoutRowBuilder gridLayout(com.grafana.foundation.cog.Builder<GridLayoutKind> gridLayoutKind) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.RowsLayoutRowSpecBuilder().build();
		}
    GridLayoutKind gridLayoutKindResource = gridLayoutKind.build();
    GridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKind gridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKind = new GridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKind();
        gridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKind.gridLayoutKind = gridLayoutKindResource;
        this.internal.spec.layout = gridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKind;
        return this;
    }
    
    public RowsLayoutRowBuilder autoGridLayout(com.grafana.foundation.cog.Builder<AutoGridLayoutKind> autoGridLayoutKind) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.RowsLayoutRowSpecBuilder().build();
		}
    AutoGridLayoutKind autoGridLayoutKindResource = autoGridLayoutKind.build();
    GridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKind gridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKind = new GridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKind();
        gridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKind.autoGridLayoutKind = autoGridLayoutKindResource;
        this.internal.spec.layout = gridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKind;
        return this;
    }
    
    public RowsLayoutRowBuilder tabsLayout(com.grafana.foundation.cog.Builder<TabsLayoutKind> tabsLayoutKind) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.RowsLayoutRowSpecBuilder().build();
		}
    TabsLayoutKind tabsLayoutKindResource = tabsLayoutKind.build();
    GridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKind gridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKind = new GridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKind();
        gridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKind.tabsLayoutKind = tabsLayoutKindResource;
        this.internal.spec.layout = gridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKind;
        return this;
    }
    
    public RowsLayoutRowBuilder rowsLayout(com.grafana.foundation.cog.Builder<RowsLayoutKind> rowsLayoutKind) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.RowsLayoutRowSpecBuilder().build();
		}
    RowsLayoutKind rowsLayoutKindResource = rowsLayoutKind.build();
    GridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKind gridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKind = new GridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKind();
        gridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKind.rowsLayoutKind = rowsLayoutKindResource;
        this.internal.spec.layout = gridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKind;
        return this;
    }
    public RowsLayoutRowKind build() {
        return this.internal;
    }
}
