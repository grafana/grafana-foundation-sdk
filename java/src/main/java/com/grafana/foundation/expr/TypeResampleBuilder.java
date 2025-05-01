// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.expr;

import com.grafana.foundation.dashboard.DataSourceRef;

public class TypeResampleBuilder implements com.grafana.foundation.cog.Builder<com.grafana.foundation.cog.variants.Dataquery> {
    protected final TypeResample internal;
    
    public TypeResampleBuilder() {
        this.internal = new TypeResample();
        this.internal.type = "resample";
    }
    public TypeResampleBuilder datasource(DataSourceRef datasource) {
        this.internal.datasource = datasource;
        return this;
    }
    
    public TypeResampleBuilder downsampler(TypeResampleDownsampler downsampler) {
        this.internal.downsampler = downsampler;
        return this;
    }
    
    public TypeResampleBuilder expression(String expression) {
        if (!(expression.length() >= 1)) {
            throw new IllegalArgumentException("expression.length() must be >= 1");
        }
        this.internal.expression = expression;
        return this;
    }
    
    public TypeResampleBuilder hide(Boolean hide) {
        this.internal.hide = hide;
        return this;
    }
    
    public TypeResampleBuilder intervalMs(Double intervalMs) {
        this.internal.intervalMs = intervalMs;
        return this;
    }
    
    public TypeResampleBuilder maxDataPoints(Long maxDataPoints) {
        this.internal.maxDataPoints = maxDataPoints;
        return this;
    }
    
    public TypeResampleBuilder queryType(String queryType) {
        this.internal.queryType = queryType;
        return this;
    }
    
    public TypeResampleBuilder refId(String refId) {
        this.internal.refId = refId;
        return this;
    }
    
    public TypeResampleBuilder resultAssertions(com.grafana.foundation.cog.Builder<ExprTypeResampleResultAssertions> resultAssertions) {
    ExprTypeResampleResultAssertions resultAssertionsResource = resultAssertions.build();
        this.internal.resultAssertions = resultAssertionsResource;
        return this;
    }
    
    public TypeResampleBuilder timeRange(com.grafana.foundation.cog.Builder<ExprTypeResampleTimeRange> timeRange) {
    ExprTypeResampleTimeRange timeRangeResource = timeRange.build();
        this.internal.timeRange = timeRangeResource;
        return this;
    }
    
    public TypeResampleBuilder upsampler(TypeResampleUpsampler upsampler) {
        this.internal.upsampler = upsampler;
        return this;
    }
    
    public TypeResampleBuilder window(String window) {
        if (!(window.length() >= 1)) {
            throw new IllegalArgumentException("window.length() must be >= 1");
        }
        this.internal.window = window;
        return this;
    }
    public TypeResample build() {
        return this.internal;
    }
}
