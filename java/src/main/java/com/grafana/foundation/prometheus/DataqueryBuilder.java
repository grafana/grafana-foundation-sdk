// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.prometheus;

import java.util.List;
import java.util.LinkedList;
import com.grafana.foundation.common.DataSourceRef;

public class DataqueryBuilder implements com.grafana.foundation.cog.Builder<com.grafana.foundation.cog.variants.Dataquery> {
    protected final Dataquery internal;
    
    public DataqueryBuilder() {
        this.internal = new Dataquery();
    }
    public DataqueryBuilder adhocFilters(List<com.grafana.foundation.cog.Builder<AdhocFilters>> adhocFilters) {
        List<AdhocFilters> adhocFiltersResources = new LinkedList<>();
        for (com.grafana.foundation.cog.Builder<AdhocFilters> r1 : adhocFilters) {
                AdhocFilters adhocFiltersDepth1 = r1.build();
                adhocFiltersResources.add(adhocFiltersDepth1); 
        }
        this.internal.adhocFilters = adhocFiltersResources;
        return this;
    }
    
    public DataqueryBuilder datasource(DataSourceRef datasource) {
        this.internal.datasource = datasource;
        return this;
    }
    
    public DataqueryBuilder editorMode(QueryEditorMode editorMode) {
        this.internal.editorMode = editorMode;
        return this;
    }
    
    public DataqueryBuilder exemplar(Boolean exemplar) {
        this.internal.exemplar = exemplar;
        return this;
    }
    
    public DataqueryBuilder expr(String expr) {
        this.internal.expr = expr;
        return this;
    }
    
    public DataqueryBuilder format(PromQueryFormat format) {
        this.internal.format = format;
        return this;
    }
    
    public DataqueryBuilder groupByKeys(List<String> groupByKeys) {
        this.internal.groupByKeys = groupByKeys;
        return this;
    }
    
    public DataqueryBuilder hide(Boolean hide) {
        this.internal.hide = hide;
        return this;
    }
    
    public DataqueryBuilder instant() {
        this.internal.instant = true;
        this.internal.range = false;
        return this;
    }
    
    public DataqueryBuilder intervalFactor(Long intervalFactor) {
        this.internal.intervalFactor = intervalFactor;
        return this;
    }
    
    public DataqueryBuilder intervalMs(Double intervalMs) {
        this.internal.intervalMs = intervalMs;
        return this;
    }
    
    public DataqueryBuilder legendFormat(String legendFormat) {
        this.internal.legendFormat = legendFormat;
        return this;
    }
    
    public DataqueryBuilder maxDataPoints(Long maxDataPoints) {
        this.internal.maxDataPoints = maxDataPoints;
        return this;
    }
    
    public DataqueryBuilder queryType(String queryType) {
        this.internal.queryType = queryType;
        return this;
    }
    
    public DataqueryBuilder range() {
        this.internal.range = true;
        this.internal.instant = false;
        return this;
    }
    
    public DataqueryBuilder refId(String refId) {
        this.internal.refId = refId;
        return this;
    }
    
    public DataqueryBuilder resultAssertions(com.grafana.foundation.cog.Builder<ResultAssertions> resultAssertions) {
    ResultAssertions resultAssertionsResource = resultAssertions.build();
        this.internal.resultAssertions = resultAssertionsResource;
        return this;
    }
    
    public DataqueryBuilder scopes(List<com.grafana.foundation.cog.Builder<Scopes>> scopes) {
        List<Scopes> scopesResources = new LinkedList<>();
        for (com.grafana.foundation.cog.Builder<Scopes> r1 : scopes) {
                Scopes scopesDepth1 = r1.build();
                scopesResources.add(scopesDepth1); 
        }
        this.internal.scopes = scopesResources;
        return this;
    }
    
    public DataqueryBuilder timeRange(com.grafana.foundation.cog.Builder<TimeRange> timeRange) {
    TimeRange timeRangeResource = timeRange.build();
        this.internal.timeRange = timeRangeResource;
        return this;
    }
    
    public DataqueryBuilder interval(String interval) {
        this.internal.interval = interval;
        return this;
    }
    
    public DataqueryBuilder rangeAndInstant() {
        this.internal.range = true;
        this.internal.instant = true;
        return this;
    }
    public Dataquery build() {
        return this.internal;
    }
}
