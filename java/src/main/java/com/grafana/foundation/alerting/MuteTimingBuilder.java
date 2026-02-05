// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.alerting;

import java.util.List;
import java.util.LinkedList;

public class MuteTimingBuilder implements com.grafana.foundation.cog.Builder<MuteTiming> {
    protected final MuteTiming internal;
    
    public MuteTimingBuilder() {
        this.internal = new MuteTiming();
    }
    public MuteTimingBuilder name(String name) {
        this.internal.name = name;
        return this;
    }
    
    public MuteTimingBuilder timeIntervals(List<com.grafana.foundation.cog.Builder<TimeInterval>> timeIntervals) {
        List<TimeInterval> timeIntervalsResources = new LinkedList<>();
        for (com.grafana.foundation.cog.Builder<TimeInterval> r1 : timeIntervals) {
                TimeInterval timeIntervalsDepth1 = r1.build();
                timeIntervalsResources.add(timeIntervalsDepth1); 
        }
        this.internal.timeIntervals = timeIntervalsResources;
        return this;
    }
    public MuteTiming build() {
        return this.internal;
    }
}
