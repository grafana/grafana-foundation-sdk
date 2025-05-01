// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboard;

import java.util.List;
import java.util.LinkedList;

public class TimePickerBuilder implements com.grafana.foundation.cog.Builder<TimePickerConfig> {
    protected final TimePickerConfig internal;
    
    public TimePickerBuilder() {
        this.internal = new TimePickerConfig();
    }
    public TimePickerBuilder hidden(Boolean hidden) {
        this.internal.hidden = hidden;
        return this;
    }
    
    public TimePickerBuilder refreshIntervals(List<String> refreshIntervals) {
        this.internal.refreshIntervals = refreshIntervals;
        return this;
    }
    
    public TimePickerBuilder quickRanges(List<com.grafana.foundation.cog.Builder<TimeOption>> quickRanges) {
        List<TimeOption> quickRangesResources = new LinkedList<>();
        for (com.grafana.foundation.cog.Builder<TimeOption> r1 : quickRanges) {
                TimeOption quickRangesDepth1 = r1.build();
                quickRangesResources.add(quickRangesDepth1); 
        }
        this.internal.quickRanges = quickRangesResources;
        return this;
    }
    
    public TimePickerBuilder nowDelay(String nowDelay) {
        this.internal.nowDelay = nowDelay;
        return this;
    }
    public TimePickerConfig build() {
        return this.internal;
    }
}
