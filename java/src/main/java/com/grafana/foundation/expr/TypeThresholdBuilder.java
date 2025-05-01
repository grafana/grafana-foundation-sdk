// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.expr;

import java.util.List;
import java.util.LinkedList;
import com.grafana.foundation.dashboard.DataSourceRef;

public class TypeThresholdBuilder implements com.grafana.foundation.cog.Builder<com.grafana.foundation.cog.variants.Dataquery> {
    protected final TypeThreshold internal;
    
    public TypeThresholdBuilder() {
        this.internal = new TypeThreshold();
        this.internal.type = "threshold";
    }
    public TypeThresholdBuilder conditions(List<com.grafana.foundation.cog.Builder<ExprTypeThresholdConditions>> conditions) {
        List<ExprTypeThresholdConditions> conditionsResources = new LinkedList<>();
        for (com.grafana.foundation.cog.Builder<ExprTypeThresholdConditions> r1 : conditions) {
                ExprTypeThresholdConditions conditionsDepth1 = r1.build();
                conditionsResources.add(conditionsDepth1); 
        }
        this.internal.conditions = conditionsResources;
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
    ExprTypeThresholdResultAssertions resultAssertionsResource = resultAssertions.build();
        this.internal.resultAssertions = resultAssertionsResource;
        return this;
    }
    
    public TypeThresholdBuilder timeRange(com.grafana.foundation.cog.Builder<ExprTypeThresholdTimeRange> timeRange) {
    ExprTypeThresholdTimeRange timeRangeResource = timeRange.build();
        this.internal.timeRange = timeRangeResource;
        return this;
    }
    public TypeThreshold build() {
        return this.internal;
    }
}
