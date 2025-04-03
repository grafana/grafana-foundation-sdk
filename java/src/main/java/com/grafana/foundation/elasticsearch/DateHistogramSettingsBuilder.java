// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.elasticsearch;


public class DateHistogramSettingsBuilder implements com.grafana.foundation.cog.Builder<DateHistogramSettings> {
    protected final DateHistogramSettings internal;
    
    public DateHistogramSettingsBuilder() {
        this.internal = new DateHistogramSettings();
    }
    public DateHistogramSettingsBuilder interval(String interval) {
        this.internal.interval = interval;
        return this;
    }
    
    public DateHistogramSettingsBuilder minDocCount(String minDocCount) {
        this.internal.minDocCount = minDocCount;
        return this;
    }
    
    public DateHistogramSettingsBuilder trimEdges(String trimEdges) {
        this.internal.trimEdges = trimEdges;
        return this;
    }
    
    public DateHistogramSettingsBuilder offset(String offset) {
        this.internal.offset = offset;
        return this;
    }
    
    public DateHistogramSettingsBuilder timeZone(String timeZone) {
        this.internal.timeZone = timeZone;
        return this;
    }
    public DateHistogramSettings build() {
        return this.internal;
    }
}
