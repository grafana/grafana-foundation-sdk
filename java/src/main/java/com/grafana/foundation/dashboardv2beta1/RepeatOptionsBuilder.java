// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.dashboardv2beta1;


public class RepeatOptionsBuilder implements com.grafana.foundation.cog.Builder<RepeatOptions> {
    protected final RepeatOptions internal;
    
    public RepeatOptionsBuilder() {
        this.internal = new RepeatOptions();
    }
    public RepeatOptionsBuilder value(String value) {
        this.internal.value = value;
        return this;
    }
    
    public RepeatOptionsBuilder direction(RepeatOptionsDirection direction) {
        this.internal.direction = direction;
        return this;
    }
    
    public RepeatOptionsBuilder maxPerRow(Long maxPerRow) {
        this.internal.maxPerRow = maxPerRow;
        return this;
    }
    public RepeatOptions build() {
        return this.internal;
    }
}
