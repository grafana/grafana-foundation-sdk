// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.expr;

import com.grafana.foundation.dashboard.DataSourceRef;

public class TypeSqlBuilder implements com.grafana.foundation.cog.Builder<com.grafana.foundation.cog.variants.Dataquery> {
    protected final TypeSql internal;
    
    public TypeSqlBuilder() {
        this.internal = new TypeSql();
        this.internal.type = "sql";
    }
    public TypeSqlBuilder datasource(DataSourceRef datasource) {
        this.internal.datasource = datasource;
        return this;
    }
    
    public TypeSqlBuilder expression(String expression) {
        if (!(expression.length() >= 1)) {
            throw new IllegalArgumentException("expression.length() must be >= 1");
        }
        this.internal.expression = expression;
        return this;
    }
    
    public TypeSqlBuilder hide(Boolean hide) {
        this.internal.hide = hide;
        return this;
    }
    
    public TypeSqlBuilder intervalMs(Double intervalMs) {
        this.internal.intervalMs = intervalMs;
        return this;
    }
    
    public TypeSqlBuilder maxDataPoints(Long maxDataPoints) {
        this.internal.maxDataPoints = maxDataPoints;
        return this;
    }
    
    public TypeSqlBuilder queryType(String queryType) {
        this.internal.queryType = queryType;
        return this;
    }
    
    public TypeSqlBuilder refId(String refId) {
        this.internal.refId = refId;
        return this;
    }
    
    public TypeSqlBuilder resultAssertions(com.grafana.foundation.cog.Builder<ExprTypeSqlResultAssertions> resultAssertions) {
        this.internal.resultAssertions = resultAssertions.build();
        return this;
    }
    
    public TypeSqlBuilder timeRange(com.grafana.foundation.cog.Builder<ExprTypeSqlTimeRange> timeRange) {
        this.internal.timeRange = timeRange.build();
        return this;
    }
    public TypeSql build() {
        return this.internal;
    }
}
