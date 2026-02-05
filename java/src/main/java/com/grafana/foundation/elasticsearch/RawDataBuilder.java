// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;


public class RawDataBuilder implements com.grafana.foundation.cog.Builder<RawData> {
    protected final RawData internal;
    
    public RawDataBuilder() {
        this.internal = new RawData();
    }
    public RawDataBuilder id(String id) {
        this.internal.id = id;
        return this;
    }
    
    public RawDataBuilder settings(com.grafana.foundation.cog.Builder<ElasticsearchRawDataSettings> settings) {
    ElasticsearchRawDataSettings settingsResource = settings.build();
        this.internal.settings = settingsResource;
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
