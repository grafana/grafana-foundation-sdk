// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.expr;


public class ExprTypeSqlTimeRangeBuilder implements com.grafana.foundation.cog.Builder<ExprTypeSqlTimeRange> {
    protected final ExprTypeSqlTimeRange internal;
    
    public ExprTypeSqlTimeRangeBuilder() {
        this.internal = new ExprTypeSqlTimeRange();
    }
    public ExprTypeSqlTimeRangeBuilder from(String from) {
    this.internal.from = from;
        return this;
    }
    
    public ExprTypeSqlTimeRangeBuilder to(String to) {
    this.internal.to = to;
        return this;
    }
    public ExprTypeSqlTimeRange build() {
        return this.internal;
    }
}
