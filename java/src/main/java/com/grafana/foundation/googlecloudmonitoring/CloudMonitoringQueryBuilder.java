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
        this.internal.timeSeriesList = timeSeriesList.build();
        return this;
    }
    
    public CloudMonitoringQueryBuilder timeSeriesQuery(com.grafana.foundation.cog.Builder<TimeSeriesQuery> timeSeriesQuery) {
        this.internal.timeSeriesQuery = timeSeriesQuery.build();
        return this;
    }
    
    public CloudMonitoringQueryBuilder sloQuery(com.grafana.foundation.cog.Builder<SLOQuery> sloQuery) {
        this.internal.sloQuery = sloQuery.build();
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
