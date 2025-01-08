// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;


public class MapLayerOptionsBuilder implements com.grafana.foundation.cog.Builder<MapLayerOptions> {
    protected final MapLayerOptions internal;
    
    public MapLayerOptionsBuilder() {
        this.internal = new MapLayerOptions();
    }
    public MapLayerOptionsBuilder type(String type) {
    this.internal.type = type;
        return this;
    }
    
    public MapLayerOptionsBuilder name(String name) {
    this.internal.name = name;
        return this;
    }
    
    public MapLayerOptionsBuilder config(Object config) {
    this.internal.config = config;
        return this;
    }
    
    public MapLayerOptionsBuilder location(com.grafana.foundation.cog.Builder<FrameGeometrySource> location) {
    this.internal.location = location.build();
        return this;
    }
    
    public MapLayerOptionsBuilder filterData(Object filterData) {
    this.internal.filterData = filterData;
        return this;
    }
    
    public MapLayerOptionsBuilder opacity(Long opacity) {
    this.internal.opacity = opacity;
        return this;
    }
    
    public MapLayerOptionsBuilder tooltip(Boolean tooltip) {
    this.internal.tooltip = tooltip;
        return this;
    }
    public MapLayerOptions build() {
        return this.internal;
    }
}
