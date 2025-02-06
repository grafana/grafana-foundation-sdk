// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.expr;

import java.util.List;
import com.grafana.foundation.dashboard.DataSourceRef;

public class TypeThresholdBuilder implements com.grafana.foundation.cog.Builder<com.grafana.foundation.cog.variants.Dataquery> {
    protected final TypeThreshold internal;
    
    public TypeThresholdBuilder() {
        this.internal = new TypeThreshold();
        this.internal.type = "threshold";
    }
    public TypeThresholdBuilder conditions(com.grafana.foundation.cog.Builder<List<ExprTypeThresholdConditions>> conditions) {
        this.internal.conditions = conditions.build();
        return this;
    }
    
    public TypeThresholdBuilder datasource(DataSourceRef datasource) {
        this.internal.datasource = datasource;
        return this;
    }
    
    public TypeThresholdBuilder expression(String expression) {
        if (!(expression.length() >= 1)) {
            throw new IllegalArgumentException("expression.length() must be >= 1");
        }
        this.internal.expression = expression;
        return this;
    }
    
    public TypeThresholdBuilder hide(Boolean hide) {
        this.internal.hide = hide;
        return this;
    }
    
    public TypeThresholdBuilder intervalMs(Double intervalMs) {
        this.internal.intervalMs = intervalMs;
        return this;
    }
    
    public TypeThresholdBuilder maxDataPoints(Long maxDataPoints) {
        this.internal.maxDataPoints = maxDataPoints;
        return this;
    }
    
    public TypeThresholdBuilder queryType(String queryType) {
        this.internal.queryType = queryType;
        return this;
    }
    
    public TypeThresholdBuilder refId(String refId) {
        this.internal.refId = refId;
        return this;
    }
    
    public TypeThresholdBuilder resultAssertions(com.grafana.foundation.cog.Builder<ExprTypeThresholdResultAssertions> resultAssertions) {
        this.internal.resultAssertions = resultAssertions.build();
        return this;
    }
    
    public TypeThresholdBuilder timeRange(com.grafana.foundation.cog.Builder<ExprTypeThresholdTimeRange> timeRange) {
        this.internal.timeRange = timeRange.build();
        return this;
    }
    public TypeThreshold build() {
        return this.internal;
    }
}
