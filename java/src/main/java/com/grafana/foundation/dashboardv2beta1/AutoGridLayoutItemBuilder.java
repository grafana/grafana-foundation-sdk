// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;


public class AutoGridLayoutItemBuilder implements com.grafana.foundation.cog.Builder<AutoGridLayoutItemKind> {
    protected final AutoGridLayoutItemKind internal;
    
    public AutoGridLayoutItemBuilder(String name) {
        this.internal = new AutoGridLayoutItemKind();
        this.internal.kind = "AutoGridLayoutItem";
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.AutoGridLayoutItemSpecBuilder().build();
		}
		if (this.internal.spec.element == null) {
			this.internal.spec.element = new com.grafana.foundation.dashboardv2beta1.ElementReferenceBuilder().build();
		}
        this.internal.spec.element.kind = "ElementReference";
        this.internal.spec.element.name = name;
    }
    public AutoGridLayoutItemBuilder element(com.grafana.foundation.cog.Builder<ElementReference> element) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.AutoGridLayoutItemSpecBuilder().build();
		}
    ElementReference elementResource = element.build();
        this.internal.spec.element = elementResource;
        return this;
    }
    
    public AutoGridLayoutItemBuilder repeat(com.grafana.foundation.cog.Builder<AutoGridRepeatOptions> repeat) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.AutoGridLayoutItemSpecBuilder().build();
		}
    AutoGridRepeatOptions repeatResource = repeat.build();
        this.internal.spec.repeat = repeatResource;
        return this;
    }
    
    public AutoGridLayoutItemBuilder conditionalRendering(com.grafana.foundation.cog.Builder<ConditionalRenderingGroupKind> conditionalRendering) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.AutoGridLayoutItemSpecBuilder().build();
		}
    ConditionalRenderingGroupKind conditionalRenderingResource = conditionalRendering.build();
        this.internal.spec.conditionalRendering = conditionalRenderingResource;
        return this;
    }
    
    public AutoGridLayoutItemBuilder name(String name) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.AutoGridLayoutItemSpecBuilder().build();
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
