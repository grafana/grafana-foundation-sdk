// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.alerting;


public class MonthRangeBuilder implements com.grafana.foundation.cog.Builder<MonthRange> {
    protected final MonthRange internal;
    
    public MonthRangeBuilder() {
        this.internal = new MonthRange();
    }
    public MonthRangeBuilder begin(Integer begin) {
        this.internal.begin = begin;
        return this;
    }
    
    public MonthRangeBuilder end(Integer end) {
        this.internal.end = end;
        return this;
    }
    public MonthRange build() {
        return this.internal;
    }
}
