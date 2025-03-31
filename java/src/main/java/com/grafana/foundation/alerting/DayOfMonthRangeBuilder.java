// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.alerting;


public class DayOfMonthRangeBuilder implements com.grafana.foundation.cog.Builder<DayOfMonthRange> {
    protected final DayOfMonthRange internal;
    
    public DayOfMonthRangeBuilder() {
        this.internal = new DayOfMonthRange();
    }
    public DayOfMonthRangeBuilder begin(Integer begin) {
        this.internal.begin = begin;
        return this;
    }
    
    public DayOfMonthRangeBuilder end(Integer end) {
        this.internal.end = end;
        return this;
    }
    public DayOfMonthRange build() {
        return this.internal;
    }
}
