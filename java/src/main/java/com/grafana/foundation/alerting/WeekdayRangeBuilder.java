// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.alerting;


public class WeekdayRangeBuilder implements com.grafana.foundation.cog.Builder<WeekdayRange> {
    protected final WeekdayRange internal;
    
    public WeekdayRangeBuilder() {
        this.internal = new WeekdayRange();
    }
    public WeekdayRangeBuilder begin(Integer begin) {
        this.internal.begin = begin;
        return this;
    }
    
    public WeekdayRangeBuilder end(Integer end) {
        this.internal.end = end;
        return this;
    }
    public WeekdayRange build() {
        return this.internal;
    }
}
