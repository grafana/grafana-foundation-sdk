// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.azuremonitor;

import java.util.List;

public class AzureLogsQueryBuilder implements com.grafana.foundation.cog.Builder<AzureLogsQuery> {
    protected final AzureLogsQuery internal;
    
    public AzureLogsQueryBuilder() {
        this.internal = new AzureLogsQuery();
    }
    public AzureLogsQueryBuilder query(String query) {
        this.internal.query = query;
        return this;
    }
    
    public AzureLogsQueryBuilder resultFormat(ResultFormat resultFormat) {
        this.internal.resultFormat = resultFormat;
        return this;
    }
    
    public AzureLogsQueryBuilder resources(List<String> resources) {
        this.internal.resources = resources;
        return this;
    }
    
    public AzureLogsQueryBuilder intersectTime(Boolean intersectTime) {
        this.internal.intersectTime = intersectTime;
        return this;
    }
    
    public AzureLogsQueryBuilder workspace(String workspace) {
        this.internal.workspace = workspace;
        return this;
    }
    
    public AzureLogsQueryBuilder resource(String resource) {
        this.internal.resource = resource;
        return this;
    }
    public AzureLogsQuery build() {
        return this.internal;
    }
}
