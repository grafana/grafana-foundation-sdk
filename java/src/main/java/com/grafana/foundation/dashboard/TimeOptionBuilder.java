// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboard;


public class TimeOptionBuilder implements com.grafana.foundation.cog.Builder<TimeOption> {
    protected final TimeOption internal;
    
    public TimeOptionBuilder() {
        this.internal = new TimeOption();
    }
    public TimeOptionBuilder display(String display) {
        this.internal.display = display;
        return this;
    }
    
    public TimeOptionBuilder from(String from) {
        this.internal.from = from;
        return this;
    }
    
    public TimeOptionBuilder to(String to) {
        this.internal.to = to;
        return this;
    }
    public TimeOption build() {
        return this.internal;
    }
}
