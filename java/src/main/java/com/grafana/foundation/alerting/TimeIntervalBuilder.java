// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.alerting;

import java.util.List;

public class TimeIntervalBuilder implements com.grafana.foundation.cog.Builder<TimeInterval> {
    protected final TimeInterval internal;
    
    public TimeIntervalBuilder() {
        this.internal = new TimeInterval();
    }
    public TimeIntervalBuilder name(String name) {
        this.internal.name = name;
        return this;
    }
    
    public TimeIntervalBuilder timeIntervals(com.grafana.foundation.cog.Builder<List<TimeIntervalItem>> timeIntervals) {
        this.internal.timeIntervals = timeIntervals.build();
        return this;
    }
    public TimeInterval build() {
        return this.internal;
    }
}
