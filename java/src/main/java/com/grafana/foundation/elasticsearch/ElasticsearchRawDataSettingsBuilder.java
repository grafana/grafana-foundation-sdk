// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;


public class ElasticsearchRawDataSettingsBuilder implements com.grafana.foundation.cog.Builder<ElasticsearchRawDataSettings> {
    protected final ElasticsearchRawDataSettings internal;
    
    public ElasticsearchRawDataSettingsBuilder() {
        this.internal = new ElasticsearchRawDataSettings();
    }
    public ElasticsearchRawDataSettingsBuilder size(String size) {
        this.internal.size = size;
        return this;
    }
    public ElasticsearchRawDataSettings build() {
        return this.internal;
    }
}
