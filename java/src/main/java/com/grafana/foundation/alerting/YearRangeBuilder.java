// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.alerting;


public class YearRangeBuilder implements com.grafana.foundation.cog.Builder<YearRange> {
    protected final YearRange internal;
    
    public YearRangeBuilder() {
        this.internal = new YearRange();
    }
    public YearRangeBuilder begin(Integer begin) {
        this.internal.begin = begin;
        return this;
    }
    
    public YearRangeBuilder end(Integer end) {
        this.internal.end = end;
        return this;
    }
    public YearRange build() {
        return this.internal;
    }
}
