// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;


public class ElasticsearchGeoHashGridSettingsBuilder implements com.grafana.foundation.cog.Builder<ElasticsearchGeoHashGridSettings> {
    protected final ElasticsearchGeoHashGridSettings internal;
    
    public ElasticsearchGeoHashGridSettingsBuilder() {
        this.internal = new ElasticsearchGeoHashGridSettings();
    }
    public ElasticsearchGeoHashGridSettingsBuilder precision(String precision) {
        this.internal.precision = precision;
        return this;
    }
    public ElasticsearchGeoHashGridSettings build() {
        return this.internal;
    }
}
