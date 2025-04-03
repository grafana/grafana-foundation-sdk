// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.cloudwatch;

import java.util.Map;
import com.grafana.foundation.dashboard.DataSourceRef;
import java.util.List;

public class CloudWatchMetricsQueryBuilder implements com.grafana.foundation.cog.Builder<com.grafana.foundation.cog.variants.Dataquery> {
    protected final CloudWatchMetricsQuery internal;
    
    public CloudWatchMetricsQueryBuilder() {
        this.internal = new CloudWatchMetricsQuery();
    }
    public CloudWatchMetricsQueryBuilder queryMode(CloudWatchQueryMode queryMode) {
        this.internal.queryMode = queryMode;
        return this;
    }
    
    public CloudWatchMetricsQueryBuilder metricQueryType(MetricQueryType metricQueryType) {
        this.internal.metricQueryType = metricQueryType;
        return this;
    }
    
    public CloudWatchMetricsQueryBuilder metricEditorMode(MetricEditorMode metricEditorMode) {
        this.internal.metricEditorMode = metricEditorMode;
        return this;
    }
    
    public CloudWatchMetricsQueryBuilder id(String id) {
        this.internal.id = id;
        return this;
    }
    
    public CloudWatchMetricsQueryBuilder alias(String alias) {
        this.internal.alias = alias;
        return this;
    }
    
    public CloudWatchMetricsQueryBuilder label(String label) {
        this.internal.label = label;
        return this;
    }
    
    public CloudWatchMetricsQueryBuilder expression(String expression) {
        this.internal.expression = expression;
        return this;
    }
    
    public CloudWatchMetricsQueryBuilder sqlExpression(String sqlExpression) {
        this.internal.sqlExpression = sqlExpression;
        return this;
    }
    
    public CloudWatchMetricsQueryBuilder refId(String refId) {
        this.internal.refId = refId;
        return this;
    }
    
    public CloudWatchMetricsQueryBuilder hide(Boolean hide) {
        this.internal.hide = hide;
        return this;
    }
    
    public CloudWatchMetricsQueryBuilder queryType(String queryType) {
        this.internal.queryType = queryType;
        return this;
    }
    
    public CloudWatchMetricsQueryBuilder region(String region) {
        this.internal.region = region;
        return this;
    }
    
    public CloudWatchMetricsQueryBuilder namespace(String namespace) {
        this.internal.namespace = namespace;
        return this;
    }
    
    public CloudWatchMetricsQueryBuilder metricName(String metricName) {
        this.internal.metricName = metricName;
        return this;
    }
    
    public CloudWatchMetricsQueryBuilder dimensions(Map<String, StringOrArrayOfString> dimensions) {
        this.internal.dimensions = dimensions;
        return this;
    }
    
    public CloudWatchMetricsQueryBuilder matchExact(Boolean matchExact) {
        this.internal.matchExact = matchExact;
        return this;
    }
    
    public CloudWatchMetricsQueryBuilder period(String period) {
        this.internal.period = period;
        return this;
    }
    
    public CloudWatchMetricsQueryBuilder accountId(String accountId) {
        this.internal.accountId = accountId;
        return this;
    }
    
    public CloudWatchMetricsQueryBuilder statistic(String statistic) {
        this.internal.statistic = statistic;
        return this;
    }
    
    public CloudWatchMetricsQueryBuilder sql(com.grafana.foundation.cog.Builder<SQLExpression> sql) {
        this.internal.sql = sql.build();
        return this;
    }
    
    public CloudWatchMetricsQueryBuilder datasource(DataSourceRef datasource) {
        this.internal.datasource = datasource;
        return this;
    }
    
    public CloudWatchMetricsQueryBuilder statistics(List<String> statistics) {
        this.internal.statistics = statistics;
        return this;
    }
    public CloudWatchMetricsQuery build() {
        return this.internal;
    }
}
