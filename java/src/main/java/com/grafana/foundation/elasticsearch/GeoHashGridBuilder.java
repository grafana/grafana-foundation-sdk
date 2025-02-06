// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;


public class GeoHashGridBuilder implements com.grafana.foundation.cog.Builder<GeoHashGrid> {
    protected final GeoHashGrid internal;
    
    public GeoHashGridBuilder() {
        this.internal = new GeoHashGrid();
        this.internal.type = "geohash_grid";
    }
    public GeoHashGridBuilder field(String field) {
        this.internal.field = field;
        return this;
    }
    
    public GeoHashGridBuilder id(String id) {
        this.internal.id = id;
        return this;
    }
    
    public GeoHashGridBuilder settings(com.grafana.foundation.cog.Builder<ElasticsearchGeoHashGridSettings> settings) {
        this.internal.settings = settings.build();
        return this;
    }
    public GeoHashGrid build() {
        return this.internal;
    }
}
