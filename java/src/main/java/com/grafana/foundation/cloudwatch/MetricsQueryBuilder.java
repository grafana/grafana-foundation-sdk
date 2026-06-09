// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.cloudwatch;

import java.util.Map;
import com.grafana.foundation.common.DataSourceRef;
import java.util.List;

public class MetricsQueryBuilder implements com.grafana.foundation.cog.Builder<MetricsQuery> {
    protected final MetricsQuery internal;
    
    public MetricsQueryBuilder() {
        this.internal = new MetricsQuery();
    }
    public MetricsQueryBuilder queryMode(QueryMode queryMode) {
        this.internal.queryMode = queryMode;
        return this;
    }
    
    public MetricsQueryBuilder metricQueryType(MetricQueryType metricQueryType) {
        this.internal.metricQueryType = metricQueryType;
        return this;
    }
    
    public MetricsQueryBuilder metricEditorMode(MetricEditorMode metricEditorMode) {
        this.internal.metricEditorMode = metricEditorMode;
        return this;
    }
    
    public MetricsQueryBuilder id(String id) {
        this.internal.id = id;
        return this;
    }
    
    public MetricsQueryBuilder alias(String alias) {
        this.internal.alias = alias;
        return this;
    }
    
    public MetricsQueryBuilder label(String label) {
        this.internal.label = label;
        return this;
    }
    
    public MetricsQueryBuilder expression(String expression) {
        this.internal.expression = expression;
        return this;
    }
    
    public MetricsQueryBuilder sqlExpression(String sqlExpression) {
        this.internal.sqlExpression = sqlExpression;
        return this;
    }
    
    public MetricsQueryBuilder refId(String refId) {
        this.internal.refId = refId;
        return this;
    }
    
    public MetricsQueryBuilder hide(Boolean hide) {
        this.internal.hide = hide;
        return this;
    }
    
    public MetricsQueryBuilder queryType(String queryType) {
        this.internal.queryType = queryType;
        return this;
    }
    
    public MetricsQueryBuilder region(String region) {
        this.internal.region = region;
        return this;
    }
    
    public MetricsQueryBuilder namespace(String namespace) {
        this.internal.namespace = namespace;
        return this;
    }
    
    public MetricsQueryBuilder metricName(String metricName) {
        this.internal.metricName = metricName;
        return this;
    }
    
    public MetricsQueryBuilder dimensions(Map<String, StringOrArrayOfString> dimensions) {
        this.internal.dimensions = dimensions;
        return this;
    }
    
    public MetricsQueryBuilder matchExact(Boolean matchExact) {
        this.internal.matchExact = matchExact;
        return this;
    }
    
    public MetricsQueryBuilder period(String period) {
        this.internal.period = period;
        return this;
    }
    
    public MetricsQueryBuilder accountId(String accountId) {
        this.internal.accountId = accountId;
        return this;
    }
    
    public MetricsQueryBuilder statistic(String statistic) {
        this.internal.statistic = statistic;
        return this;
    }
    
    public MetricsQueryBuilder sql(com.grafana.foundation.cog.Builder<SQLExpression> sql) {
    SQLExpression sqlResource = sql.build();
        this.internal.sql = sqlResource;
        return this;
    }
    
    public MetricsQueryBuilder datasource(DataSourceRef datasource) {
        this.internal.datasource = datasource;
        return this;
    }
    
    public MetricsQueryBuilder statistics(List<String> statistics) {
        this.internal.statistics = statistics;
        return this;
    }
    public MetricsQuery build() {
        return this.internal;
    }
}
