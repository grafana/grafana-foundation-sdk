// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;


public class TimeRangeOptionBuilder implements com.grafana.foundation.cog.Builder<TimeRangeOption> {
    protected final TimeRangeOption internal;
    
    public TimeRangeOptionBuilder() {
        this.internal = new TimeRangeOption();
    }
    public TimeRangeOptionBuilder display(String display) {
        this.internal.display = display;
        return this;
    }
    
    public TimeRangeOptionBuilder from(String from) {
        this.internal.from = from;
        return this;
    }
    
    public TimeRangeOptionBuilder to(String to) {
        this.internal.to = to;
        return this;
    }
    public TimeRangeOption build() {
        return this.internal;
    }
}
