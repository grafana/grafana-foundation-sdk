// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.alerting;

import java.util.List;

public class MuteTimingBuilder implements com.grafana.foundation.cog.Builder<MuteTiming> {
    protected final MuteTiming internal;
    
    public MuteTimingBuilder() {
        this.internal = new MuteTiming();
    }
    public MuteTimingBuilder name(String name) {
    this.internal.name = name;
        return this;
    }
    
    public MuteTimingBuilder timeIntervals(com.grafana.foundation.cog.Builder<List<TimeInterval>> timeIntervals) {
    this.internal.timeIntervals = timeIntervals.build();
        return this;
    }
    public MuteTiming build() {
        return this.internal;
    }
}
