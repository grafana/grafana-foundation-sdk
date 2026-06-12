// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.cloudwatch;

import java.util.List;
import java.util.LinkedList;
import com.grafana.foundation.common.DataSourceRef;

/**
 * Shape of a CloudWatch Logs query
 */
public class LogsQueryBuilder implements com.grafana.foundation.cog.Builder<LogsQuery> {
    protected final LogsQuery internal;
    
    public LogsQueryBuilder() {
        this.internal = new LogsQuery();
    }
    public LogsQueryBuilder queryMode(QueryMode queryMode) {
        this.internal.queryMode = queryMode;
        return this;
    }
    
    public LogsQueryBuilder id(String id) {
        this.internal.id = id;
        return this;
    }
    
    public LogsQueryBuilder region(String region) {
        this.internal.region = region;
        return this;
    }
    
    public LogsQueryBuilder expression(String expression) {
        this.internal.expression = expression;
        return this;
    }
    
    public LogsQueryBuilder statsGroups(List<String> statsGroups) {
        this.internal.statsGroups = statsGroups;
        return this;
    }
    
    public LogsQueryBuilder logGroups(List<com.grafana.foundation.cog.Builder<LogGroup>> logGroups) {
        List<LogGroup> logGroupsResources = new LinkedList<>();
        for (com.grafana.foundation.cog.Builder<LogGroup> r1 : logGroups) {
                LogGroup logGroupsDepth1 = r1.build();
                logGroupsResources.add(logGroupsDepth1); 
        }
        this.internal.logGroups = logGroupsResources;
        return this;
    }
    
    public LogsQueryBuilder logGroupNames(List<String> logGroupNames) {
        this.internal.logGroupNames = logGroupNames;
        return this;
    }
    
    public LogsQueryBuilder refId(String refId) {
        this.internal.refId = refId;
        return this;
    }
    
    public LogsQueryBuilder hide(Boolean hide) {
        this.internal.hide = hide;
        return this;
    }
    
    public LogsQueryBuilder queryType(String queryType) {
        this.internal.queryType = queryType;
        return this;
    }
    
    public LogsQueryBuilder queryLanguage(LogsQueryLanguage queryLanguage) {
        this.internal.queryLanguage = queryLanguage;
        return this;
    }
    
    public LogsQueryBuilder datasource(DataSourceRef datasource) {
        this.internal.datasource = datasource;
        return this;
    }
    public LogsQuery build() {
        return this.internal;
    }
}
