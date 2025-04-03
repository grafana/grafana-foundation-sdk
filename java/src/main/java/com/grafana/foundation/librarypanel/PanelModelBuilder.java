// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.librarypanel;

import java.util.List;
import com.grafana.foundation.cog.variants.Dataquery;
import com.grafana.foundation.dashboard.DataSourceRef;
import com.grafana.foundation.dashboard.DashboardLink;
import com.grafana.foundation.dashboard.DataTransformerConfig;
import com.grafana.foundation.dashboard.FieldConfigSource;

public class PanelModelBuilder implements com.grafana.foundation.cog.Builder<PanelModel> {
    protected final PanelModel internal;
    
    public PanelModelBuilder() {
        this.internal = new PanelModel();
    }
    public PanelModelBuilder type(String type) {
        if (!(type.length() >= 1)) {
            throw new IllegalArgumentException("type.length() must be >= 1");
        }
        this.internal.type = type;
        return this;
    }
    
    public PanelModelBuilder pluginVersion(String pluginVersion) {
        this.internal.pluginVersion = pluginVersion;
        return this;
    }
    
    public PanelModelBuilder targets(com.grafana.foundation.cog.Builder<List<Dataquery>> targets) {
        this.internal.targets = targets.build();
        return this;
    }
    
    public PanelModelBuilder title(String title) {
        this.internal.title = title;
        return this;
    }
    
    public PanelModelBuilder description(String description) {
        this.internal.description = description;
        return this;
    }
    
    public PanelModelBuilder transparent(Boolean transparent) {
        this.internal.transparent = transparent;
        return this;
    }
    
    public PanelModelBuilder datasource(DataSourceRef datasource) {
        this.internal.datasource = datasource;
        return this;
    }
    
    public PanelModelBuilder links(com.grafana.foundation.cog.Builder<List<DashboardLink>> links) {
        this.internal.links = links.build();
        return this;
    }
    
    public PanelModelBuilder repeat(String repeat) {
        this.internal.repeat = repeat;
        return this;
    }
    
    public PanelModelBuilder repeatDirection(PanelModelRepeatDirection repeatDirection) {
        this.internal.repeatDirection = repeatDirection;
        return this;
    }
    
    public PanelModelBuilder maxPerRow(Double maxPerRow) {
        this.internal.maxPerRow = maxPerRow;
        return this;
    }
    
    public PanelModelBuilder maxDataPoints(Double maxDataPoints) {
        this.internal.maxDataPoints = maxDataPoints;
        return this;
    }
    
    public PanelModelBuilder transformations(List<DataTransformerConfig> transformations) {
        this.internal.transformations = transformations;
        return this;
    }
    
    public PanelModelBuilder interval(String interval) {
        this.internal.interval = interval;
        return this;
    }
    
    public PanelModelBuilder timeFrom(String timeFrom) {
        this.internal.timeFrom = timeFrom;
        return this;
    }
    
    public PanelModelBuilder timeShift(String timeShift) {
        this.internal.timeShift = timeShift;
        return this;
    }
    
    public PanelModelBuilder hideTimeOverride(Boolean hideTimeOverride) {
        this.internal.hideTimeOverride = hideTimeOverride;
        return this;
    }
    
    public PanelModelBuilder cacheTimeout(String cacheTimeout) {
        this.internal.cacheTimeout = cacheTimeout;
        return this;
    }
    
    public PanelModelBuilder queryCachingTTL(Double queryCachingTTL) {
        this.internal.queryCachingTTL = queryCachingTTL;
        return this;
    }
    
    public PanelModelBuilder options(Object options) {
        this.internal.options = options;
        return this;
    }
    
    public PanelModelBuilder fieldConfig(FieldConfigSource fieldConfig) {
        this.internal.fieldConfig = fieldConfig;
        return this;
    }
    public PanelModel build() {
        return this.internal;
    }
}
