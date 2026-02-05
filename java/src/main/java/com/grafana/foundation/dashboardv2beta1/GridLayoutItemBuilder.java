// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;


public class GridLayoutItemBuilder implements com.grafana.foundation.cog.Builder<GridLayoutItemKind> {
    protected final GridLayoutItemKind internal;
    
    public GridLayoutItemBuilder(String name) {
        this.internal = new GridLayoutItemKind();
        this.internal.kind = "GridLayoutItem";
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.GridLayoutItemSpec();
		}
		if (this.internal.spec.element == null) {
			this.internal.spec.element = new com.grafana.foundation.dashboardv2beta1.ElementReferenceBuilder().build();
		}
        this.internal.spec.element.kind = "ElementReference";
        this.internal.spec.element.name = name;
    }
    public GridLayoutItemBuilder x(Long x) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.GridLayoutItemSpec();
		}
        this.internal.spec.x = x;
        return this;
    }
    
    public GridLayoutItemBuilder y(Long y) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.GridLayoutItemSpec();
		}
        this.internal.spec.y = y;
        return this;
    }
    
    public GridLayoutItemBuilder width(Long width) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.GridLayoutItemSpec();
		}
        this.internal.spec.width = width;
        return this;
    }
    
    public GridLayoutItemBuilder height(Long height) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.GridLayoutItemSpec();
		}
        this.internal.spec.height = height;
        return this;
    }
    
    public GridLayoutItemBuilder repeat(com.grafana.foundation.cog.Builder<RepeatOptions> repeat) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.GridLayoutItemSpec();
		}
    RepeatOptions repeatResource = repeat.build();
        this.internal.spec.repeat = repeatResource;
        return this;
    }
    
    public GridLayoutItemBuilder element(String name) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.GridLayoutItemSpec();
		}
		if (this.internal.spec.element == null) {
			this.internal.spec.element = new com.grafana.foundation.dashboardv2beta1.ElementReferenceBuilder().build();
		}
        this.internal.spec.element.name = name;
        return this;
    }
    public GridLayoutItemKind build() {
        return this.internal;
    }
}
