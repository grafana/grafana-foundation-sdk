// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;

import java.util.Map;

public class AnnotationQuerySpecBuilder implements com.grafana.foundation.cog.Builder<AnnotationQuerySpec> {
    protected final AnnotationQuerySpec internal;
    
    public AnnotationQuerySpecBuilder() {
        this.internal = new AnnotationQuerySpec();
    }
    public AnnotationQuerySpecBuilder query(com.grafana.foundation.cog.Builder<DataQueryKind> query) {
    DataQueryKind queryResource = query.build();
        this.internal.query = queryResource;
        return this;
    }
    
    public AnnotationQuerySpecBuilder enable(Boolean enable) {
        this.internal.enable = enable;
        return this;
    }
    
    public AnnotationQuerySpecBuilder hide(Boolean hide) {
        this.internal.hide = hide;
        return this;
    }
    
    public AnnotationQuerySpecBuilder iconColor(String iconColor) {
        this.internal.iconColor = iconColor;
        return this;
    }
    
    public AnnotationQuerySpecBuilder name(String name) {
        this.internal.name = name;
        return this;
    }
    
    public AnnotationQuerySpecBuilder builtIn(Boolean builtIn) {
        this.internal.builtIn = builtIn;
        return this;
    }
    
    public AnnotationQuerySpecBuilder filter(com.grafana.foundation.cog.Builder<AnnotationPanelFilter> filter) {
    AnnotationPanelFilter filterResource = filter.build();
        this.internal.filter = filterResource;
        return this;
    }
    
    public AnnotationQuerySpecBuilder placement(String placement) {
        this.internal.placement = placement;
        return this;
    }
    
    public AnnotationQuerySpecBuilder legacyOptions(Map<String, Object> legacyOptions) {
        this.internal.legacyOptions = legacyOptions;
        return this;
    }
    public AnnotationQuerySpec build() {
        return this.internal;
    }
}
