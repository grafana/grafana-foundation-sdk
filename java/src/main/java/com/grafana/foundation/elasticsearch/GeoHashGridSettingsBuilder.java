// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;


public class GeoHashGridSettingsBuilder implements com.grafana.foundation.cog.Builder<GeoHashGridSettings> {
    protected final GeoHashGridSettings internal;
    
    public GeoHashGridSettingsBuilder() {
        this.internal = new GeoHashGridSettings();
    }
    public GeoHashGridSettingsBuilder precision(String precision) {
        this.internal.precision = precision;
        return this;
    }
    public GeoHashGridSettings build() {
        return this.internal;
    }
}
