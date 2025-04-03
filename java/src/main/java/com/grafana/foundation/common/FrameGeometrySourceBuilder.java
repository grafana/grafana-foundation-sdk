// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.common;


public class FrameGeometrySourceBuilder implements com.grafana.foundation.cog.Builder<FrameGeometrySource> {
    protected final FrameGeometrySource internal;
    
    public FrameGeometrySourceBuilder() {
        this.internal = new FrameGeometrySource();
    }
    public FrameGeometrySourceBuilder mode(FrameGeometrySourceMode mode) {
        this.internal.mode = mode;
        return this;
    }
    
    public FrameGeometrySourceBuilder geohash(String geohash) {
        this.internal.geohash = geohash;
        return this;
    }
    
    public FrameGeometrySourceBuilder latitude(String latitude) {
        this.internal.latitude = latitude;
        return this;
    }
    
    public FrameGeometrySourceBuilder longitude(String longitude) {
        this.internal.longitude = longitude;
        return this;
    }
    
    public FrameGeometrySourceBuilder wkt(String wkt) {
        this.internal.wkt = wkt;
        return this;
    }
    
    public FrameGeometrySourceBuilder lookup(String lookup) {
        this.internal.lookup = lookup;
        return this;
    }
    
    public FrameGeometrySourceBuilder gazetteer(String gazetteer) {
        this.internal.gazetteer = gazetteer;
        return this;
    }
    public FrameGeometrySource build() {
        return this.internal;
    }
}
