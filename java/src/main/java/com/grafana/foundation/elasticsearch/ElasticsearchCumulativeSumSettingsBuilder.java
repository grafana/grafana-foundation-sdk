// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;


public class ElasticsearchCumulativeSumSettingsBuilder implements com.grafana.foundation.cog.Builder<ElasticsearchCumulativeSumSettings> {
    protected final ElasticsearchCumulativeSumSettings internal;
    
    public ElasticsearchCumulativeSumSettingsBuilder() {
        this.internal = new ElasticsearchCumulativeSumSettings();
    }
    public ElasticsearchCumulativeSumSettingsBuilder format(String format) {
    this.internal.format = format;
        return this;
    }
    public ElasticsearchCumulativeSumSettings build() {
        return this.internal;
    }
}
