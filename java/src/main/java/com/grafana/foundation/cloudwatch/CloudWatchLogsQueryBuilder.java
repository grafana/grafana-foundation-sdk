// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.cloudwatch;

import java.util.List;
import com.grafana.foundation.dashboard.DataSourceRef;

public class CloudWatchLogsQueryBuilder implements com.grafana.foundation.cog.Builder<com.grafana.foundation.cog.variants.Dataquery> {
    protected final CloudWatchLogsQuery internal;
    
    public CloudWatchLogsQueryBuilder() {
        this.internal = new CloudWatchLogsQuery();
    }
    public CloudWatchLogsQueryBuilder queryMode(CloudWatchQueryMode queryMode) {
        this.internal.queryMode = queryMode;
        return this;
    }
    
    public CloudWatchLogsQueryBuilder id(String id) {
        this.internal.id = id;
        return this;
    }
    
    public CloudWatchLogsQueryBuilder region(String region) {
        this.internal.region = region;
        return this;
    }
    
    public CloudWatchLogsQueryBuilder expression(String expression) {
        this.internal.expression = expression;
        return this;
    }
    
    public CloudWatchLogsQueryBuilder statsGroups(List<String> statsGroups) {
        this.internal.statsGroups = statsGroups;
        return this;
    }
    
    public CloudWatchLogsQueryBuilder logGroups(com.grafana.foundation.cog.Builder<List<LogGroup>> logGroups) {
        this.internal.logGroups = logGroups.build();
        return this;
    }
    
    public CloudWatchLogsQueryBuilder refId(String refId) {
        this.internal.refId = refId;
        return this;
    }
    
    public CloudWatchLogsQueryBuilder hide(Boolean hide) {
        this.internal.hide = hide;
        return this;
    }
    
    public CloudWatchLogsQueryBuilder queryType(String queryType) {
        this.internal.queryType = queryType;
        return this;
    }
    
    public CloudWatchLogsQueryBuilder logGroupNames(List<String> logGroupNames) {
        this.internal.logGroupNames = logGroupNames;
        return this;
    }
    
    public CloudWatchLogsQueryBuilder datasource(DataSourceRef datasource) {
        this.internal.datasource = datasource;
        return this;
    }
    public CloudWatchLogsQuery build() {
        return this.internal;
    }
}
