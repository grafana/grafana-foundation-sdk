// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import java.util.Map;

public class AnnotationQueryBuilder implements com.grafana.foundation.cog.Builder<AnnotationQueryKind> {
    protected final AnnotationQueryKind internal;
    
    public AnnotationQueryBuilder() {
        this.internal = new AnnotationQueryKind();
        this.internal.kind = "AnnotationQuery";
    }
    public AnnotationQueryBuilder query(com.grafana.foundation.cog.Builder<DataQueryKind> query) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.AnnotationQuerySpec();
		}
    DataQueryKind queryResource = query.build();
        this.internal.spec.query = queryResource;
        return this;
    }
    
    public AnnotationQueryBuilder enable(Boolean enable) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.AnnotationQuerySpec();
		}
        this.internal.spec.enable = enable;
        return this;
    }
    
    public AnnotationQueryBuilder hide(Boolean hide) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.AnnotationQuerySpec();
		}
        this.internal.spec.hide = hide;
        return this;
    }
    
    public AnnotationQueryBuilder iconColor(String iconColor) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.AnnotationQuerySpec();
		}
        this.internal.spec.iconColor = iconColor;
        return this;
    }
    
    public AnnotationQueryBuilder name(String name) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.AnnotationQuerySpec();
		}
        this.internal.spec.name = name;
        return this;
    }
    
    public AnnotationQueryBuilder builtIn(Boolean builtIn) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.AnnotationQuerySpec();
		}
        this.internal.spec.builtIn = builtIn;
        return this;
    }
    
    public AnnotationQueryBuilder filter(com.grafana.foundation.cog.Builder<AnnotationPanelFilter> filter) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.AnnotationQuerySpec();
		}
    AnnotationPanelFilter filterResource = filter.build();
        this.internal.spec.filter = filterResource;
        return this;
    }
    
    public AnnotationQueryBuilder placement(String placement) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.AnnotationQuerySpec();
		}
        this.internal.spec.placement = placement;
        return this;
    }
    
    public AnnotationQueryBuilder legacyOptions(Map<String, Object> legacyOptions) {
		if (this.internal.spec == null) {
			this.internal.spec = new com.grafana.foundation.dashboardv2beta1.AnnotationQuerySpec();
		}
        this.internal.spec.legacyOptions = legacyOptions;
        return this;
    }
    public AnnotationQueryKind build() {
        return this.internal;
    }
}
