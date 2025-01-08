// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.geomap;


public class MapViewConfigBuilder implements com.grafana.foundation.cog.Builder<MapViewConfig> {
    protected final MapViewConfig internal;
    
    public MapViewConfigBuilder() {
        this.internal = new MapViewConfig();
    }
    public MapViewConfigBuilder id(String id) {
    this.internal.id = id;
        return this;
    }
    
    public MapViewConfigBuilder lat(Long lat) {
    this.internal.lat = lat;
        return this;
    }
    
    public MapViewConfigBuilder lon(Long lon) {
    this.internal.lon = lon;
        return this;
    }
    
    public MapViewConfigBuilder zoom(Long zoom) {
    this.internal.zoom = zoom;
        return this;
    }
    
    public MapViewConfigBuilder minZoom(Long minZoom) {
    this.internal.minZoom = minZoom;
        return this;
    }
    
    public MapViewConfigBuilder maxZoom(Long maxZoom) {
    this.internal.maxZoom = maxZoom;
        return this;
    }
    
    public MapViewConfigBuilder padding(Long padding) {
    this.internal.padding = padding;
        return this;
    }
    
    public MapViewConfigBuilder allLayers(Boolean allLayers) {
    this.internal.allLayers = allLayers;
        return this;
    }
    
    public MapViewConfigBuilder lastOnly(Boolean lastOnly) {
    this.internal.lastOnly = lastOnly;
        return this;
    }
    
    public MapViewConfigBuilder layer(String layer) {
    this.internal.layer = layer;
        return this;
    }
    
    public MapViewConfigBuilder shared(Boolean shared) {
    this.internal.shared = shared;
        return this;
    }
    public MapViewConfig build() {
        return this.internal;
    }
}
