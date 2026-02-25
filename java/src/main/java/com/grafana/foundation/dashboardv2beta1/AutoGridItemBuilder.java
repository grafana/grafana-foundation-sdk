// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;


public class AutoGridItemBuilder implements com.grafana.foundation.cog.Builder<AutoGridLayoutItemKind> {
    protected final AutoGridLayoutItemKind internal;
    
    public AutoGridItemBuilder() {
        this.internal = new AutoGridLayoutItemKind();
        this.internal.kind = "AutoGridLayoutItem";
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.AutoGridLayoutItemSpec();
		}
		if (this.internal.spec.element == null) {
			this.internal.spec.element = new com.grafana.foundation.dashboardv2beta1.ElementReferenceBuilder().build();
		}
        this.internal.spec.element.kind = "ElementReference";
    }
    public AutoGridItemBuilder repeat(com.grafana.foundation.cog.Builder<AutoGridRepeatOptions> repeat) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.AutoGridLayoutItemSpec();
		}
    AutoGridRepeatOptions repeatResource = repeat.build();
        this.internal.spec.repeat = repeatResource;
        return this;
    }
    
    public AutoGridItemBuilder conditionalRendering(com.grafana.foundation.cog.Builder<ConditionalRenderingGroupKind> conditionalRendering) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.AutoGridLayoutItemSpec();
		}
    ConditionalRenderingGroupKind conditionalRenderingResource = conditionalRendering.build();
        this.internal.spec.conditionalRendering = conditionalRenderingResource;
        return this;
    }
    
    public AutoGridItemBuilder name(String name) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.AutoGridLayoutItemSpec();
		}
		if (this.internal.spec.element == null) {
			this.internal.spec.element = new com.grafana.foundation.dashboardv2beta1.ElementReferenceBuilder().build();
		}
        this.internal.spec.element.name = name;
        return this;
    }
    public AutoGridLayoutItemKind build() {
        return this.internal;
    }
}
