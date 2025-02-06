// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboard;

import java.util.List;

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
    
    public TimePickerBuilder timeOptions(List<String> timeOptions) {
        this.internal.timeOptions = timeOptions;
        return this;
    }
    
    public TimePickerBuilder quickRanges(com.grafana.foundation.cog.Builder<List<TimeOption>> quickRanges) {
        this.internal.quickRanges = quickRanges.build();
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
