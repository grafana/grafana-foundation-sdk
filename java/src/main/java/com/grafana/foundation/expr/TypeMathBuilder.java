// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.expr;

import com.grafana.foundation.dashboard.DataSourceRef;

public class TypeMathBuilder implements com.grafana.foundation.cog.Builder<com.grafana.foundation.cog.variants.Dataquery> {
    protected final TypeMath internal;
    
    public TypeMathBuilder() {
        this.internal = new TypeMath();
        this.internal.type = "math";
    }
    public TypeMathBuilder datasource(DataSourceRef datasource) {
        this.internal.datasource = datasource;
        return this;
    }
    
    public TypeMathBuilder expression(String expression) {
        if (!(expression.length() >= 1)) {
            throw new IllegalArgumentException("expression.length() must be >= 1");
        }
        this.internal.expression = expression;
        return this;
    }
    
    public TypeMathBuilder hide(Boolean hide) {
        this.internal.hide = hide;
        return this;
    }
    
    public TypeMathBuilder intervalMs(Double intervalMs) {
        this.internal.intervalMs = intervalMs;
        return this;
    }
    
    public TypeMathBuilder maxDataPoints(Long maxDataPoints) {
        this.internal.maxDataPoints = maxDataPoints;
        return this;
    }
    
    public TypeMathBuilder queryType(String queryType) {
        this.internal.queryType = queryType;
        return this;
    }
    
    public TypeMathBuilder refId(String refId) {
        this.internal.refId = refId;
        return this;
    }
    
    public TypeMathBuilder resultAssertions(com.grafana.foundation.cog.Builder<ExprTypeMathResultAssertions> resultAssertions) {
    ExprTypeMathResultAssertions resultAssertionsResource = resultAssertions.build();
        this.internal.resultAssertions = resultAssertionsResource;
        return this;
    }
    
    public TypeMathBuilder timeRange(com.grafana.foundation.cog.Builder<ExprTypeMathTimeRange> timeRange) {
    ExprTypeMathTimeRange timeRangeResource = timeRange.build();
        this.internal.timeRange = timeRangeResource;
        return this;
    }
    public TypeMath build() {
        return this.internal;
    }
}
