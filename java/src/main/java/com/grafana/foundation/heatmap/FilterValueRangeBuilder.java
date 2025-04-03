// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.heatmap;


public class FilterValueRangeBuilder implements com.grafana.foundation.cog.Builder<FilterValueRange> {
    protected final FilterValueRange internal;
    
    public FilterValueRangeBuilder() {
        this.internal = new FilterValueRange();
    }
    public FilterValueRangeBuilder le(Float le) {
        this.internal.le = le;
        return this;
    }
    
    public FilterValueRangeBuilder ge(Float ge) {
        this.internal.ge = ge;
        return this;
    }
    public FilterValueRange build() {
        return this.internal;
    }
}
