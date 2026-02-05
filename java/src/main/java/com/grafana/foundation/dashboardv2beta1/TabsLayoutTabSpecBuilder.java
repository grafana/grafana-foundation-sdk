// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;


public class TabsLayoutTabSpecBuilder implements com.grafana.foundation.cog.Builder<TabsLayoutTabSpec> {
    protected final TabsLayoutTabSpec internal;
    
    public TabsLayoutTabSpecBuilder() {
        this.internal = new TabsLayoutTabSpec();
    }
    public TabsLayoutTabSpecBuilder title(String title) {
        this.internal.title = title;
        return this;
    }
    
    public TabsLayoutTabSpecBuilder layout(GridLayoutKindOrRowsLayoutKindOrAutoGridLayoutKindOrTabsLayoutKind layout) {
        this.internal.layout = layout;
        return this;
    }
    
    public TabsLayoutTabSpecBuilder conditionalRendering(com.grafana.foundation.cog.Builder<ConditionalRenderingGroupKind> conditionalRendering) {
    ConditionalRenderingGroupKind conditionalRenderingResource = conditionalRendering.build();
        this.internal.conditionalRendering = conditionalRenderingResource;
        return this;
    }
    
    public TabsLayoutTabSpecBuilder repeat(com.grafana.foundation.cog.Builder<TabRepeatOptions> repeat) {
    TabRepeatOptions repeatResource = repeat.build();
        this.internal.repeat = repeatResource;
        return this;
    }
    public TabsLayoutTabSpec build() {
        return this.internal;
    }
}
