// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;


public class ElasticsearchDateHistogramSettingsBuilder implements com.grafana.foundation.cog.Builder<ElasticsearchDateHistogramSettings> {
    protected final ElasticsearchDateHistogramSettings internal;
    
    public ElasticsearchDateHistogramSettingsBuilder() {
        this.internal = new ElasticsearchDateHistogramSettings();
    }
    public ElasticsearchDateHistogramSettingsBuilder interval(String interval) {
    this.internal.interval = interval;
        return this;
    }
    
    public ElasticsearchDateHistogramSettingsBuilder minDocCount(String minDocCount) {
    this.internal.minDocCount = minDocCount;
        return this;
    }
    
    public ElasticsearchDateHistogramSettingsBuilder trimEdges(String trimEdges) {
    this.internal.trimEdges = trimEdges;
        return this;
    }
    
    public ElasticsearchDateHistogramSettingsBuilder offset(String offset) {
    this.internal.offset = offset;
        return this;
    }
    
    public ElasticsearchDateHistogramSettingsBuilder timeZone(String timeZone) {
    this.internal.timeZone = timeZone;
        return this;
    }
    public ElasticsearchDateHistogramSettings build() {
        return this.internal;
    }
}
