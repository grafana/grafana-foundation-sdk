// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboard;

import java.util.List;

public class TimePickerBuilder implements com.grafana.foundation.cog.Builder<TimePicker> {
    protected final TimePicker internal;
    
    public TimePickerBuilder() {
        this.internal = new TimePicker();
    }
    public TimePickerBuilder hidden(Boolean hidden) {
    this.internal.hidden = hidden;
        return this;
    }
    
    public TimePickerBuilder refreshIntervals(List<String> refreshIntervals) {
    this.internal.refreshIntervals = refreshIntervals;
        return this;
    }
    
    public TimePickerBuilder collapse(Boolean collapse) {
    this.internal.collapse = collapse;
        return this;
    }
    
    public TimePickerBuilder enable(Boolean enable) {
    this.internal.enable = enable;
        return this;
    }
    
    public TimePickerBuilder timeOptions(List<String> timeOptions) {
    this.internal.timeOptions = timeOptions;
        return this;
    }
    public TimePicker build() {
        return this.internal;
    }
}
