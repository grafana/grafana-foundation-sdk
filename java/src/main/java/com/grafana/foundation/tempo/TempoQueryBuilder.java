// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.tempo;

import com.grafana.foundation.dashboard.DataSourceRef;
import java.util.List;

public class TempoQueryBuilder implements com.grafana.foundation.cog.Builder<com.grafana.foundation.cog.variants.Dataquery> {
    protected final TempoQuery internal;
    
    public TempoQueryBuilder() {
        this.internal = new TempoQuery();
    }
    public TempoQueryBuilder refId(String refId) {
    this.internal.refId = refId;
        return this;
    }
    
    public TempoQueryBuilder hide(Boolean hide) {
    this.internal.hide = hide;
        return this;
    }
    
    public TempoQueryBuilder queryType(String queryType) {
    this.internal.queryType = queryType;
        return this;
    }
    
    public TempoQueryBuilder query(String query) {
    this.internal.query = query;
        return this;
    }
    
    public TempoQueryBuilder search(String search) {
    this.internal.search = search;
        return this;
    }
    
    public TempoQueryBuilder serviceName(String serviceName) {
    this.internal.serviceName = serviceName;
        return this;
    }
    
    public TempoQueryBuilder spanName(String spanName) {
    this.internal.spanName = spanName;
        return this;
    }
    
    public TempoQueryBuilder minDuration(String minDuration) {
    this.internal.minDuration = minDuration;
        return this;
    }
    
    public TempoQueryBuilder maxDuration(String maxDuration) {
    this.internal.maxDuration = maxDuration;
        return this;
    }
    
    public TempoQueryBuilder serviceMapQuery(String serviceMapQuery) {
    this.internal.serviceMapQuery = serviceMapQuery;
        return this;
    }
    
    public TempoQueryBuilder serviceMapIncludeNamespace(Boolean serviceMapIncludeNamespace) {
    this.internal.serviceMapIncludeNamespace = serviceMapIncludeNamespace;
        return this;
    }
    
    public TempoQueryBuilder limit(Long limit) {
    this.internal.limit = limit;
        return this;
    }
    
    public TempoQueryBuilder datasource(DataSourceRef datasource) {
    this.internal.datasource = datasource;
        return this;
    }
    
    public TempoQueryBuilder filters(com.grafana.foundation.cog.Builder<List<TraceqlFilter>> filters) {
    this.internal.filters = filters.build();
        return this;
    }
    public TempoQuery build() {
        return this.internal;
    }
}
