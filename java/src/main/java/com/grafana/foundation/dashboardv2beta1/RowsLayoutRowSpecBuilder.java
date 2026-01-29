// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;


public class RowsLayoutRowSpecBuilder implements com.grafana.foundation.cog.Builder<RowsLayoutRowSpec> {
    protected final RowsLayoutRowSpec internal;
    
    public RowsLayoutRowSpecBuilder() {
        this.internal = new RowsLayoutRowSpec();
    }
    public RowsLayoutRowSpecBuilder title(String title) {
        this.internal.title = title;
        return this;
    }
    
    public RowsLayoutRowSpecBuilder collapse(Boolean collapse) {
        this.internal.collapse = collapse;
        return this;
    }
    
    public RowsLayoutRowSpecBuilder hideHeader(Boolean hideHeader) {
        this.internal.hideHeader = hideHeader;
        return this;
    }
    
    public RowsLayoutRowSpecBuilder fillScreen(Boolean fillScreen) {
        this.internal.fillScreen = fillScreen;
        return this;
    }
    
    public RowsLayoutRowSpecBuilder conditionalRendering(com.grafana.foundation.cog.Builder<ConditionalRenderingGroupKind> conditionalRendering) {
    ConditionalRenderingGroupKind conditionalRenderingResource = conditionalRendering.build();
        this.internal.conditionalRendering = conditionalRenderingResource;
        return this;
    }
    
    public RowsLayoutRowSpecBuilder repeat(com.grafana.foundation.cog.Builder<RowRepeatOptions> repeat) {
    RowRepeatOptions repeatResource = repeat.build();
        this.internal.repeat = repeatResource;
        return this;
    }
    
    public RowsLayoutRowSpecBuilder layout(GridLayoutKindOrAutoGridLayoutKindOrTabsLayoutKindOrRowsLayoutKind layout) {
        this.internal.layout = layout;
        return this;
    }
    public RowsLayoutRowSpec build() {
        return this.internal;
    }
}
