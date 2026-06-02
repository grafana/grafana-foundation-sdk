// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.azuremonitor;

import java.util.List;

public class LogsQueryBuilder implements com.grafana.foundation.cog.Builder<LogsQuery> {
    protected final LogsQuery internal;
    
    public LogsQueryBuilder() {
        this.internal = new LogsQuery();
    }
    public LogsQueryBuilder query(String query) {
        this.internal.query = query;
        return this;
    }
    
    public LogsQueryBuilder resultFormat(ResultFormat resultFormat) {
        this.internal.resultFormat = resultFormat;
        return this;
    }
    
    public LogsQueryBuilder resources(List<String> resources) {
        this.internal.resources = resources;
        return this;
    }
    
    public LogsQueryBuilder dashboardTime(Boolean dashboardTime) {
        this.internal.dashboardTime = dashboardTime;
        return this;
    }
    
    public LogsQueryBuilder timeColumn(String timeColumn) {
        this.internal.timeColumn = timeColumn;
        return this;
    }
    
    public LogsQueryBuilder basicLogsQuery(Boolean basicLogsQuery) {
        this.internal.basicLogsQuery = basicLogsQuery;
        return this;
    }
    
    public LogsQueryBuilder workspace(String workspace) {
        this.internal.workspace = workspace;
        return this;
    }
    
    public LogsQueryBuilder resource(String resource) {
        this.internal.resource = resource;
        return this;
    }
    
    public LogsQueryBuilder intersectTime(Boolean intersectTime) {
        this.internal.intersectTime = intersectTime;
        return this;
    }
    public LogsQuery build() {
        return this.internal;
    }
}
