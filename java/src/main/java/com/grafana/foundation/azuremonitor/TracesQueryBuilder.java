// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.azuremonitor;

import java.util.List;
import java.util.LinkedList;

public class TracesQueryBuilder implements com.grafana.foundation.cog.Builder<TracesQuery> {
    protected final TracesQuery internal;
    
    public TracesQueryBuilder() {
        this.internal = new TracesQuery();
    }
    public TracesQueryBuilder resultFormat(ResultFormat resultFormat) {
        this.internal.resultFormat = resultFormat;
        return this;
    }
    
    public TracesQueryBuilder resources(List<String> resources) {
        this.internal.resources = resources;
        return this;
    }
    
    public TracesQueryBuilder operationId(String operationId) {
        this.internal.operationId = operationId;
        return this;
    }
    
    public TracesQueryBuilder traceTypes(List<String> traceTypes) {
        this.internal.traceTypes = traceTypes;
        return this;
    }
    
    public TracesQueryBuilder filters(List<com.grafana.foundation.cog.Builder<TracesFilter>> filters) {
        List<TracesFilter> filtersResources = new LinkedList<>();
        for (com.grafana.foundation.cog.Builder<TracesFilter> r1 : filters) {
                TracesFilter filtersDepth1 = r1.build();
                filtersResources.add(filtersDepth1); 
        }
        this.internal.filters = filtersResources;
        return this;
    }
    
    public TracesQueryBuilder query(String query) {
        this.internal.query = query;
        return this;
    }
    public TracesQuery build() {
        return this.internal;
    }
}
