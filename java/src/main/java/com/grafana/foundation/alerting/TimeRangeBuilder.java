// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.alerting;


public class TimeRangeBuilder implements com.grafana.foundation.cog.Builder<TimeRange> {
    protected final TimeRange internal;
    
    public TimeRangeBuilder() {
        this.internal = new TimeRange();
    }
    public TimeRangeBuilder from(String from) {
        this.internal.from = from;
        return this;
    }
    
    public TimeRangeBuilder to(String to) {
        this.internal.to = to;
        return this;
    }
    public TimeRange build() {
        return this.internal;
    }
}
