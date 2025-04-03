// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.cloudwatch;

import java.util.Map;
import java.util.List;

public class MetricStatBuilder implements com.grafana.foundation.cog.Builder<MetricStat> {
    protected final MetricStat internal;
    
    public MetricStatBuilder() {
        this.internal = new MetricStat();
    }
    public MetricStatBuilder region(String region) {
        this.internal.region = region;
        return this;
    }
    
    public MetricStatBuilder namespace(String namespace) {
        this.internal.namespace = namespace;
        return this;
    }
    
    public MetricStatBuilder metricName(String metricName) {
        this.internal.metricName = metricName;
        return this;
    }
    
    public MetricStatBuilder dimensions(Map<String, StringOrArrayOfString> dimensions) {
        this.internal.dimensions = dimensions;
        return this;
    }
    
    public MetricStatBuilder matchExact(Boolean matchExact) {
        this.internal.matchExact = matchExact;
        return this;
    }
    
    public MetricStatBuilder period(String period) {
        this.internal.period = period;
        return this;
    }
    
    public MetricStatBuilder accountId(String accountId) {
        this.internal.accountId = accountId;
        return this;
    }
    
    public MetricStatBuilder statistic(String statistic) {
        this.internal.statistic = statistic;
        return this;
    }
    
    public MetricStatBuilder statistics(List<String> statistics) {
        this.internal.statistics = statistics;
        return this;
    }
    public MetricStat build() {
        return this.internal;
    }
}
