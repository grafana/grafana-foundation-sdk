// Code generated - EDITING IS FUTILE. DO NOT EDIT.

package com.grafana.foundation.googlecloudmonitoring;

import com.grafana.foundation.dashboard.DataSourceRef;

public class CloudMonitoringQueryBuilder implements com.grafana.foundation.cog.Builder<com.grafana.foundation.cog.variants.Dataquery> {
    protected final CloudMonitoringQuery internal;
    
    public CloudMonitoringQueryBuilder() {
        this.internal = new CloudMonitoringQuery();
    }
    public CloudMonitoringQueryBuilder refId(String refId) {
        this.internal.refId = refId;
        return this;
    }
    
    public CloudMonitoringQueryBuilder hide(Boolean hide) {
        this.internal.hide = hide;
        return this;
    }
    
    public CloudMonitoringQueryBuilder queryType(String queryType) {
        this.internal.queryType = queryType;
        return this;
    }
    
    public CloudMonitoringQueryBuilder aliasBy(String aliasBy) {
        this.internal.aliasBy = aliasBy;
        return this;
    }
    
    public CloudMonitoringQueryBuilder timeSeriesList(com.grafana.foundation.cog.Builder<TimeSeriesList> timeSeriesList) {
    TimeSeriesList timeSeriesListResource = timeSeriesList.build();
        this.internal.timeSeriesList = timeSeriesListResource;
        return this;
    }
    
    public CloudMonitoringQueryBuilder timeSeriesQuery(com.grafana.foundation.cog.Builder<TimeSeriesQuery> timeSeriesQuery) {
    TimeSeriesQuery timeSeriesQueryResource = timeSeriesQuery.build();
        this.internal.timeSeriesQuery = timeSeriesQueryResource;
        return this;
    }
    
    public CloudMonitoringQueryBuilder sloQuery(com.grafana.foundation.cog.Builder<SLOQuery> sloQuery) {
    SLOQuery sloQueryResource = sloQuery.build();
        this.internal.sloQuery = sloQueryResource;
        return this;
    }
    
    public CloudMonitoringQueryBuilder datasource(DataSourceRef datasource) {
        this.internal.datasource = datasource;
        return this;
    }
    
    public CloudMonitoringQueryBuilder intervalMs(Double intervalMs) {
        this.internal.intervalMs = intervalMs;
        return this;
    }
    public CloudMonitoringQuery build() {
        return this.internal;
    }
}
