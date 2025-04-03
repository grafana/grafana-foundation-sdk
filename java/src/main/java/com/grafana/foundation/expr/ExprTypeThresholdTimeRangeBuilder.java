// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.expr;


public class ExprTypeThresholdTimeRangeBuilder implements com.grafana.foundation.cog.Builder<ExprTypeThresholdTimeRange> {
    protected final ExprTypeThresholdTimeRange internal;
    
    public ExprTypeThresholdTimeRangeBuilder() {
        this.internal = new ExprTypeThresholdTimeRange();
    }
    public ExprTypeThresholdTimeRangeBuilder from(String from) {
        this.internal.from = from;
        return this;
    }
    
    public ExprTypeThresholdTimeRangeBuilder to(String to) {
        this.internal.to = to;
        return this;
    }
    public ExprTypeThresholdTimeRange build() {
        return this.internal;
    }
}
