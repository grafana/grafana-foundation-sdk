// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.expr;


public class ExprTypeReduceTimeRangeBuilder implements com.grafana.foundation.cog.Builder<ExprTypeReduceTimeRange> {
    protected final ExprTypeReduceTimeRange internal;
    
    public ExprTypeReduceTimeRangeBuilder() {
        this.internal = new ExprTypeReduceTimeRange();
    }
    public ExprTypeReduceTimeRangeBuilder from(String from) {
        this.internal.from = from;
        return this;
    }
    
    public ExprTypeReduceTimeRangeBuilder to(String to) {
        this.internal.to = to;
        return this;
    }
    public ExprTypeReduceTimeRange build() {
        return this.internal;
    }
}
