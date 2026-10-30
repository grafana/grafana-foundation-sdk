// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.alertingv1beta1;


public class ExpressionBuilder implements com.grafana.foundation.cog.Builder<Expression> {
    protected final Expression internal;
    
    public ExpressionBuilder() {
        this.internal = new Expression();
    }
    public ExpressionBuilder queryType(String queryType) {
        this.internal.queryType = queryType;
        return this;
    }
    
    public ExpressionBuilder relativeTimeRange(com.grafana.foundation.cog.Builder<RelativeTimeRange> relativeTimeRange) {
    RelativeTimeRange relativeTimeRangeResource = relativeTimeRange.build();
        this.internal.relativeTimeRange = relativeTimeRangeResource;
        return this;
    }
    
    public ExpressionBuilder datasourceUID(String datasourceUID) {
        if (!datasourceUID.matches("^[a-zA-Z0-9_-]+$")) {
            throw new IllegalArgumentException("datasourceUID must match regex ^[a-zA-Z0-9_-]+$");
        }
        this.internal.datasourceUID = datasourceUID;
        return this;
    }
    
    public ExpressionBuilder model(Object model) {
        this.internal.model = model;
        return this;
    }
    
    public ExpressionBuilder source(Boolean source) {
        this.internal.source = source;
        return this;
    }
    public Expression build() {
        return this.internal;
    }
}
