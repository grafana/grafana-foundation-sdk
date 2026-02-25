// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;


public class RowBuilder implements com.grafana.foundation.cog.Builder<RowsLayoutRowKind> {
    protected final RowsLayoutRowKind internal;
    
    public RowBuilder() {
        this.internal = new RowsLayoutRowKind();
        this.internal.kind = "RowsLayoutRow";
    }
    public RowBuilder title(String title) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.RowsLayoutRowSpec();
		}
        this.internal.spec.title = title;
        return this;
    }
    
    public RowBuilder collapse(Boolean collapse) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.RowsLayoutRowSpec();
		}
        this.internal.spec.collapse = collapse;
        return this;
    }
    
    public RowBuilder hideHeader(Boolean hideHeader) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.RowsLayoutRowSpec();
		}
        this.internal.spec.hideHeader = hideHeader;
        return this;
    }
    
    public RowBuilder fillScreen(Boolean fillScreen) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.RowsLayoutRowSpec();
		}
        this.internal.spec.fillScreen = fillScreen;
        return this;
    }
    
    public RowBuilder conditionalRendering(com.grafana.foundation.cog.Builder<ConditionalRenderingGroupKind> conditionalRendering) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.RowsLayoutRowSpec();
		}
    ConditionalRenderingGroupKind conditionalRenderingResource = conditionalRendering.build();
        this.internal.spec.conditionalRendering = conditionalRenderingResource;
        return this;
    }
    
    public RowBuilder repeat(com.grafana.foundation.cog.Builder<RowRepeatOptions> repeat) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.RowsLayoutRowSpec();
		}
    RowRepeatOptions repeatResource = repeat.build();
        this.internal.spec.repeat = repeatResource;
        return this;
    }
    
    public RowBuilder gridLayout(com.grafana.foundation.cog.Builder<GridLayoutKind> gridLayoutKind) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.RowsLayoutRowSpec();
		}
    GridLayoutKind gridLayoutKindResource = gridLayoutKind.build();
    GridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKind gridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKind = new GridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKind();
        gridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKind.gridLayoutKind = gridLayoutKindResource;
        this.internal.spec.layout = gridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKind;
        return this;
    }
    
    public RowBuilder autoGridLayout(com.grafana.foundation.cog.Builder<AutoGridLayoutKind> autoGridLayoutKind) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.RowsLayoutRowSpec();
		}
    AutoGridLayoutKind autoGridLayoutKindResource = autoGridLayoutKind.build();
    GridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKind gridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKind = new GridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKind();
        gridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKind.autoGridLayoutKind = autoGridLayoutKindResource;
        this.internal.spec.layout = gridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKind;
        return this;
    }
    
    public RowBuilder tabsLayout(com.grafana.foundation.cog.Builder<TabsLayoutKind> tabsLayoutKind) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.RowsLayoutRowSpec();
		}
    TabsLayoutKind tabsLayoutKindResource = tabsLayoutKind.build();
    GridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKind gridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKind = new GridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKind();
        gridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKind.tabsLayoutKind = tabsLayoutKindResource;
        this.internal.spec.layout = gridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKind;
        return this;
    }
    
    public RowBuilder rowsLayout(com.grafana.foundation.cog.Builder<RowsLayoutKind> rowsLayoutKind) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.RowsLayoutRowSpec();
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
