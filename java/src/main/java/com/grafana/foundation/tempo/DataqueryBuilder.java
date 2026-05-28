// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.tempo;

import java.util.List;
import java.util.LinkedList;
import com.grafana.foundation.common.DataSourceRef;

public class DataqueryBuilder implements com.grafana.foundation.cog.Builder<com.grafana.foundation.cog.variants.Dataquery> {
    protected final Dataquery internal;
    
    public DataqueryBuilder() {
        this.internal = new Dataquery();
    }
    public DataqueryBuilder refId(String refId) {
        this.internal.refId = refId;
        return this;
    }
    
    public DataqueryBuilder hide(Boolean hide) {
        this.internal.hide = hide;
        return this;
    }
    
    public DataqueryBuilder queryType(String queryType) {
        this.internal.queryType = queryType;
        return this;
    }
    
    public DataqueryBuilder query(String query) {
        this.internal.query = query;
        return this;
    }
    
    public DataqueryBuilder search(String search) {
        this.internal.search = search;
        return this;
    }
    
    public DataqueryBuilder serviceName(String serviceName) {
        this.internal.serviceName = serviceName;
        return this;
    }
    
    public DataqueryBuilder spanName(String spanName) {
        this.internal.spanName = spanName;
        return this;
    }
    
    public DataqueryBuilder minDuration(String minDuration) {
        this.internal.minDuration = minDuration;
        return this;
    }
    
    public DataqueryBuilder maxDuration(String maxDuration) {
        this.internal.maxDuration = maxDuration;
        return this;
    }
    
    public DataqueryBuilder serviceMapQuery(StringOrArrayOfString serviceMapQuery) {
        this.internal.serviceMapQuery = serviceMapQuery;
        return this;
    }
    
    public DataqueryBuilder serviceMapIncludeNamespace(Boolean serviceMapIncludeNamespace) {
        this.internal.serviceMapIncludeNamespace = serviceMapIncludeNamespace;
        return this;
    }
    
    public DataqueryBuilder limit(Long limit) {
        this.internal.limit = limit;
        return this;
    }
    
    public DataqueryBuilder spss(Long spss) {
        this.internal.spss = spss;
        return this;
    }
    
    public DataqueryBuilder filters(List<com.grafana.foundation.cog.Builder<TraceqlFilter>> filters) {
        List<TraceqlFilter> filtersResources = new LinkedList<>();
        for (com.grafana.foundation.cog.Builder<TraceqlFilter> r1 : filters) {
                TraceqlFilter filtersDepth1 = r1.build();
                filtersResources.add(filtersDepth1); 
        }
        this.internal.filters = filtersResources;
        return this;
    }
    
    public DataqueryBuilder groupBy(List<com.grafana.foundation.cog.Builder<TraceqlFilter>> groupBy) {
        List<TraceqlFilter> groupByResources = new LinkedList<>();
        for (com.grafana.foundation.cog.Builder<TraceqlFilter> r1 : groupBy) {
                TraceqlFilter groupByDepth1 = r1.build();
                groupByResources.add(groupByDepth1); 
        }
        this.internal.groupBy = groupByResources;
        return this;
    }
    
    public DataqueryBuilder tableType(SearchTableType tableType) {
        this.internal.tableType = tableType;
        return this;
    }
    
    public DataqueryBuilder step(String step) {
        this.internal.step = step;
        return this;
    }
    
    public DataqueryBuilder exemplars(Long exemplars) {
        this.internal.exemplars = exemplars;
        return this;
    }
    
    public DataqueryBuilder datasource(DataSourceRef datasource) {
        this.internal.datasource = datasource;
        return this;
    }
    
    public DataqueryBuilder metricsQueryType(MetricsQueryType metricsQueryType) {
        this.internal.metricsQueryType = metricsQueryType;
        return this;
    }
    public Dataquery build() {
        return this.internal;
    }
}
