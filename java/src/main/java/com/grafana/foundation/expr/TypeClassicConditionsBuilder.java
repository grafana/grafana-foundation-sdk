// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.expr;

import java.util.List;
import java.util.LinkedList;
import com.grafana.foundation.common.DataSourceRef;

public class TypeClassicConditionsBuilder implements com.grafana.foundation.cog.Builder<com.grafana.foundation.cog.variants.Dataquery> {
    protected final TypeClassicConditions internal;
    
    public TypeClassicConditionsBuilder() {
        this.internal = new TypeClassicConditions();
        this.internal.type = "classic_conditions";
    }
    public TypeClassicConditionsBuilder conditions(List<com.grafana.foundation.cog.Builder<ExprTypeClassicConditionsConditions>> conditions) {
        List<ExprTypeClassicConditionsConditions> conditionsResources = new LinkedList<>();
        for (com.grafana.foundation.cog.Builder<ExprTypeClassicConditionsConditions> r1 : conditions) {
                ExprTypeClassicConditionsConditions conditionsDepth1 = r1.build();
                conditionsResources.add(conditionsDepth1); 
        }
        this.internal.conditions = conditionsResources;
        return this;
    }
    
    public TypeClassicConditionsBuilder datasource(DataSourceRef datasource) {
        this.internal.datasource = datasource;
        return this;
    }
    
    public TypeClassicConditionsBuilder hide(Boolean hide) {
        this.internal.hide = hide;
        return this;
    }
    
    public TypeClassicConditionsBuilder intervalMs(Double intervalMs) {
        this.internal.intervalMs = intervalMs;
        return this;
    }
    
    public TypeClassicConditionsBuilder maxDataPoints(Long maxDataPoints) {
        this.internal.maxDataPoints = maxDataPoints;
        return this;
    }
    
    public TypeClassicConditionsBuilder queryType(String queryType) {
        this.internal.queryType = queryType;
        return this;
    }
    
    public TypeClassicConditionsBuilder refId(String refId) {
        this.internal.refId = refId;
        return this;
    }
    
    public TypeClassicConditionsBuilder resultAssertions(com.grafana.foundation.cog.Builder<ExprTypeClassicConditionsResultAssertions> resultAssertions) {
    ExprTypeClassicConditionsResultAssertions resultAssertionsResource = resultAssertions.build();
        this.internal.resultAssertions = resultAssertionsResource;
        return this;
    }
    
    public TypeClassicConditionsBuilder timeRange(com.grafana.foundation.cog.Builder<ExprTypeClassicConditionsTimeRange> timeRange) {
    ExprTypeClassicConditionsTimeRange timeRangeResource = timeRange.build();
        this.internal.timeRange = timeRangeResource;
        return this;
    }
    public TypeClassicConditions build() {
        return this.internal;
    }
}
