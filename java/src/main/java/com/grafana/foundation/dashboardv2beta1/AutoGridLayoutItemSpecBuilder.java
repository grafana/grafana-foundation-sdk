// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;


public class AutoGridLayoutItemSpecBuilder implements com.grafana.foundation.cog.Builder<AutoGridLayoutItemSpec> {
    protected final AutoGridLayoutItemSpec internal;
    
    public AutoGridLayoutItemSpecBuilder() {
        this.internal = new AutoGridLayoutItemSpec();
    }
    public AutoGridLayoutItemSpecBuilder element(com.grafana.foundation.cog.Builder<ElementReference> element) {
    ElementReference elementResource = element.build();
        this.internal.element = elementResource;
        return this;
    }
    
    public AutoGridLayoutItemSpecBuilder repeat(com.grafana.foundation.cog.Builder<AutoGridRepeatOptions> repeat) {
    AutoGridRepeatOptions repeatResource = repeat.build();
        this.internal.repeat = repeatResource;
        return this;
    }
    
    public AutoGridLayoutItemSpecBuilder conditionalRendering(com.grafana.foundation.cog.Builder<ConditionalRenderingGroupKind> conditionalRendering) {
    ConditionalRenderingGroupKind conditionalRenderingResource = conditionalRendering.build();
        this.internal.conditionalRendering = conditionalRenderingResource;
        return this;
    }
    public AutoGridLayoutItemSpec build() {
        return this.internal;
    }
}
