// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.expr;

import com.grafana.foundation.dashboard.DataSourceRef;

public class TypeReduceBuilder implements com.grafana.foundation.cog.Builder<com.grafana.foundation.cog.variants.Dataquery> {
    protected final TypeReduce internal;
    
    public TypeReduceBuilder() {
        this.internal = new TypeReduce();
        this.internal.type = "reduce";
    }
    public TypeReduceBuilder datasource(DataSourceRef datasource) {
        this.internal.datasource = datasource;
        return this;
    }
    
    public TypeReduceBuilder expression(String expression) {
        if (!(expression.length() >= 1)) {
            throw new IllegalArgumentException("expression.length() must be >= 1");
        }
        this.internal.expression = expression;
        return this;
    }
    
    public TypeReduceBuilder hide(Boolean hide) {
        this.internal.hide = hide;
        return this;
    }
    
    public TypeReduceBuilder intervalMs(Double intervalMs) {
        this.internal.intervalMs = intervalMs;
        return this;
    }
    
    public TypeReduceBuilder maxDataPoints(Long maxDataPoints) {
        this.internal.maxDataPoints = maxDataPoints;
        return this;
    }
    
    public TypeReduceBuilder queryType(String queryType) {
        this.internal.queryType = queryType;
        return this;
    }
    
    public TypeReduceBuilder reducer(TypeReduceReducer reducer) {
        this.internal.reducer = reducer;
        return this;
    }
    
    public TypeReduceBuilder refId(String refId) {
        this.internal.refId = refId;
        return this;
    }
    
    public TypeReduceBuilder resultAssertions(com.grafana.foundation.cog.Builder<ExprTypeReduceResultAssertions> resultAssertions) {
        this.internal.resultAssertions = resultAssertions.build();
        return this;
    }
    
    public TypeReduceBuilder settings(com.grafana.foundation.cog.Builder<ExprTypeReduceSettings> settings) {
        this.internal.settings = settings.build();
        return this;
    }
    
    public TypeReduceBuilder timeRange(com.grafana.foundation.cog.Builder<ExprTypeReduceTimeRange> timeRange) {
        this.internal.timeRange = timeRange.build();
        return this;
    }
    public TypeReduce build() {
        return this.internal;
    }
}
