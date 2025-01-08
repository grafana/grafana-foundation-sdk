// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.expr;

import java.util.List;
import com.grafana.foundation.dashboard.DataSourceRef;

public class TypeClassicConditionsBuilder implements com.grafana.foundation.cog.Builder<com.grafana.foundation.cog.variants.Dataquery> {
    protected final TypeClassicConditions internal;
    
    public TypeClassicConditionsBuilder() {
        this.internal = new TypeClassicConditions();
    this.internal.type = "classic_conditions";
    }
    public TypeClassicConditionsBuilder conditions(com.grafana.foundation.cog.Builder<List<ExprTypeClassicConditionsConditions>> conditions) {
    this.internal.conditions = conditions.build();
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
    this.internal.resultAssertions = resultAssertions.build();
        return this;
    }
    
    public TypeClassicConditionsBuilder timeRange(com.grafana.foundation.cog.Builder<ExprTypeClassicConditionsTimeRange> timeRange) {
    this.internal.timeRange = timeRange.build();
        return this;
    }
    public TypeClassicConditions build() {
        return this.internal;
    }
}
