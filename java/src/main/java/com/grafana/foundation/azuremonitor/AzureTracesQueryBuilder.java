// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.azuremonitor;

import java.util.List;
import java.util.LinkedList;

public class AzureTracesQueryBuilder implements com.grafana.foundation.cog.Builder<AzureTracesQuery> {
    protected final AzureTracesQuery internal;
    
    public AzureTracesQueryBuilder() {
        this.internal = new AzureTracesQuery();
    }
    public AzureTracesQueryBuilder resultFormat(ResultFormat resultFormat) {
        this.internal.resultFormat = resultFormat;
        return this;
    }
    
    public AzureTracesQueryBuilder resources(List<String> resources) {
        this.internal.resources = resources;
        return this;
    }
    
    public AzureTracesQueryBuilder operationId(String operationId) {
        this.internal.operationId = operationId;
        return this;
    }
    
    public AzureTracesQueryBuilder traceTypes(List<String> traceTypes) {
        this.internal.traceTypes = traceTypes;
        return this;
    }
    
    public AzureTracesQueryBuilder filters(List<com.grafana.foundation.cog.Builder<AzureTracesFilter>> filters) {
        List<AzureTracesFilter> filtersResources = new LinkedList<>();
        for (com.grafana.foundation.cog.Builder<AzureTracesFilter> r1 : filters) {
                AzureTracesFilter filtersDepth1 = r1.build();
                filtersResources.add(filtersDepth1); 
        }
        this.internal.filters = filtersResources;
        return this;
    }
    
    public AzureTracesQueryBuilder query(String query) {
        this.internal.query = query;
        return this;
    }
    public AzureTracesQuery build() {
        return this.internal;
    }
}
