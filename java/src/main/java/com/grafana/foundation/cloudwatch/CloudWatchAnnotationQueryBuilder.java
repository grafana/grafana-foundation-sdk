// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.cloudwatch;

import java.util.Map;
import com.grafana.foundation.dashboard.DataSourceRef;
import java.util.List;

public class CloudWatchAnnotationQueryBuilder implements com.grafana.foundation.cog.Builder<com.grafana.foundation.cog.variants.Dataquery> {
    protected final CloudWatchAnnotationQuery internal;
    
    public CloudWatchAnnotationQueryBuilder() {
        this.internal = new CloudWatchAnnotationQuery();
    }
    public CloudWatchAnnotationQueryBuilder queryMode(CloudWatchQueryMode queryMode) {
    this.internal.queryMode = queryMode;
        return this;
    }
    
    public CloudWatchAnnotationQueryBuilder prefixMatching(Boolean prefixMatching) {
    this.internal.prefixMatching = prefixMatching;
        return this;
    }
    
    public CloudWatchAnnotationQueryBuilder actionPrefix(String actionPrefix) {
    this.internal.actionPrefix = actionPrefix;
        return this;
    }
    
    public CloudWatchAnnotationQueryBuilder refId(String refId) {
    this.internal.refId = refId;
        return this;
    }
    
    public CloudWatchAnnotationQueryBuilder hide(Boolean hide) {
    this.internal.hide = hide;
        return this;
    }
    
    public CloudWatchAnnotationQueryBuilder queryType(String queryType) {
    this.internal.queryType = queryType;
        return this;
    }
    
    public CloudWatchAnnotationQueryBuilder region(String region) {
    this.internal.region = region;
        return this;
    }
    
    public CloudWatchAnnotationQueryBuilder namespace(String namespace) {
    this.internal.namespace = namespace;
        return this;
    }
    
    public CloudWatchAnnotationQueryBuilder metricName(String metricName) {
    this.internal.metricName = metricName;
        return this;
    }
    
    public CloudWatchAnnotationQueryBuilder dimensions(Map<String, StringOrArrayOfString> dimensions) {
    this.internal.dimensions = dimensions;
        return this;
    }
    
    public CloudWatchAnnotationQueryBuilder matchExact(Boolean matchExact) {
    this.internal.matchExact = matchExact;
        return this;
    }
    
    public CloudWatchAnnotationQueryBuilder period(String period) {
    this.internal.period = period;
        return this;
    }
    
    public CloudWatchAnnotationQueryBuilder accountId(String accountId) {
    this.internal.accountId = accountId;
        return this;
    }
    
    public CloudWatchAnnotationQueryBuilder statistic(String statistic) {
    this.internal.statistic = statistic;
        return this;
    }
    
    public CloudWatchAnnotationQueryBuilder alarmNamePrefix(String alarmNamePrefix) {
    this.internal.alarmNamePrefix = alarmNamePrefix;
        return this;
    }
    
    public CloudWatchAnnotationQueryBuilder datasource(DataSourceRef datasource) {
    this.internal.datasource = datasource;
        return this;
    }
    
    public CloudWatchAnnotationQueryBuilder statistics(List<String> statistics) {
    this.internal.statistics = statistics;
        return this;
    }
    public CloudWatchAnnotationQuery build() {
        return this.internal;
    }
}
