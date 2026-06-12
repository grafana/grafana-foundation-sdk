// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboard;

import com.grafana.foundation.common.DataSourceRef;
import com.grafana.foundation.cog.variants.Dataquery;
import java.util.Map;
import java.util.HashMap;

/**
 * TODO docs
 * FROM: AnnotationQuery in grafana-data/src/types/annotations.ts
 */
public class AnnotationQueryBuilder implements com.grafana.foundation.cog.Builder<AnnotationQuery> {
    protected final AnnotationQuery internal;
    
    public AnnotationQueryBuilder() {
        this.internal = new AnnotationQuery();
    }
    public AnnotationQueryBuilder name(String name) {
        this.internal.name = name;
        return this;
    }
    
    public AnnotationQueryBuilder datasource(DataSourceRef datasource) {
        this.internal.datasource = datasource;
        return this;
    }
    
    public AnnotationQueryBuilder enable(Boolean enable) {
        this.internal.enable = enable;
        return this;
    }
    
    public AnnotationQueryBuilder hide(Boolean hide) {
        this.internal.hide = hide;
        return this;
    }
    
    public AnnotationQueryBuilder iconColor(String iconColor) {
        this.internal.iconColor = iconColor;
        return this;
    }
    
    public AnnotationQueryBuilder filter(com.grafana.foundation.cog.Builder<AnnotationPanelFilter> filter) {
    AnnotationPanelFilter filterResource = filter.build();
        this.internal.filter = filterResource;
        return this;
    }
    
    public AnnotationQueryBuilder target(com.grafana.foundation.cog.Builder<Dataquery> target) {
    Dataquery targetResource = target.build();
        this.internal.target = targetResource;
        return this;
    }
    
    public AnnotationQueryBuilder type(String type) {
        this.internal.type = type;
        return this;
    }
    
    public AnnotationQueryBuilder builtIn(Double builtIn) {
        this.internal.builtIn = builtIn;
        return this;
    }
    
    public AnnotationQueryBuilder placement(AnnotationQueryPlacement placement) {
        this.internal.placement = placement;
        return this;
    }
    
    public AnnotationQueryBuilder expr(String expr) {
        this.internal.expr = expr;
        return this;
    }
    
    public AnnotationQueryBuilder textFormat(String textFormat) {
        this.internal.textFormat = textFormat;
        return this;
    }
    
    public AnnotationQueryBuilder titleFormat(String titleFormat) {
        this.internal.titleFormat = titleFormat;
        return this;
    }
    
    public AnnotationQueryBuilder tagKeys(String tagKeys) {
        this.internal.tagKeys = tagKeys;
        return this;
    }
    
    public AnnotationQueryBuilder step(String step) {
        this.internal.step = step;
        return this;
    }
    
    public AnnotationQueryBuilder useValueForTime(Boolean useValueForTime) {
        this.internal.useValueForTime = useValueForTime;
        return this;
    }
    
    public AnnotationQueryBuilder mappings(Map<String, com.grafana.foundation.cog.Builder<AnnotationEventFieldMapping>> mappings) {
        Map<String, AnnotationEventFieldMapping> mappingsResources = new HashMap<>();
        for (var entry1 : mappings.entrySet()) {
                AnnotationEventFieldMapping mappingsDepth1 = entry1.getValue().build();
                mappingsResources.put(entry1.getKey(), mappingsDepth1);           
        }
        this.internal.mappings = mappingsResources;
        return this;
    }
    public AnnotationQuery build() {
        return this.internal;
    }
}
