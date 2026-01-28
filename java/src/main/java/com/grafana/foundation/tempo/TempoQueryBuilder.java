// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.tempo;

import java.util.List;
import java.util.LinkedList;
import com.grafana.foundation.common.DataSourceRef;

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
    
    public TempoQueryBuilder spss(Long spss) {
        this.internal.spss = spss;
        return this;
    }
    
    public TempoQueryBuilder filters(List<com.grafana.foundation.cog.Builder<TraceqlFilter>> filters) {
        List<TraceqlFilter> filtersResources = new LinkedList<>();
        for (com.grafana.foundation.cog.Builder<TraceqlFilter> r1 : filters) {
                TraceqlFilter filtersDepth1 = r1.build();
                filtersResources.add(filtersDepth1); 
        }
        this.internal.filters = filtersResources;
        return this;
    }
    
    public TempoQueryBuilder groupBy(List<com.grafana.foundation.cog.Builder<TraceqlFilter>> groupBy) {
        List<TraceqlFilter> groupByResources = new LinkedList<>();
        for (com.grafana.foundation.cog.Builder<TraceqlFilter> r1 : groupBy) {
                TraceqlFilter groupByDepth1 = r1.build();
                groupByResources.add(groupByDepth1); 
        }
        this.internal.groupBy = groupByResources;
        return this;
    }
    
    public TempoQueryBuilder datasource(DataSourceRef datasource) {
        this.internal.datasource = datasource;
        return this;
    }
    
    public TempoQueryBuilder tableType(SearchTableType tableType) {
        this.internal.tableType = tableType;
        return this;
    }
    public TempoQuery build() {
        return this.internal;
    }
}
