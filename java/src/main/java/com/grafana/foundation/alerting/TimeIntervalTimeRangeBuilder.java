// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.alerting;


public class TimeIntervalTimeRangeBuilder implements com.grafana.foundation.cog.Builder<TimeIntervalTimeRange> {
    protected final TimeIntervalTimeRange internal;
    
    public TimeIntervalTimeRangeBuilder() {
        this.internal = new TimeIntervalTimeRange();
    }
    public TimeIntervalTimeRangeBuilder endTime(String endTime) {
        this.internal.endTime = endTime;
        return this;
    }
    
    public TimeIntervalTimeRangeBuilder startTime(String startTime) {
        this.internal.startTime = startTime;
        return this;
    }
    public TimeIntervalTimeRange build() {
        return this.internal;
    }
}
