// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;


public class HistogramSettingsBuilder implements com.grafana.foundation.cog.Builder<HistogramSettings> {
    protected final HistogramSettings internal;
    
    public HistogramSettingsBuilder() {
        this.internal = new HistogramSettings();
    }
    public HistogramSettingsBuilder interval(String interval) {
    this.internal.interval = interval;
        return this;
    }
    
    public HistogramSettingsBuilder minDocCount(String minDocCount) {
    this.internal.minDocCount = minDocCount;
        return this;
    }
    public HistogramSettings build() {
        return this.internal;
    }
}
