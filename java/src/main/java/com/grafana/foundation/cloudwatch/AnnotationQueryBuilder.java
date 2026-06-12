// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.cloudwatch;

import java.util.Map;
import com.grafana.foundation.common.DataSourceRef;
import java.util.List;

/**
 * Shape of a CloudWatch Annotation query
 * TS type is CloudWatchDefaultQuery = Omit<CloudWatchLogsQuery, 'queryMode'> & CloudWatchMetricsQuery, declared in veneer
 * #CloudWatchDefaultQuery: #CloudWatchLogsQuery & #CloudWatchMetricsQuery @cuetsy(kind="type")
 */
public class AnnotationQueryBuilder implements com.grafana.foundation.cog.Builder<AnnotationQuery> {
    protected final AnnotationQuery internal;
    
    public AnnotationQueryBuilder() {
        this.internal = new AnnotationQuery();
    }
    public AnnotationQueryBuilder queryMode(QueryMode queryMode) {
        this.internal.queryMode = queryMode;
        return this;
    }
    
    public AnnotationQueryBuilder prefixMatching(Boolean prefixMatching) {
        this.internal.prefixMatching = prefixMatching;
        return this;
    }
    
    public AnnotationQueryBuilder actionPrefix(String actionPrefix) {
        this.internal.actionPrefix = actionPrefix;
        return this;
    }
    
    public AnnotationQueryBuilder refId(String refId) {
        this.internal.refId = refId;
        return this;
    }
    
    public AnnotationQueryBuilder hide(Boolean hide) {
        this.internal.hide = hide;
        return this;
    }
    
    public AnnotationQueryBuilder queryType(String queryType) {
        this.internal.queryType = queryType;
        return this;
    }
    
    public AnnotationQueryBuilder region(String region) {
        this.internal.region = region;
        return this;
    }
    
    public AnnotationQueryBuilder namespace(String namespace) {
        this.internal.namespace = namespace;
        return this;
    }
    
    public AnnotationQueryBuilder metricName(String metricName) {
        this.internal.metricName = metricName;
        return this;
    }
    
    public AnnotationQueryBuilder dimensions(Map<String, StringOrArrayOfString> dimensions) {
        this.internal.dimensions = dimensions;
        return this;
    }
    
    public AnnotationQueryBuilder matchExact(Boolean matchExact) {
        this.internal.matchExact = matchExact;
        return this;
    }
    
    public AnnotationQueryBuilder period(String period) {
        this.internal.period = period;
        return this;
    }
    
    public AnnotationQueryBuilder accountId(String accountId) {
        this.internal.accountId = accountId;
        return this;
    }
    
    public AnnotationQueryBuilder statistic(String statistic) {
        this.internal.statistic = statistic;
        return this;
    }
    
    public AnnotationQueryBuilder alarmNamePrefix(String alarmNamePrefix) {
        this.internal.alarmNamePrefix = alarmNamePrefix;
        return this;
    }
    
    public AnnotationQueryBuilder datasource(DataSourceRef datasource) {
        this.internal.datasource = datasource;
        return this;
    }
    
    public AnnotationQueryBuilder statistics(List<String> statistics) {
        this.internal.statistics = statistics;
        return this;
    }
    public AnnotationQuery build() {
        return this.internal;
    }
}
