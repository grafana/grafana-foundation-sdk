// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.expr;


public class ExprTypeClassicConditionsTimeRangeBuilder implements com.grafana.foundation.cog.Builder<ExprTypeClassicConditionsTimeRange> {
    protected final ExprTypeClassicConditionsTimeRange internal;
    
    public ExprTypeClassicConditionsTimeRangeBuilder() {
        this.internal = new ExprTypeClassicConditionsTimeRange();
    }
    public ExprTypeClassicConditionsTimeRangeBuilder from(String from) {
        this.internal.from = from;
        return this;
    }
    
    public ExprTypeClassicConditionsTimeRangeBuilder to(String to) {
        this.internal.to = to;
        return this;
    }
    public ExprTypeClassicConditionsTimeRange build() {
        return this.internal;
    }
}
