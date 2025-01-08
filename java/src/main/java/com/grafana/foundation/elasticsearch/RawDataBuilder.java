// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;


public class RawDataBuilder implements com.grafana.foundation.cog.Builder<RawData> {
    protected final RawData internal;
    
    public RawDataBuilder() {
        this.internal = new RawData();
    this.internal.type = "raw_data";
    }
    public RawDataBuilder id(String id) {
    this.internal.id = id;
        return this;
    }
    
    public RawDataBuilder settings(com.grafana.foundation.cog.Builder<ElasticsearchRawDataSettings> settings) {
    this.internal.settings = settings.build();
        return this;
    }
    
    public RawDataBuilder hide(Boolean hide) {
    this.internal.hide = hide;
        return this;
    }
    public RawData build() {
        return this.internal;
    }
}
