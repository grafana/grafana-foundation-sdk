// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.expr;


public class ExprTypeMathTimeRangeBuilder implements com.grafana.foundation.cog.Builder<ExprTypeMathTimeRange> {
    protected final ExprTypeMathTimeRange internal;
    
    public ExprTypeMathTimeRangeBuilder() {
        this.internal = new ExprTypeMathTimeRange();
    }
    public ExprTypeMathTimeRangeBuilder from(String from) {
        this.internal.from = from;
        return this;
    }
    
    public ExprTypeMathTimeRangeBuilder to(String to) {
        this.internal.to = to;
        return this;
    }
    public ExprTypeMathTimeRange build() {
        return this.internal;
    }
}
