// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.expr;


public class ExprTypeResampleTimeRangeBuilder implements com.grafana.foundation.cog.Builder<ExprTypeResampleTimeRange> {
    protected final ExprTypeResampleTimeRange internal;
    
    public ExprTypeResampleTimeRangeBuilder() {
        this.internal = new ExprTypeResampleTimeRange();
    }
    public ExprTypeResampleTimeRangeBuilder from(String from) {
    this.internal.from = from;
        return this;
    }
    
    public ExprTypeResampleTimeRangeBuilder to(String to) {
    this.internal.to = to;
        return this;
    }
    public ExprTypeResampleTimeRange build() {
        return this.internal;
    }
}
