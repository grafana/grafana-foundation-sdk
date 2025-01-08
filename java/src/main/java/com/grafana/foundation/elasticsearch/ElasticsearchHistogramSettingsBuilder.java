// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;


public class ElasticsearchHistogramSettingsBuilder implements com.grafana.foundation.cog.Builder<ElasticsearchHistogramSettings> {
    protected final ElasticsearchHistogramSettings internal;
    
    public ElasticsearchHistogramSettingsBuilder() {
        this.internal = new ElasticsearchHistogramSettings();
    }
    public ElasticsearchHistogramSettingsBuilder interval(String interval) {
    this.internal.interval = interval;
        return this;
    }
    
    public ElasticsearchHistogramSettingsBuilder minDocCount(String minDocCount) {
    this.internal.minDocCount = minDocCount;
        return this;
    }
    public ElasticsearchHistogramSettings build() {
        return this.internal;
    }
}
